package tournament

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace"
	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	"github.com/ethereum-optimism/optimism/op-challenger/metrics"
	"github.com/ethereum-optimism/optimism/op-service/clock"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// actor is a function that performs an action on a tournament game
type actor func(ctx context.Context) error

// TournamentInfo defines the interface for retrieving tournament game information
type TournamentInfo interface {
	// GetStatus returns the current status of the tournament game
	GetStatus(context.Context) (types.GameStatus, error)

	// GetTotalParticipants returns the total number of participants in the tournament
	GetTotalParticipants(context.Context) (uint64, error)

	// GetCurrentRound returns the current round of the tournament
	GetCurrentRound(context.Context) (uint64, error)

	// GetMatchCount returns the total number of matches in the tournament
	GetMatchCount(context.Context) (uint64, error)

	// HasParticipated checks if an address has participated in the tournament
	HasParticipated(context.Context, common.Address) (bool, error)
}

// SyncValidator validates that the local node is sufficiently synced to play the game
type SyncValidator interface {
	// ValidateNodeSynced checks that the local node is sufficiently up to date to play the game.
	// It returns types.ErrNotInSync if the node is too far behind.
	ValidateNodeSynced(ctx context.Context, gameL1Head eth.BlockID) error
}

// L1HeaderSource provides access to L1 headers
type L1HeaderSource interface {
	HeaderByHash(context.Context, common.Hash) (*gethTypes.Header, error)
}

// TxSender sends transactions to the network
type TxSender interface {
	From() common.Address
	SendAndWaitSimple(txPurpose string, txs ...txmgr.TxCandidate) error
}

// TraceProvider provides access to execution traces
type TraceProvider interface {
	// GetTrace returns the trace at the specified index
	GetTrace(ctx context.Context, idx uint64) (common.Hash, error)

	// GenerateProof generates a proof for the specified trace index
	GenerateProof(ctx context.Context, idx uint64) ([]byte, error)
}

// Validator validates the prestate of a tournament game
type Validator interface {
	Validate(ctx context.Context) error
}

// TournamentContract defines the interface for interacting with the tournament game contract
type TournamentContract interface {
	TournamentInfo

	// JoinTournament joins the tournament with a counter-claim
	JoinTournament(ctx context.Context, claim common.Hash) error

	// ResolveMatch resolves a match in the tournament
	ResolveMatch(ctx context.Context, matchIndex uint64, winnerIndex uint64) error

	// ClaimBond claims the bond as the tournament winner
	ClaimBond(ctx context.Context) error

	// GetL1Head returns the L1 head hash at the time the tournament was created
	GetL1Head(ctx context.Context) (common.Hash, error)

	// GetRootClaim returns the root claim of the tournament
	GetRootClaim(ctx context.Context) (common.Hash, error)

	// GetMatch returns information about a match in the tournament
	GetMatch(ctx context.Context, matchIndex uint64) (Match, error)

	// GetNode returns information about a node in the tournament tree
	GetNode(ctx context.Context, nodeIndex uint64) (Node, error)
}

// Match represents a match in the tournament
type Match struct {
	NodeA     uint64
	NodeB     uint64
	Winner    uint64
	StartTime uint64
	Resolved  bool
}

// Node represents a node in the tournament tree
type Node struct {
	Participant common.Address
	Claim       common.Hash
	HasJoined   bool
	BondAmount  uint64
}

// TournamentPlayer is responsible for participating in tournament-based dispute games
type TournamentPlayer struct {
	act                actor
	loader             TournamentInfo
	logger             log.Logger
	syncValidator      SyncValidator
	prestateValidators []Validator
	status             types.GameStatus
	gameL1Head         eth.BlockID
	txSender           TxSender
	traceProvider      TraceProvider
}

var actNoop = func(ctx context.Context) error {
	return nil
}

// NewTournamentPlayer creates a new TournamentPlayer
func NewTournamentPlayer(
	ctx context.Context,
	systemClock clock.Clock,
	logger log.Logger,
	m metrics.Metricer,
	dir string,
	addr common.Address,
	txSender TxSender,
	loader TournamentContract,
	syncValidator SyncValidator,
	validators []Validator,
	traceProvider TraceProvider,
	l1HeaderSource L1HeaderSource,
) (*TournamentPlayer, error) {
	logger = logger.New("tournament", addr)

	status, err := loader.GetStatus(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tournament status: %w", err)
	}

	if status != types.GameStatusInProgress {
		logger.Info("Tournament already resolved", "status", status)
		// Tournament is already complete so skip creating the trace provider, loading game inputs etc.
		return &TournamentPlayer{
			logger:             logger,
			loader:             loader,
			prestateValidators: validators,
			status:             status,
			txSender:           txSender,
			// Act function does nothing because the game is already complete
			act: actNoop,
		}, nil
	}

	l1HeadHash, err := loader.GetL1Head(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load tournament L1 head: %w", err)
	}

	l1Header, err := l1HeaderSource.HeaderByHash(ctx, l1HeadHash)
	if err != nil {
		return nil, fmt.Errorf("failed to load L1 header %v: %w", l1HeadHash, err)
	}

	l1Head := eth.HeaderBlockID(l1Header)

	// Check if we've already participated in this tournament
	hasParticipated, err := loader.HasParticipated(ctx, txSender.From())
	if err != nil {
		return nil, fmt.Errorf("failed to check if we've participated: %w", err)
	}

	agent := NewAgent(
		systemClock,
		loader,
		txSender,
		traceProvider,
		logger,
		hasParticipated,
	)

	return &TournamentPlayer{
		act:                agent.Act,
		loader:             loader,
		logger:             logger,
		status:             status,
		gameL1Head:         l1Head,
		syncValidator:      syncValidator,
		prestateValidators: validators,
		txSender:           txSender,
		traceProvider:      traceProvider,
	}, nil
}

// ValidatePrestate validates the prestate of the tournament game
func (t *TournamentPlayer) ValidatePrestate(ctx context.Context) error {
	for _, validator := range t.prestateValidators {
		if err := validator.Validate(ctx); err != nil {
			return fmt.Errorf("failed to validate prestate: %w", err)
		}
	}
	return nil
}

// Status returns the current status of the tournament game
func (t *TournamentPlayer) Status() types.GameStatus {
	return t.status
}

// ProgressGame progresses the tournament game by performing necessary actions
func (t *TournamentPlayer) ProgressGame(ctx context.Context) types.GameStatus {
	if t.status != types.GameStatusInProgress {
		// Tournament is already complete so don't try to perform further actions.
		t.logger.Trace("Skipping completed tournament")
		return t.status
	}

	if err := t.syncValidator.ValidateNodeSynced(ctx, t.gameL1Head); errors.Is(err, trace.ErrNotInSync) {
		t.logger.Warn("Local node not sufficiently up to date", "err", err)
		return t.status
	} else if err != nil {
		t.logger.Error("Could not check local node was in sync", "err", err)
		return t.status
	}

	t.logger.Trace("Checking if actions are required for tournament")
	if err := t.act(ctx); err != nil {
		t.logger.Error("Error when acting on tournament", "err", err)
	}

	status, err := t.loader.GetStatus(ctx)
	if err != nil {
		t.logger.Error("Unable to retrieve tournament status", "err", err)
		return types.GameStatusInProgress
	}

	t.logTournamentStatus(ctx, status)
	t.status = status

	if status != types.GameStatusInProgress {
		// Release the agent as we will no longer need to act on this tournament.
		t.act = actNoop
	}

	return status
}

// logTournamentStatus logs the current status of the tournament
func (t *TournamentPlayer) logTournamentStatus(ctx context.Context, status types.GameStatus) {
	if status == types.GameStatusInProgress {
		participants, err := t.loader.GetTotalParticipants(ctx)
		if err != nil {
			t.logger.Error("Failed to get participant count for in progress tournament", "err", err)
			return
		}

		currentRound, err := t.loader.GetCurrentRound(ctx)
		if err != nil {
			t.logger.Error("Failed to get current round for in progress tournament", "err", err)
			return
		}

		matchCount, err := t.loader.GetMatchCount(ctx)
		if err != nil {
			t.logger.Error("Failed to get match count for in progress tournament", "err", err)
			return
		}

		t.logger.Info("Tournament info", "participants", participants, "round", currentRound, "matches", matchCount, "status", status)
		return
	}

	t.logger.Info("Tournament resolved", "status", status)
}
