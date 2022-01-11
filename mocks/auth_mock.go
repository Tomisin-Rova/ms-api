// Code generated by MockGen. DO NOT EDIT.
// Source: ./protos/pb/auth/auth.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	auth "ms.api/protos/pb/auth"
)

// MockAuthServiceClient is a mock of AuthServiceClient interface.
type MockAuthServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceClientMockRecorder
}

// MockAuthServiceClientMockRecorder is the mock recorder for MockAuthServiceClient.
type MockAuthServiceClientMockRecorder struct {
	mock *MockAuthServiceClient
}

// NewMockAuthServiceClient creates a new mock instance.
func NewMockAuthServiceClient(ctrl *gomock.Controller) *MockAuthServiceClient {
	mock := &MockAuthServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceClient) EXPECT() *MockAuthServiceClientMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthServiceClient) Login(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceClientMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServiceClient)(nil).Login), varargs...)
}

// RefreshToken mocks base method.
func (m *MockAuthServiceClient) RefreshToken(ctx context.Context, in *auth.RefreshTokenRequest, opts ...grpc.CallOption) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RefreshToken", varargs...)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthServiceClientMockRecorder) RefreshToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthServiceClient)(nil).RefreshToken), varargs...)
}

// Signup mocks base method.
func (m *MockAuthServiceClient) Signup(ctx context.Context, in *auth.SignupRequest, opts ...grpc.CallOption) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Signup", varargs...)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthServiceClientMockRecorder) Signup(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuthServiceClient)(nil).Signup), varargs...)
}

// StaffLogin mocks base method.
func (m *MockAuthServiceClient) StaffLogin(ctx context.Context, in *auth.StaffLoginRequest, opts ...grpc.CallOption) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StaffLogin", varargs...)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StaffLogin indicates an expected call of StaffLogin.
func (mr *MockAuthServiceClientMockRecorder) StaffLogin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaffLogin", reflect.TypeOf((*MockAuthServiceClient)(nil).StaffLogin), varargs...)
}

// ValidateToken mocks base method.
func (m *MockAuthServiceClient) ValidateToken(ctx context.Context, in *auth.ValidateTokenRequest, opts ...grpc.CallOption) (*auth.ValidateTokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateToken", varargs...)
	ret0, _ := ret[0].(*auth.ValidateTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthServiceClientMockRecorder) ValidateToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthServiceClient)(nil).ValidateToken), varargs...)
}

// MockAuthServiceServer is a mock of AuthServiceServer interface.
type MockAuthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceServerMockRecorder
}

// MockAuthServiceServerMockRecorder is the mock recorder for MockAuthServiceServer.
type MockAuthServiceServerMockRecorder struct {
	mock *MockAuthServiceServer
}

// NewMockAuthServiceServer creates a new mock instance.
func NewMockAuthServiceServer(ctrl *gomock.Controller) *MockAuthServiceServer {
	mock := &MockAuthServiceServer{ctrl: ctrl}
	mock.recorder = &MockAuthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceServer) EXPECT() *MockAuthServiceServerMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthServiceServer) Login(arg0 context.Context, arg1 *auth.LoginRequest) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceServerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServiceServer)(nil).Login), arg0, arg1)
}

// RefreshToken mocks base method.
func (m *MockAuthServiceServer) RefreshToken(arg0 context.Context, arg1 *auth.RefreshTokenRequest) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", arg0, arg1)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthServiceServerMockRecorder) RefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthServiceServer)(nil).RefreshToken), arg0, arg1)
}

// Signup mocks base method.
func (m *MockAuthServiceServer) Signup(arg0 context.Context, arg1 *auth.SignupRequest) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", arg0, arg1)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthServiceServerMockRecorder) Signup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuthServiceServer)(nil).Signup), arg0, arg1)
}

// StaffLogin mocks base method.
func (m *MockAuthServiceServer) StaffLogin(arg0 context.Context, arg1 *auth.StaffLoginRequest) (*auth.TokenPairResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaffLogin", arg0, arg1)
	ret0, _ := ret[0].(*auth.TokenPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StaffLogin indicates an expected call of StaffLogin.
func (mr *MockAuthServiceServerMockRecorder) StaffLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaffLogin", reflect.TypeOf((*MockAuthServiceServer)(nil).StaffLogin), arg0, arg1)
}

// ValidateToken mocks base method.
func (m *MockAuthServiceServer) ValidateToken(arg0 context.Context, arg1 *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", arg0, arg1)
	ret0, _ := ret[0].(*auth.ValidateTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthServiceServerMockRecorder) ValidateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthServiceServer)(nil).ValidateToken), arg0, arg1)
}
