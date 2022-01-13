// Code generated by MockGen. DO NOT EDIT.
// Source: ./libs/validator/email/email.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailValidator is a mock of EmailValidator interface.
type MockEmailValidator struct {
	ctrl     *gomock.Controller
	recorder *MockEmailValidatorMockRecorder
}

// MockEmailValidatorMockRecorder is the mock recorder for MockEmailValidator.
type MockEmailValidatorMockRecorder struct {
	mock *MockEmailValidator
}

// NewMockEmailValidator creates a new mock instance.
func NewMockEmailValidator(ctrl *gomock.Controller) *MockEmailValidator {
	mock := &MockEmailValidator{ctrl: ctrl}
	mock.recorder = &MockEmailValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailValidator) EXPECT() *MockEmailValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockEmailValidator) Validate(email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockEmailValidatorMockRecorder) Validate(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockEmailValidator)(nil).Validate), email)
}