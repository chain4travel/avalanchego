// Code generated by MockGen. DO NOT EDIT.
// Source: cache_internal_state.go

// Package platformvm is a generated GoMock package.
package platformvm

import (
	reflect "reflect"
	time "time"

	database "github.com/chain4travel/caminogo/database"
	ids "github.com/chain4travel/caminogo/ids"
	avax "github.com/chain4travel/caminogo/vms/components/avax"
	status "github.com/chain4travel/caminogo/vms/platformvm/status"
	gomock "github.com/golang/mock/gomock"
)

// MockInternalState is a mock of InternalState interface.
type MockInternalState struct {
	ctrl     *gomock.Controller
	recorder *MockInternalStateMockRecorder
}

// MockInternalStateMockRecorder is the mock recorder for MockInternalState.
type MockInternalStateMockRecorder struct {
	mock *MockInternalState
}

// NewMockInternalState creates a new mock instance.
func NewMockInternalState(ctrl *gomock.Controller) *MockInternalState {
	mock := &MockInternalState{ctrl: ctrl}
	mock.recorder = &MockInternalStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInternalState) EXPECT() *MockInternalStateMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *MockInternalState) Abort() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort.
func (mr *MockInternalStateMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockInternalState)(nil).Abort))
}

// AddBlock mocks base method.
func (m *MockInternalState) AddBlock(block Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBlock", block)
}

// AddBlock indicates an expected call of AddBlock.
func (mr *MockInternalStateMockRecorder) AddBlock(block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlock", reflect.TypeOf((*MockInternalState)(nil).AddBlock), block)
}

// AddChain mocks base method.
func (m *MockInternalState) AddChain(createChainTx *Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChain", createChainTx)
}

// AddChain indicates an expected call of AddChain.
func (mr *MockInternalStateMockRecorder) AddChain(createChainTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChain", reflect.TypeOf((*MockInternalState)(nil).AddChain), createChainTx)
}

// AddCurrentStaker mocks base method.
func (m *MockInternalState) AddCurrentStaker(tx *Tx, potentialReward uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddCurrentStaker", tx, potentialReward)
}

// AddCurrentStaker indicates an expected call of AddCurrentStaker.
func (mr *MockInternalStateMockRecorder) AddCurrentStaker(tx, potentialReward interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCurrentStaker", reflect.TypeOf((*MockInternalState)(nil).AddCurrentStaker), tx, potentialReward)
}

// AddDepositOffer mocks base method.
func (m *MockInternalState) AddDepositOffer(offer *depositOffer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddDepositOffer", offer)
}

// AddDepositOffer indicates an expected call of AddDepositOffer.
func (mr *MockInternalStateMockRecorder) AddDepositOffer(offer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDepositOffer", reflect.TypeOf((*MockInternalState)(nil).AddDepositOffer), offer)
}

// AddPendingStaker mocks base method.
func (m *MockInternalState) AddPendingStaker(tx *Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPendingStaker", tx)
}

// AddPendingStaker indicates an expected call of AddPendingStaker.
func (mr *MockInternalStateMockRecorder) AddPendingStaker(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPendingStaker", reflect.TypeOf((*MockInternalState)(nil).AddPendingStaker), tx)
}

// AddRewardUTXO mocks base method.
func (m *MockInternalState) AddRewardUTXO(txID ids.ID, utxo *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRewardUTXO", txID, utxo)
}

// AddRewardUTXO indicates an expected call of AddRewardUTXO.
func (mr *MockInternalStateMockRecorder) AddRewardUTXO(txID, utxo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRewardUTXO", reflect.TypeOf((*MockInternalState)(nil).AddRewardUTXO), txID, utxo)
}

// AddSubnet mocks base method.
func (m *MockInternalState) AddSubnet(createSubnetTx *Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSubnet", createSubnetTx)
}

// AddSubnet indicates an expected call of AddSubnet.
func (mr *MockInternalStateMockRecorder) AddSubnet(createSubnetTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnet", reflect.TypeOf((*MockInternalState)(nil).AddSubnet), createSubnetTx)
}

// AddTx mocks base method.
func (m *MockInternalState) AddTx(tx *Tx, status status.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTx", tx, status)
}

// AddTx indicates an expected call of AddTx.
func (mr *MockInternalStateMockRecorder) AddTx(tx, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTx", reflect.TypeOf((*MockInternalState)(nil).AddTx), tx, status)
}

// AddUTXO mocks base method.
func (m *MockInternalState) AddUTXO(utxo *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUTXO", utxo)
}

// AddUTXO indicates an expected call of AddUTXO.
func (mr *MockInternalStateMockRecorder) AddUTXO(utxo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUTXO", reflect.TypeOf((*MockInternalState)(nil).AddUTXO), utxo)
}

// Close mocks base method.
func (m *MockInternalState) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockInternalStateMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockInternalState)(nil).Close))
}

// Commit mocks base method.
func (m *MockInternalState) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockInternalStateMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockInternalState)(nil).Commit))
}

// CommitBatch mocks base method.
func (m *MockInternalState) CommitBatch() (database.Batch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBatch")
	ret0, _ := ret[0].(database.Batch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitBatch indicates an expected call of CommitBatch.
func (mr *MockInternalStateMockRecorder) CommitBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBatch", reflect.TypeOf((*MockInternalState)(nil).CommitBatch))
}

// CurrentStakerChainState mocks base method.
func (m *MockInternalState) CurrentStakerChainState() currentStakerChainState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentStakerChainState")
	ret0, _ := ret[0].(currentStakerChainState)
	return ret0
}

// CurrentStakerChainState indicates an expected call of CurrentStakerChainState.
func (mr *MockInternalStateMockRecorder) CurrentStakerChainState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentStakerChainState", reflect.TypeOf((*MockInternalState)(nil).CurrentStakerChainState))
}

// DeleteCurrentStaker mocks base method.
func (m *MockInternalState) DeleteCurrentStaker(tx *Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCurrentStaker", tx)
}

// DeleteCurrentStaker indicates an expected call of DeleteCurrentStaker.
func (mr *MockInternalStateMockRecorder) DeleteCurrentStaker(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentStaker", reflect.TypeOf((*MockInternalState)(nil).DeleteCurrentStaker), tx)
}

// DeletePendingStaker mocks base method.
func (m *MockInternalState) DeletePendingStaker(tx *Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePendingStaker", tx)
}

// DeletePendingStaker indicates an expected call of DeletePendingStaker.
func (mr *MockInternalStateMockRecorder) DeletePendingStaker(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingStaker", reflect.TypeOf((*MockInternalState)(nil).DeletePendingStaker), tx)
}

// DeleteUTXO mocks base method.
func (m *MockInternalState) DeleteUTXO(utxoID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUTXO", utxoID)
}

// DeleteUTXO indicates an expected call of DeleteUTXO.
func (mr *MockInternalStateMockRecorder) DeleteUTXO(utxoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUTXO", reflect.TypeOf((*MockInternalState)(nil).DeleteUTXO), utxoID)
}

// DepositOffersChainState mocks base method.
func (m *MockInternalState) DepositOffersChainState() depositOffersChainState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositOffersChainState")
	ret0, _ := ret[0].(depositOffersChainState)
	return ret0
}

// DepositOffersChainState indicates an expected call of DepositOffersChainState.
func (mr *MockInternalStateMockRecorder) DepositOffersChainState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositOffersChainState", reflect.TypeOf((*MockInternalState)(nil).DepositOffersChainState))
}

// GetBlock mocks base method.
func (m *MockInternalState) GetBlock(blockID ids.ID) (Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", blockID)
	ret0, _ := ret[0].(Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockInternalStateMockRecorder) GetBlock(blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockInternalState)(nil).GetBlock), blockID)
}

// GetChains mocks base method.
func (m *MockInternalState) GetChains(subnetID ids.ID) ([]*Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChains", subnetID)
	ret0, _ := ret[0].([]*Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChains indicates an expected call of GetChains.
func (mr *MockInternalStateMockRecorder) GetChains(subnetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChains", reflect.TypeOf((*MockInternalState)(nil).GetChains), subnetID)
}

// GetCurrentSupply mocks base method.
func (m *MockInternalState) GetCurrentSupply() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSupply")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetCurrentSupply indicates an expected call of GetCurrentSupply.
func (mr *MockInternalStateMockRecorder) GetCurrentSupply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSupply", reflect.TypeOf((*MockInternalState)(nil).GetCurrentSupply))
}

// GetLastAccepted mocks base method.
func (m *MockInternalState) GetLastAccepted() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastAccepted")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// GetLastAccepted indicates an expected call of GetLastAccepted.
func (mr *MockInternalStateMockRecorder) GetLastAccepted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAccepted", reflect.TypeOf((*MockInternalState)(nil).GetLastAccepted))
}

// GetRewardUTXOs mocks base method.
func (m *MockInternalState) GetRewardUTXOs(txID ids.ID) ([]*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewardUTXOs", txID)
	ret0, _ := ret[0].([]*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRewardUTXOs indicates an expected call of GetRewardUTXOs.
func (mr *MockInternalStateMockRecorder) GetRewardUTXOs(txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewardUTXOs", reflect.TypeOf((*MockInternalState)(nil).GetRewardUTXOs), txID)
}

// GetStartTime mocks base method.
func (m *MockInternalState) GetStartTime(nodeID ids.ShortID) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartTime", nodeID)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStartTime indicates an expected call of GetStartTime.
func (mr *MockInternalStateMockRecorder) GetStartTime(nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartTime", reflect.TypeOf((*MockInternalState)(nil).GetStartTime), nodeID)
}

// GetSubnets mocks base method.
func (m *MockInternalState) GetSubnets() ([]*Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnets")
	ret0, _ := ret[0].([]*Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnets indicates an expected call of GetSubnets.
func (mr *MockInternalStateMockRecorder) GetSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnets", reflect.TypeOf((*MockInternalState)(nil).GetSubnets))
}

// GetTimestamp mocks base method.
func (m *MockInternalState) GetTimestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *MockInternalStateMockRecorder) GetTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*MockInternalState)(nil).GetTimestamp))
}

// GetTx mocks base method.
func (m *MockInternalState) GetTx(txID ids.ID) (*Tx, status.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", txID)
	ret0, _ := ret[0].(*Tx)
	ret1, _ := ret[1].(status.Status)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTx indicates an expected call of GetTx.
func (mr *MockInternalStateMockRecorder) GetTx(txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockInternalState)(nil).GetTx), txID)
}

// GetUTXO mocks base method.
func (m *MockInternalState) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTXO", utxoID)
	ret0, _ := ret[0].(*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUTXO indicates an expected call of GetUTXO.
func (mr *MockInternalStateMockRecorder) GetUTXO(utxoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTXO", reflect.TypeOf((*MockInternalState)(nil).GetUTXO), utxoID)
}

// GetUptime mocks base method.
func (m *MockInternalState) GetUptime(nodeID ids.ShortID) (time.Duration, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUptime", nodeID)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUptime indicates an expected call of GetUptime.
func (mr *MockInternalStateMockRecorder) GetUptime(nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUptime", reflect.TypeOf((*MockInternalState)(nil).GetUptime), nodeID)
}

// GetValidatorWeightDiffs mocks base method.
func (m *MockInternalState) GetValidatorWeightDiffs(height uint64, subnetID ids.ID) (map[ids.ShortID]*ValidatorWeightDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorWeightDiffs", height, subnetID)
	ret0, _ := ret[0].(map[ids.ShortID]*ValidatorWeightDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorWeightDiffs indicates an expected call of GetValidatorWeightDiffs.
func (mr *MockInternalStateMockRecorder) GetValidatorWeightDiffs(height, subnetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorWeightDiffs", reflect.TypeOf((*MockInternalState)(nil).GetValidatorWeightDiffs), height, subnetID)
}

// PendingStakerChainState mocks base method.
func (m *MockInternalState) PendingStakerChainState() pendingStakerChainState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingStakerChainState")
	ret0, _ := ret[0].(pendingStakerChainState)
	return ret0
}

// PendingStakerChainState indicates an expected call of PendingStakerChainState.
func (mr *MockInternalStateMockRecorder) PendingStakerChainState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingStakerChainState", reflect.TypeOf((*MockInternalState)(nil).PendingStakerChainState))
}

// SetCurrentStakerChainState mocks base method.
func (m *MockInternalState) SetCurrentStakerChainState(arg0 currentStakerChainState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentStakerChainState", arg0)
}

// SetCurrentStakerChainState indicates an expected call of SetCurrentStakerChainState.
func (mr *MockInternalStateMockRecorder) SetCurrentStakerChainState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentStakerChainState", reflect.TypeOf((*MockInternalState)(nil).SetCurrentStakerChainState), arg0)
}

// SetCurrentSupply mocks base method.
func (m *MockInternalState) SetCurrentSupply(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentSupply", arg0)
}

// SetCurrentSupply indicates an expected call of SetCurrentSupply.
func (mr *MockInternalStateMockRecorder) SetCurrentSupply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentSupply", reflect.TypeOf((*MockInternalState)(nil).SetCurrentSupply), arg0)
}

// SetDepositOffersChainState mocks base method.
func (m *MockInternalState) SetDepositOffersChainState(arg0 depositOffersChainState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDepositOffersChainState", arg0)
}

// SetDepositOffersChainState indicates an expected call of SetDepositOffersChainState.
func (mr *MockInternalStateMockRecorder) SetDepositOffersChainState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDepositOffersChainState", reflect.TypeOf((*MockInternalState)(nil).SetDepositOffersChainState), arg0)
}

// SetHeight mocks base method.
func (m *MockInternalState) SetHeight(height uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHeight", height)
}

// SetHeight indicates an expected call of SetHeight.
func (mr *MockInternalStateMockRecorder) SetHeight(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeight", reflect.TypeOf((*MockInternalState)(nil).SetHeight), height)
}

// SetLastAccepted mocks base method.
func (m *MockInternalState) SetLastAccepted(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastAccepted", arg0)
}

// SetLastAccepted indicates an expected call of SetLastAccepted.
func (mr *MockInternalStateMockRecorder) SetLastAccepted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastAccepted", reflect.TypeOf((*MockInternalState)(nil).SetLastAccepted), arg0)
}

// SetPendingStakerChainState mocks base method.
func (m *MockInternalState) SetPendingStakerChainState(arg0 pendingStakerChainState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPendingStakerChainState", arg0)
}

// SetPendingStakerChainState indicates an expected call of SetPendingStakerChainState.
func (mr *MockInternalStateMockRecorder) SetPendingStakerChainState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPendingStakerChainState", reflect.TypeOf((*MockInternalState)(nil).SetPendingStakerChainState), arg0)
}

// SetTimestamp mocks base method.
func (m *MockInternalState) SetTimestamp(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimestamp", arg0)
}

// SetTimestamp indicates an expected call of SetTimestamp.
func (mr *MockInternalStateMockRecorder) SetTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimestamp", reflect.TypeOf((*MockInternalState)(nil).SetTimestamp), arg0)
}

// SetUptime mocks base method.
func (m *MockInternalState) SetUptime(nodeID ids.ShortID, upDuration time.Duration, lastUpdated time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUptime", nodeID, upDuration, lastUpdated)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUptime indicates an expected call of SetUptime.
func (mr *MockInternalStateMockRecorder) SetUptime(nodeID, upDuration, lastUpdated interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUptime", reflect.TypeOf((*MockInternalState)(nil).SetUptime), nodeID, upDuration, lastUpdated)
}

// UTXOIDs mocks base method.
func (m *MockInternalState) UTXOIDs(addr []byte, previous ids.ID, limit int) ([]ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UTXOIDs", addr, previous, limit)
	ret0, _ := ret[0].([]ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UTXOIDs indicates an expected call of UTXOIDs.
func (mr *MockInternalStateMockRecorder) UTXOIDs(addr, previous, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UTXOIDs", reflect.TypeOf((*MockInternalState)(nil).UTXOIDs), addr, previous, limit)
}
