// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Fantom-foundation/go-lachesis/src/posnode (interfaces: Consensus)

// Package posnode is a generated GoMock package.
package posnode

import (
	hash "github.com/Fantom-foundation/go-lachesis/src/hash"
	inter "github.com/Fantom-foundation/go-lachesis/src/inter"
	idx "github.com/Fantom-foundation/go-lachesis/src/inter/idx"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConsensus is a mock of Consensus interface
type MockConsensus struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusMockRecorder
}

// MockConsensusMockRecorder is the mock recorder for MockConsensus
type MockConsensusMockRecorder struct {
	mock *MockConsensus
}

// NewMockConsensus creates a new mock instance
func NewMockConsensus(ctrl *gomock.Controller) *MockConsensus {
	mock := &MockConsensus{ctrl: ctrl}
	mock.recorder = &MockConsensusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsensus) EXPECT() *MockConsensusMockRecorder {
	return m.recorder
}

// CurrentSuperFrameN mocks base method
func (m *MockConsensus) CurrentSuperFrameN() idx.SuperFrame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentSuperFrameN")
	ret0, _ := ret[0].(idx.SuperFrame)
	return ret0
}

// CurrentSuperFrameN indicates an expected call of CurrentSuperFrameN
func (mr *MockConsensusMockRecorder) CurrentSuperFrameN() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentSuperFrameN", reflect.TypeOf((*MockConsensus)(nil).CurrentSuperFrameN))
}

// Prepare mocks base method
func (m *MockConsensus) Prepare(e *inter.Event) *inter.Event {
	m.ctrl.T.Helper()
	return e
}

// GetGenesisHash mocks base method
func (m *MockConsensus) GetGenesisHash() hash.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesisHash")
	ret0, _ := ret[0].(hash.Hash)
	return ret0
}

// GetGenesisHash indicates an expected call of GetGenesisHash
func (mr *MockConsensusMockRecorder) GetGenesisHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesisHash", reflect.TypeOf((*MockConsensus)(nil).GetGenesisHash))
}

// ProcessEvent mocks base method
func (m *MockConsensus) ProcessEvent(arg0 *inter.Event) error {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PushEvent", arg0)
	return nil
}

// ProcessEvent indicates an expected call of ProcessEvent
func (mr *MockConsensusMockRecorder) ProcessEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessEvent", reflect.TypeOf((*MockConsensus)(nil).ProcessEvent), arg0)
}

// StakeOf mocks base method
func (m *MockConsensus) StakeOf(arg0 hash.Peer) inter.Stake {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StakeOf", arg0)
	ret0, _ := ret[0].(inter.Stake)
	return ret0
}

// StakeOf indicates an expected call of StakeOf
func (mr *MockConsensusMockRecorder) StakeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StakeOf", reflect.TypeOf((*MockConsensus)(nil).StakeOf), arg0)
}

// SuperFrameMembers mocks base method
func (m *MockConsensus) SuperFrameMembers() []hash.Peer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuperFrameMembers")
	ret0, _ := ret[0].([]hash.Peer)
	return ret0
}

// SuperFrameMembers indicates an expected call of SuperFrameMembers
func (mr *MockConsensusMockRecorder) SuperFrameMembers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuperFrameMembers", reflect.TypeOf((*MockConsensus)(nil).SuperFrameMembers))
}
