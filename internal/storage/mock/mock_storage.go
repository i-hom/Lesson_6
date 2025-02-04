// Code generated by MockGen. DO NOT EDIT.
// Source: .\storage\storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	"github.com/Shemistan/Lesson_6/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIStorage) Add(arg0 *model.User) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIStorageMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIStorage)(nil).Add), arg0)
}

// Auth mocks base method.
func (m *MockIStorage) Auth(arg0, arg1 string) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", arg0, arg1)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth.
func (mr *MockIStorageMockRecorder) Auth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockIStorage)(nil).Auth), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIStorage) Delete(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIStorageMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIStorage)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockIStorage) Get(arg0 uint32) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIStorageMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIStorage)(nil).Get), arg0)
}

// GetAll mocks base method.
func (m *MockIStorage) GetAll() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIStorageMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIStorage)(nil).GetAll))
}

// Update mocks base method.
func (m *MockIStorage) Update(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIStorageMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIStorage)(nil).Update), arg0)
}
