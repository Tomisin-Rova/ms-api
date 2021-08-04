// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	identityService "ms.api/protos/pb/identityService"

	mock "github.com/stretchr/testify/mock"

	types "ms.api/protos/pb/types"
)

// IdentityServiceClient is an autogenerated mock type for the IdentityServiceClient type
type IdentityServiceClient struct {
	mock.Mock
}

// CreateTransactionPassword provides a mock function with given fields: ctx, in, opts
func (_m *IdentityServiceClient) CreateTransactionPassword(ctx context.Context, in *identityService.CreateTransactionPasswordRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *identityService.CreateTransactionPasswordRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *identityService.CreateTransactionPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestResetTransactionPassword provides a mock function with given fields: ctx, in, opts
func (_m *IdentityServiceClient) RequestResetTransactionPassword(ctx context.Context, in *identityService.RequestResetTransactionPasswordRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *identityService.RequestResetTransactionPasswordRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *identityService.RequestResetTransactionPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetTransactionPassword provides a mock function with given fields: ctx, in, opts
func (_m *IdentityServiceClient) ResetTransactionPassword(ctx context.Context, in *identityService.ResetTransactionPasswordRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *identityService.ResetTransactionPasswordRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *identityService.ResetTransactionPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeviceTokens provides a mock function with given fields: ctx, in, opts
func (_m *IdentityServiceClient) UpdateDeviceTokens(ctx context.Context, in *identityService.UpdateDeviceTokensRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *identityService.UpdateDeviceTokensRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *identityService.UpdateDeviceTokensRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
