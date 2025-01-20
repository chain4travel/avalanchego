// Copyright (C) 2022-2025, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************
// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"crypto"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/staking"
)

func TestBuild(t *testing.T) {
	require := require.New(t)

	parentID := ids.ID{1}
	timestamp := time.Unix(123, 0)
	pChainHeight := uint64(2)
	innerBlockBytes := []byte{3}
	chainID := ids.ID{4}

	tlsCert, err := staking.NewTLSCert()
	require.NoError(err)

	cert := tlsCert.Leaf
	key := tlsCert.PrivateKey.(crypto.Signer)

	builtBlock, err := Build(
		parentID,
		timestamp,
		pChainHeight,
		ids.EmptyNodeID,
		cert,
		innerBlockBytes,
		chainID,
		key,
	)
	require.NoError(err)

	require.Equal(parentID, builtBlock.ParentID())
	require.Equal(pChainHeight, builtBlock.PChainHeight())
	require.Equal(timestamp, builtBlock.Timestamp())
	require.Equal(innerBlockBytes, builtBlock.Block())

	err = builtBlock.Verify(true, chainID)
	require.NoError(err)

	err = builtBlock.Verify(false, chainID)
	require.Error(err)
}

func TestBuildUnsigned(t *testing.T) {
	parentID := ids.ID{1}
	timestamp := time.Unix(123, 0)
	pChainHeight := uint64(2)
	innerBlockBytes := []byte{3}

	require := require.New(t)

	builtBlock, err := BuildUnsigned(parentID, timestamp, pChainHeight, innerBlockBytes)
	require.NoError(err)

	require.Equal(parentID, builtBlock.ParentID())
	require.Equal(pChainHeight, builtBlock.PChainHeight())
	require.Equal(timestamp, builtBlock.Timestamp())
	require.Equal(innerBlockBytes, builtBlock.Block())
	require.Equal(ids.EmptyNodeID, builtBlock.Proposer())

	err = builtBlock.Verify(false, ids.Empty)
	require.NoError(err)

	err = builtBlock.Verify(true, ids.Empty)
	require.Error(err)
}

func TestBuildHeader(t *testing.T) {
	require := require.New(t)

	chainID := ids.ID{1}
	parentID := ids.ID{2}
	bodyID := ids.ID{3}

	builtHeader, err := BuildHeader(
		chainID,
		parentID,
		bodyID,
	)
	require.NoError(err)

	require.Equal(chainID, builtHeader.ChainID())
	require.Equal(parentID, builtHeader.ParentID())
	require.Equal(bodyID, builtHeader.BodyID())
}

func TestBuildOption(t *testing.T) {
	require := require.New(t)

	parentID := ids.ID{1}
	innerBlockBytes := []byte{3}

	builtOption, err := BuildOption(parentID, innerBlockBytes)
	require.NoError(err)

	require.Equal(parentID, builtOption.ParentID())
	require.Equal(innerBlockBytes, builtOption.Block())
}
