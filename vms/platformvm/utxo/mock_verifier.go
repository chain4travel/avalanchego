// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/platformvm/utxo (interfaces: Verifier)

// Package utxo is a generated GoMock package.
package utxo

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	avax "github.com/ava-labs/avalanchego/vms/components/avax"
	verify "github.com/ava-labs/avalanchego/vms/components/verify"
	locked "github.com/ava-labs/avalanchego/vms/platformvm/locked"
	state "github.com/ava-labs/avalanchego/vms/platformvm/state"
	txs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
	gomock "go.uber.org/mock/gomock"
)

// MockVerifier is a mock of Verifier interface.
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier.
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance.
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// Unlock mocks base method.
func (m *MockVerifier) Unlock(arg0 state.Chain, arg1 []ids.ID, arg2 locked.State) ([]*avax.TransferableInput, []*avax.TransferableOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*avax.TransferableInput)
	ret1, _ := ret[1].([]*avax.TransferableOutput)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Unlock indicates an expected call of Unlock.
func (mr *MockVerifierMockRecorder) Unlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockVerifier)(nil).Unlock), arg0, arg1, arg2)
}

// VerifyLock mocks base method.
func (m *MockVerifier) VerifyLock(arg0 txs.UnsignedTx, arg1 avax.UTXOGetter, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 uint64, arg6 uint64, arg7 ids.ID, arg8 locked.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLock", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyLock indicates an expected call of VerifyLock.
func (mr *MockVerifierMockRecorder) VerifyLock(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLock", reflect.TypeOf((*MockVerifier)(nil).VerifyLock), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// VerifySpend mocks base method.
func (m *MockVerifier) VerifySpend(arg0 txs.UnsignedTx, arg1 avax.UTXOGetter, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpend", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpend indicates an expected call of VerifySpend.
func (mr *MockVerifierMockRecorder) VerifySpend(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpend", reflect.TypeOf((*MockVerifier)(nil).VerifySpend), arg0, arg1, arg2, arg3, arg4, arg5)
}

// VerifySpendUTXOs mocks base method.
func (m *MockVerifier) VerifySpendUTXOs(arg0 avax.UTXOGetter, arg1 txs.UnsignedTx, arg2 []*avax.UTXO, arg3 []*avax.TransferableInput, arg4 []*avax.TransferableOutput, arg5 []verify.Verifiable, arg6 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpendUTXOs", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpendUTXOs indicates an expected call of VerifySpendUTXOs.
func (mr *MockVerifierMockRecorder) VerifySpendUTXOs(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpendUTXOs", reflect.TypeOf((*MockVerifier)(nil).VerifySpendUTXOs), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// VerifyUnlockDeposit mocks base method.
func (m *MockVerifier) VerifyUnlockDeposit(arg0 avax.UTXOGetter, arg1 txs.UnsignedTx, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 uint64, arg6 ids.ID, arg7 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUnlockDeposit", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUnlockDeposit indicates an expected call of VerifyUnlockDeposit.
func (mr *MockVerifierMockRecorder) VerifyUnlockDeposit(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUnlockDeposit", reflect.TypeOf((*MockVerifier)(nil).VerifyUnlockDeposit), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
