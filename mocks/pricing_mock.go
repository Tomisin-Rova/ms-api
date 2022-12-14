// Code generated by MockGen. DO NOT EDIT.
// Source: ./protos/pb/pricing/pricing.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	pricing "ms.api/protos/pb/pricing"
	types "ms.api/protos/pb/types"
)

// MockPricingServiceClient is a mock of PricingServiceClient interface.
type MockPricingServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockPricingServiceClientMockRecorder
}

// MockPricingServiceClientMockRecorder is the mock recorder for MockPricingServiceClient.
type MockPricingServiceClientMockRecorder struct {
	mock *MockPricingServiceClient
}

// NewMockPricingServiceClient creates a new mock instance.
func NewMockPricingServiceClient(ctrl *gomock.Controller) *MockPricingServiceClient {
	mock := &MockPricingServiceClient{ctrl: ctrl}
	mock.recorder = &MockPricingServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPricingServiceClient) EXPECT() *MockPricingServiceClientMockRecorder {
	return m.recorder
}

// GetCurrencies mocks base method.
func (m *MockPricingServiceClient) GetCurrencies(ctx context.Context, in *pricing.GetCurrenciesRequest, opts ...grpc.CallOption) (*pricing.GetCurrenciesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCurrencies", varargs...)
	ret0, _ := ret[0].(*pricing.GetCurrenciesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencies indicates an expected call of GetCurrencies.
func (mr *MockPricingServiceClientMockRecorder) GetCurrencies(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencies", reflect.TypeOf((*MockPricingServiceClient)(nil).GetCurrencies), varargs...)
}

// GetCurrency mocks base method.
func (m *MockPricingServiceClient) GetCurrency(ctx context.Context, in *pricing.GetCurrencyRequest, opts ...grpc.CallOption) (*types.Currency, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCurrency", varargs...)
	ret0, _ := ret[0].(*types.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrency indicates an expected call of GetCurrency.
func (mr *MockPricingServiceClientMockRecorder) GetCurrency(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrency", reflect.TypeOf((*MockPricingServiceClient)(nil).GetCurrency), varargs...)
}

// GetExchangeRate mocks base method.
func (m *MockPricingServiceClient) GetExchangeRate(ctx context.Context, in *pricing.GetExchangeRateRequest, opts ...grpc.CallOption) (*pricing.GetExchangeRateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExchangeRate", varargs...)
	ret0, _ := ret[0].(*pricing.GetExchangeRateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate.
func (mr *MockPricingServiceClientMockRecorder) GetExchangeRate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockPricingServiceClient)(nil).GetExchangeRate), varargs...)
}

// GetFees mocks base method.
func (m *MockPricingServiceClient) GetFees(ctx context.Context, in *pricing.GetFeesRequest, opts ...grpc.CallOption) (*pricing.GetFeesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFees", varargs...)
	ret0, _ := ret[0].(*pricing.GetFeesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFees indicates an expected call of GetFees.
func (mr *MockPricingServiceClientMockRecorder) GetFees(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFees", reflect.TypeOf((*MockPricingServiceClient)(nil).GetFees), varargs...)
}

// UpdateFX mocks base method.
func (m *MockPricingServiceClient) UpdateFX(ctx context.Context, in *pricing.UpdateFXRequest, opts ...grpc.CallOption) (*types.DefaultResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFX", varargs...)
	ret0, _ := ret[0].(*types.DefaultResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFX indicates an expected call of UpdateFX.
func (mr *MockPricingServiceClientMockRecorder) UpdateFX(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFX", reflect.TypeOf((*MockPricingServiceClient)(nil).UpdateFX), varargs...)
}

// UpdateFees mocks base method.
func (m *MockPricingServiceClient) UpdateFees(ctx context.Context, in *pricing.UpdateFeesRequests, opts ...grpc.CallOption) (*types.DefaultResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFees", varargs...)
	ret0, _ := ret[0].(*types.DefaultResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFees indicates an expected call of UpdateFees.
func (mr *MockPricingServiceClientMockRecorder) UpdateFees(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFees", reflect.TypeOf((*MockPricingServiceClient)(nil).UpdateFees), varargs...)
}

// MockPricingServiceServer is a mock of PricingServiceServer interface.
type MockPricingServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockPricingServiceServerMockRecorder
}

// MockPricingServiceServerMockRecorder is the mock recorder for MockPricingServiceServer.
type MockPricingServiceServerMockRecorder struct {
	mock *MockPricingServiceServer
}

// NewMockPricingServiceServer creates a new mock instance.
func NewMockPricingServiceServer(ctrl *gomock.Controller) *MockPricingServiceServer {
	mock := &MockPricingServiceServer{ctrl: ctrl}
	mock.recorder = &MockPricingServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPricingServiceServer) EXPECT() *MockPricingServiceServerMockRecorder {
	return m.recorder
}

// GetCurrencies mocks base method.
func (m *MockPricingServiceServer) GetCurrencies(arg0 context.Context, arg1 *pricing.GetCurrenciesRequest) (*pricing.GetCurrenciesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencies", arg0, arg1)
	ret0, _ := ret[0].(*pricing.GetCurrenciesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencies indicates an expected call of GetCurrencies.
func (mr *MockPricingServiceServerMockRecorder) GetCurrencies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencies", reflect.TypeOf((*MockPricingServiceServer)(nil).GetCurrencies), arg0, arg1)
}

// GetCurrency mocks base method.
func (m *MockPricingServiceServer) GetCurrency(arg0 context.Context, arg1 *pricing.GetCurrencyRequest) (*types.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrency", arg0, arg1)
	ret0, _ := ret[0].(*types.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrency indicates an expected call of GetCurrency.
func (mr *MockPricingServiceServerMockRecorder) GetCurrency(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrency", reflect.TypeOf((*MockPricingServiceServer)(nil).GetCurrency), arg0, arg1)
}

// GetExchangeRate mocks base method.
func (m *MockPricingServiceServer) GetExchangeRate(arg0 context.Context, arg1 *pricing.GetExchangeRateRequest) (*pricing.GetExchangeRateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", arg0, arg1)
	ret0, _ := ret[0].(*pricing.GetExchangeRateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate.
func (mr *MockPricingServiceServerMockRecorder) GetExchangeRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockPricingServiceServer)(nil).GetExchangeRate), arg0, arg1)
}

// GetFees mocks base method.
func (m *MockPricingServiceServer) GetFees(arg0 context.Context, arg1 *pricing.GetFeesRequest) (*pricing.GetFeesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFees", arg0, arg1)
	ret0, _ := ret[0].(*pricing.GetFeesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFees indicates an expected call of GetFees.
func (mr *MockPricingServiceServerMockRecorder) GetFees(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFees", reflect.TypeOf((*MockPricingServiceServer)(nil).GetFees), arg0, arg1)
}

// UpdateFX mocks base method.
func (m *MockPricingServiceServer) UpdateFX(arg0 context.Context, arg1 *pricing.UpdateFXRequest) (*types.DefaultResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFX", arg0, arg1)
	ret0, _ := ret[0].(*types.DefaultResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFX indicates an expected call of UpdateFX.
func (mr *MockPricingServiceServerMockRecorder) UpdateFX(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFX", reflect.TypeOf((*MockPricingServiceServer)(nil).UpdateFX), arg0, arg1)
}

// UpdateFees mocks base method.
func (m *MockPricingServiceServer) UpdateFees(arg0 context.Context, arg1 *pricing.UpdateFeesRequests) (*types.DefaultResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFees", arg0, arg1)
	ret0, _ := ret[0].(*types.DefaultResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFees indicates an expected call of UpdateFees.
func (mr *MockPricingServiceServerMockRecorder) UpdateFees(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFees", reflect.TypeOf((*MockPricingServiceServer)(nil).UpdateFees), arg0, arg1)
}
