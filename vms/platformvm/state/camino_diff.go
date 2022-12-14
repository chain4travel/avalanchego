// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/dao"
	"github.com/ava-labs/avalanchego/vms/platformvm/deposit"
	"github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
)

func NewCaminoDiff(
	parentID ids.ID,
	stateVersions Versions,
) (Diff, error) {
	parentState, ok := stateVersions.GetState(parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, parentID)
	}
	return &diff{
		parentID:      parentID,
		stateVersions: stateVersions,
		timestamp:     parentState.GetTimestamp(),
		caminoDiff:    newCaminoDiff(),
	}, nil
}

func (d *diff) LockedUTXOs(txIDs ids.Set, addresses ids.ShortSet, lockState locked.State) ([]*avax.UTXO, error) {
	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}

	retUtxos, err := parentState.LockedUTXOs(txIDs, addresses, lockState)
	if err != nil {
		return nil, err
	}

	// Apply modifiedUTXO's
	// Step 1: remove / update existing UTXOs
	remaining := ids.NewSet(len(d.modifiedUTXOs))
	for k := range d.modifiedUTXOs {
		remaining.Add(k)
	}
	for i := len(retUtxos) - 1; i >= 0; i-- {
		if utxo, exists := d.modifiedUTXOs[retUtxos[i].InputID()]; exists {
			if utxo.utxo == nil {
				retUtxos = append(retUtxos[:i], retUtxos[i+1:]...)
			} else {
				retUtxos[i] = utxo.utxo
			}
			delete(remaining, utxo.utxoID)
		}
	}

	// Step 2: Append new UTXOs
	for utxoID := range remaining {
		utxo := d.modifiedUTXOs[utxoID].utxo
		if utxo != nil {
			if lockedOut, ok := utxo.Out.(*locked.Out); ok && lockedOut.IDs.Match(lockState, txIDs) {
				retUtxos = append(retUtxos, utxo)
			}
		}
	}

	return retUtxos, nil
}

func (d *diff) CaminoGenesisState() (*genesis.Camino, error) {
	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}
	return parentState.CaminoGenesisState()
}

func (d *diff) SetAddressStates(address ids.ShortID, states uint64) {
	d.caminoDiff.modifiedAddressStates[address] = states
}

func (d *diff) GetAddressStates(address ids.ShortID) (uint64, error) {
	if states, ok := d.caminoDiff.modifiedAddressStates[address]; ok {
		return states, nil
	}

	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}

	return parentState.GetAddressStates(address)
}

func (d *diff) AddDepositOffer(offer *deposit.Offer) {
	d.caminoDiff.modifiedDepositOffers[offer.ID] = offer
}

func (d *diff) GetDepositOffer(offerID ids.ID) (*deposit.Offer, error) {
	if offer, ok := d.caminoDiff.modifiedDepositOffers[offerID]; ok {
		return offer, nil
	}

	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}

	return parentState.GetDepositOffer(offerID)
}

func (d *diff) GetAllDepositOffers() ([]*deposit.Offer, error) {
	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}

	parentOffers, err := parentState.GetAllDepositOffers()
	if err != nil {
		return nil, err
	}

	var offers []*deposit.Offer

	for _, offer := range d.caminoDiff.modifiedDepositOffers {
		if offer != nil {
			offers = append(offers, offer)
		}
	}

	for _, offer := range parentOffers {
		if _, ok := d.caminoDiff.modifiedDepositOffers[offer.ID]; !ok {
			offers = append(offers, offer)
		}
	}

	return offers, nil
}

func (d *diff) UpdateDeposit(depositTxID ids.ID, deposit *deposit.Deposit) {
	d.caminoDiff.modifiedDeposits[depositTxID] = deposit
}

func (d *diff) GetDeposit(depositTxID ids.ID) (*deposit.Deposit, error) {
	if deposit, ok := d.caminoDiff.modifiedDeposits[depositTxID]; ok {
		return deposit, nil
	}

	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}

	return parentState.GetDeposit(depositTxID)
}

// Voting
func (d *diff) GetAllProposals() ([]*ProposalLookup, error) {
	proposals := make([]*ProposalLookup, len(d.caminoDiff.modifiedProposalLookups))
	i := 0
	for _, proposal := range d.caminoDiff.modifiedProposalLookups {
		proposals[i] = proposal
		i++
	}

	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}

	parentProposals, err := parentState.GetAllProposals()
	if err != nil {
		return nil, err
	}

	return append(proposals, parentProposals...), nil
}

func (d *diff) GetProposalLookup(proposalID ids.ID) (*ProposalLookup, error) {
	if proposal, ok := d.caminoDiff.modifiedProposalLookups[proposalID]; ok {
		return proposal, nil
	}

	parentState, ok := d.stateVersions.GetState(d.parentID)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrMissingParentState, d.parentID)
	}

	return parentState.GetProposalLookup(proposalID)
}

func (d *diff) AddProposal(propsalID ids.ID, proposal *dao.Proposal, state dao.ProposalState) {
	d.caminoDiff.modifiedProposalLookups[propsalID] = &ProposalLookup{
		proposal, make(map[ids.ID]*dao.Vote), state,
	}
}

func (d *diff) AddProposalLookup(propsalID ids.ID, lookup *ProposalLookup) {
	d.caminoDiff.modifiedProposalLookups[propsalID] = lookup
}

func (d *diff) ArchiveProposal(proposalID ids.ID) error {
	proposal, err := d.GetProposalLookup(proposalID)
	if err != nil {
		return err
	}
	for k := range proposal.Votes {
		delete(proposal.Votes, k)
	}

	d.caminoDiff.modifiedProposalLookups[proposalID] = proposal
	return nil
}

func (d *diff) SetProposalState(proposalID ids.ID, state dao.ProposalState) error {
	proposal, err := d.GetProposalLookup(proposalID)
	if err != nil {
		return err
	}

	proposal.State = state

	d.caminoDiff.modifiedProposalLookups[proposalID] = proposal

	return nil
}

func (d *diff) AddVote(proposalID ids.ID, voteID ids.ID, vote *dao.Vote) error {
	proposal, err := d.GetProposalLookup(proposalID)
	if err != nil {
		return err
	}

	proposal.Votes[voteID] = vote

	d.caminoDiff.modifiedProposalLookups[proposalID] = proposal
	return nil
}

// Finally apply all changes
func (d *diff) ApplyCaminoState(baseState State) {
	for k, v := range d.caminoDiff.modifiedAddressStates {
		baseState.SetAddressStates(k, v)
	}

	for _, depositOffer := range d.caminoDiff.modifiedDepositOffers {
		baseState.AddDepositOffer(depositOffer)
	}

	for depositTxID, deposit := range d.caminoDiff.modifiedDeposits {
		baseState.UpdateDeposit(depositTxID, deposit)
	}

	for k, v := range d.caminoDiff.modifiedProposalLookups {
		baseState.AddProposalLookup(k, v)
	}
}
