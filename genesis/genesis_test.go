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
// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/utils/perms"
	"github.com/ava-labs/avalanchego/vms/platformvm/genesis"
)

var (
	//go:embed genesis_test.json
	customGenesisConfigJSON  []byte
	invalidGenesisConfigJSON = []byte(`{
		"networkID": 9999}}}}
	}`)
)

func TestValidateConfig(t *testing.T) {
	tests := map[string]struct {
		networkID uint32
		config    *Config
		err       string
	}{
		"camino": {
			networkID: 1000,
			config:    &CaminoConfig,
		},
		"columbus": {
			networkID: 1001,
			config:    &ColumbusConfig,
		},
		"kopernikus": {
			networkID: 1002,
			config:    &KopernikusConfig,
		},
		"local": {
			networkID: 12345,
			config:    &LocalConfig,
		},
		"camino (networkID mismatch)": {
			networkID: 999,
			config:    &CaminoConfig,
			err:       "networkID 999 specified but genesis config contains networkID 1000",
		},
		"invalid start time": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.StartTime = 999999999999999
				return &thisConfig
			}(),
			err: "start time cannot be in the future",
		},
		"no initial supply": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.Allocations = []Allocation{}
				return &thisConfig
			}(),
			err: "initial supply must be > 0",
		},
		"no initial stakers": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakers = []Staker{}
				return &thisConfig
			}(),
			err: "initial stakers must be > 0",
		},
		"invalid initial stake duration": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakeDuration = 0
				return &thisConfig
			}(),
			err: "initial stake duration must be > 0",
		},
		"invalid stake offset": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakeDurationOffset = 100000000
				return &thisConfig
			}(),
			err: "initial stake duration is 31536000 but need at least 400000000 with offset of 100000000",
		},
		"empty initial staked funds": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakedFunds = []ids.ShortID(nil)
				return &thisConfig
			}(),
			err: "initial staked funds cannot be empty",
		},
		"duplicate initial staked funds": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakedFunds = append(thisConfig.InitialStakedFunds, thisConfig.InitialStakedFunds[0])
				return &thisConfig
			}(),
			err: "duplicated in initial staked funds",
		},
		"initial staked funds not in allocations": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakedFunds = append(thisConfig.InitialStakedFunds, ids.GenerateTestShortID())
				return &thisConfig
			}(),
			err: "does not have an allocation to stake",
		},
		"empty C-Chain genesis": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.CChainGenesis = ""
				return &thisConfig
			}(),
			err: "C-Chain genesis cannot be empty",
		},
		"empty message": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.Message = ""
				return &thisConfig
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require := require.New(t)

			err := validateConfig(test.networkID, test.config)
			if len(test.err) > 0 {
				require.Error(err)
				require.Contains(err.Error(), test.err)
				return
			}
			require.NoError(err)
		})
	}
}

func TestGenesisFromFile(t *testing.T) {
	tests := map[string]struct {
		networkID       uint32
		customConfig    []byte
		missingFilepath string
		err             string
		expected        string
	}{
		"camino": {
			networkID:    constants.CaminoID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network camino (1000)",
		},
		"columbus": {
			networkID:    constants.ColumbusID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network columbus (1001)",
		},
		"columbus (with custom specified)": {
			networkID:    constants.ColumbusID,
			customConfig: localGenesisConfigJSON, // won't load
			err:          "cannot override genesis config for standard network columbus (1001)",
		},
		"local": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"local (with custom specified)": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"custom": {
			networkID:    9999,
			customConfig: customGenesisConfigJSON,
			expected:     "04068ee6f095f09017f3d914f4f1a25c04d7c1a1518f3db48f97040a7d9f005c",
		},
		"custom (networkID mismatch)": {
			networkID:    9999,
			customConfig: localGenesisConfigJSON,
			err:          "networkID 9999 specified but genesis config contains networkID 12345",
		},
		"custom (invalid format)": {
			networkID:    9999,
			customConfig: invalidGenesisConfigJSON,
			err:          "unable to load provided genesis config",
		},
		"custom (missing filepath)": {
			networkID:       9999,
			missingFilepath: "missing.json",
			err:             "unable to load provided genesis config",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// test loading of genesis from file

			require := require.New(t)
			var customFile string
			if len(test.customConfig) > 0 {
				customFile = filepath.Join(t.TempDir(), "config.json")
				require.NoError(perms.WriteFile(customFile, test.customConfig, perms.ReadWrite))
			}

			if len(test.missingFilepath) > 0 {
				customFile = test.missingFilepath
			}

			genesisBytes, _, err := FromFile(test.networkID, customFile)
			if len(test.err) > 0 {
				require.Error(err)
				require.Contains(err.Error(), test.err)
				return
			}
			require.NoError(err)

			genesisHash := fmt.Sprintf("%x", hashing.ComputeHash256(genesisBytes))
			require.Equal(test.expected, genesisHash, "genesis hash mismatch")

			_, err = genesis.Parse(genesisBytes)
			require.NoError(err)
		})
	}
}

func TestGenesisFromFlag(t *testing.T) {
	tests := map[string]struct {
		networkID    uint32
		customConfig []byte
		err          string
		expected     string
	}{
		"camino": {
			networkID: constants.CaminoID,
			err:       "cannot override genesis config for standard network camino (1000)",
		},
		"columbus": {
			networkID: constants.ColumbusID,
			err:       "cannot override genesis config for standard network columbus (1001)",
		},
		"kopernikus": {
			networkID: constants.KopernikusID,
			err:       "cannot override genesis config for standard network kopernikus (1002)",
		},
		"local": {
			networkID: constants.LocalID,
			err:       "cannot override genesis config for standard network local (12345)",
		},
		"local (with custom specified)": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"custom": {
			networkID:    9999,
			customConfig: customGenesisConfigJSON,
			expected:     "04068ee6f095f09017f3d914f4f1a25c04d7c1a1518f3db48f97040a7d9f005c",
		},
		"custom (networkID mismatch)": {
			networkID:    9999,
			customConfig: localGenesisConfigJSON,
			err:          "networkID 9999 specified but genesis config contains networkID 12345",
		},
		"custom (invalid format)": {
			networkID:    9999,
			customConfig: invalidGenesisConfigJSON,
			err:          "unable to load genesis content from flag",
		},
		"custom (missing content)": {
			networkID: 9999,
			err:       "unable to load genesis content from flag",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// test loading of genesis content from flag/env-var

			require := require.New(t)
			var genBytes []byte
			if len(test.customConfig) == 0 {
				// try loading a default config
				var err error
				switch test.networkID {
				case constants.CaminoID:
					genBytes, err = json.Marshal(&CaminoConfig)
					require.NoError(err)
				case constants.TestnetID:
					genBytes, err = json.Marshal(&ColumbusConfig)
					require.NoError(err)
				case constants.KopernikusID:
					genBytes, err = json.Marshal(&KopernikusConfig)
					require.NoError(err)
				case constants.LocalID:
					genBytes, err = json.Marshal(&LocalConfig)
					require.NoError(err)
				default:
					genBytes = make([]byte, 0)
				}
			} else {
				genBytes = test.customConfig
			}
			content := base64.StdEncoding.EncodeToString(genBytes)

			genesisBytes, _, err := FromFlag(test.networkID, content)
			if len(test.err) > 0 {
				require.Error(err)
				require.Contains(err.Error(), test.err)
				return
			}
			require.NoError(err)

			genesisHash := fmt.Sprintf("%x", hashing.ComputeHash256(genesisBytes))
			require.Equal(test.expected, genesisHash, "genesis hash mismatch")

			_, err = genesis.Parse(genesisBytes)
			require.NoError(err)
		})
	}
}

func TestGenesis(t *testing.T) {
	tests := []struct {
		networkID  uint32
		expectedID string
	}{
		{
			networkID:  constants.MainnetID,
			expectedID: "2eP2teFA41Pb61nQyQYJY38onkCqkrJDPSVeQz8v9ahbEma18S",
		},
		{
			networkID:  constants.FujiID,
			expectedID: "2fAKF9ph6o4jxr12QVWmbww8dVfkNKepBt547QFEJ5eyD9x2Z9",
		},
		{
			networkID:  constants.CaminoID,
			expectedID: "VLRhFeu2tE4ZJDPoAMXE1XMA1P1P5HCY2gtyJJvyVQh19EWTM",
		},
		{
			networkID:  constants.ColumbusID,
			expectedID: "2DWrwmjXoJ9i9pVVf7AZfCv1mFk6LtngrKU71uQU15YUCcZEQ9",
		},
		{
			networkID:  constants.KopernikusID,
			expectedID: "2qmJHpvmxmqY2vaLbHTTT5bLGzvJxETgkmSe8SzZFS35TDoMdM",
		},
		{
			networkID:  constants.LocalID,
			expectedID: "294HrmVEniYX2mrvLjww9AmpV9NZnFhEi6eRrCsAzxZ2PkvdVb",
		},
	}
	for _, test := range tests {
		t.Run(constants.NetworkIDToNetworkName[test.networkID], func(t *testing.T) {
			require := require.New(t)

			config := GetConfig(test.networkID)
			genesisBytes, _, err := FromConfig(config)
			require.NoError(err)

			var genesisID ids.ID = hashing.ComputeHash256Array(genesisBytes)
			require.Equal(test.expectedID, genesisID.String())
		})
	}
}

func TestVMGenesis(t *testing.T) {
	type vmTest struct {
		vmID       ids.ID
		expectedID string
	}
	tests := []struct {
		networkID uint32
		vmTest    []vmTest
	}{
		{
			networkID: constants.CaminoID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "yMQo4UEa2Gkk6aSmifkUuBsystV1iu1NppatvoYz6yCDnRjiq",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "RinAZCjd5Dm4wk1FBWiXiiSW2VZkjzgNyR7nNBRkuCvG9zRkJ",
				},
			},
		},
		{
			networkID: constants.ColumbusID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "sKRk9Q6u3NEVdGtSvsMhjoGdzc7AnRZ2wP5YcoJL4xJB7UX4J",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2avDgyQpecb3vQvMpUvJv75Hphn6o3Ms79FD4Rjqf7jJ1ZY7Qe",
				},
			},
		},
		{
			networkID: constants.KopernikusID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "2hQXfgx6aeM7ZPthZJTdtUK6rG1uhdvfYqFhBw7ajziYMM7toX",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2oBwpcTG6ZLViVnigvZ2MLgmz4mL2JYGR89ymRnUM6wNpN8cvq",
				},
			},
		},
		{
			networkID: constants.LocalID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "2eNy1mUFdmaxXNj1eQHUe7Np4gju9sJsEtWQ4MX3ToiNKuADed",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2CA6j5zYzasynPsFeNoqWkmTCt3VScMvXUZHbfDJ8k3oGzAPtU",
				},
			},
		},
	}

	for _, test := range tests {
		for _, vmTest := range test.vmTest {
			name := fmt.Sprintf("%s-%s",
				constants.NetworkIDToNetworkName[test.networkID],
				vmTest.vmID,
			)
			t.Run(name, func(t *testing.T) {
				require := require.New(t)

				config := GetConfig(test.networkID)
				genesisBytes, _, err := FromConfig(config)
				require.NoError(err)

				genesisTx, err := VMGenesis(genesisBytes, vmTest.vmID)
				require.NoError(err)

				require.Equal(
					vmTest.expectedID,
					genesisTx.ID().String(),
					"%s genesisID with networkID %d mismatch",
					vmTest.vmID,
					test.networkID,
				)
			})
		}
	}
}

func TestAVAXAssetID(t *testing.T) {
	tests := []struct {
		networkID  uint32
		expectedID string
	}{
		{
			networkID:  constants.CaminoID,
			expectedID: "2LgYQ5nWZgiwYFVpSoDPma5e3A4GsehYsbCge9eH8Z1149Ca5b",
		},
		{
			networkID:  constants.ColumbusID,
			expectedID: "21aEnvVjGVrjiHkv1QrTvU7F19kXCZUPbgFpe8HrSsXMTpBPP6",
		},
		{
			networkID:  constants.KopernikusID,
			expectedID: "227CxttTm32vqdUFsYjysNdXdrFahsmwqCwjjFteuQpmsEeCbv",
		},
		{
			networkID:  constants.LocalID,
			expectedID: "2fombhL7aGPwj3KH4bfrmJwW6PVnMobf9Y2fn9GwxiAAJyFDbe",
		},
	}

	for _, test := range tests {
		t.Run(constants.NetworkIDToNetworkName[test.networkID], func(t *testing.T) {
			require := require.New(t)

			config := GetConfig(test.networkID)
			_, avaxAssetID, err := FromConfig(config)
			require.NoError(err)

			require.Equal(
				test.expectedID,
				avaxAssetID.String(),
				"AVAX assetID with networkID %d mismatch",
				test.networkID,
			)
		})
	}
}
