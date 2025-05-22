package tournament

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/log"
)

func TestRISCVTraceProvider(t *testing.T) {
	// Create a logger with a discard handler
	handler := log.DiscardHandler()
	logger := log.New()
	logger = log.NewLogger(handler)

	// Create a default configuration
	config := DefaultRISCVMachineConfig()

	// Create a new RISC-V trace provider
	provider := NewRISCVTraceProvider(logger, config)

	// Test GetTrace
	ctx := context.Background()
	trace, err := provider.GetTrace(ctx, 1)
	if err != nil {
		t.Fatalf("Failed to get trace: %v", err)
	}
	if trace.Hex() == "0x0000000000000000000000000000000000000000000000000000000000000000" {
		t.Fatalf("Expected non-zero trace, got: %s", trace.Hex())
	}

	// Test GenerateProof
	proof, err := provider.GenerateProof(ctx, 1)
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}
	if len(proof) == 0 {
		t.Fatalf("Expected non-empty proof, got empty proof")
	}

	// Test Step
	initialState := []byte{0x01, 0x02, 0x03}
	newState, err := provider.Step(ctx, initialState)
	if err != nil {
		t.Fatalf("Failed to execute step: %v", err)
	}
	if newState[0] != 0x02 {
		t.Fatalf("Expected first byte to be incremented, got: %v", newState[0])
	}

	// Test GenerateTrace
	input := []byte("test-input")
	traces, err := provider.GenerateTrace(ctx, input)
	if err != nil {
		t.Fatalf("Failed to generate trace: %v", err)
	}
	if len(traces) != 5 {
		t.Fatalf("Expected 5 trace steps, got: %d", len(traces))
	}

	// Test VerifyProof
	valid, err := provider.VerifyProof(ctx, trace, proof)
	if err != nil {
		t.Fatalf("Failed to verify proof: %v", err)
	}
	if !valid {
		t.Fatalf("Expected proof to be valid")
	}

	// Test machine configuration
	if provider.GetMachineConfig().RamSize != config.RamSize {
		t.Fatalf("Expected RAM size to be %d, got: %d", config.RamSize, provider.GetMachineConfig().RamSize)
	}

	// Test setting machine configuration
	newConfig := &RISCVMachineConfig{
		RamSize: 128 * 1024 * 1024,
	}
	provider.SetMachineConfig(newConfig)
	if provider.GetMachineConfig().RamSize != newConfig.RamSize {
		t.Fatalf("Expected RAM size to be %d, got: %d", newConfig.RamSize, provider.GetMachineConfig().RamSize)
	}

	// Test Cartesi Machine initialization
	err = provider.InitializeCartesiMachine()
	if err != nil {
		t.Fatalf("Failed to initialize Cartesi Machine: %v", err)
	}

	// Test Cartesi Machine execution
	err = provider.ExecuteCartesiMachine(1000)
	if err != nil {
		t.Fatalf("Failed to execute Cartesi Machine: %v", err)
	}

	// Test getting Cartesi Machine state
	state, err := provider.GetCartesiMachineState()
	if err != nil {
		t.Fatalf("Failed to get Cartesi Machine state: %v", err)
	}
	if len(state) == 0 {
		t.Fatalf("Expected non-empty state, got empty state")
	}

	// Test setting Cartesi Machine state
	err = provider.SetCartesiMachineState([]byte("new-state"))
	if err != nil {
		t.Fatalf("Failed to set Cartesi Machine state: %v", err)
	}

	// Test Cartesi Machine shutdown
	err = provider.ShutdownCartesiMachine()
	if err != nil {
		t.Fatalf("Failed to shut down Cartesi Machine: %v", err)
	}
}
