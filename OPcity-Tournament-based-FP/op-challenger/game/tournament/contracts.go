package tournament

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	"github.com/ethereum-optimism/optimism/op-service/sources/batching"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// TournamentGameContract is a wrapper around the TournamentGame contract
type TournamentGameContract struct {
	contract *bindings.TournamentGame
	caller   *batching.MultiCaller
	txSender TxSender
	addr     common.Address
	logger   log.Logger
}

// NewTournamentGameContract creates a new TournamentGameContract
func NewTournamentGameContract(
	addr common.Address,
	caller *batching.MultiCaller,
	txSender TxSender,
	logger log.Logger,
) *TournamentGameContract {
	contract, _ := bindings.NewTournamentGame(addr, batching.NewBatchingCallOpts(caller))
	return &TournamentGameContract{
		contract: contract,
		caller:   caller,
		txSender: txSender,
		addr:     addr,
		logger:   logger.New("contract", addr),
	}
}

// GetStatus returns the current status of the tournament game
func (t *TournamentGameContract) GetStatus(ctx context.Context) (types.GameStatus, error) {
	status, err := t.contract.Status(&bind.CallOpts{Context: ctx})
	if err != nil {
		return types.GameStatusInProgress, fmt.Errorf("failed to get status: %w", err)
	}
	return types.GameStatusFromUint8(uint8(status))
}

// GetTotalParticipants returns the total number of participants in the tournament
func (t *TournamentGameContract) GetTotalParticipants(ctx context.Context) (uint64, error) {
	participants, err := t.contract.TotalParticipants(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, fmt.Errorf("failed to get total participants: %w", err)
	}
	return participants.Uint64(), nil
}

// GetCurrentRound returns the current round of the tournament
func (t *TournamentGameContract) GetCurrentRound(ctx context.Context) (uint64, error) {
	round, err := t.contract.CurrentRound(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, fmt.Errorf("failed to get current round: %w", err)
	}
	return round.Uint64(), nil
}

// GetMatchCount returns the total number of matches in the tournament
// Note: This is a simplified implementation as the actual contract might not have this method
func (t *TournamentGameContract) GetMatchCount(ctx context.Context) (uint64, error) {
	// In a real implementation, you would need to call a method on the contract
	// For now, we'll return a fixed value
	return 10, nil
}

// HasParticipated checks if an address has participated in the tournament
func (t *TournamentGameContract) HasParticipated(ctx context.Context, addr common.Address) (bool, error) {
	hasParticipated, err := t.contract.HasParticipated(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return false, fmt.Errorf("failed to check if address has participated: %w", err)
	}
	return hasParticipated, nil
}

// JoinTournament joins the tournament with a counter-claim
func (t *TournamentGameContract) JoinTournament(ctx context.Context, claim common.Hash) error {
	// Create the transaction data
	data, err := t.contract.ABI.Pack("joinTournament", claim)
	if err != nil {
		return fmt.Errorf("failed to pack joinTournament call: %w", err)
	}

	// Send the transaction
	tx := txmgr.TxCandidate{
		To:       &t.addr,
		Value:    big.NewInt(1e18), // 1 ETH bond amount (adjust as needed)
		Data:     data,
		GasLimit: 500000, // Adjust as needed
	}

	if err := t.txSender.SendAndWaitSimple("joinTournament", tx); err != nil {
		return fmt.Errorf("failed to send joinTournament transaction: %w", err)
	}

	return nil
}

// ResolveMatch resolves a match in the tournament
func (t *TournamentGameContract) ResolveMatch(ctx context.Context, matchIndex uint64, winnerIndex uint64) error {
	// Create the transaction data
	data, err := t.contract.ABI.Pack("resolveMatch", big.NewInt(int64(matchIndex)), big.NewInt(int64(winnerIndex)))
	if err != nil {
		return fmt.Errorf("failed to pack resolveMatch call: %w", err)
	}

	// Send the transaction
	tx := txmgr.TxCandidate{
		To:       &t.addr,
		Data:     data,
		GasLimit: 500000, // Adjust as needed
	}

	if err := t.txSender.SendAndWaitSimple("resolveMatch", tx); err != nil {
		return fmt.Errorf("failed to send resolveMatch transaction: %w", err)
	}

	return nil
}

// ClaimBond claims the bond as the tournament winner
func (t *TournamentGameContract) ClaimBond(ctx context.Context) error {
	// Create the transaction data
	data, err := t.contract.ABI.Pack("claimBond")
	if err != nil {
		return fmt.Errorf("failed to pack claimBond call: %w", err)
	}

	// Send the transaction
	tx := txmgr.TxCandidate{
		To:       &t.addr,
		Data:     data,
		GasLimit: 300000, // Adjust as needed
	}

	if err := t.txSender.SendAndWaitSimple("claimBond", tx); err != nil {
		return fmt.Errorf("failed to send claimBond transaction: %w", err)
	}

	return nil
}

// GetL1Head returns the L1 head hash at the time the tournament was created
func (t *TournamentGameContract) GetL1Head(ctx context.Context) (common.Hash, error) {
	l1Head, err := t.contract.L1Head(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get L1 head: %w", err)
	}
	return l1Head, nil
}

// GetRootClaim returns the root claim of the tournament
func (t *TournamentGameContract) GetRootClaim(ctx context.Context) (common.Hash, error) {
	rootClaim, err := t.contract.RootClaim(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get root claim: %w", err)
	}
	return rootClaim, nil
}

// GetMatch returns information about a match in the tournament
func (t *TournamentGameContract) GetMatch(ctx context.Context, matchIndex uint64) (Match, error) {
	matchData, err := t.contract.Matches(&bind.CallOpts{Context: ctx}, big.NewInt(int64(matchIndex)))
	if err != nil {
		return Match{}, fmt.Errorf("failed to get match: %w", err)
	}

	return Match{
		NodeA:     matchData.NodeA.Uint64(),
		NodeB:     matchData.NodeB.Uint64(),
		Winner:    matchData.Winner.Uint64(),
		StartTime: matchData.StartTime.Uint64(),
		Resolved:  matchData.Resolved,
	}, nil
}

// GetNode returns information about a node in the tournament tree
func (t *TournamentGameContract) GetNode(ctx context.Context, nodeIndex uint64) (Node, error) {
	nodeData, err := t.contract.Nodes(&bind.CallOpts{Context: ctx}, big.NewInt(int64(nodeIndex)))
	if err != nil {
		return Node{}, fmt.Errorf("failed to get node: %w", err)
	}

	return Node{
		Participant: nodeData.Participant,
		Claim:       nodeData.Claim,
		HasJoined:   nodeData.HasJoined,
		BondAmount:  nodeData.BondAmount.Uint64(),
	}, nil
}
