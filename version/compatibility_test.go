// Copyright (C) 2022-2024, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************
// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCompatibility(t *testing.T) {
	v := &Application{
		Name:  Client,
		Major: 1,
		Minor: 4,
		Patch: 3,
	}
	minCompatable := &Application{
		Name:  Client,
		Major: 1,
		Minor: 4,
		Patch: 0,
	}
	minCompatableTime := time.Unix(9000, 0)
	prevMinCompatable := &Application{
		Name:  Client,
		Major: 1,
		Minor: 3,
		Patch: 0,
	}

	compatibility := NewCompatibility(
		v,
		minCompatable,
		minCompatableTime,
		prevMinCompatable,
	).(*compatibility)
	require.Equal(t, v, compatibility.Version())

	tests := []struct {
		peer        *Application
		time        time.Time
		expectedErr error
	}{
		{
			peer: &Application{
				Name:  LegacyAppName,
				Major: 1,
				Minor: 5,
				Patch: 0,
			},
			time: minCompatableTime,
		},
		{
			peer: &Application{
				Name:  Client,
				Major: 1,
				Minor: 5,
				Patch: 0,
			},
			time: minCompatableTime,
		},
		{
			peer: &Application{
				Name:  Client,
				Major: 1,
				Minor: 3,
				Patch: 5,
			},
			time: time.Unix(8500, 0),
		},
		{
			peer: &Application{
				Name:  Client,
				Major: 0,
				Minor: 1,
				Patch: 0,
			},
			time:        minCompatableTime,
			expectedErr: errIncompatible,
		},
		{
			peer: &Application{
				Name:  Client,
				Major: 1,
				Minor: 3,
				Patch: 5,
			},
			time:        minCompatableTime,
			expectedErr: errIncompatible,
		},
		{
			peer: &Application{
				Name:  Client,
				Major: 1,
				Minor: 2,
				Patch: 5,
			},
			time:        time.Unix(8500, 0),
			expectedErr: errIncompatible,
		},
		{
			peer: &Application{
				Name:  Client,
				Major: 1,
				Minor: 1,
				Patch: 5,
			},
			time:        time.Unix(7500, 0),
			expectedErr: errIncompatible,
		},
	}
	for _, test := range tests {
		peer := test.peer
		compatibility.clock.Set(test.time)
		t.Run(fmt.Sprintf("%s-%s", peer, test.time), func(t *testing.T) {
			err := compatibility.Compatible(peer)
			require.ErrorIs(t, err, test.expectedErr)
		})
	}
}
