package multiphase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/cannon"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum-optimism/optimism/op-challenger/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	ErrInstructionVerificationFailed = errors.New("instruction verification failed")
	ErrWitnessGenerationFailed       = errors.New("witness generation failed")
	ErrFaultNotFound                 = errors.New("fault not found in segment")
)

// Phase2Verifier performs fine-grained verification of specific instructions
type Phase2Verifier struct {
	client             *ethclient.Client
	lazyLoadingManager *LazyLoadingManager
	logger             log.Logger
}

// NewPhase2Verifier creates a new instance of Phase2Verifier
func NewPhase2Verifier(client *ethclient.Client, llm *LazyLoadingManager) *Phase2Verifier {
	return &Phase2Verifier{
		client:             client,
		lazyLoadingManager: llm,
		logger:             log.New("component", "Phase2Verifier"),
	}
}

// VerifyAndFindFault identifies the exact instruction causing a fault
func (v *Phase2Verifier) VerifyAndFindFault(
	ctx context.Context,
	disputeID *big.Int,
	segmentStart *big.Int,
	segmentEnd *big.Int,
	segmentStateRoot [32]byte,
) (*big.Int, []byte, error) {
	// Perform binary search on the instructions in the segment
	currentStart := new(big.Int).Set(segmentStart)
	currentEnd := new(big.Int).Set(segmentEnd)

	v.logger.Info("Starting fault verification",
		"disputeID", disputeID,
		"segmentStart", segmentStart,
		"segmentEnd", segmentEnd,
		"segmentStateRoot", segmentStateRoot.Hex())

	for currentStart.Cmp(currentEnd) < 0 {
		// If we're down to one instruction, verify it
		if new(big.Int).Sub(currentEnd, currentStart).Cmp(big.NewInt(1)) <= 0 {
			v.logger.Info("Found fault at instruction", "index", currentStart)
			return currentStart, v.generateWitness(ctx, disputeID, currentStart)
		}

		// Find midpoint
		midpoint := new(big.Int).Add(currentStart, currentEnd)
		midpoint.Div(midpoint, big.NewInt(2))

		v.logger.Debug("Checking midpoint", "midpoint", midpoint)

		// Load state at midpoint
		midStateRootBytes, err := v.lazyLoadingManager.LoadState(ctx, disputeID, midpoint)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to load state at midpoint: %w", err)
		}

		midStateRoot := common.BytesToHash(midStateRootBytes)

		// Verify instruction execution from start to midpoint
		valid, err := v.VerifyInstructionExecution(
			ctx,
			disputeID,
			currentStart,
			midpoint,
			segmentStateRoot,
			midStateRoot,
		)
		if err != nil {
			return nil, nil, err
		}

		// Determine which half contains the fault
		if !valid {
			// Fault is in the first half
			v.logger.Debug("Fault found in first half", "start", currentStart, "end", midpoint)
			currentEnd = midpoint
		} else {
			// Fault is in the second half
			v.logger.Debug("Fault found in second half", "start", midpoint, "end", currentEnd)
			currentStart = midpoint
		}
	}

	return nil, nil, ErrFaultNotFound
}

// VerifyInstructionExecution verifies the execution of instructions in a segment
func (v *Phase2Verifier) VerifyInstructionExecution(
	ctx context.Context,
	disputeID *big.Int,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
	expectedEndStateRoot [32]byte,
) (bool, error) {
	// Create a temporary directory for execution
	tempDir, err := os.MkdirTemp("", "cannon-verify-*")
	if err != nil {
		return false, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	v.logger.Debug("Created temporary directory for verification", "dir", tempDir)

	// Load the state at the start index
	startStateBytes, err := v.lazyLoadingManager.LoadState(ctx, disputeID, startIndex)
	if err != nil {
		return false, fmt.Errorf("failed to load start state: %w", err)
	}

	// Verify the start state matches the expected start state root
	if common.BytesToHash(startStateBytes) != startStateRoot {
		return false, fmt.Errorf("start state mismatch: expected %s, got %s",
			startStateRoot.Hex(), common.BytesToHash(startStateBytes).Hex())
	}

	// Create a Cannon executor to run the instructions
	vmConfig := vm.Config{
		// VM Configuration
		VmType:          types.TraceTypeCannon,
		VmBin:           "/usr/local/bin/cannon", // Path to cannon binary
		SnapshotFreq:    1000,                    // Snapshot frequency
		InfoFreq:        1000,                    // Info frequency
		DebugInfo:       true,                    // Enable debug info
		BinarySnapshots: true,                    // Use binary snapshots

		// Host Configuration
		Server:   "/usr/local/bin/op-program", // Path to op-program binary
		L1:       "http://localhost:8545",     // Default L1 RPC URL
		Networks: []string{"optimism-goerli"}, // Default network
	}

	// Create local game inputs
	localInputs := utils.LocalGameInputs{
		L1Head:           common.Hash{},
		L2Head:           common.Hash{},
		L2OutputRoot:     common.Hash{},
		L2SequenceNumber: startIndex,
		L2Claim:          common.BytesToHash(startStateBytes),
	}

	// Create the executor
	logger := log.New("component", "CannonExecutor")
	metricer := metrics.NoopMetrics.ToTypedVmMetrics("cannon")
	oracleServer := vm.NewOpProgramServerExecutor(logger)
	executor := vm.NewExecutor(logger, metricer, vmConfig, oracleServer, "", localInputs)

	v.logger.Debug("Executing instructions", "start", startIndex, "end", endIndex)

	// Execute the instructions from start to end
	if err := executor.DoGenerateProof(ctx, tempDir, startIndex.Uint64(), endIndex.Uint64()); err != nil {
		return false, fmt.Errorf("failed to execute instructions: %w", err)
	}

	// Convert the final state to a proof
	stateConverter := cannon.NewStateConverter(vmConfig)
	proof, _, _, err := stateConverter.ConvertStateToProof(ctx, vm.FinalStatePath(tempDir, vmConfig.BinarySnapshots))
	if err != nil {
		return false, fmt.Errorf("failed to convert state to proof: %w", err)
	}

	v.logger.Debug("Verification completed",
		"expectedRoot", expectedEndStateRoot.Hex(),
		"actualRoot", proof.ClaimValue.Hex())

	// Compare the resulting state root with the expected one
	if proof.ClaimValue != expectedEndStateRoot {
		return false, nil // Verification failed, but not an error
	}

	return true, nil
}

// generateWitness generates witness data for a disputed instruction
func (v *Phase2Verifier) generateWitness(
	ctx context.Context,
	disputeID *big.Int,
	instructionIndex *big.Int,
) ([]byte, error) {
	// Create a temporary directory for execution
	tempDir, err := os.MkdirTemp("", "cannon-witness-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	v.logger.Info("Generating witness for instruction", "index", instructionIndex, "dir", tempDir)

	// Load the state at the instruction index
	stateBytes, err := v.lazyLoadingManager.LoadState(ctx, disputeID, instructionIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to load state: %w", err)
	}

	// Create a Cannon executor configuration
	vmConfig := vm.Config{
		// VM Configuration
		VmType:          types.TraceTypeCannon,
		VmBin:           "/usr/local/bin/cannon", // Path to cannon binary
		SnapshotFreq:    1000,                    // Snapshot frequency
		InfoFreq:        1000,                    // Info frequency
		DebugInfo:       true,                    // Enable debug info
		BinarySnapshots: true,                    // Use binary snapshots

		// Host Configuration
		Server:   "/usr/local/bin/op-program", // Path to op-program binary
		L1:       "http://localhost:8545",     // Default L1 RPC URL
		Networks: []string{"optimism-goerli"}, // Default network
	}

	// Create the state converter
	stateConverter := cannon.NewStateConverter(vmConfig)

	// Create local game inputs
	localInputs := utils.LocalGameInputs{
		L1Head:           common.Hash{},
		L2Head:           common.Hash{},
		L2OutputRoot:     common.Hash{},
		L2SequenceNumber: instructionIndex,
		L2Claim:          common.BytesToHash(stateBytes),
	}

	// Create the executor
	logger := log.New("component", "CannonExecutor")
	metricer := metrics.NoopMetrics.ToTypedVmMetrics("cannon")
	oracleServer := vm.NewOpProgramServerExecutor(logger)
	executor := vm.NewExecutor(logger, metricer, vmConfig, oracleServer, "", localInputs)

	// Create necessary directories
	if err := os.MkdirAll(filepath.Join(tempDir, vm.SnapsDir), 0755); err != nil {
		return nil, fmt.Errorf("failed to create snapshot directory: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(tempDir, vm.PreimagesDir), 0755); err != nil {
		return nil, fmt.Errorf("failed to create preimages directory: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(tempDir, utils.ProofsDir), 0755); err != nil {
		return nil, fmt.Errorf("failed to create proofs directory: %w", err)
	}

	// Execute the instruction and generate the witness
	v.logger.Debug("Executing instruction to generate witness", "index", instructionIndex)
	if err := executor.GenerateProof(ctx, tempDir, instructionIndex.Uint64()); err != nil {
		return nil, fmt.Errorf("failed to execute instruction: %w", err)
	}

	// Get the proof data
	proof, step, exited, err := stateConverter.ConvertStateToProof(ctx, vm.FinalStatePath(tempDir, vmConfig.BinarySnapshots))
	if err != nil {
		return nil, fmt.Errorf("failed to convert state to proof: %w", err)
	}

	v.logger.Info("Generated witness", "step", step, "exited", exited)

	// Create the witness data structure
	witnessData := struct {
		InstructionIndex *big.Int      `json:"instructionIndex"`
		StateRoot        common.Hash   `json:"stateRoot"`
		StateData        hexutil.Bytes `json:"stateData"`
		ProofData        hexutil.Bytes `json:"proofData"`
		OracleKey        hexutil.Bytes `json:"oracleKey,omitempty"`
		OracleValue      hexutil.Bytes `json:"oracleValue,omitempty"`
		OracleOffset     uint32        `json:"oracleOffset,omitempty"`
		Step             uint64        `json:"step"`
		Exited           bool          `json:"exited"`
	}{
		InstructionIndex: instructionIndex,
		StateRoot:        proof.ClaimValue,
		StateData:        proof.StateData,
		ProofData:        proof.ProofData,
		OracleKey:        proof.OracleKey,
		OracleValue:      proof.OracleValue,
		OracleOffset:     proof.OracleOffset,
		Step:             step,
		Exited:           exited,
	}

	// Serialize the witness data to JSON
	witnessBytes, err := json.Marshal(witnessData)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize witness data: %w", err)
	}

	return witnessBytes, nil
}
