// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/platformvm/state (interfaces: Diff)

// Package state is a generated GoMock package.
package state

import (
	reflect "reflect"
	time "time"

	ids "github.com/ava-labs/avalanchego/ids"
	set "github.com/ava-labs/avalanchego/utils/set"
	avax "github.com/ava-labs/avalanchego/vms/components/avax"
	multisig "github.com/ava-labs/avalanchego/vms/components/multisig"
	config "github.com/ava-labs/avalanchego/vms/platformvm/config"
	deposit "github.com/ava-labs/avalanchego/vms/platformvm/deposit"
	locked "github.com/ava-labs/avalanchego/vms/platformvm/locked"
	status "github.com/ava-labs/avalanchego/vms/platformvm/status"
	txs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
	gomock "github.com/golang/mock/gomock"
)

// MockDiff is a mock of Diff interface.
type MockDiff struct {
	ctrl     *gomock.Controller
	recorder *MockDiffMockRecorder
}

// MockDiffMockRecorder is the mock recorder for MockDiff.
type MockDiffMockRecorder struct {
	mock *MockDiff
}

// NewMockDiff creates a new mock instance.
func NewMockDiff(ctrl *gomock.Controller) *MockDiff {
	mock := &MockDiff{ctrl: ctrl}
	mock.recorder = &MockDiffMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiff) EXPECT() *MockDiffMockRecorder {
	return m.recorder
}

// AddChain mocks base method.
func (m *MockDiff) AddChain(arg0 *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChain", arg0)
}

// AddChain indicates an expected call of AddChain.
func (mr *MockDiffMockRecorder) AddChain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChain", reflect.TypeOf((*MockDiff)(nil).AddChain), arg0)
}

// AddDepositOffer mocks base method.
func (m *MockDiff) AddDepositOffer(arg0 *deposit.Offer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddDepositOffer", arg0)
}

// AddDepositOffer indicates an expected call of AddDepositOffer.
func (mr *MockDiffMockRecorder) AddDepositOffer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDepositOffer", reflect.TypeOf((*MockDiff)(nil).AddDepositOffer), arg0)
}

// AddRewardUTXO mocks base method.
func (m *MockDiff) AddRewardUTXO(arg0 ids.ID, arg1 *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRewardUTXO", arg0, arg1)
}

// AddRewardUTXO indicates an expected call of AddRewardUTXO.
func (mr *MockDiffMockRecorder) AddRewardUTXO(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRewardUTXO", reflect.TypeOf((*MockDiff)(nil).AddRewardUTXO), arg0, arg1)
}

// AddSubnet mocks base method.
func (m *MockDiff) AddSubnet(arg0 *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSubnet", arg0)
}

// AddSubnet indicates an expected call of AddSubnet.
func (mr *MockDiffMockRecorder) AddSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnet", reflect.TypeOf((*MockDiff)(nil).AddSubnet), arg0)
}

// AddSubnetTransformation mocks base method.
func (m *MockDiff) AddSubnetTransformation(arg0 *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSubnetTransformation", arg0)
}

// AddSubnetTransformation indicates an expected call of AddSubnetTransformation.
func (mr *MockDiffMockRecorder) AddSubnetTransformation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnetTransformation", reflect.TypeOf((*MockDiff)(nil).AddSubnetTransformation), arg0)
}

// AddTx mocks base method.
func (m *MockDiff) AddTx(arg0 *txs.Tx, arg1 status.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTx", arg0, arg1)
}

// AddTx indicates an expected call of AddTx.
func (mr *MockDiffMockRecorder) AddTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTx", reflect.TypeOf((*MockDiff)(nil).AddTx), arg0, arg1)
}

// AddUTXO mocks base method.
func (m *MockDiff) AddUTXO(arg0 *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUTXO", arg0)
}

// AddUTXO indicates an expected call of AddUTXO.
func (mr *MockDiffMockRecorder) AddUTXO(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUTXO", reflect.TypeOf((*MockDiff)(nil).AddUTXO), arg0)
}

// Apply mocks base method.
func (m *MockDiff) Apply(arg0 State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Apply", arg0)
}

// Apply indicates an expected call of Apply.
func (mr *MockDiffMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockDiff)(nil).Apply), arg0)
}

// ApplyCaminoState mocks base method.
func (m *MockDiff) ApplyCaminoState(arg0 State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyCaminoState", arg0)
}

// ApplyCaminoState indicates an expected call of ApplyCaminoState.
func (mr *MockDiffMockRecorder) ApplyCaminoState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyCaminoState", reflect.TypeOf((*MockDiff)(nil).ApplyCaminoState), arg0)
}

// CaminoConfig mocks base method.
func (m *MockDiff) CaminoConfig() (*CaminoConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaminoConfig")
	ret0, _ := ret[0].(*CaminoConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CaminoConfig indicates an expected call of CaminoConfig.
func (mr *MockDiffMockRecorder) CaminoConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaminoConfig", reflect.TypeOf((*MockDiff)(nil).CaminoConfig))
}

// Config mocks base method.
func (m *MockDiff) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockDiffMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockDiff)(nil).Config))
}

// DeleteCurrentDelegator mocks base method.
func (m *MockDiff) DeleteCurrentDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCurrentDelegator", arg0)
}

// DeleteCurrentDelegator indicates an expected call of DeleteCurrentDelegator.
func (mr *MockDiffMockRecorder) DeleteCurrentDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentDelegator", reflect.TypeOf((*MockDiff)(nil).DeleteCurrentDelegator), arg0)
}

// DeleteCurrentValidator mocks base method.
func (m *MockDiff) DeleteCurrentValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCurrentValidator", arg0)
}

// DeleteCurrentValidator indicates an expected call of DeleteCurrentValidator.
func (mr *MockDiffMockRecorder) DeleteCurrentValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentValidator", reflect.TypeOf((*MockDiff)(nil).DeleteCurrentValidator), arg0)
}

// DeletePendingDelegator mocks base method.
func (m *MockDiff) DeletePendingDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePendingDelegator", arg0)
}

// DeletePendingDelegator indicates an expected call of DeletePendingDelegator.
func (mr *MockDiffMockRecorder) DeletePendingDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingDelegator", reflect.TypeOf((*MockDiff)(nil).DeletePendingDelegator), arg0)
}

// DeletePendingValidator mocks base method.
func (m *MockDiff) DeletePendingValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePendingValidator", arg0)
}

// DeletePendingValidator indicates an expected call of DeletePendingValidator.
func (mr *MockDiffMockRecorder) DeletePendingValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingValidator", reflect.TypeOf((*MockDiff)(nil).DeletePendingValidator), arg0)
}

// DeleteUTXO mocks base method.
func (m *MockDiff) DeleteUTXO(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUTXO", arg0)
}

// DeleteUTXO indicates an expected call of DeleteUTXO.
func (mr *MockDiffMockRecorder) DeleteUTXO(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUTXO", reflect.TypeOf((*MockDiff)(nil).DeleteUTXO), arg0)
}

// GetAddressStates mocks base method.
func (m *MockDiff) GetAddressStates(arg0 ids.ShortID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressStates", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressStates indicates an expected call of GetAddressStates.
func (mr *MockDiffMockRecorder) GetAddressStates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressStates", reflect.TypeOf((*MockDiff)(nil).GetAddressStates), arg0)
}

// GetAllDepositOffers mocks base method.
func (m *MockDiff) GetAllDepositOffers() ([]*deposit.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDepositOffers")
	ret0, _ := ret[0].([]*deposit.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDepositOffers indicates an expected call of GetAllDepositOffers.
func (mr *MockDiffMockRecorder) GetAllDepositOffers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDepositOffers", reflect.TypeOf((*MockDiff)(nil).GetAllDepositOffers))
}

// GetChains mocks base method.
func (m *MockDiff) GetChains(arg0 ids.ID) ([]*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChains", arg0)
	ret0, _ := ret[0].([]*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChains indicates an expected call of GetChains.
func (mr *MockDiffMockRecorder) GetChains(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChains", reflect.TypeOf((*MockDiff)(nil).GetChains), arg0)
}

// GetClaimable mocks base method.
func (m *MockDiff) GetClaimable(arg0 ids.ID) (*Claimable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaimable", arg0)
	ret0, _ := ret[0].(*Claimable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClaimable indicates an expected call of GetClaimable.
func (mr *MockDiffMockRecorder) GetClaimable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaimable", reflect.TypeOf((*MockDiff)(nil).GetClaimable), arg0)
}

// GetCurrentDelegatorIterator mocks base method.
func (m *MockDiff) GetCurrentDelegatorIterator(arg0 ids.ID, arg1 ids.NodeID) (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentDelegatorIterator", arg0, arg1)
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentDelegatorIterator indicates an expected call of GetCurrentDelegatorIterator.
func (mr *MockDiffMockRecorder) GetCurrentDelegatorIterator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentDelegatorIterator", reflect.TypeOf((*MockDiff)(nil).GetCurrentDelegatorIterator), arg0, arg1)
}

// GetCurrentStakerIterator mocks base method.
func (m *MockDiff) GetCurrentStakerIterator() (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentStakerIterator")
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentStakerIterator indicates an expected call of GetCurrentStakerIterator.
func (mr *MockDiffMockRecorder) GetCurrentStakerIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentStakerIterator", reflect.TypeOf((*MockDiff)(nil).GetCurrentStakerIterator))
}

// GetCurrentSupply mocks base method.
func (m *MockDiff) GetCurrentSupply(arg0 ids.ID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSupply", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentSupply indicates an expected call of GetCurrentSupply.
func (mr *MockDiffMockRecorder) GetCurrentSupply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSupply", reflect.TypeOf((*MockDiff)(nil).GetCurrentSupply), arg0)
}

// GetCurrentValidator mocks base method.
func (m *MockDiff) GetCurrentValidator(arg0 ids.ID, arg1 ids.NodeID) (*Staker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentValidator", arg0, arg1)
	ret0, _ := ret[0].(*Staker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentValidator indicates an expected call of GetCurrentValidator.
func (mr *MockDiffMockRecorder) GetCurrentValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentValidator", reflect.TypeOf((*MockDiff)(nil).GetCurrentValidator), arg0, arg1)
}

// GetDeposit mocks base method.
func (m *MockDiff) GetDeposit(arg0 ids.ID) (*deposit.Deposit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeposit", arg0)
	ret0, _ := ret[0].(*deposit.Deposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeposit indicates an expected call of GetDeposit.
func (mr *MockDiffMockRecorder) GetDeposit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeposit", reflect.TypeOf((*MockDiff)(nil).GetDeposit), arg0)
}

// GetDepositOffer mocks base method.
func (m *MockDiff) GetDepositOffer(arg0 ids.ID) (*deposit.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDepositOffer", arg0)
	ret0, _ := ret[0].(*deposit.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDepositOffer indicates an expected call of GetDepositOffer.
func (mr *MockDiffMockRecorder) GetDepositOffer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDepositOffer", reflect.TypeOf((*MockDiff)(nil).GetDepositOffer), arg0)
}

// GetMultisigAlias mocks base method.
func (m *MockDiff) GetMultisigAlias(arg0 ids.ShortID) (*multisig.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultisigAlias", arg0)
	ret0, _ := ret[0].(*multisig.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultisigAlias indicates an expected call of GetMultisigAlias.
func (mr *MockDiffMockRecorder) GetMultisigAlias(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultisigAlias", reflect.TypeOf((*MockDiff)(nil).GetMultisigAlias), arg0)
}

// GetNotDistributedValidatorReward mocks base method.
func (m *MockDiff) GetNotDistributedValidatorReward() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotDistributedValidatorReward")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotDistributedValidatorReward indicates an expected call of GetNotDistributedValidatorReward.
func (mr *MockDiffMockRecorder) GetNotDistributedValidatorReward() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotDistributedValidatorReward", reflect.TypeOf((*MockDiff)(nil).GetNotDistributedValidatorReward))
}

// GetPendingDelegatorIterator mocks base method.
func (m *MockDiff) GetPendingDelegatorIterator(arg0 ids.ID, arg1 ids.NodeID) (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingDelegatorIterator", arg0, arg1)
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingDelegatorIterator indicates an expected call of GetPendingDelegatorIterator.
func (mr *MockDiffMockRecorder) GetPendingDelegatorIterator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingDelegatorIterator", reflect.TypeOf((*MockDiff)(nil).GetPendingDelegatorIterator), arg0, arg1)
}

// GetPendingStakerIterator mocks base method.
func (m *MockDiff) GetPendingStakerIterator() (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingStakerIterator")
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingStakerIterator indicates an expected call of GetPendingStakerIterator.
func (mr *MockDiffMockRecorder) GetPendingStakerIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingStakerIterator", reflect.TypeOf((*MockDiff)(nil).GetPendingStakerIterator))
}

// GetPendingValidator mocks base method.
func (m *MockDiff) GetPendingValidator(arg0 ids.ID, arg1 ids.NodeID) (*Staker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingValidator", arg0, arg1)
	ret0, _ := ret[0].(*Staker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingValidator indicates an expected call of GetPendingValidator.
func (mr *MockDiffMockRecorder) GetPendingValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingValidator", reflect.TypeOf((*MockDiff)(nil).GetPendingValidator), arg0, arg1)
}

// GetDeferredStakerIterator mocks base method.
func (m *MockDiff) GetDeferredStakerIterator() (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeferredStakerIterator")
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeferredStakerIterator indicates an expected call of GetDeferredStakerIterator.
func (mr *MockDiffMockRecorder) GetDeferredStakerIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeferredStakerIterator", reflect.TypeOf((*MockDiff)(nil).GetDeferredStakerIterator))
}

// GetDeferredValidator mocks base method.
func (m *MockDiff) GetDeferredValidator(arg0 ids.ID, arg1 ids.NodeID) (*Staker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeferredValidator", arg0, arg1)
	ret0, _ := ret[0].(*Staker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeferredValidator indicates an expected call of GetDeferredValidator.
func (mr *MockDiffMockRecorder) GetDeferredValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeferredValidator", reflect.TypeOf((*MockDiff)(nil).GetDeferredValidator), arg0, arg1)
}

// DeleteDeferredValidator mocks base method.
func (m *MockDiff) DeleteDeferredValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteDeferredValidator", arg0)
}

// DeleteDeferredValidator indicates an expected call of DeleteDeferredValidator.
func (mr *MockDiffMockRecorder) DeleteDeferredValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeferredValidator", reflect.TypeOf((*MockDiff)(nil).DeleteDeferredValidator), arg0)
}

// PutDeferredValidator mocks base method.
func (m *MockDiff) PutDeferredValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutDeferredValidator", arg0)
}

// PutDeferredValidator indicates an expected call of PutDeferredValidator.
func (mr *MockDiffMockRecorder) PutDeferredValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutDeferredValidator", reflect.TypeOf((*MockDiff)(nil).PutDeferredValidator), arg0)
}

// GetRewardUTXOs mocks base method.
func (m *MockDiff) GetRewardUTXOs(arg0 ids.ID) ([]*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewardUTXOs", arg0)
	ret0, _ := ret[0].([]*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRewardUTXOs indicates an expected call of GetRewardUTXOs.
func (mr *MockDiffMockRecorder) GetRewardUTXOs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewardUTXOs", reflect.TypeOf((*MockDiff)(nil).GetRewardUTXOs), arg0)
}

// GetShortIDLink mocks base method.
func (m *MockDiff) GetShortIDLink(arg0 ids.ShortID, arg1 ShortLinkKey) (ids.ShortID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortIDLink", arg0, arg1)
	ret0, _ := ret[0].(ids.ShortID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShortIDLink indicates an expected call of GetShortIDLink.
func (mr *MockDiffMockRecorder) GetShortIDLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortIDLink", reflect.TypeOf((*MockDiff)(nil).GetShortIDLink), arg0, arg1)
}

// GetSubnetTransformation mocks base method.
func (m *MockDiff) GetSubnetTransformation(arg0 ids.ID) (*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetTransformation", arg0)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetTransformation indicates an expected call of GetSubnetTransformation.
func (mr *MockDiffMockRecorder) GetSubnetTransformation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetTransformation", reflect.TypeOf((*MockDiff)(nil).GetSubnetTransformation), arg0)
}

// GetSubnets mocks base method.
func (m *MockDiff) GetSubnets() ([]*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnets")
	ret0, _ := ret[0].([]*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnets indicates an expected call of GetSubnets.
func (mr *MockDiffMockRecorder) GetSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnets", reflect.TypeOf((*MockDiff)(nil).GetSubnets))
}

// GetTimestamp mocks base method.
func (m *MockDiff) GetTimestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *MockDiffMockRecorder) GetTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*MockDiff)(nil).GetTimestamp))
}

// GetTx mocks base method.
func (m *MockDiff) GetTx(arg0 ids.ID) (*txs.Tx, status.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", arg0)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(status.Status)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTx indicates an expected call of GetTx.
func (mr *MockDiffMockRecorder) GetTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockDiff)(nil).GetTx), arg0)
}

// GetUTXO mocks base method.
func (m *MockDiff) GetUTXO(arg0 ids.ID) (*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTXO", arg0)
	ret0, _ := ret[0].(*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUTXO indicates an expected call of GetUTXO.
func (mr *MockDiffMockRecorder) GetUTXO(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTXO", reflect.TypeOf((*MockDiff)(nil).GetUTXO), arg0)
}

// LockedUTXOs mocks base method.
func (m *MockDiff) LockedUTXOs(arg0 set.Set[ids.ID], arg1 set.Set[ids.ShortID], arg2 locked.State) ([]*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockedUTXOs", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockedUTXOs indicates an expected call of LockedUTXOs.
func (mr *MockDiffMockRecorder) LockedUTXOs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockedUTXOs", reflect.TypeOf((*MockDiff)(nil).LockedUTXOs), arg0, arg1, arg2)
}

// PutCurrentDelegator mocks base method.
func (m *MockDiff) PutCurrentDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCurrentDelegator", arg0)
}

// PutCurrentDelegator indicates an expected call of PutCurrentDelegator.
func (mr *MockDiffMockRecorder) PutCurrentDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCurrentDelegator", reflect.TypeOf((*MockDiff)(nil).PutCurrentDelegator), arg0)
}

// PutCurrentValidator mocks base method.
func (m *MockDiff) PutCurrentValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCurrentValidator", arg0)
}

// PutCurrentValidator indicates an expected call of PutCurrentValidator.
func (mr *MockDiffMockRecorder) PutCurrentValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCurrentValidator", reflect.TypeOf((*MockDiff)(nil).PutCurrentValidator), arg0)
}

// PutPendingDelegator mocks base method.
func (m *MockDiff) PutPendingDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPendingDelegator", arg0)
}

// PutPendingDelegator indicates an expected call of PutPendingDelegator.
func (mr *MockDiffMockRecorder) PutPendingDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPendingDelegator", reflect.TypeOf((*MockDiff)(nil).PutPendingDelegator), arg0)
}

// PutPendingValidator mocks base method.
func (m *MockDiff) PutPendingValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPendingValidator", arg0)
}

// PutPendingValidator indicates an expected call of PutPendingValidator.
func (mr *MockDiffMockRecorder) PutPendingValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPendingValidator", reflect.TypeOf((*MockDiff)(nil).PutPendingValidator), arg0)
}

// SetAddressStates mocks base method.
func (m *MockDiff) SetAddressStates(arg0 ids.ShortID, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAddressStates", arg0, arg1)
}

// SetAddressStates indicates an expected call of SetAddressStates.
func (mr *MockDiffMockRecorder) SetAddressStates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAddressStates", reflect.TypeOf((*MockDiff)(nil).SetAddressStates), arg0, arg1)
}

// SetClaimable mocks base method.
func (m *MockDiff) SetClaimable(arg0 ids.ID, arg1 *Claimable) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClaimable", arg0, arg1)
}

// SetClaimable indicates an expected call of SetClaimable.
func (mr *MockDiffMockRecorder) SetClaimable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClaimable", reflect.TypeOf((*MockDiff)(nil).SetClaimable), arg0, arg1)
}

// SetCurrentSupply mocks base method.
func (m *MockDiff) SetCurrentSupply(arg0 ids.ID, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentSupply", arg0, arg1)
}

// SetCurrentSupply indicates an expected call of SetCurrentSupply.
func (mr *MockDiffMockRecorder) SetCurrentSupply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentSupply", reflect.TypeOf((*MockDiff)(nil).SetCurrentSupply), arg0, arg1)
}

// SetMultisigAliasRaw mocks base method.
func (m *MockDiff) SetMultisigAliasRaw(arg0 *multisig.AliasRaw) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMultisigAliasRaw", arg0)
}

// SetMultisigAliasRaw indicates an expected call of SetMultisigAliasRaw.
func (mr *MockDiffMockRecorder) SetMultisigAliasRaw(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMultisigAliasRaw", reflect.TypeOf((*MockDiff)(nil).SetMultisigAliasRaw), arg0)
}

// SetNotDistributedValidatorReward mocks base method.
func (m *MockDiff) SetNotDistributedValidatorReward(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNotDistributedValidatorReward", arg0)
}

// SetNotDistributedValidatorReward indicates an expected call of SetNotDistributedValidatorReward.
func (mr *MockDiffMockRecorder) SetNotDistributedValidatorReward(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNotDistributedValidatorReward", reflect.TypeOf((*MockDiff)(nil).SetNotDistributedValidatorReward), arg0)
}

// SetShortIDLink mocks base method.
func (m *MockDiff) SetShortIDLink(arg0 ids.ShortID, arg1 ShortLinkKey, arg2 *ids.ShortID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetShortIDLink", arg0, arg1, arg2)
}

// SetShortIDLink indicates an expected call of SetShortIDLink.
func (mr *MockDiffMockRecorder) SetShortIDLink(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShortIDLink", reflect.TypeOf((*MockDiff)(nil).SetShortIDLink), arg0, arg1, arg2)
}

// SetTimestamp mocks base method.
func (m *MockDiff) SetTimestamp(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimestamp", arg0)
}

// SetTimestamp indicates an expected call of SetTimestamp.
func (mr *MockDiffMockRecorder) SetTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimestamp", reflect.TypeOf((*MockDiff)(nil).SetTimestamp), arg0)
}

// UpdateDeposit mocks base method.
func (m *MockDiff) UpdateDeposit(arg0 ids.ID, arg1 *deposit.Deposit) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeposit", arg0, arg1)
}

// UpdateDeposit indicates an expected call of UpdateDeposit.
func (mr *MockDiffMockRecorder) UpdateDeposit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeposit", reflect.TypeOf((*MockDiff)(nil).UpdateDeposit), arg0, arg1)
}
