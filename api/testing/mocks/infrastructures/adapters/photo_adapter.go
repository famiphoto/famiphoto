// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/adapters/photo_adapter.go

// Package mock_adapters is a generated GoMock package.
package mock_adapters

import (
	context "context"
	reflect "reflect"

	entities "github.com/famiphoto/famiphoto/api/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockPhotoAdapter is a mock of PhotoAdapter interface.
type MockPhotoAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoAdapterMockRecorder
}

// MockPhotoAdapterMockRecorder is the mock recorder for MockPhotoAdapter.
type MockPhotoAdapterMockRecorder struct {
	mock *MockPhotoAdapter
}

// NewMockPhotoAdapter creates a new mock instance.
func NewMockPhotoAdapter(ctrl *gomock.Controller) *MockPhotoAdapter {
	mock := &MockPhotoAdapter{ctrl: ctrl}
	mock.recorder = &MockPhotoAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoAdapter) EXPECT() *MockPhotoAdapterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockPhotoAdapter) Upsert(ctx context.Context, photo *entities.Photo) (*entities.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, photo)
	ret0, _ := ret[0].(*entities.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockPhotoAdapterMockRecorder) Upsert(ctx, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockPhotoAdapter)(nil).Upsert), ctx, photo)
}
