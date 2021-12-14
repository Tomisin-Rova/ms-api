// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "ms.api/protos/pb/types"

	verification "ms.api/protos/pb/verification"
)

// VerificationServiceClient is an autogenerated mock type for the VerificationServiceClient type
type VerificationServiceClient struct {
	mock.Mock
}

// RequestOTP provides a mock function with given fields: ctx, in, opts
func (_m *VerificationServiceClient) RequestOTP(ctx context.Context, in *verification.RequestOTPRequest, opts ...grpc.CallOption) (*types.DefaultResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.DefaultResponse
	if rf, ok := ret.Get(0).(func(context.Context, *verification.RequestOTPRequest, ...grpc.CallOption) *types.DefaultResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.DefaultResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *verification.RequestOTPRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyOTP provides a mock function with given fields: ctx, in, opts
func (_m *VerificationServiceClient) VerifyOTP(ctx context.Context, in *verification.VerifyOTPRequest, opts ...grpc.CallOption) (*types.DefaultResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.DefaultResponse
	if rf, ok := ret.Get(0).(func(context.Context, *verification.VerifyOTPRequest, ...grpc.CallOption) *types.DefaultResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.DefaultResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *verification.VerifyOTPRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
