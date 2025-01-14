// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/platformvm/utxo (interfaces: Handler)

// Package utxo is a generated GoMock package.
package utxo

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	secp256k1 "github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	avax "github.com/ava-labs/avalanchego/vms/components/avax"
	verify "github.com/ava-labs/avalanchego/vms/components/verify"
	locked "github.com/ava-labs/avalanchego/vms/platformvm/locked"
	state "github.com/ava-labs/avalanchego/vms/platformvm/state"
	txs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
	secp256k1fx "github.com/ava-labs/avalanchego/vms/secp256k1fx"
	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockHandler) Authorize(arg0 state.Chain, arg1 ids.ID, arg2 []*secp256k1.PrivateKey) (verify.Verifiable, []*secp256k1.PrivateKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", arg0, arg1, arg2)
	ret0, _ := ret[0].(verify.Verifiable)
	ret1, _ := ret[1].([]*secp256k1.PrivateKey)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Authorize indicates an expected call of Authorize.
func (mr *MockHandlerMockRecorder) Authorize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockHandler)(nil).Authorize), arg0, arg1, arg2)
}

// Lock mocks base method.
func (m *MockHandler) Lock(arg0 []*secp256k1.PrivateKey, arg1, arg2 uint64, arg3 locked.State, arg4, arg5 *secp256k1fx.OutputOwners, arg6 uint64) ([]*avax.TransferableInput, []*avax.TransferableOutput, [][]*secp256k1.PrivateKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]*avax.TransferableInput)
	ret1, _ := ret[1].([]*avax.TransferableOutput)
	ret2, _ := ret[2].([][]*secp256k1.PrivateKey)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Lock indicates an expected call of Lock.
func (mr *MockHandlerMockRecorder) Lock(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockHandler)(nil).Lock), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Spend mocks base method.
func (m *MockHandler) Spend(arg0 []*secp256k1.PrivateKey, arg1, arg2 uint64, arg3 ids.ShortID) ([]*avax.TransferableInput, []*avax.TransferableOutput, []*avax.TransferableOutput, [][]*secp256k1.PrivateKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Spend", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*avax.TransferableInput)
	ret1, _ := ret[1].([]*avax.TransferableOutput)
	ret2, _ := ret[2].([]*avax.TransferableOutput)
	ret3, _ := ret[3].([][]*secp256k1.PrivateKey)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// Spend indicates an expected call of Spend.
func (mr *MockHandlerMockRecorder) Spend(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spend", reflect.TypeOf((*MockHandler)(nil).Spend), arg0, arg1, arg2, arg3)
}

// Unlock mocks base method.
func (m *MockHandler) Unlock(arg0 state.Chain, arg1 []ids.ID, arg2 locked.State) ([]*avax.TransferableInput, []*avax.TransferableOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*avax.TransferableInput)
	ret1, _ := ret[1].([]*avax.TransferableOutput)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Unlock indicates an expected call of Unlock.
func (mr *MockHandlerMockRecorder) Unlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockHandler)(nil).Unlock), arg0, arg1, arg2)
}

// UnlockDeposit mocks base method.
func (m *MockHandler) UnlockDeposit(arg0 state.Chain, arg1 []*secp256k1.PrivateKey, arg2 []ids.ID) ([]*avax.TransferableInput, []*avax.TransferableOutput, [][]*secp256k1.PrivateKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockDeposit", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*avax.TransferableInput)
	ret1, _ := ret[1].([]*avax.TransferableOutput)
	ret2, _ := ret[2].([][]*secp256k1.PrivateKey)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// UnlockDeposit indicates an expected call of UnlockDeposit.
func (mr *MockHandlerMockRecorder) UnlockDeposit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockDeposit", reflect.TypeOf((*MockHandler)(nil).UnlockDeposit), arg0, arg1, arg2)
}

// VerifyLock mocks base method.
func (m *MockHandler) VerifyLock(arg0 txs.UnsignedTx, arg1 avax.UTXOGetter, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 uint64, arg6 ids.ID, arg7 locked.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLock", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyLock indicates an expected call of VerifyLock.
func (mr *MockHandlerMockRecorder) VerifyLock(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLock", reflect.TypeOf((*MockHandler)(nil).VerifyLock), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// VerifySpend mocks base method.
func (m *MockHandler) VerifySpend(arg0 txs.UnsignedTx, arg1 avax.UTXOGetter, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpend", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpend indicates an expected call of VerifySpend.
func (mr *MockHandlerMockRecorder) VerifySpend(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpend", reflect.TypeOf((*MockHandler)(nil).VerifySpend), arg0, arg1, arg2, arg3, arg4, arg5)
}

// VerifySpendUTXOs mocks base method.
func (m *MockHandler) VerifySpendUTXOs(arg0 txs.UnsignedTx, arg1 []*avax.UTXO, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpendUTXOs", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpendUTXOs indicates an expected call of VerifySpendUTXOs.
func (mr *MockHandlerMockRecorder) VerifySpendUTXOs(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpendUTXOs", reflect.TypeOf((*MockHandler)(nil).VerifySpendUTXOs), arg0, arg1, arg2, arg3, arg4, arg5)
}

// VerifyUnlockDeposit mocks base method.
func (m *MockHandler) VerifyUnlockDeposit(arg0 state.Chain, arg1 txs.UnsignedTx, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 uint64, arg6 ids.ID) (map[ids.ID]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUnlockDeposit", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(map[ids.ID]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyUnlockDeposit indicates an expected call of VerifyUnlockDeposit.
func (mr *MockHandlerMockRecorder) VerifyUnlockDeposit(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUnlockDeposit", reflect.TypeOf((*MockHandler)(nil).VerifyUnlockDeposit), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
