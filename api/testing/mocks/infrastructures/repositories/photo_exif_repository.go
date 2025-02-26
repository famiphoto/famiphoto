// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/photo_exif_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	dbmodels "github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	gomock "github.com/golang/mock/gomock"
)

// MockPhotoExifRepository is a mock of PhotoExifRepository interface.
type MockPhotoExifRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoExifRepositoryMockRecorder
}

// MockPhotoExifRepositoryMockRecorder is the mock recorder for MockPhotoExifRepository.
type MockPhotoExifRepositoryMockRecorder struct {
	mock *MockPhotoExifRepository
}

// NewMockPhotoExifRepository creates a new mock instance.
func NewMockPhotoExifRepository(ctrl *gomock.Controller) *MockPhotoExifRepository {
	mock := &MockPhotoExifRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoExifRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoExifRepository) EXPECT() *MockPhotoExifRepositoryMockRecorder {
	return m.recorder
}

// GetPhotoExifByPhotoIDTagID mocks base method.
func (m *MockPhotoExifRepository) GetPhotoExifByPhotoIDTagID(ctx context.Context, photoID, tagID int64) (*dbmodels.PhotoExif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoExifByPhotoIDTagID", ctx, photoID, tagID)
	ret0, _ := ret[0].(*dbmodels.PhotoExif)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoExifByPhotoIDTagID indicates an expected call of GetPhotoExifByPhotoIDTagID.
func (mr *MockPhotoExifRepositoryMockRecorder) GetPhotoExifByPhotoIDTagID(ctx, photoID, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoExifByPhotoIDTagID", reflect.TypeOf((*MockPhotoExifRepository)(nil).GetPhotoExifByPhotoIDTagID), ctx, photoID, tagID)
}

// Insert mocks base method.
func (m *MockPhotoExifRepository) Insert(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, exif)
	ret0, _ := ret[0].(*dbmodels.PhotoExif)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockPhotoExifRepositoryMockRecorder) Insert(ctx, exif interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPhotoExifRepository)(nil).Insert), ctx, exif)
}

// Update mocks base method.
func (m *MockPhotoExifRepository) Update(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, exif)
	ret0, _ := ret[0].(*dbmodels.PhotoExif)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPhotoExifRepositoryMockRecorder) Update(ctx, exif interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPhotoExifRepository)(nil).Update), ctx, exif)
}
