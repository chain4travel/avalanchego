// Copyright (C) 2022-2024, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/snowtest"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
	"github.com/ava-labs/avalanchego/vms/platformvm/test/generate"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

func TestDepositTxSyntacticVerify(t *testing.T) {
	ctx := snowtest.Context(t, snowtest.PChainID)
	owner1 := secp256k1fx.OutputOwners{Threshold: 1, Addrs: []ids.ShortID{{1}}}

	tests := map[string]struct {
		tx          *DepositTx
		expectedErr error
	}{
		"Nil tx": {
			expectedErr: ErrNilTx,
		},
		"Bad reward owner": {
			tx: &DepositTx{
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
				}},
				RewardsOwner: (*secp256k1fx.OutputOwners)(nil),
			},
			expectedErr: errInvalidRewardOwner,
		},
		"To big total deposit amount": {
			tx: &DepositTx{
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
					Outs: []*avax.TransferableOutput{
						generate.Out(ctx.AVAXAssetID, math.MaxUint64, owner1, locked.ThisTxID, ids.Empty),
						generate.Out(ctx.AVAXAssetID, math.MaxUint64, owner1, locked.ThisTxID, ids.Empty),
					},
				}},
				RewardsOwner: &secp256k1fx.OutputOwners{},
			},
			expectedErr: errTooBigDeposit,
		},
		"V1, bad deposit creator auth": {
			tx: &DepositTx{
				UpgradeVersionID: codec.UpgradeVersion1,
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
				}},
				RewardsOwner:       &secp256k1fx.OutputOwners{},
				DepositCreatorAuth: (*secp256k1fx.Input)(nil),
			},
			expectedErr: errBadDepositCreatorAuth,
		},
		"V1, bad deposit offer owner auth": {
			tx: &DepositTx{
				UpgradeVersionID: codec.UpgradeVersion1,
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
				}},
				RewardsOwner:          &secp256k1fx.OutputOwners{},
				DepositCreatorAuth:    &secp256k1fx.Input{},
				DepositOfferOwnerAuth: (*secp256k1fx.Input)(nil),
			},
			expectedErr: errBadOfferOwnerAuth,
		},
		"OK: v0": {
			tx: &DepositTx{
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
				}},
				RewardsOwner: &secp256k1fx.OutputOwners{},
			},
		},
		"OK: v1": {
			tx: &DepositTx{
				UpgradeVersionID: codec.UpgradeVersion1,
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
				}},
				RewardsOwner:          &secp256k1fx.OutputOwners{},
				DepositCreatorAddress: ids.ShortID{1},
				DepositCreatorAuth:    &secp256k1fx.Input{},
				DepositOfferOwnerAuth: &secp256k1fx.Input{},
			},
		},
		"OK: v1, empty creator addr": {
			tx: &DepositTx{
				UpgradeVersionID: codec.UpgradeVersion1,
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
				}},
				RewardsOwner:          &secp256k1fx.OutputOwners{},
				DepositCreatorAuth:    &secp256k1fx.Input{},
				DepositOfferOwnerAuth: &secp256k1fx.Input{},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.tx.SyntacticVerify(ctx)
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}
