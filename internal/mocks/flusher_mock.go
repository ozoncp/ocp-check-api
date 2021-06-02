// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-check-api/internal/flusher (interfaces: CheckFlusher,TestFlusher)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-check-api/internal/models"
)

// MockCheckFlusher is a mock of CheckFlusher interface.
type MockCheckFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockCheckFlusherMockRecorder
}

// MockCheckFlusherMockRecorder is the mock recorder for MockCheckFlusher.
type MockCheckFlusherMockRecorder struct {
	mock *MockCheckFlusher
}

// NewMockCheckFlusher creates a new mock instance.
func NewMockCheckFlusher(ctrl *gomock.Controller) *MockCheckFlusher {
	mock := &MockCheckFlusher{ctrl: ctrl}
	mock.recorder = &MockCheckFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckFlusher) EXPECT() *MockCheckFlusherMockRecorder {
	return m.recorder
}

// Flush mocks base method.
func (m *MockCheckFlusher) Flush(arg0 []models.Check) []models.Check {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].([]models.Check)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockCheckFlusherMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockCheckFlusher)(nil).Flush), arg0)
}

// MockTestFlusher is a mock of TestFlusher interface.
type MockTestFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockTestFlusherMockRecorder
}

// MockTestFlusherMockRecorder is the mock recorder for MockTestFlusher.
type MockTestFlusherMockRecorder struct {
	mock *MockTestFlusher
}

// NewMockTestFlusher creates a new mock instance.
func NewMockTestFlusher(ctrl *gomock.Controller) *MockTestFlusher {
	mock := &MockTestFlusher{ctrl: ctrl}
	mock.recorder = &MockTestFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestFlusher) EXPECT() *MockTestFlusherMockRecorder {
	return m.recorder
}

// Flush mocks base method.
func (m *MockTestFlusher) Flush(arg0 []models.Test) []models.Test {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].([]models.Test)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockTestFlusherMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockTestFlusher)(nil).Flush), arg0)
}
