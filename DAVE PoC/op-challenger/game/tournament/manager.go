// Package tournament provides functionality for interacting with tournament-based dispute games.
package tournament

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	"github.com/ethereum-optimism/optimism/op-service/clock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// TournamentFactoryContract defines the interface for interacting with the tournament factory contract
type TournamentFactoryContract interface {
	// GetActiveTournaments returns the addresses of all active tournaments
	GetActiveTournaments(ctx context.Context) ([]common.Address, error)

	// GetTournamentCount returns the total number of tournaments created
	GetTournamentCount(ctx context.Context) (uint64, error)

	// GetTournamentAt returns the address of the tournament at the specified index
	GetTournamentAt(ctx context.Context, index uint64) (common.Address, error)
}

// TournamentCreator creates new TournamentPlayer instances
type TournamentCreator interface {
	// CreateTournamentPlayer creates a new TournamentPlayer for the specified tournament address
	CreateTournamentPlayer(ctx context.Context, addr common.Address) (*TournamentPlayer, error)
}

// TournamentState represents the state of a tournament
type TournamentState struct {
	Address      common.Address
	Status       types.GameStatus
	Participants uint64
	CurrentRound uint64
	MatchCount   uint64
	RootClaim    common.Hash
	L1Head       common.Hash
	UpdatedAt    time.Time
}

// Manager is responsible for managing tournament state and interactions
type Manager struct {
	ctx              context.Context
	cancel           context.CancelFunc
	logger           log.Logger
	clock            clock.Clock
	factory          TournamentFactoryContract
	creator          TournamentCreator
	players          map[common.Address]*TournamentPlayer
	tournamentStates sync.Map // map[common.Address]TournamentState
	pollInterval     time.Duration
	wg               sync.WaitGroup
}

// NewManager creates a new tournament manager
func NewManager(
	ctx context.Context,
	logger log.Logger,
	clock clock.Clock,
	factory TournamentFactoryContract,
	creator TournamentCreator,
	pollInterval time.Duration,
) *Manager {
	ctx, cancel := context.WithCancel(ctx)
	return &Manager{
		ctx:          ctx,
		cancel:       cancel,
		logger:       logger.New("component", "TournamentManager"),
		clock:        clock,
		factory:      factory,
		creator:      creator,
		players:      make(map[common.Address]*TournamentPlayer),
		pollInterval: pollInterval,
	}
}

// Start starts the tournament manager
func (m *Manager) Start() error {
	m.logger.Info("Starting tournament manager")

	// Start the tournament monitoring loop
	m.wg.Add(1)
	go m.monitorTournaments()

	return nil
}

// Stop stops the tournament manager
func (m *Manager) Stop() {
	m.logger.Info("Stopping tournament manager")
	m.cancel()
	m.wg.Wait()
}

// GetTournamentState returns the state of a tournament
func (m *Manager) GetTournamentState(ctx context.Context, addr common.Address) (TournamentState, error) {
	// Check if we have the state cached
	if state, ok := m.tournamentStates.Load(addr); ok {
		return state.(TournamentState), nil
	}

	// If not cached, fetch the state from the blockchain
	state, err := m.fetchTournamentState(ctx, addr)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to fetch tournament state: %w", err)
	}

	// Cache the state
	m.tournamentStates.Store(addr, state)

	return state, nil
}

// GetActiveTournaments returns the addresses of all active tournaments
func (m *Manager) GetActiveTournaments(ctx context.Context) ([]common.Address, error) {
	return m.factory.GetActiveTournaments(ctx)
}

// GetAllTournaments returns the addresses of all tournaments
func (m *Manager) GetAllTournaments(ctx context.Context) ([]common.Address, error) {
	count, err := m.factory.GetTournamentCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get tournament count: %w", err)
	}

	tournaments := make([]common.Address, 0, count)
	for i := uint64(0); i < count; i++ {
		addr, err := m.factory.GetTournamentAt(ctx, i)
		if err != nil {
			m.logger.Error("Failed to get tournament at index", "index", i, "err", err)
			continue
		}
		tournaments = append(tournaments, addr)
	}

	return tournaments, nil
}

// GetTournamentPlayer returns the TournamentPlayer for the specified tournament address
func (m *Manager) GetTournamentPlayer(ctx context.Context, addr common.Address) (*TournamentPlayer, error) {
	// Check if we already have a player for this tournament
	if player, ok := m.players[addr]; ok {
		return player, nil
	}

	// Create a new player
	player, err := m.creator.CreateTournamentPlayer(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create tournament player: %w", err)
	}

	// Store the player
	m.players[addr] = player

	return player, nil
}

// ProgressTournament progresses the specified tournament
func (m *Manager) ProgressTournament(ctx context.Context, addr common.Address) (types.GameStatus, error) {
	player, err := m.GetTournamentPlayer(ctx, addr)
	if err != nil {
		return types.GameStatusInProgress, fmt.Errorf("failed to get tournament player: %w", err)
	}

	status := player.ProgressGame(ctx)

	// Update the cached state
	m.updateTournamentState(ctx, addr)

	return status, nil
}

// ValidateTournamentPrestate validates the prestate of the specified tournament
func (m *Manager) ValidateTournamentPrestate(ctx context.Context, addr common.Address) error {
	player, err := m.GetTournamentPlayer(ctx, addr)
	if err != nil {
		return fmt.Errorf("failed to get tournament player: %w", err)
	}

	return player.ValidatePrestate(ctx)
}

// monitorTournaments monitors active tournaments
func (m *Manager) monitorTournaments() {
	defer m.wg.Done()

	ticker := m.clock.NewTicker(m.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-ticker.Ch():
			if err := m.checkActiveTournaments(); err != nil {
				m.logger.Error("Failed to check active tournaments", "err", err)
			}
		}
	}
}

// checkActiveTournaments checks all active tournaments
func (m *Manager) checkActiveTournaments() error {
	ctx, cancel := context.WithTimeout(m.ctx, m.pollInterval/2)
	defer cancel()

	// Get all active tournaments
	activeTournaments, err := m.factory.GetActiveTournaments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active tournaments: %w", err)
	}

	m.logger.Debug("Checking active tournaments", "count", len(activeTournaments))

	// Process each active tournament
	for _, addr := range activeTournaments {
		if err := m.processTournament(ctx, addr); err != nil {
			m.logger.Error("Failed to process tournament", "addr", addr, "err", err)
		}
	}

	return nil
}

// processTournament processes a single tournament
func (m *Manager) processTournament(ctx context.Context, addr common.Address) error {
	// Update the tournament state
	if err := m.updateTournamentState(ctx, addr); err != nil {
		return fmt.Errorf("failed to update tournament state: %w", err)
	}

	// Progress the tournament
	if _, err := m.ProgressTournament(ctx, addr); err != nil {
		return fmt.Errorf("failed to progress tournament: %w", err)
	}

	return nil
}

// fetchTournamentState fetches the state of a tournament from the blockchain
func (m *Manager) fetchTournamentState(ctx context.Context, addr common.Address) (TournamentState, error) {
	// Get the tournament contract
	player, err := m.GetTournamentPlayer(ctx, addr)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to get tournament player: %w", err)
	}

	// Get the tournament info
	loader, ok := player.loader.(TournamentContract)
	if !ok {
		return TournamentState{}, errors.New("tournament player loader is not a TournamentContract")
	}

	// Get the tournament status
	status, err := loader.GetStatus(ctx)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to get tournament status: %w", err)
	}

	// Get the tournament participants
	participants, err := loader.GetTotalParticipants(ctx)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to get tournament participants: %w", err)
	}

	// Get the current round
	currentRound, err := loader.GetCurrentRound(ctx)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to get current round: %w", err)
	}

	// Get the match count
	matchCount, err := loader.GetMatchCount(ctx)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to get match count: %w", err)
	}

	// Get the root claim
	rootClaim, err := loader.GetRootClaim(ctx)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to get root claim: %w", err)
	}

	// Get the L1 head
	l1Head, err := loader.GetL1Head(ctx)
	if err != nil {
		return TournamentState{}, fmt.Errorf("failed to get L1 head: %w", err)
	}

	return TournamentState{
		Address:      addr,
		Status:       status,
		Participants: participants,
		CurrentRound: currentRound,
		MatchCount:   matchCount,
		RootClaim:    rootClaim,
		L1Head:       l1Head,
		UpdatedAt:    m.clock.Now(),
	}, nil
}

// updateTournamentState updates the cached state of a tournament
func (m *Manager) updateTournamentState(ctx context.Context, addr common.Address) error {
	state, err := m.fetchTournamentState(ctx, addr)
	if err != nil {
		return fmt.Errorf("failed to fetch tournament state: %w", err)
	}

	m.tournamentStates.Store(addr, state)

	return nil
}

// GetTournamentMatches returns the matches for a tournament
func (m *Manager) GetTournamentMatches(ctx context.Context, addr common.Address) ([]Match, error) {
	// Get the tournament contract
	player, err := m.GetTournamentPlayer(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get tournament player: %w", err)
	}

	// Get the tournament info
	loader, ok := player.loader.(TournamentContract)
	if !ok {
		return nil, errors.New("tournament player loader is not a TournamentContract")
	}

	// Get the match count
	matchCount, err := loader.GetMatchCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get match count: %w", err)
	}

	// Get all matches
	matches := make([]Match, 0, matchCount)
	for i := uint64(0); i < matchCount; i++ {
		match, err := loader.GetMatch(ctx, i)
		if err != nil {
			m.logger.Error("Failed to get match", "index", i, "err", err)
			continue
		}
		matches = append(matches, match)
	}

	return matches, nil
}

// GetTournamentNodes returns the nodes for a tournament
func (m *Manager) GetTournamentNodes(ctx context.Context, addr common.Address) ([]Node, error) {
	// Get the tournament contract
	player, err := m.GetTournamentPlayer(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get tournament player: %w", err)
	}

	// Get the tournament info
	loader, ok := player.loader.(TournamentContract)
	if !ok {
		return nil, errors.New("tournament player loader is not a TournamentContract")
	}

	// Get the participant count
	participantCount, err := loader.GetTotalParticipants(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get participant count: %w", err)
	}

	// Get all nodes
	nodes := make([]Node, 0, participantCount)
	for i := uint64(0); i < participantCount; i++ {
		node, err := loader.GetNode(ctx, i)
		if err != nil {
			m.logger.Error("Failed to get node", "index", i, "err", err)
			continue
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

// GetTournamentStats returns statistics for a tournament
func (m *Manager) GetTournamentStats(ctx context.Context, addr common.Address) (map[string]interface{}, error) {
	state, err := m.GetTournamentState(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get tournament state: %w", err)
	}

	matches, err := m.GetTournamentMatches(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get tournament matches: %w", err)
	}

	nodes, err := m.GetTournamentNodes(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get tournament nodes: %w", err)
	}

	// Calculate statistics
	resolvedMatches := 0
	for _, match := range matches {
		if match.Resolved {
			resolvedMatches++
		}
	}

	joinedNodes := 0
	for _, node := range nodes {
		if node.HasJoined {
			joinedNodes++
		}
	}

	return map[string]interface{}{
		"address":          state.Address,
		"status":           state.Status,
		"participants":     state.Participants,
		"currentRound":     state.CurrentRound,
		"matchCount":       state.MatchCount,
		"resolvedMatches":  resolvedMatches,
		"joinedNodes":      joinedNodes,
		"rootClaim":        state.RootClaim,
		"l1Head":           state.L1Head,
		"updatedAt":        state.UpdatedAt,
		"completionRatio":  float64(resolvedMatches) / float64(state.MatchCount),
		"participantRatio": float64(joinedNodes) / float64(state.Participants),
	}, nil
}

// TournamentCreatorFunc is a function that creates a TournamentPlayer
type TournamentCreatorFunc func(ctx context.Context, addr common.Address) (*TournamentPlayer, error)

// CreateTournamentPlayer implements the TournamentCreator interface
func (f TournamentCreatorFunc) CreateTournamentPlayer(ctx context.Context, addr common.Address) (*TournamentPlayer, error) {
	return f(ctx, addr)
}
