// Code generated by MockGen. DO NOT EDIT.
// Source: ./libs/validator/device/device_validator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "ms.api/types"
)

// MockDeviceValidator is a mock of DeviceValidator interface.
type MockDeviceValidator struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceValidatorMockRecorder
}

// MockDeviceValidatorMockRecorder is the mock recorder for MockDeviceValidator.
type MockDeviceValidatorMockRecorder struct {
	mock *MockDeviceValidator
}

// NewMockDeviceValidator creates a new mock instance.
func NewMockDeviceValidator(ctrl *gomock.Controller) *MockDeviceValidator {
	mock := &MockDeviceValidator{ctrl: ctrl}
	mock.recorder = &MockDeviceValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceValidator) EXPECT() *MockDeviceValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockDeviceValidator) Validate(deviceInput *types.DeviceInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", deviceInput)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockDeviceValidatorMockRecorder) Validate(deviceInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockDeviceValidator)(nil).Validate), deviceInput)
}
