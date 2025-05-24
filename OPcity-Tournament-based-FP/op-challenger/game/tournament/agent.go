package tournament

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-service/clock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Agent is responsible for monitoring tournaments and taking actions
type Agent struct {
	clock         clock.Clock
	loader        TournamentContract
	txSender      TxSender
	traceProvider TraceProvider
	logger        log.Logger
	hasJoined     bool
}

// NewAgent creates a new tournament agent
func NewAgent(
	clock clock.Clock,
	loader TournamentContract,
	txSender TxSender,
	traceProvider TraceProvider,
	logger log.Logger,
	hasJoined bool,
) *Agent {
	return &Agent{
		clock:         clock,
		loader:        loader,
		txSender:      txSender,
		traceProvider: traceProvider,
		logger:        logger,
		hasJoined:     hasJoined,
	}
}

// Act performs the necessary actions for the tournament
func (a *Agent) Act(ctx context.Context) error {
	// Check if the tournament is still in progress
	status, err := a.loader.GetStatus(ctx)
	if err != nil {
		return fmt.Errorf("failed to get tournament status: %w", err)
	}

	if status != 0 { // 0 = IN_PROGRESS
		a.logger.Info("Tournament is no longer in progress", "status", status)
		return nil
	}

	// If we haven't joined the tournament yet, join it
	if !a.hasJoined {
		if err := a.joinTournament(ctx); err != nil {
			return fmt.Errorf("failed to join tournament: %w", err)
		}
		a.hasJoined = true
	}

	// Check for matches that need to be resolved
	if err := a.checkAndResolveMatches(ctx); err != nil {
		return fmt.Errorf("failed to check and resolve matches: %w", err)
	}

	// Check if we can claim the bond
	if err := a.checkAndClaimBond(ctx); err != nil {
		return fmt.Errorf("failed to check and claim bond: %w", err)
	}

	return nil
}

// joinTournament joins the tournament with a counter-claim
func (a *Agent) joinTournament(ctx context.Context) error {
	// Get the root claim to generate a counter-claim
	rootClaim, err := a.loader.GetRootClaim(ctx)
	if err != nil {
		return fmt.Errorf("failed to get root claim: %w", err)
	}

	// Generate a counter-claim using the trace provider
	// For simplicity, we'll use the trace at index 1 as our counter-claim
	counterClaim, err := a.traceProvider.GetTrace(ctx, 1)
	if err != nil {
		return fmt.Errorf("failed to generate counter-claim: %w", err)
	}

	a.logger.Info("Joining tournament", "rootClaim", rootClaim.Hex(), "counterClaim", counterClaim.Hex())

	// Join the tournament with our counter-claim
	if err := a.loader.JoinTournament(ctx, counterClaim); err != nil {
		return fmt.Errorf("failed to join tournament: %w", err)
	}

	a.logger.Info("Successfully joined tournament")
	return nil
}

// checkAndResolveMatches checks for matches that need to be resolved and resolves them
func (a *Agent) checkAndResolveMatches(ctx context.Context) error {
	// Get the total number of matches
	matchCount, err := a.loader.GetMatchCount(ctx)
	if err != nil {
		return fmt.Errorf("failed to get match count: %w", err)
	}

	// Check each match
	for i := uint64(0); i < matchCount; i++ {
		match, err := a.loader.GetMatch(ctx, i)
		if err != nil {
			a.logger.Error("Failed to get match", "index", i, "err", err)
			continue
		}

		// Skip already resolved matches
		if match.Resolved {
			continue
		}

		// Check if we're a participant in this match
		nodeA, err := a.loader.GetNode(ctx, match.NodeA)
		if err != nil {
			a.logger.Error("Failed to get node A", "index", match.NodeA, "err", err)
			continue
		}

		nodeB, err := a.loader.GetNode(ctx, match.NodeB)
		if err != nil {
			a.logger.Error("Failed to get node B", "index", match.NodeB, "err", err)
			continue
		}

		// Check if we're a participant in this match
		if nodeA.Participant != a.txSender.From() && nodeB.Participant != a.txSender.From() {
			continue
		}

		// Check if the match duration has passed
		if uint64(a.clock.Now().Unix()) < match.StartTime+86400 { // Assuming 24 hours match duration
			a.logger.Debug("Match not ready to be resolved yet", "index", i)
			continue
		}

		// Determine the winner based on trace evidence
		winnerIndex, err := a.determineWinner(ctx, match, nodeA, nodeB)
		if err != nil {
			a.logger.Error("Failed to determine winner", "match", i, "err", err)
			continue
		}

		a.logger.Info("Resolving match", "index", i, "winner", winnerIndex)

		// Resolve the match
		if err := a.loader.ResolveMatch(ctx, i, winnerIndex); err != nil {
			a.logger.Error("Failed to resolve match", "index", i, "err", err)
			continue
		}

		a.logger.Info("Successfully resolved match", "index", i, "winner", winnerIndex)
	}

	return nil
}

// determineWinner determines the winner of a match based on trace evidence
func (a *Agent) determineWinner(ctx context.Context, match Match, nodeA Node, nodeB Node) (uint64, error) {
	// Generate proofs for both claims
	proofA, err := a.generateProofForClaim(ctx, nodeA.Claim)
	if err != nil {
		return 0, fmt.Errorf("failed to generate proof for node A: %w", err)
	}

	proofB, err := a.generateProofForClaim(ctx, nodeB.Claim)
	if err != nil {
		return 0, fmt.Errorf("failed to generate proof for node B: %w", err)
	}

	// Compare the proofs and determine the winner
	// This is a simplified example - in a real implementation, you would need to
	// implement proper verification logic based on your specific requirements
	if len(proofA) > len(proofB) {
		return match.NodeA, nil
	} else {
		return match.NodeB, nil
	}
}

// generateProofForClaim generates a proof for a claim
func (a *Agent) generateProofForClaim(ctx context.Context, claim common.Hash) ([]byte, error) {
	// In a real implementation, you would need to find the trace index that corresponds to the claim
	// For simplicity, we'll use a fixed index here
	traceIndex := uint64(1)

	// Generate the proof
	proof, err := a.traceProvider.GenerateProof(ctx, traceIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to generate proof: %w", err)
	}

	return proof, nil
}

// checkAndClaimBond checks if we can claim the bond and claims it if possible
func (a *Agent) checkAndClaimBond(ctx context.Context) error {
	// Check if the tournament is resolved
	status, err := a.loader.GetStatus(ctx)
	if err != nil {
		return fmt.Errorf("failed to get tournament status: %w", err)
	}

	// If the tournament is not resolved, we can't claim the bond
	if status == 0 { // 0 = IN_PROGRESS
		return nil
	}

	// Try to claim the bond
	// The contract will revert if we're not the winner
	if err := a.loader.ClaimBond(ctx); err != nil {
		a.logger.Debug("Failed to claim bond, likely not the winner", "err", err)
		return nil
	}

	a.logger.Info("Successfully claimed bond")
	return nil
}
