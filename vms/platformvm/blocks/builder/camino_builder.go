// Copyright (C) 2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/timer"
	blockexecutor "github.com/ava-labs/avalanchego/vms/platformvm/blocks/executor"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	txBuilder "github.com/ava-labs/avalanchego/vms/platformvm/txs/builder"
	txexecutor "github.com/ava-labs/avalanchego/vms/platformvm/txs/executor"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/mempool"
)

type caminoBuilder struct {
	builder
	txBuilder txBuilder.CaminoBuilder
}

func CaminoNew(
	mempool mempool.Mempool,
	txBuilder txBuilder.CaminoBuilder,
	txExecutorBackend *txexecutor.Backend,
	blkManager blockexecutor.Manager,
	toEngine chan<- common.Message,
	appSender common.AppSender,
) Builder {
	builder := &caminoBuilder{
		builder: builder{
			Mempool:           mempool,
			txExecutorBackend: txExecutorBackend,
			blkManager:        blkManager,
			toEngine:          toEngine,
			txBuilder:         txBuilder,
		},
		txBuilder: txBuilder,
	}

	builder.timer = timer.NewTimer(builder.setNextBuildBlockTime)

	builder.Network = NewCaminoNetwork(
		txExecutorBackend.Ctx,
		builder,
		appSender,
		builder.txBuilder,
	)

	go txExecutorBackend.Ctx.Log.RecoverAndPanic(builder.timer.Dispatch)
	return builder
}

func getNextPendingStakerToRemove(
	chainTimestamp time.Time,
	shouldRewardNextCurrentStaker bool,
	nextCurrentStaker *state.Staker,
	preferredState state.Chain,
) (ids.ID, bool, error) {
	pendingStakerIterator, err := preferredState.GetPendingStakerIterator()
	if err != nil {
		return ids.Empty, false, err
	}
	defer pendingStakerIterator.Release()

	if pendingStakerIterator.Next() {
		pendingStaker := pendingStakerIterator.Value()
		if shouldRewardNextCurrentStaker && !nextCurrentStaker.EndTime.After(pendingStaker.EndTime) {
			return nextCurrentStaker.TxID, shouldRewardNextCurrentStaker, nil
		}
		return pendingStaker.TxID, chainTimestamp.Equal(pendingStaker.EndTime), nil
	}

	return nextCurrentStaker.TxID, shouldRewardNextCurrentStaker, nil
}
