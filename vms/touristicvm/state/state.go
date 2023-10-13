// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"fmt"
	"github.com/ava-labs/avalanchego/vms/touristicvm/locked"
	"math"
	"time"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/cache/metercacher"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/touristicvm/blocks"
	"github.com/ava-labs/avalanchego/vms/touristicvm/status"
	"github.com/ava-labs/avalanchego/vms/touristicvm/txs"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	txCacheSize = 2048
)

var (
	_ State = &state{}

	// These are prefixes for db keys.
	// It's important to set different prefixes for each separate database objects.
	singletonStatePrefix = []byte("singleton")
	blockStatePrefix     = []byte("block")
	txPrefix             = []byte("tx")
	utxoPrefix           = []byte("utxo")
)

type ReadOnlyChain interface {
	avax.UTXOGetter

	GetTx(txID ids.ID) (*txs.Tx, status.Status, error)
	GetBlock(blkID ids.ID) (blocks.Block, error)
	GetTimestamp() time.Time
}
type Chain interface {
	ReadOnlyChain
	avax.UTXOAdder
	avax.UTXODeleter
	Chequebook

	AddTx(tx *txs.Tx, status status.Status)
	SetTimestamp(t time.Time)

	GetCurrentSupply() (uint64, error)
	SetCurrentSupply(cs uint64)

	LockedUTXOs(ids.ShortID) ([]*avax.UTXO, error)
}

// State is a wrapper around avax.SingleTonState and BlockState
// State also exposes a few methods needed for managing database commits and close.
type State interface {
	// MetadataState is defined in avalanchego,
	// it is used to understand if db is initialized already.
	BlockState
	Chain
	avax.UTXOReader

	GetLastAccepted() ids.ID
	SetLastAccepted(blkID ids.ID)

	IsInitialized() (bool, error)
	SetInitialized() error

	Abort()
	Commit() error
	CommitBatch() (database.Batch, error)
	Close() error
	Load() error
}
type state struct {
	blockState
	chequebookState
	baseDB *versiondb.Database

	// cache of blockID -> Block
	// If the block isn't known, nil is cached.

	addedTxs map[ids.ID]*txAndStatus            // map of txID -> {*txs.Tx, Status}
	txCache  cache.Cacher[ids.ID, *txAndStatus] // txID -> {*txs.Tx, Status}. If the entry is nil, it isn't in the database
	txDB     database.Database

	utxoDB    database.Database
	utxoState avax.UTXOState

	currentSupply, persistedCurrentSupply uint64
	modifiedUTXOs                         map[ids.ID]*avax.UTXO // map of modified UTXOID -> *UTXO if the UTXO is nil, it has been removed
	// The persisted fields represent the current database value
	timestamp, persistedTimestamp time.Time
	// [lastAccepted] is the most recently accepted block.
	lastAccepted, persistedLastAccepted ids.ID

	singletonDB MetadataState
}

func (s *state) GetBlock(blkID ids.ID) (blocks.Block, error) {
	//TODO implement me
	panic("implement me")
}

type txBytesAndStatus struct {
	Tx     []byte        `serialize:"true"`
	Status status.Status `serialize:"true"`
}
type txAndStatus struct {
	tx     *txs.Tx
	status status.Status
}

func (s *state) GetCurrentSupply() (uint64, error) {
	return s.currentSupply, nil
}

func (s *state) SetCurrentSupply(cs uint64) {
	s.currentSupply = cs
}

func (s *state) GetTx(txID ids.ID) (*txs.Tx, status.Status, error) {
	if tx, exists := s.addedTxs[txID]; exists {
		return tx.tx, tx.status, nil
	}
	if tx, cached := s.txCache.Get(txID); cached {
		if tx == nil {
			return nil, status.Unknown, database.ErrNotFound
		}
		return tx.tx, tx.status, nil
	}
	txBytes, err := s.txDB.Get(txID[:])
	if err == database.ErrNotFound {
		s.txCache.Put(txID, nil)
		return nil, status.Unknown, database.ErrNotFound
	} else if err != nil {
		return nil, status.Unknown, err
	}

	stx := txBytesAndStatus{}
	if _, err := txs.GenesisCodec.Unmarshal(txBytes, &stx); err != nil {
		return nil, status.Unknown, err
	}

	tx, err := txs.Parse(txs.GenesisCodec, stx.Tx)
	if err != nil {
		return nil, status.Unknown, err
	}

	ptx := &txAndStatus{
		tx:     tx,
		status: stx.Status,
	}

	s.txCache.Put(txID, ptx)
	return ptx.tx, ptx.status, nil
}

func (s *state) AddTx(tx *txs.Tx, status status.Status) {
	s.addedTxs[tx.ID()] = &txAndStatus{
		tx:     tx,
		status: status,
	}
}

func (s *state) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	if utxo, exists := s.modifiedUTXOs[utxoID]; exists {
		if utxo == nil {
			return nil, database.ErrNotFound
		}
		return utxo, nil
	}
	return s.utxoState.GetUTXO(utxoID)
}

func (s *state) GetLastAccepted() ids.ID {
	return s.lastAccepted
}

func (s *state) GetTimestamp() time.Time {
	return s.timestamp
}

func (s *state) UTXOIDs(addr []byte, start ids.ID, limit int) ([]ids.ID, error) {
	return s.utxoState.UTXOIDs(addr, start, limit)

}

func (s *state) AddUTXO(utxo *avax.UTXO) {
	s.modifiedUTXOs[utxo.InputID()] = utxo
}

func (s *state) DeleteUTXO(utxoID ids.ID) {
	s.modifiedUTXOs[utxoID] = nil
}

func (s *state) SetLastAccepted(lastAccepted ids.ID) {
	s.lastAccepted = lastAccepted
}

func (s *state) SetTimestamp(t time.Time) {
	s.timestamp = t
}
func (s *state) LockedUTXOs(address ids.ShortID) ([]*avax.UTXO, error) {
	retUtxos := []*avax.UTXO{}
	utxoIDs, err := s.UTXOIDs(address.Bytes(), ids.ID{}, math.MaxInt)
	if err != nil {
		return nil, err
	}
	for _, utxoID := range utxoIDs {
		utxo, err := s.GetUTXO(utxoID)
		if err != nil {
			return nil, err
		}
		if utxo == nil {
			continue
		}
		if _, ok := utxo.Out.(*locked.Out); ok {
			retUtxos = append(retUtxos, utxo)
		}
	}
	return retUtxos, nil
}

func NewState(db database.Database, metricsReg prometheus.Registerer) (State, error) {
	// create a new baseDB
	baseDB := versiondb.New(db)

	// create a prefixed "blockDB" from baseDB
	blockDB := prefixdb.New(blockStatePrefix, baseDB)
	// create a prefixed "singletonDB" from baseDB
	singletonDB := prefixdb.New(singletonStatePrefix, baseDB)

	txCache, err := metercacher.New[ids.ID, *txAndStatus](
		"tx_cache",
		metricsReg,
		&cache.LRU[ids.ID, *txAndStatus]{Size: txCacheSize},
	)
	if err != nil {
		return nil, err
	}
	utxoDB := prefixdb.New(utxoPrefix, baseDB)
	utxoState, err := avax.NewMeteredUTXOState(utxoDB, txs.GenesisCodec, metricsReg)
	if err != nil {
		return nil, err
	}

	// return state with created sub state components
	return &state{
		blockState: blockState{
			addedBlocks: make(map[ids.ID]blkWrapper),
			blockCache:  &cache.LRU[ids.ID, *blkWrapper]{Size: blockCacheSize},
			blockDB:     blockDB,
		},
		singletonDB: MetadataState{singletonDB},
		baseDB:      baseDB,

		addedTxs: make(map[ids.ID]*txAndStatus),
		txDB:     prefixdb.New(txPrefix, baseDB),
		txCache:  txCache,

		modifiedUTXOs: make(map[ids.ID]*avax.UTXO),
		utxoDB:        utxoDB,
		utxoState:     utxoState,
		chequebookState: chequebookState{
			lastCheque: make(map[IssuerAgent]map[ids.ShortID]Cheque),
		},
	}, nil
}

// Commit commits pending operations to baseDB
func (s *state) Commit() error {
	defer s.Abort()
	batch, err := s.CommitBatch()
	if err != nil {
		return err
	}
	return batch.Write()
}
func (s *state) CommitBatch() (database.Batch, error) {
	if err := s.write(); err != nil {
		return nil, err
	}
	return s.baseDB.CommitBatch()
}

func (s *state) writeBlocks() error {
	for blkID, stateBlk := range s.addedBlocks {
		var (
			blkID = blkID
			stBlk = stateBlk
		)

		// Note: blocks to be stored are verified, so it's safe to marshal them with GenesisCodec
		blockBytes, err := blocks.GenesisCodec.Marshal(blocks.Version, &stBlk)
		if err != nil {
			return fmt.Errorf("failed to marshal block %s to store: %w", blkID, err)
		}

		delete(s.addedBlocks, blkID)
		s.blockCache.Put(blkID, &stBlk)
		if err := s.blockDB.Put(blkID[:], blockBytes); err != nil {
			return fmt.Errorf("failed to write block %s: %w", blkID, err)
		}
	}
	return nil
}

func (s *state) writeTXs() error {
	for txID, txStatus := range s.addedTxs {
		txID := txID

		stx := txBytesAndStatus{
			Tx:     txStatus.tx.Bytes(),
			Status: txStatus.status,
		}

		// Note that we're serializing a [txBytesAndStatus] here, not a
		// *txs.Tx, so we don't use [txs.Codec].
		txBytes, err := txs.GenesisCodec.Marshal(txs.Version, &stx)
		if err != nil {
			return fmt.Errorf("failed to serialize tx: %w", err)
		}

		delete(s.addedTxs, txID)
		s.txCache.Put(txID, txStatus)
		if err := s.txDB.Put(txID[:], txBytes); err != nil {
			return fmt.Errorf("failed to add tx: %w", err)
		}
	}
	return nil
}

func (s *state) writeUTXOs() error {
	for utxoID, utxo := range s.modifiedUTXOs {
		delete(s.modifiedUTXOs, utxoID)

		if utxo == nil {
			if err := s.utxoState.DeleteUTXO(utxoID); err != nil {
				return fmt.Errorf("failed to delete UTXO: %w", err)
			}
			continue
		}
		if err := s.utxoState.PutUTXO(utxo); err != nil {
			return fmt.Errorf("failed to add UTXO: %w", err)
		}
	}
	return nil
}
func (s *state) write() error {
	errs := wrappers.Errs{}
	errs.Add(
		s.writeBlocks(),
		s.writeTXs(),
		s.writeUTXOs(),
		s.writeMetadata(),
	)
	return errs.Err
}
func (s *state) Abort() {
	s.baseDB.Abort()
}

// Close closes the underlying base database
func (s *state) Close() error {
	errs := wrappers.Errs{}
	errs.Add(
		s.baseDB.Close(),
		s.txDB.Close(),
		s.utxoDB.Close(),
		s.singletonDB.Close(),
	)
	return errs.Err
}

// Load pulls data previously stored on disk that is expected to be in memory.
func (s *state) Load() error {
	errs := wrappers.Errs{}
	errs.Add(
		s.loadMetadata(),
	)
	return errs.Err
}

func (s *state) IsInitialized() (bool, error) {
	return s.singletonDB.IsInitialized()
}

func (s *state) SetInitialized() error {
	return s.singletonDB.SetInitialized()
}

func (s *state) loadMetadata() error {
	timestamp, err := database.GetTimestamp(s.singletonDB, timestampKey)
	if err != nil {
		return err
	}
	s.persistedTimestamp = timestamp
	s.SetTimestamp(timestamp)

	currentSupply, err := database.GetUInt64(s.singletonDB, currentSupplyKey)
	if err != nil {
		return err
	}
	s.persistedCurrentSupply = currentSupply
	s.SetCurrentSupply(currentSupply)

	lastAccepted, err := database.GetID(s.singletonDB, lastAcceptedKey)
	if err != nil {
		return err
	}
	s.persistedLastAccepted = lastAccepted
	s.lastAccepted = lastAccepted
	return nil
}

func (s *state) writeMetadata() error {
	if !s.persistedTimestamp.Equal(s.timestamp) {
		if err := database.PutTimestamp(s.singletonDB, timestampKey, s.timestamp); err != nil {
			return fmt.Errorf("failed to write timestamp: %w", err)
		}
		s.persistedTimestamp = s.timestamp
	}
	if s.persistedCurrentSupply != s.currentSupply {
		if err := database.PutUInt64(s.singletonDB, currentSupplyKey, s.currentSupply); err != nil {
			return fmt.Errorf("failed to write current supply: %w", err)
		}
		s.persistedCurrentSupply = s.currentSupply
	}
	if s.persistedLastAccepted != s.lastAccepted {
		if err := database.PutID(s.singletonDB, lastAcceptedKey, s.lastAccepted); err != nil {
			return fmt.Errorf("failed to write last accepted: %w", err)
		}
		s.persistedLastAccepted = s.lastAccepted
	}
	return nil
}
