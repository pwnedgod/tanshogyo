// Code generated by MockGen. DO NOT EDIT.
// Source: authedlister.go

// Package usecasemock is a generated GoMock package.
package usecasemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	usecase "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
)

// MockProductAuthedLister is a mock of ProductAuthedLister interface.
type MockProductAuthedLister struct {
	ctrl     *gomock.Controller
	recorder *MockProductAuthedListerMockRecorder
}

// MockProductAuthedListerMockRecorder is the mock recorder for MockProductAuthedLister.
type MockProductAuthedListerMockRecorder struct {
	mock *MockProductAuthedLister
}

// NewMockProductAuthedLister creates a new mock instance.
func NewMockProductAuthedLister(ctrl *gomock.Controller) *MockProductAuthedLister {
	mock := &MockProductAuthedLister{ctrl: ctrl}
	mock.recorder = &MockProductAuthedListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductAuthedLister) EXPECT() *MockProductAuthedListerMockRecorder {
	return m.recorder
}

// AuthedList mocks base method.
func (m *MockProductAuthedLister) AuthedList(ctx context.Context, userId string, limit, offset int) (usecase.ProductList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthedList", ctx, userId, limit, offset)
	ret0, _ := ret[0].(usecase.ProductList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthedList indicates an expected call of AuthedList.
func (mr *MockProductAuthedListerMockRecorder) AuthedList(ctx, userId, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthedList", reflect.TypeOf((*MockProductAuthedLister)(nil).AuthedList), ctx, userId, limit, offset)
}
