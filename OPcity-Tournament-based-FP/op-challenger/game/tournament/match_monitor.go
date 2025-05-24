// Package tournament provides functionality for interacting with tournament-based dispute games.
package tournament

import (
	"container/heap"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	"github.com/ethereum-optimism/optimism/op-service/clock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// MatchPriority represents the priority level of a match
type MatchPriority int

const (
	// PriorityLow represents a low priority match
	PriorityLow MatchPriority = iota
	// PriorityMedium represents a medium priority match
	PriorityMedium
	// PriorityHigh represents a high priority match
	PriorityHigh
	// PriorityCritical represents a critical priority match that needs immediate attention
	PriorityCritical
)

// String returns a string representation of the match priority
func (p MatchPriority) String() string {
	switch p {
	case PriorityLow:
		return "Low"
	case PriorityMedium:
		return "Medium"
	case PriorityHigh:
		return "High"
	case PriorityCritical:
		return "Critical"
	default:
		return "Unknown"
	}
}

// MatchInfo contains information about a match including its priority
type MatchInfo struct {
	TournamentAddress common.Address // Address of the tournament contract
	MatchIndex        uint64         // Index of the match in the tournament
	Match             Match          // Match data from the contract
	Priority          MatchPriority  // Priority level of the match
	Deadline          time.Time      // Deadline for responding to the match
	IsOurs            bool           // Whether we are a participant in this match
}

// priorityQueue implements a priority queue for MatchInfo
type priorityQueue []*MatchInfo

// Len returns the length of the priority queue
func (pq priorityQueue) Len() int { return len(pq) }

// Less returns whether the item at index i has higher priority than the item at index j
func (pq priorityQueue) Less(i, j int) bool {
	// Higher priority comes first
	if pq[i].Priority != pq[j].Priority {
		return pq[i].Priority > pq[j].Priority
	}
	// Earlier deadline comes first
	return pq[i].Deadline.Before(pq[j].Deadline)
}

// Swap swaps the items at indices i and j
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an item to the priority queue
func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*MatchInfo)
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority item from the priority queue
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// MatchMonitor monitors match deadlines and prioritizes responses
type MatchMonitor struct {
	ctx            context.Context
	cancel         context.CancelFunc
	logger         log.Logger
	clock          clock.Clock
	manager        *Manager
	txSender       TxSender
	matches        map[common.Address]map[uint64]*MatchInfo // Tournament address -> match index -> match info
	priorityQueue  priorityQueue                            // Priority queue for matches
	mu             sync.Mutex                               // Mutex for thread safety
	pollInterval   time.Duration                            // Interval for polling match updates
	deadlineBuffer time.Duration                            // Buffer time before deadline to consider a match critical
	wg             sync.WaitGroup                           // WaitGroup for goroutines
}

// NewMatchMonitor creates a new MatchMonitor
func NewMatchMonitor(
	ctx context.Context,
	logger log.Logger,
	clock clock.Clock,
	manager *Manager,
	txSender TxSender,
	pollInterval time.Duration,
	deadlineBuffer time.Duration,
) *MatchMonitor {
	ctx, cancel := context.WithCancel(ctx)
	return &MatchMonitor{
		ctx:            ctx,
		cancel:         cancel,
		logger:         logger.New("component", "MatchMonitor"),
		clock:          clock,
		manager:        manager,
		txSender:       txSender,
		matches:        make(map[common.Address]map[uint64]*MatchInfo),
		priorityQueue:  make(priorityQueue, 0),
		pollInterval:   pollInterval,
		deadlineBuffer: deadlineBuffer,
	}
}

// Start starts the match monitor
func (m *MatchMonitor) Start() error {
	m.logger.Info("Starting match monitor")

	// Initialize the priority queue
	heap.Init(&m.priorityQueue)

	// Start the match monitoring loop
	m.wg.Add(1)
	go m.monitorMatches()

	return nil
}

// Stop stops the match monitor
func (m *MatchMonitor) Stop() {
	m.logger.Info("Stopping match monitor")
	m.cancel()
	m.wg.Wait()
}

// monitorMatches monitors matches across all tournaments
func (m *MatchMonitor) monitorMatches() {
	defer m.wg.Done()

	ticker := m.clock.NewTicker(m.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-ticker.Ch():
			if err := m.updateMatches(); err != nil {
				m.logger.Error("Failed to update matches", "err", err)
			}
		}
	}
}

// updateMatches updates the match information for all active tournaments
func (m *MatchMonitor) updateMatches() error {
	ctx, cancel := context.WithTimeout(m.ctx, m.pollInterval/2)
	defer cancel()

	// Get all active tournaments
	activeTournaments, err := m.manager.GetActiveTournaments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active tournaments: %w", err)
	}

	m.logger.Debug("Updating matches for active tournaments", "count", len(activeTournaments))

	// Process each active tournament
	for _, tournamentAddr := range activeTournaments {
		if err := m.updateTournamentMatches(ctx, tournamentAddr); err != nil {
			m.logger.Error("Failed to update tournament matches", "addr", tournamentAddr, "err", err)
		}
	}

	return nil
}

// updateTournamentMatches updates the match information for a specific tournament
func (m *MatchMonitor) updateTournamentMatches(ctx context.Context, tournamentAddr common.Address) error {
	// Get the tournament state
	state, err := m.manager.GetTournamentState(ctx, tournamentAddr)
	if err != nil {
		return fmt.Errorf("failed to get tournament state: %w", err)
	}

	// Skip if the tournament is not in progress
	if state.Status != types.GameStatusInProgress {
		return nil
	}

	// Get all matches for the tournament
	matches, err := m.manager.GetTournamentMatches(ctx, tournamentAddr)
	if err != nil {
		return fmt.Errorf("failed to get tournament matches: %w", err)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Initialize the tournament's match map if it doesn't exist
	if _, ok := m.matches[tournamentAddr]; !ok {
		m.matches[tournamentAddr] = make(map[uint64]*MatchInfo)
	}

	// Get our address
	ourAddress := m.txSender.From()

	// Update each match
	for i, match := range matches {
		// Skip resolved matches
		if match.Resolved {
			// Remove from our tracking if it exists
			if _, ok := m.matches[tournamentAddr][uint64(i)]; ok {
				delete(m.matches[tournamentAddr], uint64(i))
			}
			continue
		}

		// Check if we're a participant in this match
		nodeA, err := m.getNode(ctx, tournamentAddr, match.NodeA)
		if err != nil {
			m.logger.Error("Failed to get node A", "tournament", tournamentAddr, "index", match.NodeA, "err", err)
			continue
		}

		nodeB, err := m.getNode(ctx, tournamentAddr, match.NodeB)
		if err != nil {
			m.logger.Error("Failed to get node B", "tournament", tournamentAddr, "index", match.NodeB, "err", err)
			continue
		}

		isOurs := nodeA.Participant == ourAddress || nodeB.Participant == ourAddress

		// Calculate the deadline
		deadline := time.Unix(int64(match.StartTime), 0).Add(24 * time.Hour) // Assuming 24 hours match duration

		// Calculate the priority based on deadline and participation
		priority := m.calculatePriority(deadline, isOurs)

		// Create or update the match info
		matchInfo := &MatchInfo{
			TournamentAddress: tournamentAddr,
			MatchIndex:        uint64(i),
			Match:             match,
			Priority:          priority,
			Deadline:          deadline,
			IsOurs:            isOurs,
		}

		// Update our tracking
		m.matches[tournamentAddr][uint64(i)] = matchInfo

		// Add to priority queue if it's our match
		if isOurs {
			heap.Push(&m.priorityQueue, matchInfo)
		}
	}

	return nil
}

// getNode gets a node from a tournament
func (m *MatchMonitor) getNode(ctx context.Context, tournamentAddr common.Address, nodeIndex uint64) (Node, error) {
	// Get the tournament player
	player, err := m.manager.GetTournamentPlayer(ctx, tournamentAddr)
	if err != nil {
		return Node{}, fmt.Errorf("failed to get tournament player: %w", err)
	}

	// Get the tournament contract
	loader, ok := player.loader.(TournamentContract)
	if !ok {
		return Node{}, fmt.Errorf("tournament player loader is not a TournamentContract")
	}

	// Get the node
	return loader.GetNode(ctx, nodeIndex)
}

// calculatePriority calculates the priority of a match based on its deadline and our participation
func (m *MatchMonitor) calculatePriority(deadline time.Time, isOurs bool) MatchPriority {
	now := m.clock.Now()

	// If it's not our match, it's low priority
	if !isOurs {
		return PriorityLow
	}

	// Calculate time until deadline
	timeUntilDeadline := deadline.Sub(now)

	// If the deadline has passed, it's low priority (too late)
	if timeUntilDeadline <= 0 {
		return PriorityLow
	}

	// If the deadline is very close, it's critical priority
	if timeUntilDeadline <= m.deadlineBuffer {
		return PriorityCritical
	}

	// If the deadline is within 6 hours, it's high priority
	if timeUntilDeadline <= 6*time.Hour {
		return PriorityHigh
	}

	// If the deadline is within 12 hours, it's medium priority
	if timeUntilDeadline <= 12*time.Hour {
		return PriorityMedium
	}

	// Otherwise, it's low priority
	return PriorityLow
}

// GetHighestPriorityMatch returns the highest priority match
func (m *MatchMonitor) GetHighestPriorityMatch() (*MatchInfo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.priorityQueue.Len() == 0 {
		return nil, nil
	}

	// Peek at the highest priority match without removing it
	return m.priorityQueue[0], nil
}

// PopHighestPriorityMatch removes and returns the highest priority match
func (m *MatchMonitor) PopHighestPriorityMatch() (*MatchInfo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.priorityQueue.Len() == 0 {
		return nil, nil
	}

	// Remove and return the highest priority match
	return heap.Pop(&m.priorityQueue).(*MatchInfo), nil
}

// GetMatchesByPriority returns all matches grouped by priority
func (m *MatchMonitor) GetMatchesByPriority() map[MatchPriority][]*MatchInfo {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make(map[MatchPriority][]*MatchInfo)

	// Initialize the priority groups
	result[PriorityLow] = make([]*MatchInfo, 0)
	result[PriorityMedium] = make([]*MatchInfo, 0)
	result[PriorityHigh] = make([]*MatchInfo, 0)
	result[PriorityCritical] = make([]*MatchInfo, 0)

	// Group matches by priority
	for _, tournamentMatches := range m.matches {
		for _, matchInfo := range tournamentMatches {
			result[matchInfo.Priority] = append(result[matchInfo.Priority], matchInfo)
		}
	}

	return result
}

// GetMatchesForTournament returns all matches for a specific tournament
func (m *MatchMonitor) GetMatchesForTournament(tournamentAddr common.Address) []*MatchInfo {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]*MatchInfo, 0)

	if tournamentMatches, ok := m.matches[tournamentAddr]; ok {
		for _, matchInfo := range tournamentMatches {
			result = append(result, matchInfo)
		}
	}

	return result
}

// GetMatchInfo returns information about a specific match
func (m *MatchMonitor) GetMatchInfo(tournamentAddr common.Address, matchIndex uint64) (*MatchInfo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if tournamentMatches, ok := m.matches[tournamentAddr]; ok {
		if matchInfo, ok := tournamentMatches[matchIndex]; ok {
			return matchInfo, nil
		}
	}

	return nil, fmt.Errorf("match not found: tournament=%s, index=%d", tournamentAddr.Hex(), matchIndex)
}

// GetMatchCount returns the total number of matches being monitored
func (m *MatchMonitor) GetMatchCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	count := 0
	for _, tournamentMatches := range m.matches {
		count += len(tournamentMatches)
	}

	return count
}

// GetOurMatchCount returns the number of matches where we are a participant
func (m *MatchMonitor) GetOurMatchCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	count := 0
	for _, tournamentMatches := range m.matches {
		for _, matchInfo := range tournamentMatches {
			if matchInfo.IsOurs {
				count++
			}
		}
	}

	return count
}

// ResolveMatch resolves a match
func (m *MatchMonitor) ResolveMatch(ctx context.Context, tournamentAddr common.Address, matchIndex uint64, winnerIndex uint64) error {
	// Get the tournament player
	player, err := m.manager.GetTournamentPlayer(ctx, tournamentAddr)
	if err != nil {
		return fmt.Errorf("failed to get tournament player: %w", err)
	}

	// Get the tournament contract
	loader, ok := player.loader.(TournamentContract)
	if !ok {
		return fmt.Errorf("tournament player loader is not a TournamentContract")
	}

	// Resolve the match
	if err := loader.ResolveMatch(ctx, matchIndex, winnerIndex); err != nil {
		return fmt.Errorf("failed to resolve match: %w", err)
	}

	// Update our tracking
	m.mu.Lock()
	defer m.mu.Unlock()

	if tournamentMatches, ok := m.matches[tournamentAddr]; ok {
		if matchInfo, ok := tournamentMatches[matchIndex]; ok {
			// Mark the match as resolved
			matchInfo.Match.Resolved = true
			// Remove from our tracking
			delete(tournamentMatches, matchIndex)
		}
	}

	return nil
}

// ProcessHighestPriorityMatch processes the highest priority match
func (m *MatchMonitor) ProcessHighestPriorityMatch(ctx context.Context) error {
	// Get the highest priority match
	matchInfo, err := m.PopHighestPriorityMatch()
	if err != nil {
		return fmt.Errorf("failed to get highest priority match: %w", err)
	}

	if matchInfo == nil {
		m.logger.Debug("No matches to process")
		return nil
	}

	m.logger.Info("Processing highest priority match",
		"tournament", matchInfo.TournamentAddress,
		"match", matchInfo.MatchIndex,
		"priority", matchInfo.Priority.String(),
		"deadline", matchInfo.Deadline)

	// Get the tournament player
	player, err := m.manager.GetTournamentPlayer(ctx, matchInfo.TournamentAddress)
	if err != nil {
		return fmt.Errorf("failed to get tournament player: %w", err)
	}

	// Get the tournament contract
	loader, ok := player.loader.(TournamentContract)
	if !ok {
		return fmt.Errorf("tournament player loader is not a TournamentContract")
	}

	// Get the nodes involved in the match
	nodeA, err := loader.GetNode(ctx, matchInfo.Match.NodeA)
	if err != nil {
		return fmt.Errorf("failed to get node A: %w", err)
	}

	nodeB, err := loader.GetNode(ctx, matchInfo.Match.NodeB)
	if err != nil {
		return fmt.Errorf("failed to get node B: %w", err)
	}

	// Determine which node is ours
	ourAddress := m.txSender.From()
	var ourNode, theirNode Node
	var ourIndex, theirIndex uint64

	if nodeA.Participant == ourAddress {
		ourNode = nodeA
		theirNode = nodeB
		ourIndex = matchInfo.Match.NodeA
		theirIndex = matchInfo.Match.NodeB
	} else if nodeB.Participant == ourAddress {
		ourNode = nodeB
		theirNode = nodeA
		ourIndex = matchInfo.Match.NodeB
		theirIndex = matchInfo.Match.NodeA
	} else {
		return fmt.Errorf("we are not a participant in this match")
	}

	// Determine the winner based on trace evidence
	// This is a simplified example - in a real implementation, you would need to
	// implement proper verification logic based on your specific requirements
	winnerIndex, err := m.determineWinner(ctx, matchInfo.Match, ourNode, theirNode, ourIndex, theirIndex)
	if err != nil {
		return fmt.Errorf("failed to determine winner: %w", err)
	}

	// Resolve the match
	if err := m.ResolveMatch(ctx, matchInfo.TournamentAddress, matchInfo.MatchIndex, winnerIndex); err != nil {
		return fmt.Errorf("failed to resolve match: %w", err)
	}

	m.logger.Info("Successfully processed match",
		"tournament", matchInfo.TournamentAddress,
		"match", matchInfo.MatchIndex,
		"winner", winnerIndex)

	return nil
}

// determineWinner determines the winner of a match based on trace evidence
func (m *MatchMonitor) determineWinner(
	ctx context.Context,
	match Match,
	ourNode Node,
	theirNode Node,
	ourIndex uint64,
	theirIndex uint64,
) (uint64, error) {
	// Get the tournament player
	player, err := m.manager.GetTournamentPlayer(ctx, ourNode.Participant)
	if err != nil {
		return 0, fmt.Errorf("failed to get tournament player: %w", err)
	}

	// Get the trace provider
	traceProvider, ok := player.traceProvider.(*TournamentTraceProvider)
	if !ok {
		return 0, fmt.Errorf("tournament player trace provider is not a TournamentTraceProvider")
	}

	// Generate proofs for both claims
	ourProof, err := m.generateProofForClaim(ctx, traceProvider, ourNode.Claim)
	if err != nil {
		return 0, fmt.Errorf("failed to generate proof for our node: %w", err)
	}

	theirProof, err := m.generateProofForClaim(ctx, traceProvider, theirNode.Claim)
	if err != nil {
		return 0, fmt.Errorf("failed to generate proof for their node: %w", err)
	}

	// Compare the proofs and determine the winner
	// This is a simplified example - in a real implementation, you would need to
	// implement proper verification logic based on your specific requirements
	if len(ourProof) > len(theirProof) {
		return ourIndex, nil
	} else {
		return theirIndex, nil
	}
}

// generateProofForClaim generates a proof for a claim
func (m *MatchMonitor) generateProofForClaim(
	ctx context.Context,
	traceProvider *TournamentTraceProvider,
	claim common.Hash,
) ([]byte, error) {
	// In a real implementation, you would need to find the trace index that corresponds to the claim
	// For simplicity, we'll use a fixed index here
	traceIndex := uint64(1)

	// Generate the proof
	proof, err := traceProvider.GenerateProof(ctx, traceIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to generate proof: %w", err)
	}

	return proof, nil
}
