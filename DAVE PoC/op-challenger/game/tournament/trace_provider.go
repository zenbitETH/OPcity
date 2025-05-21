package tournament

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// TournamentTraceProvider provides traces for tournament games
type TournamentTraceProvider struct {
	logger        log.Logger
	traceAccessor trace.TraceAccessor
}

// NewTournamentTraceProvider creates a new TournamentTraceProvider
func NewTournamentTraceProvider(
	logger log.Logger,
	traceAccessor trace.TraceAccessor,
) *TournamentTraceProvider {
	return &TournamentTraceProvider{
		logger:        logger,
		traceAccessor: traceAccessor,
	}
}

// GetTrace returns the trace at the specified index
func (t *TournamentTraceProvider) GetTrace(ctx context.Context, idx uint64) (common.Hash, error) {
	value, err := t.traceAccessor.Get(ctx, idx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get trace at index %d: %w", idx, err)
	}
	return value, nil
}

// GenerateProof generates a proof for the specified trace index
func (t *TournamentTraceProvider) GenerateProof(ctx context.Context, idx uint64) ([]byte, error) {
	// In a real implementation, you would generate a proof using the appropriate
	// proof generation mechanism (e.g., Cannon, Asterisc, etc.)
	// For simplicity, we'll just return the trace value as bytes
	value, err := t.traceAccessor.Get(ctx, idx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate proof for index %d: %w", idx, err)
	}
	return value.Bytes(), nil
}
