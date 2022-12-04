// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/cache/metercacher"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/linkeddb"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/deposit"
	"github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
	"github.com/ava-labs/avalanchego/vms/platformvm/status"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	addressStateCacheSize = 1024
	depositsCacheSize     = 1024
)

var (
	_ CaminoState = (*caminoState)(nil)

	addressStatePrefix  = []byte("addressState")
	depositOffersPrefix = []byte("depositOffers")
	depositsPrefix      = []byte("deposits")
	proposalsPrefix     = []byte("proposals")
	votesPrefix         = []byte("votes")

	errWrongTxType = errors.New("unexpected tx type")
)

type CaminoApply interface {
	ApplyCaminoState(State)
}

type CaminoDiff interface {
	// Address State

	SetAddressStates(ids.ShortID, uint64)
	GetAddressStates(ids.ShortID) (uint64, error)

	// Deposit offers state

	// precondition: offer.SetID() must be called and return no error
	AddDepositOffer(offer *deposit.Offer)
	GetDepositOffer(offerID ids.ID) (*deposit.Offer, error)
	GetAllDepositOffers() ([]*deposit.Offer, error)

	// Deposits state
	UpdateDeposit(depositTxID ids.ID, deposit *deposit.Deposit)
	GetDeposit(depositTxID ids.ID) (*deposit.Deposit, error)

	// Voting / Proposals
	GetAllProposals() ([]*Proposal, error)
	GetProposal(proposalID ids.ID) (*Proposal, error)
	AddProposal(proposal *Proposal)
	ConcludeProposal(proposalID ids.ID, outcome ProposalOutcome) error

	GetAllVotes() ([]*Vote, error)
	GetVote(voteID ids.ID) (*Vote, error)
	AddVote(vote *Vote)
	// RevokeVote(voteID ids.ID) might be a good idea, specifics might be tricky
}

// For state and diff
type Camino interface {
	CaminoDiff

	LockedUTXOs(set.Set[ids.ID], set.Set[ids.ShortID], locked.State) ([]*avax.UTXO, error)
	CaminoConfig() (*CaminoConfig, error)
}

// For state only
type CaminoState interface {
	CaminoDiff

	CaminoConfig() *CaminoConfig
	SyncGenesis(*state, *genesis.State) error
	Load() error
	Write() error
}

type CaminoConfig struct {
	VerifyNodeSignature bool
	LockModeBondDeposit bool
}

type caminoDiff struct {
	modifiedAddressStates map[ids.ShortID]uint64
	modifiedDepositOffers map[ids.ID]*deposit.Offer
	modifiedDeposits      map[ids.ID]*deposit.Deposit
	modifiedProposals     map[ids.ID]*Proposal
	modifiedVotes         map[ids.ID]*Vote
}

type caminoState struct {
	*caminoDiff

	verifyNodeSignature bool
	lockModeBondDeposit bool

	// Address State
	addressStateCache cache.Cacher
	addressStateDB    database.Database

	// Deposit offers
	depositOffers     map[ids.ID]*deposit.Offer
	depositOffersList linkeddb.LinkedDB
	depositOffersDB   database.Database

	// Deposits
	depositsCache cache.Cacher
	depositsDB    database.Database

	// Proposals
	proposals    map[ids.ID]*Proposal
	proposalList linkeddb.LinkedDB
	proposalsDB  database.Database

	// Vote
	votes    map[ids.ID]*Vote
	voteList linkeddb.LinkedDB
	votesDB  database.Database
}

func newCaminoDiff() *caminoDiff {
	return &caminoDiff{
		modifiedAddressStates: make(map[ids.ShortID]uint64),
		modifiedDepositOffers: make(map[ids.ID]*deposit.Offer),
		modifiedDeposits:      make(map[ids.ID]*deposit.Deposit),
		modifiedProposals:     make(map[ids.ID]*Proposal),
		modifiedVotes:         make(map[ids.ID]*Vote),
	}

}

func newCaminoState(baseDB *versiondb.Database, metricsReg prometheus.Registerer) (*caminoState, error) {
	addressStateCache, err := metercacher.New(
		"address_state_cache",
		metricsReg,
		&cache.LRU{Size: addressStateCacheSize},
	)
	if err != nil {
		return nil, err
	}

	depositsCache, err := metercacher.New(
		"deposits_cache",
		metricsReg,
		&cache.LRU{Size: depositsCacheSize},
	)
	if err != nil {
		return nil, err
	}

	depositOffersDB := prefixdb.New(depositOffersPrefix, baseDB)
	proposalsDB := prefixdb.New(proposalsPrefix, baseDB)
	votesDB := prefixdb.New(votesPrefix, baseDB)

	return &caminoState{
		addressStateDB:    prefixdb.New(addressStatePrefix, baseDB),
		addressStateCache: addressStateCache,

		depositOffers:     make(map[ids.ID]*deposit.Offer),
		depositOffersDB:   depositOffersDB,
		depositOffersList: linkeddb.NewDefault(depositOffersDB),

		depositsCache: depositsCache,
		depositsDB:    prefixdb.New(depositsPrefix, baseDB),

		caminoDiff: newCaminoDiff(),

		proposals:    make(map[ids.ID]*Proposal),
		proposalsDB:  proposalsDB,
		proposalList: linkeddb.NewDefault(proposalsDB),

		votes:    make(map[ids.ID]*Vote),
		votesDB:  votesDB,
		voteList: linkeddb.NewDefault(votesDB),
	}, nil
}

// Return current genesis args
func (cs *caminoState) CaminoConfig() *CaminoConfig {
	return &CaminoConfig{
		VerifyNodeSignature: cs.verifyNodeSignature,
		LockModeBondDeposit: cs.lockModeBondDeposit,
	}
}

// Extract camino tag from genesis
func (cs *caminoState) SyncGenesis(s *state, g *genesis.State) error {
	cs.lockModeBondDeposit = g.Camino.LockModeBondDeposit
	cs.verifyNodeSignature = g.Camino.VerifyNodeSignature

	// adding address states

	for _, addrState := range g.Camino.AddressStates {
		cs.SetAddressStates(addrState.Address, addrState.State)
	}

	initalAdminAddressState, err := cs.GetAddressStates(g.Camino.InitialAdmin)
	if err != nil {
		return err
	}
	cs.SetAddressStates(g.Camino.InitialAdmin,
		initalAdminAddressState|txs.AddressStateRoleAdminBit)

	tx := &txs.AddAddressStateTx{
		Address: g.Camino.InitialAdmin,
		State:   txs.AddressStateRoleAdmin,
		Remove:  false,
	}
	s.AddTx(&txs.Tx{Unsigned: tx}, status.Committed)

	// adding deposit offers

	for _, genesisOffer := range g.Camino.DepositOffers {
		offer := &deposit.Offer{
			InterestRateNominator:   genesisOffer.InterestRateNominator,
			Start:                   genesisOffer.Start,
			End:                     genesisOffer.End,
			MinAmount:               genesisOffer.MinAmount,
			MinDuration:             genesisOffer.MinDuration,
			MaxDuration:             genesisOffer.MaxDuration,
			UnlockPeriodDuration:    genesisOffer.UnlockPeriodDuration,
			NoRewardsPeriodDuration: genesisOffer.NoRewardsPeriodDuration,
			Flags:                   genesisOffer.Flags,
		}
		if err := offer.SetID(); err != nil {
			return err
		}

		cs.AddDepositOffer(offer)
	}

	// adding deposits

	currentTimestamp := uint64(s.GetTimestamp().Unix())

	for _, tx := range g.Camino.Deposits {
		depositTx, ok := tx.Unsigned.(*txs.DepositTx)
		if !ok {
			return errWrongTxType
		}
		depositAmount := uint64(0)
		for _, out := range depositTx.Outs {
			newAmount, err := math.Add64(depositAmount, out.Out.Amount())
			if err != nil {
				return err
			}
			depositAmount = newAmount
		}
		cs.UpdateDeposit(
			tx.ID(),
			&deposit.Deposit{
				DepositOfferID: depositTx.DepositOfferID,
				Start:          currentTimestamp,
				Duration:       depositTx.DepositDuration,
				Amount:         depositAmount,
			},
		)
		s.AddTx(tx, status.Committed)
	}

	return nil
}

func (cs *caminoState) Load() error {

	if err := cs.loadDepositOffers(); err != nil {
		return err
	}

	if err := cs.loadProposals(); err != nil {
		return err
	}

	if err := cs.loadVotes(); err != nil {
		return err
	}

	return nil
}

func (cs *caminoState) Write() error {
	if err := cs.writeAddressStates(); err != nil {
		return err
	}
	if err := cs.writeDeposits(); err != nil {
		return err
	}

	if err := cs.writeDepositOffers(); err != nil {
		return err
	}

	if err := cs.writeProposals(); err != nil {
		return err
	}

	if err := cs.writeVotes(); err != nil {
		return err
	}

	return nil
}
