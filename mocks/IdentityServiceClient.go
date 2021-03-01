// Code generated by mockery v2.6.0. DO NOT EDIT.

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
