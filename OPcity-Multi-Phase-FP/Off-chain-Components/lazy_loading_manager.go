package multiphase

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum-optimism/optimism/op-service/sources/batching"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrStateNotFound      = errors.New("state not found")
	ErrStateLoadFailed    = errors.New("failed to load state")
	ErrCacheCapacityLimit = errors.New("cache capacity limit reached")
)

// cacheItem represents an item in the LRU cache
type cacheItem struct {
	key   string
	value []byte
}

// LazyLoadingManager optimizes memory usage during dispute resolution
type LazyLoadingManager struct {
	client       *ethclient.Client
	stateCache   map[string]*list.Element // Maps keys to list elements
	lruList      *list.List               // Tracks LRU order
	cacheMutex   sync.RWMutex
	maxCacheSize int
	cacheHits    int
	cacheMisses  int
	lastEviction time.Time
	loadCount    int64
}

// NewLazyLoadingManager creates a new instance of LazyLoadingManager
func NewLazyLoadingManager(client *ethclient.Client) *LazyLoadingManager {
	return &LazyLoadingManager{
		client:       client,
		stateCache:   make(map[string]*list.Element),
		lruList:      list.New(),
		maxCacheSize: 1000, // Configurable based on memory constraints
		lastEviction: time.Now(),
	}
}

// LoadState loads the state at the given index, using caching for efficiency
func (m *LazyLoadingManager) LoadState(
	ctx context.Context,
	disputeID *big.Int,
	index *big.Int,
) ([]byte, error) {
	// Generate cache key
	cacheKey := fmt.Sprintf("%s:%s", disputeID.String(), index.String())

	// Try to get from cache first
	m.cacheMutex.RLock()
	if element, exists := m.stateCache[cacheKey]; exists {
		// Move to front of list to mark as most recently used
		m.cacheMutex.RUnlock()

		// Need write lock to modify the list
		m.cacheMutex.Lock()
		m.lruList.MoveToFront(element)
		item := element.Value.(*cacheItem)
		state := item.value
		m.cacheHits++
		m.cacheMutex.Unlock()

		return state, nil
	}
	m.cacheMisses++
	m.cacheMutex.RUnlock()

	// If not in cache, load from PreimageOracle
	state, err := m.fetchStateFromOracle(ctx, disputeID, index)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrStateLoadFailed, err)
	}

	// Store in cache
	m.cacheMutex.Lock()
	defer m.cacheMutex.Unlock()

	// Check if we need to evict some entries
	if len(m.stateCache) >= m.maxCacheSize {
		m.evictCacheEntries()
	}

	// Add to cache and LRU list
	item := &cacheItem{
		key:   cacheKey,
		value: state,
	}
	element := m.lruList.PushFront(item)
	m.stateCache[cacheKey] = element
	m.loadCount++

	return state, nil
}

// fetchStateFromOracle fetches state from the PreimageOracle
func (m *LazyLoadingManager) fetchStateFromOracle(
	ctx context.Context,
	disputeID *big.Int,
	index *big.Int,
) ([]byte, error) {
	// 1. Create a connection to the dispute game contract
	gameAddr := common.HexToAddress(disputeID.String())

	// 2. Create a batching caller for efficient RPC calls
	caller := batching.NewMultiCaller(m.client, batching.DefaultBatchSize)

	// 3. Create a contract instance for the dispute game
	gameContract, err := contracts.NewFaultDisputeGameContract(ctx, nil, gameAddr, caller)
	if err != nil {
		return nil, fmt.Errorf("failed to create game contract: %w", err)
	}

	// 4. Get the PreimageOracle from the game contract
	oracle, err := gameContract.GetOracle(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get oracle: %w", err)
	}

	// 5. Compute the preimage key for the state at the given index
	// The key format depends on the specific implementation, but typically includes the index
	// We'll use a simple key derivation based on the dispute ID and index
	keyBytes := append(disputeID.Bytes(), index.Bytes()...)
	key := crypto.Keccak256(keyBytes)

	// 6. Create the PreimageOracleData structure
	oracleData := types.NewPreimageOracleData(key, nil, 0)

	// 7. Check if the data exists in the oracle
	exists, err := oracle.GlobalDataExists(ctx, oracleData)
	if err != nil {
		return nil, fmt.Errorf("failed to check if data exists: %w", err)
	}

	if !exists {
		return nil, ErrStateNotFound
	}

	// 8. Get the state data from the oracle
	stateBytes, err := oracle.GetGlobalData(ctx, oracleData)
	if err != nil {
		return nil, fmt.Errorf("failed to get state data: %w", err)
	}

	// 9. Return the state data
	return stateBytes[:], nil
}

// evictCacheEntries evicts entries from the cache using an LRU algorithm
func (m *LazyLoadingManager) evictCacheEntries() {
	// Proper LRU eviction: remove the least recently used item
	// which is at the back of the list
	if m.lruList.Len() == 0 {
		return // Nothing to evict
	}

	// Get the least recently used element
	element := m.lruList.Back()
	if element != nil {
		// Remove from the list
		m.lruList.Remove(element)

		// Get the item and remove from the map
		item := element.Value.(*cacheItem)
		delete(m.stateCache, item.key)
	}

	m.lastEviction = time.Now()
}

// GetCacheStats returns statistics about the cache
func (m *LazyLoadingManager) GetCacheStats() (hits int, misses int, entries int, loadCount int64) {
	m.cacheMutex.RLock()
	defer m.cacheMutex.RUnlock()

	return m.cacheHits, m.cacheMisses, len(m.stateCache), m.loadCount
}
