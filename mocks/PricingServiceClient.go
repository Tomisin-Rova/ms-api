// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pricing "ms.api/protos/pb/pricing"

	types "ms.api/protos/pb/types"
)

// PricingServiceClient is an autogenerated mock type for the PricingServiceClient type
type PricingServiceClient struct {
	mock.Mock
}

// GetCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *PricingServiceClient) GetCurrencies(ctx context.Context, in *pricing.GetCurrenciesRequest, opts ...grpc.CallOption) (*pricing.GetCurrenciesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pricing.GetCurrenciesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pricing.GetCurrenciesRequest, ...grpc.CallOption) *pricing.GetCurrenciesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pricing.GetCurrenciesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pricing.GetCurrenciesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrency provides a mock function with given fields: ctx, in, opts
func (_m *PricingServiceClient) GetCurrency(ctx context.Context, in *pricing.GetCurrencyRequest, opts ...grpc.CallOption) (*types.Currency, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Currency
	if rf, ok := ret.Get(0).(func(context.Context, *pricing.GetCurrencyRequest, ...grpc.CallOption) *types.Currency); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Currency)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pricing.GetCurrencyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExchangeRate provides a mock function with given fields: ctx, in, opts
func (_m *PricingServiceClient) GetExchangeRate(ctx context.Context, in *pricing.GetExchangeRateRequest, opts ...grpc.CallOption) (*pricing.GetExchangeRateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pricing.GetExchangeRateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pricing.GetExchangeRateRequest, ...grpc.CallOption) *pricing.GetExchangeRateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pricing.GetExchangeRateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pricing.GetExchangeRateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFees provides a mock function with given fields: ctx, in, opts
func (_m *PricingServiceClient) GetFees(ctx context.Context, in *pricing.GetFeesRequest, opts ...grpc.CallOption) (*pricing.GetFeesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pricing.GetFeesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pricing.GetFeesRequest, ...grpc.CallOption) *pricing.GetFeesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pricing.GetFeesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pricing.GetFeesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
