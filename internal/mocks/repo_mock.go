// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-check-api/internal/repo (interfaces: CheckRepo,TestRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-check-api/internal/models"
)

// MockCheckRepo is a mock of CheckRepo interface.
type MockCheckRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCheckRepoMockRecorder
}

// MockCheckRepoMockRecorder is the mock recorder for MockCheckRepo.
type MockCheckRepoMockRecorder struct {
	mock *MockCheckRepo
}

// NewMockCheckRepo creates a new mock instance.
func NewMockCheckRepo(ctrl *gomock.Controller) *MockCheckRepo {
	mock := &MockCheckRepo{ctrl: ctrl}
	mock.recorder = &MockCheckRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckRepo) EXPECT() *MockCheckRepoMockRecorder {
	return m.recorder
}

// AddChecks mocks base method.
func (m *MockCheckRepo) AddChecks(arg0 []models.Check) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChecks", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChecks indicates an expected call of AddChecks.
func (mr *MockCheckRepoMockRecorder) AddChecks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChecks", reflect.TypeOf((*MockCheckRepo)(nil).AddChecks), arg0)
}

// DescribeCheck mocks base method.
func (m *MockCheckRepo) DescribeCheck(arg0 uint64) (*models.Check, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCheck", arg0)
	ret0, _ := ret[0].(*models.Check)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCheck indicates an expected call of DescribeCheck.
func (mr *MockCheckRepoMockRecorder) DescribeCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCheck", reflect.TypeOf((*MockCheckRepo)(nil).DescribeCheck), arg0)
}

// ListChecks mocks base method.
func (m *MockCheckRepo) ListChecks(arg0, arg1 uint64) ([]models.Check, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChecks", arg0, arg1)
	ret0, _ := ret[0].([]models.Check)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChecks indicates an expected call of ListChecks.
func (mr *MockCheckRepoMockRecorder) ListChecks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChecks", reflect.TypeOf((*MockCheckRepo)(nil).ListChecks), arg0, arg1)
}

// RemoveCheck mocks base method.
func (m *MockCheckRepo) RemoveCheck(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCheck", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCheck indicates an expected call of RemoveCheck.
func (mr *MockCheckRepoMockRecorder) RemoveCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCheck", reflect.TypeOf((*MockCheckRepo)(nil).RemoveCheck), arg0)
}

// MockTestRepo is a mock of TestRepo interface.
type MockTestRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTestRepoMockRecorder
}

// MockTestRepoMockRecorder is the mock recorder for MockTestRepo.
type MockTestRepoMockRecorder struct {
	mock *MockTestRepo
}

// NewMockTestRepo creates a new mock instance.
func NewMockTestRepo(ctrl *gomock.Controller) *MockTestRepo {
	mock := &MockTestRepo{ctrl: ctrl}
	mock.recorder = &MockTestRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestRepo) EXPECT() *MockTestRepoMockRecorder {
	return m.recorder
}

// AddTests mocks base method.
func (m *MockTestRepo) AddTests(arg0 []models.Test) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTests", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTests indicates an expected call of AddTests.
func (mr *MockTestRepoMockRecorder) AddTests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTests", reflect.TypeOf((*MockTestRepo)(nil).AddTests), arg0)
}

// DescribeTest mocks base method.
func (m *MockTestRepo) DescribeTest(arg0 uint64) (*models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTest", arg0)
	ret0, _ := ret[0].(*models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTest indicates an expected call of DescribeTest.
func (mr *MockTestRepoMockRecorder) DescribeTest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTest", reflect.TypeOf((*MockTestRepo)(nil).DescribeTest), arg0)
}

// ListTests mocks base method.
func (m *MockTestRepo) ListTests(arg0, arg1 uint64) ([]models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTests", arg0, arg1)
	ret0, _ := ret[0].([]models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTests indicates an expected call of ListTests.
func (mr *MockTestRepoMockRecorder) ListTests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTests", reflect.TypeOf((*MockTestRepo)(nil).ListTests), arg0, arg1)
}

// RemoveTest mocks base method.
func (m *MockTestRepo) RemoveTest(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTest indicates an expected call of RemoveTest.
func (mr *MockTestRepoMockRecorder) RemoveTest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTest", reflect.TypeOf((*MockTestRepo)(nil).RemoveTest), arg0)
}