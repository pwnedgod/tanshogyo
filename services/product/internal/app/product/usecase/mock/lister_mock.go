// Code generated by MockGen. DO NOT EDIT.
// Source: lister.go

// Package usecasemock is a generated GoMock package.
package usecasemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	usecase "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
)

// MockProductLister is a mock of ProductLister interface.
type MockProductLister struct {
	ctrl     *gomock.Controller
	recorder *MockProductListerMockRecorder
}

// MockProductListerMockRecorder is the mock recorder for MockProductLister.
type MockProductListerMockRecorder struct {
	mock *MockProductLister
}

// NewMockProductLister creates a new mock instance.
func NewMockProductLister(ctrl *gomock.Controller) *MockProductLister {
	mock := &MockProductLister{ctrl: ctrl}
	mock.recorder = &MockProductListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductLister) EXPECT() *MockProductListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockProductLister) List(ctx context.Context, limit, offset int) (usecase.ProductList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset)
	ret0, _ := ret[0].(usecase.ProductList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductListerMockRecorder) List(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductLister)(nil).List), ctx, limit, offset)
}