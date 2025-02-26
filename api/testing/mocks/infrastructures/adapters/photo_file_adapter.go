// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/adapters/photo_file_adapter.go

// Package mock_adapters is a generated GoMock package.
package mock_adapters

import (
	context "context"
	reflect "reflect"

	entities "github.com/famiphoto/famiphoto/api/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockPhotoFileAdapter is a mock of PhotoFileAdapter interface.
type MockPhotoFileAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoFileAdapterMockRecorder
}

// MockPhotoFileAdapterMockRecorder is the mock recorder for MockPhotoFileAdapter.
type MockPhotoFileAdapterMockRecorder struct {
	mock *MockPhotoFileAdapter
}

// NewMockPhotoFileAdapter creates a new mock instance.
func NewMockPhotoFileAdapter(ctrl *gomock.Controller) *MockPhotoFileAdapter {
	mock := &MockPhotoFileAdapter{ctrl: ctrl}
	mock.recorder = &MockPhotoFileAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoFileAdapter) EXPECT() *MockPhotoFileAdapterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockPhotoFileAdapter) Upsert(ctx context.Context, photoFile *entities.PhotoFile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, photoFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockPhotoFileAdapterMockRecorder) Upsert(ctx, photoFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockPhotoFileAdapter)(nil).Upsert), ctx, photoFile)
}
