// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	cddService "ms.api/protos/pb/cddService"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "ms.api/protos/pb/types"
)

// CddServiceClient is an autogenerated mock type for the CddServiceClient type
type CddServiceClient struct {
	mock.Mock
}

// GetCDDById provides a mock function with given fields: ctx, in, opts
func (_m *CddServiceClient) GetCDDById(ctx context.Context, in *cddService.CddIdRequest, opts ...grpc.CallOption) (*types.Cdd, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Cdd
	if rf, ok := ret.Get(0).(func(context.Context, *cddService.CddIdRequest, ...grpc.CallOption) *types.Cdd); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Cdd)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cddService.CddIdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCDDByOwner provides a mock function with given fields: ctx, in, opts
func (_m *CddServiceClient) GetCDDByOwner(ctx context.Context, in *cddService.GetCDDByOwnerRequest, opts ...grpc.CallOption) (*types.Cdd, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Cdd
	if rf, ok := ret.Get(0).(func(context.Context, *cddService.GetCDDByOwnerRequest, ...grpc.CallOption) *types.Cdd); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Cdd)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cddService.GetCDDByOwnerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCDDSummaryReport provides a mock function with given fields: ctx, in, opts
func (_m *CddServiceClient) GetCDDSummaryReport(ctx context.Context, in *cddService.PersonIdRequest, opts ...grpc.CallOption) (*cddService.Cddsummary, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cddService.Cddsummary
	if rf, ok := ret.Get(0).(func(context.Context, *cddService.PersonIdRequest, ...grpc.CallOption) *cddService.Cddsummary); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cddService.Cddsummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cddService.PersonIdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}