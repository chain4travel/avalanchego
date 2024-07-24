// Code generated by MockGen. DO NOT EDIT.
// Source: vms/components/avax/atomic_utxos.go
//
// Generated by this command:
//
//	mockgen -source=vms/components/avax/atomic_utxos.go -destination=vms/components/avax/mock_atomic_utxos.go -package=avax -exclude_interfaces=
//

// Package avax is a generated GoMock package.
package avax

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	set "github.com/ava-labs/avalanchego/utils/set"
	gomock "go.uber.org/mock/gomock"
)

// MockAtomicUTXOManager is a mock of AtomicUTXOManager interface.
type MockAtomicUTXOManager struct {
	ctrl     *gomock.Controller
	recorder *MockAtomicUTXOManagerMockRecorder
}

// MockAtomicUTXOManagerMockRecorder is the mock recorder for MockAtomicUTXOManager.
type MockAtomicUTXOManagerMockRecorder struct {
	mock *MockAtomicUTXOManager
}

// NewMockAtomicUTXOManager creates a new mock instance.
func NewMockAtomicUTXOManager(ctrl *gomock.Controller) *MockAtomicUTXOManager {
	mock := &MockAtomicUTXOManager{ctrl: ctrl}
	mock.recorder = &MockAtomicUTXOManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtomicUTXOManager) EXPECT() *MockAtomicUTXOManagerMockRecorder {
	return m.recorder
}

// GetAtomicUTXOs mocks base method.
func (m *MockAtomicUTXOManager) GetAtomicUTXOs(chainID ids.ID, addrs set.Set[ids.ShortID], startAddr ids.ShortID, startUTXOID ids.ID, limit int) ([]*UTXO, ids.ShortID, ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAtomicUTXOs", chainID, addrs, startAddr, startUTXOID, limit)
	ret0, _ := ret[0].([]*UTXO)
	ret1, _ := ret[1].(ids.ShortID)
	ret2, _ := ret[2].(ids.ID)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetAtomicUTXOs indicates an expected call of GetAtomicUTXOs.
func (mr *MockAtomicUTXOManagerMockRecorder) GetAtomicUTXOs(chainID, addrs, startAddr, startUTXOID, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAtomicUTXOs", reflect.TypeOf((*MockAtomicUTXOManager)(nil).GetAtomicUTXOs), chainID, addrs, startAddr, startUTXOID, limit)
}
