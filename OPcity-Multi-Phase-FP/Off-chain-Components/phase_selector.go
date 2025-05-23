package multiphase

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts"
	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	ErrSegmentStateRootMismatch = errors.New("segment state root mismatch")
	ErrGameNotFinalized         = errors.New("dispute game not finalized")
	ErrInvalidGameStatus        = errors.New("invalid game status")
	ErrDisputeGameNotFound      = errors.New("dispute game not found")
)

// PhaseSelector is responsible for selecting the appropriate phase for dispute resolution
// based on the state of the dispute.
type PhaseSelector struct {
	log             log.Logger
	client          *ethclient.Client
	disputeFactory  *contracts.DisputeGameFactoryContract
	gameType        uint32
	maxLookupGames  int
	finalityTimeout time.Duration
}

// NewPhaseSelector creates a new PhaseSelector instance.
func NewPhaseSelector(
	logger log.Logger,
	client *ethclient.Client,
	disputeFactory *contracts.DisputeGameFactoryContract,
	gameType uint32,
	maxLookupGames int,
	finalityTimeout time.Duration,
) *PhaseSelector {
	return &PhaseSelector{
		log:             logger,
		client:          client,
		disputeFactory:  disputeFactory,
		gameType:        gameType,
		maxLookupGames:  maxLookupGames,
		finalityTimeout: finalityTimeout,
	}
}

// SelectPhase determines which phase of verification should be used based on the segment state.
func (p *PhaseSelector) SelectPhase(
	ctx context.Context,
	disputeID *big.Int,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
	endStateRoot [32]byte,
) (int, error) {
	// First, verify the segment state root
	verified, err := p.verifySegmentStateRoot(ctx, disputeID, startIndex, endIndex, startStateRoot, endStateRoot)
	if err != nil {
		return 0, fmt.Errorf("failed to verify segment state root: %w", err)
	}

	// If the segment state root is verified, use Phase 1 (coarse-grained verification)
	// Otherwise, use Phase 2 (fine-grained verification)
	if verified {
		p.log.Info("Segment state root verified, using Phase 1", 
			"disputeID", disputeID, 
			"startIndex", startIndex, 
			"endIndex", endIndex)
		return 1, nil
	}

	p.log.Info("Segment state root verification failed, using Phase 2", 
		"disputeID", disputeID, 
		"startIndex", startIndex, 
		"endIndex", endIndex)
	return 2, nil
}

// verifySegmentStateRoot verifies if the segment state root is valid by checking
// finalized dispute games from the dispute game factory.
func (p *PhaseSelector) verifySegmentStateRoot(
	ctx context.Context,
	disputeID *big.Int,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
	endStateRoot [32]byte,
) (bool, error) {
	p.log.Debug("Verifying segment state root", 
		"disputeID", disputeID, 
		"startIndex", startIndex, 
		"endIndex", endIndex, 
		"startStateRoot", common.Hash(startStateRoot), 
		"endStateRoot", common.Hash(endStateRoot))

	// Get the current block hash to use for queries
	blockHash, err := p.getCurrentBlockHash(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get current block hash: %w", err)
	}

	// Get recent games from the dispute game factory
	games, err := p.getRecentGames(ctx, blockHash)
	if err != nil {
		return false, fmt.Errorf("failed to get recent games: %w", err)
	}

	// Check if there's a finalized game that covers this segment
	for _, game := range games {
		isRelevant, err := p.isGameRelevantForSegment(ctx, game, startIndex, endIndex, startStateRoot, endStateRoot)
		if err != nil {
			p.log.Warn("Error checking game relevance", "gameIndex", game.Index, "error", err)
			continue
		}

		if !isRelevant {
			continue
		}

		// Check if the game is finalized
		isFinalized, gameStatus, err := p.isGameFinalized(ctx, game)
		if err != nil {
			p.log.Warn("Error checking game finalization", "gameIndex", game.Index, "error", err)
			continue
		}

		if isFinalized {
			p.log.Info("Found finalized game for segment", 
				"gameIndex", game.Index, 
				"gameStatus", gameStatus, 
				"startIndex", startIndex, 
				"endIndex", endIndex)
			
			// If the game is finalized with defender winning, the state root is valid
			if gameStatus == types.GameStatusDefenderWon {
				return true, nil
			} else if gameStatus == types.GameStatusChallengerWon {
				// If the challenger won, the state root is invalid
				return false, nil
			}
		}
	}

	// If we didn't find a finalized game for this segment, we need to perform verification
	p.log.Debug("No finalized game found for segment, performing verification", 
		"startIndex", startIndex, 
		"endIndex", endIndex)
	
	// Default to returning false to trigger detailed verification
	// This is a conservative approach - if we're not sure, verify in detail
	return false, nil
}

// getCurrentBlockHash gets the current block hash
func (p *PhaseSelector) getCurrentBlockHash(ctx context.Context) (common.Hash, error) {
	header, err := p.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get latest header: %w", err)
	}
	return header.Hash(), nil
}

// getRecentGames gets recent games from the dispute game factory
func (p *PhaseSelector) getRecentGames(ctx context.Context, blockHash common.Hash) ([]types.GameMetadata, error) {
	// Calculate the earliest timestamp to consider (now - finality timeout)
	earliestTimestamp := uint64(time.Now().Add(-p.finalityTimeout).Unix())
	
	// Get games at or after the earliest timestamp
	games, err := p.disputeFactory.GetGamesAtOrAfter(ctx, blockHash, earliestTimestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get games: %w", err)
	}

	// Filter games by game type
	var filteredGames []types.GameMetadata
	for _, game := range games {
		if game.GameType == p.gameType {
			filteredGames = append(filteredGames, game)
		}
		
		// Limit the number of games to check
		if len(filteredGames) >= p.maxLookupGames {
			break
		}
	}

	return filteredGames, nil
}

// isGameRelevantForSegment checks if a game is relevant for the given segment
func (p *PhaseSelector) isGameRelevantForSegment(
	ctx context.Context,
	game types.GameMetadata,
	startIndex *big.Int,
	endIndex *big.Int,
	startStateRoot [32]byte,
	endStateRoot [32]byte,
) (bool, error) {
	// Create a contract instance for the game
	gameContract, err := contracts.NewFaultDisputeGameContract(ctx, p.log, game.Proxy, p.disputeFactory.MultiCaller())
	if err != nil {
		return false, fmt.Errorf("failed to create game contract: %w", err)
	}

	// Get the game range
	prestateBlock, poststateBlock, err := gameContract.GetGameRange(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get game range: %w", err)
	}

	// Check if the game covers our segment
	if new(big.Int).SetUint64(prestateBlock).Cmp(startIndex) <= 0 &&
		new(big.Int).SetUint64(poststateBlock).Cmp(endIndex) >= 0 {
		
		// Get the absolute prestate hash
		absolutePrestate, err := gameContract.GetAbsolutePrestateHash(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to get absolute prestate hash: %w", err)
		}

		// Check if the prestate matches
		if absolutePrestate == common.Hash(startStateRoot) {
			p.log.Debug("Found relevant game for segment", 
				"gameIndex", game.Index, 
				"prestateBlock", prestateBlock, 
				"poststateBlock", poststateBlock)
			return true, nil
		}
	}

	return false, nil
}

// isGameFinalized checks if a game is finalized
func (p *PhaseSelector) isGameFinalized(ctx context.Context, game types.GameMetadata) (bool, types.GameStatus, error) {
	// Create a contract instance for the game
	gameContract, err := contracts.NewFaultDisputeGameContract(ctx, p.log, game.Proxy, p.disputeFactory.MultiCaller())
	if err != nil {
		return false, types.GameStatusInProgress, fmt.Errorf("failed to create game contract: %w", err)
	}

	// Get the game status
	status, err := gameContract.GetStatus(ctx)
	if err != nil {
		return false, types.GameStatusInProgress, fmt.Errorf("failed to get game status: %w", err)
	}

	// Check if the game is finalized (not in progress)
	isFinalized := status != types.GameStatusInProgress

	// If the game is finalized, check when it was resolved
	if isFinalized {
		resolvedAt, err := gameContract.GetResolvedAt(ctx, nil)
		if err != nil {
			return false, types.GameStatusInProgress, fmt.Errorf("failed to get resolved time: %w", err)
		}

		// Check if the game has been finalized for long enough
		if time.Since(resolvedAt) < p.finalityTimeout {
			p.log.Debug("Game is resolved but not finalized yet", 
				"gameIndex", game.Index, 
				"resolvedAt", resolvedAt, 
				"status", status)
			return false, status, nil
		}

		p.log.Info("Found finalized game", 
			"gameIndex", game.Index, 
			"resolvedAt", resolvedAt, 
			"status", status)
		return true, status, nil
	}

	return false, status, nil
}
