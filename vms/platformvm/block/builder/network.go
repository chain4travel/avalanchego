// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// TODO: consider moving the network implementation to a separate package

package builder

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/components/message"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

// We allow [recentCacheSize] to be fairly large because we only store hashes
// in the cache, not entire transactions.
const recentCacheSize = 512

var _ Network = (*network)(nil)

type Network interface {
	common.AppHandler

	// GossipTx gossips the transaction to some of the connected peers
	GossipTx(ctx context.Context, tx *txs.Tx) error
}

type network struct {
	// We embed a noop handler for all unhandled messages
	common.AppHandler

	ctx        *snow.Context
	blkBuilder *builder
	appSender  common.AppSender

	// gossip related attributes
	recentTxsLock sync.Mutex
	recentTxs     *cache.LRU[ids.ID, struct{}]
}

func NewNetwork(
	ctx *snow.Context,
	blkBuilder *builder,
	appSender common.AppSender,
) Network {
	return &network{
		AppHandler: common.NewNoOpAppHandler(ctx.Log),

		ctx:        ctx,
		blkBuilder: blkBuilder,
		appSender:  appSender,
		recentTxs:  &cache.LRU[ids.ID, struct{}]{Size: recentCacheSize},
	}
}

func (n *network) AppGossip(_ context.Context, nodeID ids.NodeID, msgBytes []byte) error {
	n.ctx.Log.Debug("called AppGossip message handler",
		zap.Stringer("nodeID", nodeID),
		zap.Int("messageLen", len(msgBytes)),
	)

	if n.blkBuilder.txExecutorBackend.Config.PartialSyncPrimaryNetwork {
		n.ctx.Log.Debug("dropping AppGossip message",
			zap.String("reason", "primary network is not being fully synced"),
		)
		return nil
	}

	msgIntf, err := message.Parse(msgBytes)
	if err != nil {
		n.ctx.Log.Debug("dropping AppGossip message",
			zap.String("reason", "failed to parse message"),
		)
		return nil
	}

	msg, ok := msgIntf.(*message.Tx)
	if !ok {
		n.ctx.Log.Debug("dropping unexpected message",
			zap.Stringer("nodeID", nodeID),
		)
		return nil
	}

	tx, err := txs.Parse(txs.Codec, msg.Tx)
	if err != nil {
		n.ctx.Log.Verbo("received invalid tx",
			zap.Stringer("nodeID", nodeID),
			zap.Binary("tx", msg.Tx),
			zap.Error(err),
		)
		return nil
	}

	txID := tx.ID()

	// We need to grab the context lock here to avoid racy behavior with
	// transaction verification + mempool modifications.
	n.ctx.Lock.Lock()
	defer n.ctx.Lock.Unlock()

	if reason := n.blkBuilder.GetDropReason(txID); reason != nil {
		// If the tx is being dropped - just ignore it
		return nil
	}

	// add to mempool
	if err := n.blkBuilder.AddUnverifiedTx(tx); err != nil {
		n.ctx.Log.Debug("tx failed verification",
			zap.Stringer("nodeID", nodeID),
			zap.Error(err),
		)
	}
	return nil
}

func (n *network) GossipTx(ctx context.Context, tx *txs.Tx) error {
	txBytes := tx.Bytes()
	msg := &message.Tx{
		Tx: txBytes,
	}
	msgBytes, err := message.Build(msg)
	if err != nil {
		return err
	}

	txID := tx.ID()
	n.gossipTx(ctx, txID, msgBytes)
	return nil
}

func (n *network) gossipTx(ctx context.Context, txID ids.ID, msgBytes []byte) {
	n.recentTxsLock.Lock()
	_, has := n.recentTxs.Get(txID)
	n.recentTxs.Put(txID, struct{}{})
	n.recentTxsLock.Unlock()

	// Don't gossip a transaction if it has been recently gossiped.
	if has {
		return
	}

	n.ctx.Log.Debug("gossiping tx",
		zap.Stringer("txID", txID),
	)

	if err := n.appSender.SendAppGossip(ctx, msgBytes); err != nil {
		n.ctx.Log.Error("failed to gossip tx",
			zap.Stringer("txID", txID),
			zap.Error(err),
		)
	}
}