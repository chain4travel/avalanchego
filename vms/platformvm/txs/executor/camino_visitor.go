// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import "github.com/ava-labs/avalanchego/vms/platformvm/txs"

// Camino Visitor implementations

// Standard
func (*StandardTxExecutor) AddAddressStateTx(*txs.AddAddressStateTx) error {
	return errWrongTxType
}

func (*StandardTxExecutor) DepositTx(*txs.DepositTx) error {
	return errWrongTxType
}

func (*StandardTxExecutor) UnlockDepositTx(*txs.UnlockDepositTx) error {
	return errWrongTxType
}

func (e *StandardTxExecutor) CreateProposalTx(*txs.CreateProposalTx) error {
	return errWrongTxType
}
func (e *StandardTxExecutor) CreateVoteTx(*txs.CreateVoteTx) error {
	return errWrongTxType
}

// Proposal
func (*ProposalTxExecutor) AddAddressStateTx(*txs.AddAddressStateTx) error {
	return errWrongTxType
}

func (*ProposalTxExecutor) DepositTx(*txs.DepositTx) error {
	return errWrongTxType
}

func (*ProposalTxExecutor) UnlockDepositTx(*txs.UnlockDepositTx) error {
	return errWrongTxType
}

func (*ProposalTxExecutor) CreateProposalTx(*txs.CreateProposalTx) error {
	return errWrongTxType
}

func (*ProposalTxExecutor) CreateVoteTx(*txs.CreateVoteTx) error {
	return errWrongTxType
}

// Atomic
func (*AtomicTxExecutor) AddAddressStateTx(*txs.AddAddressStateTx) error {
	return errWrongTxType
}

func (*AtomicTxExecutor) DepositTx(*txs.DepositTx) error {
	return errWrongTxType
}

func (*AtomicTxExecutor) UnlockDepositTx(*txs.UnlockDepositTx) error {
	return errWrongTxType
}

func (*AtomicTxExecutor) CreateProposalTx(*txs.CreateProposalTx) error {
	return errWrongTxType
}

func (*AtomicTxExecutor) CreateVoteTx(*txs.CreateVoteTx) error {
	return errWrongTxType
}

// MemPool
func (v *MempoolTxVerifier) AddAddressStateTx(tx *txs.AddAddressStateTx) error {
	return v.standardTx(tx)
}

func (v *MempoolTxVerifier) DepositTx(tx *txs.DepositTx) error {
	return v.standardTx(tx)
}

func (v *MempoolTxVerifier) UnlockDepositTx(tx *txs.UnlockDepositTx) error {
	return v.standardTx(tx)
}

func (v *MempoolTxVerifier) CreateProposalTx(tx *txs.CreateProposalTx) error {
	return v.standardTx(tx)
}

func (v *MempoolTxVerifier) CreateVoteTx(tx *txs.CreateVoteTx) error {
	return v.standardTx(tx)
}
