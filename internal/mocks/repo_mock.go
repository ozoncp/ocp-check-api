// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-check-api/internal/repo (interfaces: CheckRepo,TestRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
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

// CreateCheck mocks base method.
func (m *MockCheckRepo) CreateCheck(arg0 context.Context, arg1 models.Check) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCheck", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCheck indicates an expected call of CreateCheck.
func (mr *MockCheckRepoMockRecorder) CreateCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCheck", reflect.TypeOf((*MockCheckRepo)(nil).CreateCheck), arg0, arg1)
}

// DescribeCheck mocks base method.
func (m *MockCheckRepo) DescribeCheck(arg0 context.Context, arg1 uint64) (*models.Check, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCheck", arg0, arg1)
	ret0, _ := ret[0].(*models.Check)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCheck indicates an expected call of DescribeCheck.
func (mr *MockCheckRepoMockRecorder) DescribeCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCheck", reflect.TypeOf((*MockCheckRepo)(nil).DescribeCheck), arg0, arg1)
}

// ListChecks mocks base method.
func (m *MockCheckRepo) ListChecks(arg0 context.Context, arg1, arg2 uint64) ([]models.Check, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChecks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Check)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChecks indicates an expected call of ListChecks.
func (mr *MockCheckRepoMockRecorder) ListChecks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChecks", reflect.TypeOf((*MockCheckRepo)(nil).ListChecks), arg0, arg1, arg2)
}

// MultiCreateCheck mocks base method.
func (m *MockCheckRepo) MultiCreateCheck(arg0 context.Context, arg1 []models.Check) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiCreateCheck", arg0, arg1)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiCreateCheck indicates an expected call of MultiCreateCheck.
func (mr *MockCheckRepoMockRecorder) MultiCreateCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiCreateCheck", reflect.TypeOf((*MockCheckRepo)(nil).MultiCreateCheck), arg0, arg1)
}

// RemoveCheck mocks base method.
func (m *MockCheckRepo) RemoveCheck(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCheck indicates an expected call of RemoveCheck.
func (mr *MockCheckRepoMockRecorder) RemoveCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCheck", reflect.TypeOf((*MockCheckRepo)(nil).RemoveCheck), arg0, arg1)
}

// UpdateCheck mocks base method.
func (m *MockCheckRepo) UpdateCheck(arg0 context.Context, arg1 models.Check) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCheck", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCheck indicates an expected call of UpdateCheck.
func (mr *MockCheckRepoMockRecorder) UpdateCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCheck", reflect.TypeOf((*MockCheckRepo)(nil).UpdateCheck), arg0, arg1)
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

// CreateTest mocks base method.
func (m *MockTestRepo) CreateTest(arg0 context.Context, arg1 models.Test) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTest", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTest indicates an expected call of CreateTest.
func (mr *MockTestRepoMockRecorder) CreateTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTest", reflect.TypeOf((*MockTestRepo)(nil).CreateTest), arg0, arg1)
}

// DescribeTest mocks base method.
func (m *MockTestRepo) DescribeTest(arg0 context.Context, arg1 uint64) (*models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTest", arg0, arg1)
	ret0, _ := ret[0].(*models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTest indicates an expected call of DescribeTest.
func (mr *MockTestRepoMockRecorder) DescribeTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTest", reflect.TypeOf((*MockTestRepo)(nil).DescribeTest), arg0, arg1)
}

// ListTests mocks base method.
func (m *MockTestRepo) ListTests(arg0 context.Context, arg1, arg2 uint64) ([]models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTests", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTests indicates an expected call of ListTests.
func (mr *MockTestRepoMockRecorder) ListTests(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTests", reflect.TypeOf((*MockTestRepo)(nil).ListTests), arg0, arg1, arg2)
}

// MultiCreateTest mocks base method.
func (m *MockTestRepo) MultiCreateTest(arg0 context.Context, arg1 []models.Test) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiCreateTest", arg0, arg1)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiCreateTest indicates an expected call of MultiCreateTest.
func (mr *MockTestRepoMockRecorder) MultiCreateTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiCreateTest", reflect.TypeOf((*MockTestRepo)(nil).MultiCreateTest), arg0, arg1)
}

// RemoveTest mocks base method.
func (m *MockTestRepo) RemoveTest(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTest indicates an expected call of RemoveTest.
func (mr *MockTestRepoMockRecorder) RemoveTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTest", reflect.TypeOf((*MockTestRepo)(nil).RemoveTest), arg0, arg1)
}

// UpdateTest mocks base method.
func (m *MockTestRepo) UpdateTest(arg0 context.Context, arg1 models.Test) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTest", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTest indicates an expected call of UpdateTest.
func (mr *MockTestRepoMockRecorder) UpdateTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTest", reflect.TypeOf((*MockTestRepo)(nil).UpdateTest), arg0, arg1)
}
