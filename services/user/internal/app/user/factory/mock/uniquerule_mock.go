// Code generated by MockGen. DO NOT EDIT.
// Source: uniquerule.go

// Package factorymock is a generated GoMock package.
package factorymock

import (
	reflect "reflect"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	gomock "github.com/golang/mock/gomock"
	factory "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/factory"
)

// MockUserUniqueRuleFactory is a mock of UserUniqueRuleFactory interface.
type MockUserUniqueRuleFactory struct {
	ctrl     *gomock.Controller
	recorder *MockUserUniqueRuleFactoryMockRecorder
}

// MockUserUniqueRuleFactoryMockRecorder is the mock recorder for MockUserUniqueRuleFactory.
type MockUserUniqueRuleFactoryMockRecorder struct {
	mock *MockUserUniqueRuleFactory
}

// NewMockUserUniqueRuleFactory creates a new mock instance.
func NewMockUserUniqueRuleFactory(ctrl *gomock.Controller) *MockUserUniqueRuleFactory {
	mock := &MockUserUniqueRuleFactory{ctrl: ctrl}
	mock.recorder = &MockUserUniqueRuleFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUniqueRuleFactory) EXPECT() *MockUserUniqueRuleFactoryMockRecorder {
	return m.recorder
}

// Make mocks base method.
func (m *MockUserUniqueRuleFactory) Make(excludedId, fieldName string, otherFields ...factory.Field) validation.Rule {
	m.ctrl.T.Helper()
	varargs := []interface{}{excludedId, fieldName}
	for _, a := range otherFields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Make", varargs...)
	ret0, _ := ret[0].(validation.Rule)
	return ret0
}

// Make indicates an expected call of Make.
func (mr *MockUserUniqueRuleFactoryMockRecorder) Make(excludedId, fieldName interface{}, otherFields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{excludedId, fieldName}, otherFields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockUserUniqueRuleFactory)(nil).Make), varargs...)
}
