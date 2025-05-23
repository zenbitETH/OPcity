package multiphase

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	ErrSegmentMismatch = errors.New("segment state root mismatch")
	ErrNoDiscrepancy   = errors.New("no discrepancy found in segment")
)

// Phase1Verifier handles coarse-grained verification of large state segments
type Phase1Verifier struct {
	client             *ethclient.Client
	lazyLoadingManager *LazyLoadingManager
	logger             log.Logger
	cannonBin          string
	cannonServer       string
	dataDir            string
}

// NewPhase1Verifier creates a new instance of Phase1Verifier
func NewPhase1Verifier(client *ethclient.Client, llm *LazyLoadingManager, logger log.Logger, cannonBin, cannonServer, dataDir string) *Phase1Verifier {
	return &Phase1Verifier{
		client:             client,
		lazyLoadingManager: llm,
		logger:             logger,
		cannonBin:          cannonBin,
		cannonServer:       cannonServer,
		dataDir:            dataDir,
	}
}

// VerifyAndFindDiscrepancy identifies segments with discrepancies
func (v *Phase1Verifier) VerifyAndFindDiscrepancy(
	ctx context.Context,
	disputeID *big.Int,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
	endStateRoot [32]byte,
) (*big.Int, *big.Int, [32]byte, error) {
	// Check if segment is already small enough for direct verification
	if new(big.Int).Sub(endIndex, startIndex).Cmp(big.NewInt(1000)) <= 0 {
		// If segment is small, just return it for Phase 2 verification
		return startIndex, endIndex, startStateRoot, nil
	}

	// Perform bisection to find the discrepancy
	midpoint := new(big.Int).Add(startIndex, endIndex)
	midpoint.Div(midpoint, big.NewInt(2))

	// Load state at midpoint using lazy loading manager
	midStateRootBytes, err := v.lazyLoadingManager.LoadState(ctx, disputeID, midpoint)
	if err != nil {
		return nil, nil, [32]byte{}, fmt.Errorf("failed to load state at midpoint: %w", err)
	}

	midStateRoot := common.BytesToHash(midStateRootBytes)

	// Verify left half (start to mid)
	leftValid, err := v.VerifySegment(ctx, disputeID, startIndex, midpoint, startStateRoot, midStateRoot)
	if err != nil {
		return nil, nil, [32]byte{}, err
	}

	// Verify right half (mid to end)
	rightValid, err := v.VerifySegment(ctx, disputeID, midpoint, endIndex, midStateRoot, endStateRoot)
	if err != nil {
		return nil, nil, [32]byte{}, err
	}

	// Determine which half contains the discrepancy
	if !leftValid {
		// Recursively search the left half
		return v.VerifyAndFindDiscrepancy(ctx, disputeID, startIndex, midpoint, startStateRoot, midStateRoot)
	} else if !rightValid {
		// Recursively search the right half
		return v.VerifyAndFindDiscrepancy(ctx, disputeID, midpoint, endIndex, midStateRoot, endStateRoot)
	}

	// If both segments valid but end state roots differ, something is wrong
	return nil, nil, [32]byte{}, ErrNoDiscrepancy
}

// VerifySegment verifies if a segment's state transition is valid
func (v *Phase1Verifier) VerifySegment(
	ctx context.Context,
	disputeID *big.Int,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
	endStateRoot [32]byte,
) (bool, error) {
	// Compute the state transition for this segment
	computedEndState, err := v.computeSegmentStateTransition(ctx, disputeID, startIndex, endIndex, startStateRoot)
	if err != nil {
		return false, fmt.Errorf("failed to compute segment state transition: %w", err)
	}

	// Verify that startStateRoot transitions to endStateRoot
	// Compare the computed end state with the expected end state
	if computedEndState != endStateRoot {
		return false, nil // Return false without error to indicate invalid transition
	}

	// Return true if the transition is valid
	return true, nil
}

// computeSegmentStateTransition computes the result of applying all state transitions in a segment
func (v *Phase1Verifier) computeSegmentStateTransition(
	ctx context.Context,
	disputeID *big.Int,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
) ([32]byte, error) {
	// In a real implementation, this would:
	// 1. Use Cannon or other execution engine to compute state transitions
	// 2. Or retrieve precomputed state checkpoints from a trusted source

	// For large segments, we might use a specialized execution engine
	// that can efficiently process many state transitions at once

	// Try to get the precomputed state from the lazy loading manager
	// This is a simplified approach - in a real implementation, we would have a more
	// sophisticated mechanism to retrieve or compute state transitions
	endStateBytes, err := v.lazyLoadingManager.LoadState(ctx, disputeID, endIndex)
	if err == nil {
		// If we successfully loaded the state, use it
		return common.BytesToHash(endStateBytes), nil
	}

	// If we couldn't load the precomputed state, we need to compute it
	// This would involve executing the state transitions from startIndex to endIndex

	// Create a context with timeout for the computation
	execCtx, cancel := context.WithTimeout(ctx, defaultExecutionTimeout)
	defer cancel()

	// Execute the state transitions
	// This is where we would call into Cannon or another execution engine
	return executeCannonForSegment(execCtx, v.logger, v.cannonBin, v.cannonServer, v.dataDir, disputeID, startIndex, endIndex, startStateRoot)
}

// executeCannonForSegment executes Cannon for the given segment
func executeCannonForSegment(
	ctx context.Context,
	logger log.Logger,
	cannonBin string,
	cannonServer string,
	dataDir string,
	disputeID *big.Int,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
) ([32]byte, error) {
	// Create a unique directory for this segment execution
	segmentDir := filepath.Join(dataDir, fmt.Sprintf("segment_%s_%s_%s", disputeID.String(), startIndex.String(), endIndex.String()))
	if err := os.MkdirAll(segmentDir, 0755); err != nil {
		return [32]byte{}, fmt.Errorf("failed to create segment directory: %w", err)
	}

	// Create subdirectories for proofs, snapshots, and preimages
	proofsDir := filepath.Join(segmentDir, "proofs")
	snapshotsDir := filepath.Join(segmentDir, "snapshots")
	preimagesDir := filepath.Join(segmentDir, "preimages")

	for _, dir := range []string{proofsDir, snapshotsDir, preimagesDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return [32]byte{}, fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create a file with the initial state
	initialStatePath := filepath.Join(segmentDir, "initial_state.bin")
	if err := os.WriteFile(initialStatePath, startStateRoot[:], 0644); err != nil {
		return [32]byte{}, fmt.Errorf("failed to write initial state: %w", err)
	}

	// Prepare the final state output path
	finalStatePath := filepath.Join(segmentDir, "final_state.bin")

	// Prepare Cannon execution arguments
	// The exact arguments will depend on the Cannon VM implementation
	args := []string{
		"run",
		"--input", initialStatePath,
		"--output", finalStatePath,
		"--snapshot-at", "1000000", // Take snapshots every million steps
		"--snapshot-fmt", filepath.Join(snapshotsDir, "%d.bin.gz"),
		"--proof-at", "=" + endIndex.String(), // Generate proof at the end index
		"--proof-fmt", filepath.Join(proofsDir, "%d.json.gz"),
		"--",
		cannonServer,
		"--datadir", preimagesDir,
		"--dispute-id", disputeID.String(),
		"--start-index", startIndex.String(),
		"--end-index", endIndex.String(),
	}

	// Log the command being executed
	logger.Info("Executing Cannon VM",
		"bin", cannonBin,
		"args", args,
		"disputeID", disputeID,
		"startIndex", startIndex,
		"endIndex", endIndex)

	// Execute Cannon VM
	cmd := exec.CommandContext(ctx, cannonBin, args...)
	cmd.Dir = segmentDir

	// Capture stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Cannon execution failed",
			"error", err,
			"output", string(output),
			"disputeID", disputeID,
			"startIndex", startIndex,
			"endIndex", endIndex)
		return [32]byte{}, fmt.Errorf("cannon execution failed: %w, output: %s", err, string(output))
	}

	// Read the final state
	finalStateBytes, err := os.ReadFile(finalStatePath)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to read final state: %w", err)
	}

	// Convert to [32]byte
	var finalState [32]byte
	if len(finalStateBytes) != 32 {
		return [32]byte{}, fmt.Errorf("invalid final state length: got %d, want 32", len(finalStateBytes))
	}
	copy(finalState[:], finalStateBytes)

	logger.Info("Cannon execution completed successfully",
		"disputeID", disputeID,
		"startIndex", startIndex,
		"endIndex", endIndex,
		"finalState", common.Hash(finalState))

	return finalState, nil
}

// Default timeout for execution
var defaultExecutionTimeout = 30 * time.Second
