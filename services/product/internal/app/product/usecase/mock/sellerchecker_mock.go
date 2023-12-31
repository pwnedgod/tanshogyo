// Code generated by MockGen. DO NOT EDIT.
// Source: sellerchecker.go

// Package usecasemock is a generated GoMock package.
package usecasemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProductSellerChecker is a mock of ProductSellerChecker interface.
type MockProductSellerChecker struct {
	ctrl     *gomock.Controller
	recorder *MockProductSellerCheckerMockRecorder
}

// MockProductSellerCheckerMockRecorder is the mock recorder for MockProductSellerChecker.
type MockProductSellerCheckerMockRecorder struct {
	mock *MockProductSellerChecker
}

// NewMockProductSellerChecker creates a new mock instance.
func NewMockProductSellerChecker(ctrl *gomock.Controller) *MockProductSellerChecker {
	mock := &MockProductSellerChecker{ctrl: ctrl}
	mock.recorder = &MockProductSellerCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductSellerChecker) EXPECT() *MockProductSellerCheckerMockRecorder {
	return m.recorder
}

// CheckSeller mocks base method.
func (m *MockProductSellerChecker) CheckSeller(ctx context.Context, userId, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSeller", ctx, userId, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSeller indicates an expected call of CheckSeller.
func (mr *MockProductSellerCheckerMockRecorder) CheckSeller(ctx, userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSeller", reflect.TypeOf((*MockProductSellerChecker)(nil).CheckSeller), ctx, userId, id)
}
