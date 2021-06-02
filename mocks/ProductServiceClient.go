// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	productService "ms.api/protos/pb/productService"

	types "ms.api/protos/pb/types"
)

// ProductServiceClient is an autogenerated mock type for the ProductServiceClient type
type ProductServiceClient struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, in, opts
func (_m *ProductServiceClient) CreateAccount(ctx context.Context, in *productService.CreateAccountRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *productService.CreateAccountRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *productService.CreateAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccounts provides a mock function with given fields: ctx, in, opts
func (_m *ProductServiceClient) GetAccounts(ctx context.Context, in *productService.GetAccountRequest, opts ...grpc.CallOption) (*productService.GetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *productService.GetAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *productService.GetAccountRequest, ...grpc.CallOption) *productService.GetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*productService.GetAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *productService.GetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
