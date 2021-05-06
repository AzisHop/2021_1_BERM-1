// Code generated by MockGen. DO NOT EDIT.
// Source: post/internal/app/response (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "post/internal/app/models"
	reflect "reflect"

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

// ChangeOrderResponse mocks base method.
func (m *MockRepository) ChangeOrderResponse(arg0 models.Response, arg1 context.Context) (*models.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderResponse", arg0, arg1)
	ret0, _ := ret[0].(*models.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeOrderResponse indicates an expected call of ChangeOrderResponse.
func (mr *MockRepositoryMockRecorder) ChangeOrderResponse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderResponse", reflect.TypeOf((*MockRepository)(nil).ChangeOrderResponse), arg0, arg1)
}

// ChangeVacancyResponse mocks base method.
func (m *MockRepository) ChangeVacancyResponse(arg0 models.Response, arg1 context.Context) (*models.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeVacancyResponse", arg0, arg1)
	ret0, _ := ret[0].(*models.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeVacancyResponse indicates an expected call of ChangeVacancyResponse.
func (mr *MockRepositoryMockRecorder) ChangeVacancyResponse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeVacancyResponse", reflect.TypeOf((*MockRepository)(nil).ChangeVacancyResponse), arg0, arg1)
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 models.Response, arg1 context.Context) (uint64, error) {
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

// DeleteOrderResponse mocks base method.
func (m *MockRepository) DeleteOrderResponse(arg0 models.Response, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderResponse", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderResponse indicates an expected call of DeleteOrderResponse.
func (mr *MockRepositoryMockRecorder) DeleteOrderResponse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderResponse", reflect.TypeOf((*MockRepository)(nil).DeleteOrderResponse), arg0, arg1)
}

// DeleteVacancyResponse mocks base method.
func (m *MockRepository) DeleteVacancyResponse(arg0 models.Response, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVacancyResponse", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVacancyResponse indicates an expected call of DeleteVacancyResponse.
func (mr *MockRepositoryMockRecorder) DeleteVacancyResponse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVacancyResponse", reflect.TypeOf((*MockRepository)(nil).DeleteVacancyResponse), arg0, arg1)
}

// FindByOrderPostID mocks base method.
func (m *MockRepository) FindByOrderPostID(arg0 uint64, arg1 context.Context) ([]models.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByOrderPostID", arg0, arg1)
	ret0, _ := ret[0].([]models.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOrderPostID indicates an expected call of FindByOrderPostID.
func (mr *MockRepositoryMockRecorder) FindByOrderPostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOrderPostID", reflect.TypeOf((*MockRepository)(nil).FindByOrderPostID), arg0, arg1)
}

// FindByVacancyPostID mocks base method.
func (m *MockRepository) FindByVacancyPostID(arg0 uint64, arg1 context.Context) ([]models.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByVacancyPostID", arg0, arg1)
	ret0, _ := ret[0].([]models.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVacancyPostID indicates an expected call of FindByVacancyPostID.
func (mr *MockRepositoryMockRecorder) FindByVacancyPostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVacancyPostID", reflect.TypeOf((*MockRepository)(nil).FindByVacancyPostID), arg0, arg1)
}
