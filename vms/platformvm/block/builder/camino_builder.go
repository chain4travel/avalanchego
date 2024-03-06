// Copyright (C) 2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/timer"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	blockexecutor "github.com/ava-labs/avalanchego/vms/platformvm/block/executor"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	txBuilder "github.com/ava-labs/avalanchego/vms/platformvm/txs/builder"
	txexecutor "github.com/ava-labs/avalanchego/vms/platformvm/txs/executor"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/mempool"
)

// Overriding axax block builder methods with caminoBuilder methods
// must be done with consideration, that network uses reference to avax builder,
// not to camino builder. So it will actually call avax builder methods.

type caminoBuilder struct {
	builder
	caminoTxBuilder txBuilder.CaminoBuilder
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
		caminoTxBuilder: txBuilder,
	}

	builder.timer = timer.NewTimer(builder.setNextBuildBlockTime)

	builder.Network = NewCaminoNetwork(
		txExecutorBackend.Ctx,
		builder,
		appSender,
		builder.caminoTxBuilder,
	)

	go txExecutorBackend.Ctx.Log.RecoverAndPanic(builder.timer.Dispatch)
	return builder
}

func caminoBuildBlock(
	builder *builder,
	parentID ids.ID,
	height uint64,
	timestamp time.Time,
	parentState state.Chain,
) (block.Block, error) {
	txBuilder, ok := builder.txBuilder.(txBuilder.CaminoBuilder)
	if !ok {
		// if its not caminoBuilder, than its not our camino-node
		// there will be no deposits and we don't need to process camino-specific logic
		return nil, nil
	}

	// Ulocking expired deposits
	depositsTxIDs, shouldUnlock, err := getNextDepositsToUnlock(parentState, timestamp)
	if err != nil {
		return nil, fmt.Errorf("could not find next deposits to unlock: %w", err)
	}
	if shouldUnlock {
		unlockDepositTx, err := txBuilder.NewSystemUnlockDepositTx(depositsTxIDs)
		if err != nil {
			return nil, fmt.Errorf("could not build tx to unlock deposits: %w", err)
		}

		return block.NewBanffStandardBlock(
			timestamp,
			parentID,
			height,
			[]*txs.Tx{unlockDepositTx},
		)
	}

	// Finishing expired and early finished proposals
	expiredProposalIDs, err := getExpiredProposals(parentState, timestamp)
	if err != nil {
		return nil, fmt.Errorf("could not find expired proposals: %w", err)
	}
	earlyFinishedProposalIDs, err := parentState.GetProposalIDsToFinish()
	if err != nil {
		return nil, fmt.Errorf("could not find successful proposals: %w", err)
	}
	if len(expiredProposalIDs) > 0 || len(earlyFinishedProposalIDs) > 0 {
		finishProposalsTx, err := txBuilder.FinishProposalsTx(parentState, earlyFinishedProposalIDs, expiredProposalIDs)
		if err != nil {
			return nil, fmt.Errorf("could not build tx to finish proposals: %w", err)
		}

		// FinishProposalsTx should never be in block with addVoteTx,
		// because it can affect state of proposals.
		return block.NewBanffStandardBlock(
			timestamp,
			parentID,
			height,
			[]*txs.Tx{finishProposalsTx},
		)
	}

	return nil, nil
}

func getNextDeferredStakerToRemove(
	chainTimestamp time.Time,
	shouldRewardNextCurrentStaker bool,
	nextCurrentStaker *state.Staker,
	preferredState state.Chain,
) (ids.ID, bool, error) {
	deferredStakerIterator, err := preferredState.GetDeferredStakerIterator()
	if err != nil {
		return ids.Empty, false, err
	}
	defer deferredStakerIterator.Release()

	if deferredStakerIterator.Next() {
		deferredStaker := deferredStakerIterator.Value()
		if shouldRewardNextCurrentStaker && !nextCurrentStaker.EndTime.After(deferredStaker.EndTime) {
			return nextCurrentStaker.TxID, shouldRewardNextCurrentStaker, nil
		}
		return deferredStaker.TxID, chainTimestamp.Equal(deferredStaker.EndTime), nil
	}

	return nextCurrentStaker.TxID, shouldRewardNextCurrentStaker, nil
}

func getNextDepositsToUnlock(
	preferredState state.Chain,
	chainTime time.Time,
) ([]ids.ID, bool, error) {
	if !chainTime.Before(mockable.MaxTime) {
		return nil, false, ErrEndOfTime
	}

	nextDeposits, nextDepositsEndtime, err := preferredState.GetNextToUnlockDepositIDsAndTime(nil)
	if err == database.ErrNotFound {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return nextDeposits, nextDepositsEndtime.Equal(chainTime), nil
}

func getExpiredProposals(
	preferredState state.Chain,
	chainTime time.Time,
) ([]ids.ID, error) {
	if !chainTime.Before(mockable.MaxTime) {
		return nil, ErrEndOfTime
	}

	nextProposals, nextProposalsEndtime, err := preferredState.GetNextToExpireProposalIDsAndTime(nil)
	if err == database.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if nextProposalsEndtime.Equal(chainTime) {
		return nextProposals, nil
	}

	return nil, nil
}