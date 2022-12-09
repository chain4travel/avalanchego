// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/keystore"
	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"go.uber.org/zap"

	utilsjson "github.com/ava-labs/avalanchego/utils/json"
)

var (
	errSerializeTx       = "couldn't serialize TX: %w"
	errEncodeTx          = "couldn't encode TX as string: %w"
	errInvalidChangeAddr = "couldn't parse changeAddr: %w"
	errCreateTx          = "couldn't create tx: %w"
)

// CaminoService defines the API calls that can be made to the platform chain
type CaminoService struct {
	Service
}

type GetBalanceResponseV2 struct {
	Balances               map[ids.ID]utilsjson.Uint64 `json:"balances"`
	UnlockedOutputs        map[ids.ID]utilsjson.Uint64 `json:"unlockedOutputs"`
	BondedOutputs          map[ids.ID]utilsjson.Uint64 `json:"bondedOutputs"`
	DepositedOutputs       map[ids.ID]utilsjson.Uint64 `json:"depositedOutputs"`
	DepositedBondedOutputs map[ids.ID]utilsjson.Uint64 `json:"bondedDepositedOutputs"`
}
type GetBalanceResponseWrapper struct {
	LockModeBondDeposit bool
	GetBalanceResponse
	GetBalanceResponseV2 //nolint:govet
}

// GetConfigurationReply is the response from calling GetConfiguration.
type GetConfigurationReply struct {
	// The NetworkID
	NetworkID utilsjson.Uint32 `json:"networkID"`
	// The fee asset ID
	AssetID ids.ID `json:"assetID"`
	// The symbol of the fee asset ID
	AssetSymbol string `json:"assetSymbol"`
	// beech32HRP use in addresses
	Hrp string `json:"hrp"`
	// Primary network blockchains
	Blockchains []APIBlockchain `json:"blockchains"`
	// The minimum duration a validator has to stake
	MinStakeDuration utilsjson.Uint64 `json:"minStakeDuration"`
	// The maximum duration a validator can stake
	MaxStakeDuration utilsjson.Uint64 `json:"maxStakeDuration"`
	// The minimum amount of tokens one must bond to be a validator
	MinValidatorStake utilsjson.Uint64 `json:"minValidatorStake"`
	// The maximum amount of tokens bondable to a validator
	MaxValidatorStake utilsjson.Uint64 `json:"maxValidatorStake"`
	// The minimum delegation fee
	MinDelegationFee utilsjson.Uint32 `json:"minDelegationFee"`
	// Minimum stake, in nAVAX, that can be delegated on the primary network
	MinDelegatorStake utilsjson.Uint64 `json:"minDelegatorStake"`
	// The minimum consumption rate
	MinConsumptionRate utilsjson.Uint64 `json:"minConsumptionRate"`
	// The maximum consumption rate
	MaxConsumptionRate utilsjson.Uint64 `json:"maxConsumptionRate"`
	// The supply cap for the native token (AVAX)
	SupplyCap utilsjson.Uint64 `json:"supplyCap"`
	// The codec version used for serializing
	CodecVersion utilsjson.Uint16 `json:"codecVersion"`
}

func (response GetBalanceResponseWrapper) MarshalJSON() ([]byte, error) {
	if !response.LockModeBondDeposit {
		return json.Marshal(response.GetBalanceResponse)
	}
	return json.Marshal(response.GetBalanceResponseV2)
}

// GetBalance gets the balance of an address
func (service *CaminoService) GetBalance(_ *http.Request, args *GetBalanceRequest, response *GetBalanceResponseWrapper) error {
	caminoGenesis, err := service.vm.state.CaminoGenesisState()
	if err != nil {
		return err
	}
	response.LockModeBondDeposit = caminoGenesis.LockModeBondDeposit
	if !caminoGenesis.LockModeBondDeposit {
		return service.Service.GetBalance(nil, args, &response.GetBalanceResponse)
	}

	if args.Address != nil {
		args.Addresses = append(args.Addresses, *args.Address)
	}

	service.vm.ctx.Log.Debug("Platform: GetBalance called",
		logging.UserStrings("addresses", args.Addresses),
	)

	// Parse to address
	addrs, err := avax.ParseServiceAddresses(service.addrManager, args.Addresses)
	if err != nil {
		return err
	}

	utxos, err := avax.GetAllUTXOs(service.vm.state, addrs)
	if err != nil {
		return fmt.Errorf("couldn't get UTXO set of %v: %w", args.Addresses, err)
	}

	unlockedOutputs := map[ids.ID]utilsjson.Uint64{}
	bondedOutputs := map[ids.ID]utilsjson.Uint64{}
	depositedOutputs := map[ids.ID]utilsjson.Uint64{}
	depositedBondedOutputs := map[ids.ID]utilsjson.Uint64{}
	balances := map[ids.ID]utilsjson.Uint64{}

utxoFor:
	for _, utxo := range utxos {
		assetID := utxo.AssetID()
		switch out := utxo.Out.(type) {
		case *secp256k1fx.TransferOutput:
			unlockedOutputs[assetID] = utilsjson.SafeAdd(unlockedOutputs[assetID], utilsjson.Uint64(out.Amount()))
			balances[assetID] = utilsjson.SafeAdd(balances[assetID], utilsjson.Uint64(out.Amount()))
		case *locked.Out:
			switch out.LockState() {
			case locked.StateBonded:
				bondedOutputs[assetID] = utilsjson.SafeAdd(bondedOutputs[assetID], utilsjson.Uint64(out.Amount()))
				balances[assetID] = utilsjson.SafeAdd(balances[assetID], utilsjson.Uint64(out.Amount()))
			case locked.StateDeposited:
				depositedOutputs[assetID] = utilsjson.SafeAdd(depositedOutputs[assetID], utilsjson.Uint64(out.Amount()))
				balances[assetID] = utilsjson.SafeAdd(balances[assetID], utilsjson.Uint64(out.Amount()))
			case locked.StateDepositedBonded:
				depositedBondedOutputs[assetID] = utilsjson.SafeAdd(depositedBondedOutputs[assetID], utilsjson.Uint64(out.Amount()))
				balances[assetID] = utilsjson.SafeAdd(balances[assetID], utilsjson.Uint64(out.Amount()))
			default:
				service.vm.ctx.Log.Warn("Unexpected utxo lock state")
				continue utxoFor
			}
		default:
			service.vm.ctx.Log.Warn("unexpected output type in UTXO",
				zap.String("type", fmt.Sprintf("%T", out)),
			)
			continue utxoFor
		}

		response.UTXOIDs = append(response.UTXOIDs, &utxo.UTXOID)
	}

	response.GetBalanceResponseV2 = GetBalanceResponseV2{balances, unlockedOutputs, bondedOutputs, depositedOutputs, depositedBondedOutputs}
	return nil
}

// GetMinStake returns the minimum staking amount in nAVAX.
func (service *Service) GetConfiguration(_ *http.Request, _ *struct{}, reply *GetConfigurationReply) error {
	service.vm.ctx.Log.Debug("Platform: GetConfiguration called")

	// Fee Asset ID, NetworkID and HRP
	reply.NetworkID = utilsjson.Uint32(service.vm.ctx.NetworkID)
	reply.AssetID = service.vm.GetFeeAssetID()
	reply.AssetSymbol = constants.TokenSymbol(service.vm.ctx.NetworkID)
	reply.Hrp = constants.GetHRP(service.vm.ctx.NetworkID)

	// Blockchains of the primary network
	blockchains := &GetBlockchainsResponse{}
	if err := service.appendBlockchains(constants.PrimaryNetworkID, blockchains); err != nil {
		return err
	}
	reply.Blockchains = blockchains.Blockchains

	// Staking information
	reply.MinStakeDuration = utilsjson.Uint64(service.vm.MinStakeDuration)
	reply.MaxStakeDuration = utilsjson.Uint64(service.vm.MaxStakeDuration)

	reply.MaxValidatorStake = utilsjson.Uint64(service.vm.MaxValidatorStake)
	reply.MinValidatorStake = utilsjson.Uint64(service.vm.MinValidatorStake)

	reply.MinDelegationFee = utilsjson.Uint32(service.vm.MinDelegationFee)
	reply.MinDelegatorStake = utilsjson.Uint64(service.vm.MinDelegatorStake)

	reply.MinConsumptionRate = utilsjson.Uint64(service.vm.RewardConfig.MinConsumptionRate)
	reply.MaxConsumptionRate = utilsjson.Uint64(service.vm.RewardConfig.MaxConsumptionRate)

	reply.SupplyCap = utilsjson.Uint64(service.vm.RewardConfig.SupplyCap)

	// Codec information
	reply.CodecVersion = utilsjson.Uint16(txs.Version)

	return nil
}

type SetAddressStateArgs struct {
	api.JSONSpendHeader

	Address string `json:"address"`
	State   uint8  `json:"state"`
	Remove  bool   `json:"remove"`
}

// AddAdressState issues an AddAdressStateTx
func (service *Service) SetAddressState(_ *http.Request, args *SetAddressStateArgs, response api.JSONTxID) error {
	service.vm.ctx.Log.Debug("Platform: SetAddressState called")

	keys, err := service.getKeystoreKeys(&args.JSONSpendHeader)
	if err != nil {
		return err
	}

	tx, err := service.buildAddressStateTx(args, keys)
	if err != nil {
		return err
	}

	response.TxID = tx.ID()

	if err = service.vm.Builder.AddUnverifiedTx(tx); err != nil {
		return err
	}
	return nil
}

type GetAddressStateTxArgs struct {
	SetAddressStateArgs

	Encoding formatting.Encoding `json:"encoding"`
}

type GetAddressStateTxReply struct {
	Tx string `json:"tx"`
}

// GetAddressStateTx returnes an unsigned AddAddressStateTx
func (service *Service) GetAddressStateTx(_ *http.Request, args *GetAddressStateTxArgs, response *GetAddressStateTxReply) error {
	service.vm.ctx.Log.Debug("Platform: GetAddressStateTx called")

	keys, err := service.getFakeKeys(&args.JSONSpendHeader)
	if err != nil {
		return err
	}

	tx, err := service.buildAddressStateTx(&args.SetAddressStateArgs, keys)
	if err != nil {
		return err
	}

	bytes, err := txs.Codec.Marshal(txs.Version, tx.Unsigned)
	if err != nil {
		return fmt.Errorf(errSerializeTx, err)
	}

	if response.Tx, err = formatting.Encode(args.Encoding, bytes); err != nil {
		return fmt.Errorf(errEncodeTx, err)
	}
	return nil
}

func (service *Service) getKeystoreKeys(args *api.JSONSpendHeader) (*secp256k1fx.Keychain, error) {
	// Parse the from addresses
	fromAddrs, err := avax.ParseServiceAddresses(service.addrManager, args.From)
	if err != nil {
		return nil, err
	}

	user, err := keystore.NewUserFromKeystore(service.vm.ctx.Keystore, args.Username, args.Password)
	if err != nil {
		return nil, err
	}
	defer user.Close()

	// Get the user's keys
	privKeys, err := keystore.GetKeychain(user, fromAddrs)
	if err != nil {
		return nil, fmt.Errorf("couldn't get addresses controlled by the user: %w", err)
	}

	// Parse the change address.
	if len(privKeys.Keys) == 0 {
		return nil, errNoKeys
	}

	if err = user.Close(); err != nil {
		return nil, err
	}
	return privKeys, nil
}

func (service *Service) getFakeKeys(args *api.JSONSpendHeader) (*secp256k1fx.Keychain, error) {
	// Parse the from addresses
	fromAddrs, err := avax.ParseServiceAddresses(service.addrManager, args.From)
	if err != nil {
		return nil, err
	}

	privKeys := secp256k1fx.NewKeychain()
	for fromAddr := range fromAddrs {
		privKeys.Add(crypto.FakePrivateKey(fromAddr))
	}
	return privKeys, nil
}

func (service *Service) buildAddressStateTx(args *SetAddressStateArgs, keys *secp256k1fx.Keychain) (*txs.Tx, error) {
	var changeAddr ids.ShortID
	if len(args.ChangeAddr) > 0 {
		var err error
		if changeAddr, err = avax.ParseServiceAddress(service.addrManager, args.ChangeAddr); err != nil {
			return nil, fmt.Errorf(errInvalidChangeAddr, err)
		}
	}

	targetAddr, err := avax.ParseServiceAddress(service.addrManager, args.Address)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse param Address: %w", err)
	}

	// Create the transaction
	tx, err := service.vm.txBuilder.NewAddAddressStateTx(
		targetAddr,  // Address to change state
		args.Remove, // Add or remove State
		args.State,  // The state to change
		keys.Keys,   // Keys providing the staked tokens
		changeAddr,
	)
	if err != nil {
		return nil, fmt.Errorf(errCreateTx, err)
	}
	return tx, nil
}