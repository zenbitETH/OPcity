package multiphase

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Engine is responsible for coordinating the different phases of the dispute resolution process.
type Engine struct {
	log      log.Logger
	txMgr    txmgr.TxManager
	cfg      *Config
	contract common.Address
}

// Config contains the configuration for the multiphase engine.
type Config struct {
	// Add configuration fields as needed
}

// NewEngine creates a new multiphase game engine.
func NewEngine(log log.Logger, txMgr txmgr.TxManager, cfg *Config, contract common.Address) *Engine {
	return &Engine{
		log:      log,
		txMgr:    txMgr,
		cfg:      cfg,
		contract: contract,
	}
}

// transactionOpts creates a new transaction options object for sending transactions to the blockchain.
// It sets up the transaction with the appropriate gas price, nonce, and other parameters.
func (e *Engine) transactionOpts(ctx context.Context) (*bind.TransactOpts, error) {
	// Get the chain ID from the transaction manager
	chainID := e.txMgr.ChainID()

	// Get the sender address from the transaction manager
	from := e.txMgr.From()

	// Get the suggested gas price caps from the transaction manager
	tipCap, baseFee, blobBaseFee, err := e.txMgr.SuggestGasPriceCaps(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggested gas price caps: %w", err)
	}

	// Create a new keyed transactor with the chain ID
	opts := &bind.TransactOpts{
		From:      from,
		GasTipCap: tipCap,
		GasFeeCap: new(big.Int).Add(baseFee, tipCap),
		Context:   ctx,
	}

	// Set blob fee cap if blob transactions are supported (EIP-4844)
	if blobBaseFee != nil {
		opts.BlobFeeCap = blobBaseFee
	}

	// Log the transaction options for debugging
	e.log.Debug("Created transaction options",
		"from", opts.From.Hex(),
		"gasTipCap", opts.GasTipCap,
		"gasFeeCap", opts.GasFeeCap,
		"blobFeeCap", opts.BlobFeeCap,
	)

	return opts, nil
}

// ProcessPhase processes a specific phase of the dispute resolution.
func (e *Engine) ProcessPhase(ctx context.Context, phaseID uint64) error {
	opts, err := e.transactionOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create transaction options: %w", err)
	}

	// Use the transaction options to interact with the contract
	// Implementation will depend on the specific requirements

	e.log.Info("Processing phase", "phaseID", phaseID, "from", opts.From.Hex())
	return nil
}

// GetGameStatus retrieves the current status of the game.
func (e *Engine) GetGameStatus(ctx context.Context) (types.GameStatus, error) {
	// Implementation will depend on the specific requirements
	return types.GameStatusInProgress, nil
}

// AdvanceGame attempts to advance the game to the next phase.
func (e *Engine) AdvanceGame(ctx context.Context) error {
	opts, err := e.transactionOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create transaction options: %w", err)
	}

	// Use the transaction options to interact with the contract
	// Implementation will depend on the specific requirements

	e.log.Info("Advancing game", "from", opts.From.Hex())
	return nil
}
