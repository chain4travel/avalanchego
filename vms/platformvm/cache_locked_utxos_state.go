// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

type LockedUTXOsState interface {
	LockedUTXOsChainState() lockedUTXOsChainState
}
