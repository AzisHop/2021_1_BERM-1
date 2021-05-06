// Code generated by MockGen. DO NOT EDIT.
// Source: post/internal/app/order (interfaces: UseCase)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "post/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// ChangeOrder mocks base method.
func (m *MockUseCase) ChangeOrder(arg0 models.Order, arg1 context.Context) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrder", arg0, arg1)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeOrder indicates an expected call of ChangeOrder.
func (mr *MockUseCaseMockRecorder) ChangeOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrder", reflect.TypeOf((*MockUseCase)(nil).ChangeOrder), arg0, arg1)
}

// CloseOrder mocks base method.
func (m *MockUseCase) CloseOrder(arg0 uint64, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseOrder indicates an expected call of CloseOrder.
func (mr *MockUseCaseMockRecorder) CloseOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseOrder", reflect.TypeOf((*MockUseCase)(nil).CloseOrder), arg0, arg1)
}

// Create mocks base method.
func (m *MockUseCase) Create(arg0 models.Order, arg1 context.Context) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), arg0, arg1)
}

// DeleteExecutor mocks base method.
func (m *MockUseCase) DeleteExecutor(arg0 models.Order, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExecutor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExecutor indicates an expected call of DeleteExecutor.
func (mr *MockUseCaseMockRecorder) DeleteExecutor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExecutor", reflect.TypeOf((*MockUseCase)(nil).DeleteExecutor), arg0, arg1)
}

// DeleteOrder mocks base method.
func (m *MockUseCase) DeleteOrder(arg0 uint64, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockUseCaseMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockUseCase)(nil).DeleteOrder), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockUseCase) FindByID(arg0 uint64, arg1 context.Context) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUseCaseMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUseCase)(nil).FindByID), arg0, arg1)
}

// FindByUserID mocks base method.
func (m *MockUseCase) FindByUserID(arg0 uint64, arg1 context.Context) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", arg0, arg1)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID.
func (mr *MockUseCaseMockRecorder) FindByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockUseCase)(nil).FindByUserID), arg0, arg1)
}

// GetActualOrders mocks base method.
func (m *MockUseCase) GetActualOrders(arg0 context.Context) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActualOrders", arg0)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActualOrders indicates an expected call of GetActualOrders.
func (mr *MockUseCaseMockRecorder) GetActualOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActualOrders", reflect.TypeOf((*MockUseCase)(nil).GetActualOrders), arg0)
}

// GetArchiveOrders mocks base method.
func (m *MockUseCase) GetArchiveOrders(arg0 models.UserBasicInfo, arg1 context.Context) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchiveOrders", arg0, arg1)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArchiveOrders indicates an expected call of GetArchiveOrders.
func (mr *MockUseCaseMockRecorder) GetArchiveOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchiveOrders", reflect.TypeOf((*MockUseCase)(nil).GetArchiveOrders), arg0, arg1)
}

// SearchOrders mocks base method.
func (m *MockUseCase) SearchOrders(arg0 string, arg1 context.Context) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchOrders", arg0, arg1)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchOrders indicates an expected call of SearchOrders.
func (mr *MockUseCaseMockRecorder) SearchOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchOrders", reflect.TypeOf((*MockUseCase)(nil).SearchOrders), arg0, arg1)
}

// SelectExecutor mocks base method.
func (m *MockUseCase) SelectExecutor(arg0 models.Order, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectExecutor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectExecutor indicates an expected call of SelectExecutor.
func (mr *MockUseCaseMockRecorder) SelectExecutor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectExecutor", reflect.TypeOf((*MockUseCase)(nil).SelectExecutor), arg0, arg1)
}
