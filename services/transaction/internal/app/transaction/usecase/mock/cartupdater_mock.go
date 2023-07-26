// Code generated by MockGen. DO NOT EDIT.
// Source: cartupdater.go

// Package usecasemock is a generated GoMock package.
package usecasemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	usecase "github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/usecase"
)

// MockTransactionCartUpdater is a mock of TransactionCartUpdater interface.
type MockTransactionCartUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionCartUpdaterMockRecorder
}

// MockTransactionCartUpdaterMockRecorder is the mock recorder for MockTransactionCartUpdater.
type MockTransactionCartUpdaterMockRecorder struct {
	mock *MockTransactionCartUpdater
}

// NewMockTransactionCartUpdater creates a new mock instance.
func NewMockTransactionCartUpdater(ctrl *gomock.Controller) *MockTransactionCartUpdater {
	mock := &MockTransactionCartUpdater{ctrl: ctrl}
	mock.recorder = &MockTransactionCartUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionCartUpdater) EXPECT() *MockTransactionCartUpdaterMockRecorder {
	return m.recorder
}

// UpdateCart mocks base method.
func (m *MockTransactionCartUpdater) UpdateCart(ctx context.Context, userId string, cart usecase.Cart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, userId, cart)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockTransactionCartUpdaterMockRecorder) UpdateCart(ctx, userId, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockTransactionCartUpdater)(nil).UpdateCart), ctx, userId, cart)
}