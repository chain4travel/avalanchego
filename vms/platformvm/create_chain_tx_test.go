// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************

// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chain4travel/caminogo/ids"
	"github.com/chain4travel/caminogo/utils/constants"
	"github.com/chain4travel/caminogo/utils/crypto"
	"github.com/chain4travel/caminogo/utils/hashing"
	"github.com/chain4travel/caminogo/utils/units"
	"github.com/chain4travel/caminogo/vms/components/avax"
	"github.com/chain4travel/caminogo/vms/secp256k1fx"
)

func TestUnsignedCreateChainTxVerify(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	type test struct {
		description string
		shouldErr   bool
		subnetID    ids.ID
		genesisData []byte
		vmID        ids.ID
		fxIDs       []ids.ID
		chainName   string
		keys        []*crypto.PrivateKeySECP256K1R
		setup       func(*UnsignedCreateChainTx) *UnsignedCreateChainTx
	}

	tests := []test{
		{
			description: "tx is nil",
			shouldErr:   true,
			subnetID:    testSubnet1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
			setup:       func(*UnsignedCreateChainTx) *UnsignedCreateChainTx { return nil },
		},
		{
			description: "vm ID is empty",
			shouldErr:   true,
			subnetID:    testSubnet1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
			setup:       func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx { tx.VMID = ids.ID{}; return tx },
		},
		{
			description: "subnet ID is empty",
			shouldErr:   true,
			subnetID:    testSubnet1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
			setup:       func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx { tx.SubnetID = ids.ID{}; return tx },
		},
		{
			description: "subnet ID is platform chain's ID",
			shouldErr:   true,
			subnetID:    testSubnet1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
			setup:       func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx { tx.SubnetID = vm.ctx.ChainID; return tx },
		},
		{
			description: "chain name is too long",
			shouldErr:   true,
			subnetID:    testSubnet1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
			setup: func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx {
				tx.ChainName = string(make([]byte, maxNameLen+1))
				return tx
			},
		},
		{
			description: "chain name has invalid character",
			shouldErr:   true,
			subnetID:    testSubnet1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
			setup: func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx {
				tx.ChainName = "⌘"
				return tx
			},
		},
		{
			description: "genesis data is too long",
			shouldErr:   true,
			subnetID:    testSubnet1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
			setup: func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx {
				tx.GenesisData = make([]byte, maxGenesisLen+1)
				return tx
			},
		},
	}

	for _, test := range tests {
		tx, err := vm.newCreateChainTx(
			test.subnetID,
			test.genesisData,
			test.vmID,
			test.fxIDs,
			test.chainName,
			test.keys,
		)
		if err != nil {
			t.Fatal(err)
		}
		tx.UnsignedTx.(*UnsignedCreateChainTx).syntacticallyVerified = false
		tx.UnsignedTx = test.setup(tx.UnsignedTx.(*UnsignedCreateChainTx))
		if err := tx.UnsignedTx.(*UnsignedCreateChainTx).SyntacticVerify(vm.ctx); err != nil && !test.shouldErr {
			t.Fatalf("test '%s' shouldn't have errored but got: %s", test.description, err)
		} else if err == nil && test.shouldErr {
			t.Fatalf("test '%s' didn't error but should have", test.description)
		}
	}
}

// Ensure Execute fails when there are not enough control sigs
func TestCreateChainTxInsufficientControlSigs(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	tx, err := vm.newCreateChainTx(
		testSubnet1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{keys[0], keys[1]},
	)
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
		vm.internalState.LockedUTXOsChainState(),
	)

	// Remove a signature
	tx.Creds[0].(*secp256k1fx.Credential).Sigs = tx.Creds[0].(*secp256k1fx.Credential).Sigs[1:]
	if _, err := tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx); err == nil {
		t.Fatal("should have errored because a sig is missing")
	}
}

// Ensure Execute fails when an incorrect control signature is given
func TestCreateChainTxWrongControlSig(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	tx, err := vm.newCreateChainTx( // create a tx
		testSubnet1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Generate new, random key to sign tx with
	factory := crypto.FactorySECP256K1R{}
	key, err := factory.NewPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
		vm.internalState.LockedUTXOsChainState(),
	)

	// Replace a valid signature with one from another key
	sig, err := key.SignHash(hashing.ComputeHash256(tx.UnsignedBytes()))
	if err != nil {
		t.Fatal(err)
	}
	copy(tx.Creds[0].(*secp256k1fx.Credential).Sigs[0][:], sig)
	if _, err = tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx); err == nil {
		t.Fatal("should have failed verification because a sig is invalid")
	}
}

// Ensure Execute fails when the Subnet the blockchain specifies as
// its validator set doesn't exist
func TestCreateChainTxNoSuchSubnet(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	tx, err := vm.newCreateChainTx(
		testSubnet1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
	)
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
		vm.internalState.LockedUTXOsChainState(),
	)

	tx.UnsignedTx.(*UnsignedCreateChainTx).SubnetID = ids.GenerateTestID()
	if _, err := tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx); err == nil {
		t.Fatal("should have failed because subent doesn't exist")
	}
}

// Ensure valid tx passes semanticVerify
func TestCreateChainTxValid(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	// create a valid tx
	tx, err := vm.newCreateChainTx(
		testSubnet1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{testSubnet1ControlKeys[0], testSubnet1ControlKeys[1]},
	)
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
		vm.internalState.LockedUTXOsChainState(),
	)

	_, err = tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx)
	if err != nil {
		t.Fatalf("expected tx to pass verification but got error: %v", err)
	}
}

func TestCreateChainTxAP3FeeChange(t *testing.T) {
	ap3Time := defaultGenesisTime.Add(time.Hour)
	tests := []struct {
		name         string
		time         time.Time
		fee          uint64
		expectsError bool
	}{
		{
			name:         "pre-fork - correctly priced",
			time:         defaultGenesisTime,
			fee:          0,
			expectsError: false,
		},
		{
			name:         "post-fork - incorrectly priced",
			time:         ap3Time,
			fee:          100*defaultTxFee - 1*units.NanoAvax,
			expectsError: true,
		},
		{
			name:         "post-fork - correctly priced",
			time:         ap3Time,
			fee:          100 * defaultTxFee,
			expectsError: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			vm, _, _ := defaultVM()
			vm.ApricotPhase3Time = ap3Time

			vm.ctx.Lock.Lock()
			defer func() {
				err := vm.Shutdown()
				assert.NoError(err)
				vm.ctx.Lock.Unlock()
			}()

			ins, outs, _, signers, err := vm.spend(keys, 0, test.fee, LockStateDeposited)
			assert.NoError(err)

			subnetAuth, subnetSigners, err := vm.authorize(vm.internalState, testSubnet1.ID(), keys)
			assert.NoError(err)

			signers = append(signers, subnetSigners)

			// Create the tx

			utx := &UnsignedCreateChainTx{
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    vm.ctx.NetworkID,
					BlockchainID: vm.ctx.ChainID,
					Ins:          ins,
					Outs:         outs,
				}},
				SubnetID:   testSubnet1.ID(),
				VMID:       constants.AVMID,
				SubnetAuth: subnetAuth,
			}
			tx := &Tx{UnsignedTx: utx}
			err = tx.Sign(Codec, signers)
			assert.NoError(err)

			vs := newVersionedState(
				vm.internalState,
				vm.internalState.CurrentStakerChainState(),
				vm.internalState.PendingStakerChainState(),
				vm.internalState.LockedUTXOsChainState(),
			)
			vs.SetTimestamp(test.time)

			_, err = utx.Execute(vm, vs, tx)
			assert.Equal(test.expectsError, err != nil)
		})
	}
}

func TestCreateChainLockedInsOrLockedOuts(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	outputOwners := secp256k1fx.OutputOwners{
		Locktime:  0,
		Threshold: 1,
		Addrs:     []ids.ShortID{keys[0].PublicKey().Address()},
	}
	signers := [][]*crypto.PrivateKeySECP256K1R{{keys[0]}}

	type test struct {
		name string
		outs []*avax.TransferableOutput
		ins  []*avax.TransferableInput
		err  error
	}

	tests := []test{
		{
			name: "Locked out",
			outs: []*avax.TransferableOutput{
				generateTestOut(vm.ctx.AVAXAssetID, LockStateBonded, 10, outputOwners),
			},
			ins: []*avax.TransferableInput{},
			err: errLockedInsOrOuts,
		},
		{
			name: "Locked in",
			outs: []*avax.TransferableOutput{},
			ins: []*avax.TransferableInput{
				generateTestIn(vm.ctx.AVAXAssetID, LockStateBonded, 10),
			},
			err: errLockedInsOrOuts,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utx := &UnsignedCreateChainTx{
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    vm.ctx.NetworkID,
					BlockchainID: vm.ctx.ChainID,
					Ins:          tt.ins,
					Outs:         tt.outs,
				}},
				SubnetID:   testSubnet1.ID(),
				VMID:       constants.AVMID,
				SubnetAuth: &secp256k1fx.Input{SigIndices: []uint32{0, 1}},
			}
			tx := &Tx{UnsignedTx: utx}

			err := tx.Sign(Codec, signers)
			assert.NoError(t, err)

			// Get the preferred block (which we want to build off)
			preferred, err := vm.Preferred()
			assert.NoError(t, err)

			preferredDecision, ok := preferred.(decision)
			assert.True(t, ok)

			preferredState := preferredDecision.onAccept()
			fakedState := newVersionedState(
				preferredState,
				preferredState.CurrentStakerChainState(),
				preferredState.PendingStakerChainState(),
				preferredState.LockedUTXOsChainState(),
			)

			// Testing execute
			_, err = utx.Execute(vm, fakedState, tx)
			assert.Equal(t, tt.err, err)
		})
	}
}
