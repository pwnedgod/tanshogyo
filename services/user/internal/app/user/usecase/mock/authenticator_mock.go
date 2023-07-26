// Code generated by MockGen. DO NOT EDIT.
// Source: authenticator.go

// Package usecasemock is a generated GoMock package.
package usecasemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	usecase "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
)

// MockUserAuthenticator is a mock of UserAuthenticator interface.
type MockUserAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockUserAuthenticatorMockRecorder
}

// MockUserAuthenticatorMockRecorder is the mock recorder for MockUserAuthenticator.
type MockUserAuthenticatorMockRecorder struct {
	mock *MockUserAuthenticator
}

// NewMockUserAuthenticator creates a new mock instance.
func NewMockUserAuthenticator(ctrl *gomock.Controller) *MockUserAuthenticator {
	mock := &MockUserAuthenticator{ctrl: ctrl}
	mock.recorder = &MockUserAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAuthenticator) EXPECT() *MockUserAuthenticatorMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockUserAuthenticator) Authenticate(ctx context.Context, token string) (usecase.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, token)
	ret0, _ := ret[0].(usecase.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockUserAuthenticatorMockRecorder) Authenticate(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockUserAuthenticator)(nil).Authenticate), ctx, token)
}
