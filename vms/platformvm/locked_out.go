// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"errors"
	"fmt"

	"github.com/chain4travel/caminogo/ids"
	"github.com/chain4travel/caminogo/vms/components/avax"
)

var (
	errInvalidLockState = errors.New("invalid lockState")
	errNestedLocks      = errors.New("shouldn't nest locks")
	thisTxID            = ids.ID{'t', 'h', 'i', 's', ' ', 't', 'x', ' ', 'i', 'd'}
)

type LockState byte

const (
	LockStateUnlocked        LockState = 0b00
	LockStateDeposited       LockState = 0b01
	LockStateBonded          LockState = 0b10
	LockStateDepositedBonded LockState = 0b11
)

var lockStateStrings = map[LockState]string{
	LockStateUnlocked:        "unlocked",
	LockStateDeposited:       "deposited",
	LockStateBonded:          "bonded",
	LockStateDepositedBonded: "depositedBonded",
}

func (ls LockState) String() string {
	lockStateString, ok := lockStateStrings[ls]
	if !ok {
		return fmt.Sprintf("unknownLockState(%d)", ls)
	}
	return lockStateString
}

func (ls LockState) Verify() error {
	if ls < LockStateUnlocked || LockStateDepositedBonded < ls {
		return errInvalidLockState
	}
	return nil
}

func (ls LockState) IsBonded() bool {
	return LockStateBonded&ls == LockStateBonded
}

func (ls LockState) IsDeposited() bool {
	return LockStateDeposited&ls == LockStateDeposited
}

func (ls LockState) IsLockedWith(lockState LockState) bool {
	return ls&lockState == lockState
}

func (ls LockState) IsLocked() bool {
	return ls != LockStateUnlocked
}

type LockIDs struct {
	DepositTxID ids.ID `serialize:"true" json:"depositTxID"`
	BondTxID    ids.ID `serialize:"true" json:"bondTxID"`
}

func (lock LockIDs) LockState() LockState {
	lockState := LockStateUnlocked
	if lock.DepositTxID != ids.Empty {
		lockState = LockStateDeposited
	}
	if lock.BondTxID != ids.Empty {
		lockState |= LockStateBonded
	}
	return lockState
}

func (lock LockIDs) Lock(lockState LockState) LockIDs {
	if lockState.IsDeposited() {
		lock.DepositTxID = thisTxID
	}
	if lockState.IsBonded() {
		lock.BondTxID = thisTxID
	}
	return lock
}

func (lock LockIDs) Unlock(lockState LockState) LockIDs {
	if lockState.IsDeposited() {
		lock.DepositTxID = ids.Empty
	}
	if lockState.IsBonded() {
		lock.BondTxID = ids.Empty
	}
	return lock
}

type LockedOut struct {
	LockIDs              `serialize:"true" json:"lockIDs"`
	avax.TransferableOut `serialize:"true" json:"output"`
}

func (out *LockedOut) Addresses() [][]byte {
	if addressable, ok := out.TransferableOut.(avax.Addressable); ok {
		return addressable.Addresses()
	}
	return nil
}

func (out *LockedOut) Verify() error {
	if _, nested := out.TransferableOut.(*LockedOut); nested {
		return errNestedLocks
	}
	return out.TransferableOut.Verify()
}

type LockedIn struct {
	LockIDs             `serialize:"true" json:"lockIDs"`
	avax.TransferableIn `serialize:"true" json:"input"`
}

func (in *LockedIn) Verify() error {
	if _, nested := in.TransferableIn.(*LockedIn); nested {
		return errNestedLocks
	}
	return in.TransferableIn.Verify()
}
