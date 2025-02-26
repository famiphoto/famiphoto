// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/adapters/session_adapter.go

// Package mock_adapters is a generated GoMock package.
package mock_adapters

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionAdapter is a mock of SessionAdapter interface.
type MockSessionAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockSessionAdapterMockRecorder
}

// MockSessionAdapterMockRecorder is the mock recorder for MockSessionAdapter.
type MockSessionAdapterMockRecorder struct {
	mock *MockSessionAdapter
}

// NewMockSessionAdapter creates a new mock instance.
func NewMockSessionAdapter(ctrl *gomock.Controller) *MockSessionAdapter {
	mock := &MockSessionAdapter{ctrl: ctrl}
	mock.recorder = &MockSessionAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionAdapter) EXPECT() *MockSessionAdapterMockRecorder {
	return m.recorder
}

// DeleteSession mocks base method.
func (m *MockSessionAdapter) DeleteSession(ctx context.Context, sessionID string, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, sessionID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionAdapterMockRecorder) DeleteSession(ctx, sessionID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionAdapter)(nil).DeleteSession), ctx, sessionID, userID)
}

// DeleteSessionAll mocks base method.
func (m *MockSessionAdapter) DeleteSessionAll(ctx context.Context, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionAll", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionAll indicates an expected call of DeleteSessionAll.
func (mr *MockSessionAdapterMockRecorder) DeleteSessionAll(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionAll", reflect.TypeOf((*MockSessionAdapter)(nil).DeleteSessionAll), ctx, userID)
}

// LoadSession mocks base method.
func (m *MockSessionAdapter) LoadSession(ctx context.Context, sessionID string) (map[any]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSession", ctx, sessionID)
	ret0, _ := ret[0].(map[any]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadSession indicates an expected call of LoadSession.
func (mr *MockSessionAdapterMockRecorder) LoadSession(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSession", reflect.TypeOf((*MockSessionAdapter)(nil).LoadSession), ctx, sessionID)
}

// SaveSession mocks base method.
func (m *MockSessionAdapter) SaveSession(ctx context.Context, sessionID string, userID int64, values map[any]any, age int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSession", ctx, sessionID, userID, values, age)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSession indicates an expected call of SaveSession.
func (mr *MockSessionAdapterMockRecorder) SaveSession(ctx, sessionID, userID, values, age interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSession", reflect.TypeOf((*MockSessionAdapter)(nil).SaveSession), ctx, sessionID, userID, values, age)
}
