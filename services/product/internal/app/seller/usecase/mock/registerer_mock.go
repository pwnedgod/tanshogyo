// Code generated by MockGen. DO NOT EDIT.
// Source: registerer.go

// Package usecasemock is a generated GoMock package.
package usecasemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	usecase "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
)

// MockSellerRegisterer is a mock of SellerRegisterer interface.
type MockSellerRegisterer struct {
	ctrl     *gomock.Controller
	recorder *MockSellerRegistererMockRecorder
}

// MockSellerRegistererMockRecorder is the mock recorder for MockSellerRegisterer.
type MockSellerRegistererMockRecorder struct {
	mock *MockSellerRegisterer
}

// NewMockSellerRegisterer creates a new mock instance.
func NewMockSellerRegisterer(ctrl *gomock.Controller) *MockSellerRegisterer {
	mock := &MockSellerRegisterer{ctrl: ctrl}
	mock.recorder = &MockSellerRegistererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSellerRegisterer) EXPECT() *MockSellerRegistererMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockSellerRegisterer) Register(ctx context.Context, userId string, form usecase.SellerForm) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, userId, form)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockSellerRegistererMockRecorder) Register(ctx, userId, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSellerRegisterer)(nil).Register), ctx, userId, form)
}
