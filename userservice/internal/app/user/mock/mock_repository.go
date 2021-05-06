// Code generated by MockGen. DO NOT EDIT.
// Source: user/internal/app/user (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	models "user/internal/app/models"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Change mocks base method.
func (m *MockRepository) Change(arg0 models.ChangeUser, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Change", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Change indicates an expected call of Change.
func (mr *MockRepositoryMockRecorder) Change(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Change", reflect.TypeOf((*MockRepository)(nil).Change), arg0, arg1)
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 models.NewUser, arg1 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1)
}

// FindUserByEmail mocks base method.
func (m *MockRepository) FindUserByEmail(arg0 string, arg1 context.Context) (*models.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*models.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockRepositoryMockRecorder) FindUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockRepository)(nil).FindUserByEmail), arg0, arg1)
}

// FindUserByID mocks base method.
func (m *MockRepository) FindUserByID(arg0 uint64, arg1 context.Context) (*models.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", arg0, arg1)
	ret0, _ := ret[0].(*models.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockRepositoryMockRecorder) FindUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockRepository)(nil).FindUserByID), arg0, arg1)
}

// SetUserImg mocks base method.
func (m *MockRepository) SetUserImg(arg0 uint64, arg1 string, arg2 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserImg", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserImg indicates an expected call of SetUserImg.
func (mr *MockRepositoryMockRecorder) SetUserImg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserImg", reflect.TypeOf((*MockRepository)(nil).SetUserImg), arg0, arg1, arg2)
}
