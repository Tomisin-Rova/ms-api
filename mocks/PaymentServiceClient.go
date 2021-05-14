// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	paymentService "ms.api/protos/pb/paymentService"

	types "ms.api/protos/pb/types"
)

// PaymentServiceClient is an autogenerated mock type for the PaymentServiceClient type
type PaymentServiceClient struct {
	mock.Mock
}

// AddPayeeAccount provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) AddPayeeAccount(ctx context.Context, in *paymentService.AddPayeeAccountRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.AddPayeeAccountRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.AddPayeeAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePayee provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) CreatePayee(ctx context.Context, in *paymentService.CreatePayeeRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.CreatePayeeRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.CreatePayeeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePayment provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) CreatePayment(ctx context.Context, in *paymentService.CreatePaymentRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.CreatePaymentRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.CreatePaymentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePayeeAccount provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) DeletePayeeAccount(ctx context.Context, in *paymentService.DeletePayeeAccountRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.DeletePayeeAccountRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.DeletePayeeAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPayee provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) GetPayee(ctx context.Context, in *paymentService.GetPayeeRequest, opts ...grpc.CallOption) (*types.Payee, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Payee
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.GetPayeeRequest, ...grpc.CallOption) *types.Payee); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Payee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.GetPayeeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPayees provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) GetPayees(ctx context.Context, in *paymentService.GetPayeesRequest, opts ...grpc.CallOption) (*types.Payees, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Payees
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.GetPayeesRequest, ...grpc.CallOption) *types.Payees); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Payees)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.GetPayeesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeTransfer provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) MakeTransfer(ctx context.Context, in *paymentService.TransferRequest, opts ...grpc.CallOption) (*paymentService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *paymentService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.TransferRequest, ...grpc.CallOption) *paymentService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.TransferRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePayee provides a mock function with given fields: ctx, in, opts
func (_m *PaymentServiceClient) UpdatePayee(ctx context.Context, in *paymentService.UpdatePayeeRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *paymentService.UpdatePayeeRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paymentService.UpdatePayeeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
