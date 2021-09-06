// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "ms.api/protos/pb/types"

	verifyService "ms.api/protos/pb/verifyService"
)

// VerifyServiceClient is an autogenerated mock type for the VerifyServiceClient type
type VerifyServiceClient struct {
	mock.Mock
}

// RequestOTP provides a mock function with given fields: ctx, in, opts
func (_m *VerifyServiceClient) RequestOTP(ctx context.Context, in *verifyService.RequestOTPRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *verifyService.RequestOTPRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *verifyService.RequestOTPRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateEmail provides a mock function with given fields: ctx, in, opts
func (_m *VerifyServiceClient) ValidateEmail(ctx context.Context, in *verifyService.ValidateEmailRequest, opts ...grpc.CallOption) (*verifyService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *verifyService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *verifyService.ValidateEmailRequest, ...grpc.CallOption) *verifyService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verifyService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *verifyService.ValidateEmailRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyOTP provides a mock function with given fields: ctx, in, opts
func (_m *VerifyServiceClient) VerifyOTP(ctx context.Context, in *verifyService.VerifyOTPRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *verifyService.VerifyOTPRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *verifyService.VerifyOTPRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
