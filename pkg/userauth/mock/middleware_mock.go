// Code generated by MockGen. DO NOT EDIT.
// Source: middleware.go

// Package userauthmock is a generated GoMock package.
package userauthmock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserAuthMiddlewareFactory is a mock of UserAuthMiddlewareFactory interface.
type MockUserAuthMiddlewareFactory struct {
	ctrl     *gomock.Controller
	recorder *MockUserAuthMiddlewareFactoryMockRecorder
}

// MockUserAuthMiddlewareFactoryMockRecorder is the mock recorder for MockUserAuthMiddlewareFactory.
type MockUserAuthMiddlewareFactoryMockRecorder struct {
	mock *MockUserAuthMiddlewareFactory
}

// NewMockUserAuthMiddlewareFactory creates a new mock instance.
func NewMockUserAuthMiddlewareFactory(ctrl *gomock.Controller) *MockUserAuthMiddlewareFactory {
	mock := &MockUserAuthMiddlewareFactory{ctrl: ctrl}
	mock.recorder = &MockUserAuthMiddlewareFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAuthMiddlewareFactory) EXPECT() *MockUserAuthMiddlewareFactoryMockRecorder {
	return m.recorder
}

// Make mocks base method.
func (m *MockUserAuthMiddlewareFactory) Make() func(http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Make")
	ret0, _ := ret[0].(func(http.Handler) http.Handler)
	return ret0
}

// Make indicates an expected call of Make.
func (mr *MockUserAuthMiddlewareFactoryMockRecorder) Make() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockUserAuthMiddlewareFactory)(nil).Make))
}