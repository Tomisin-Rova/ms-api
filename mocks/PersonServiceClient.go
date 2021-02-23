// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	personService "ms.api/protos/pb/personService"

	types "ms.api/protos/pb/types"
)

// PersonServiceClient is an autogenerated mock type for the PersonServiceClient type
type PersonServiceClient struct {
	mock.Mock
}

// Me provides a mock function with given fields: ctx, in, opts
func (_m *PersonServiceClient) Me(ctx context.Context, in *personService.TokenRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *personService.TokenRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *personService.TokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// People provides a mock function with given fields: ctx, in, opts
func (_m *PersonServiceClient) People(ctx context.Context, in *personService.PeopleRequest, opts ...grpc.CallOption) (*types.Persons, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Persons
	if rf, ok := ret.Get(0).(func(context.Context, *personService.PeopleRequest, ...grpc.CallOption) *types.Persons); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Persons)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *personService.PeopleRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Person provides a mock function with given fields: ctx, in, opts
func (_m *PersonServiceClient) Person(ctx context.Context, in *personService.PersonRequest, opts ...grpc.CallOption) (*types.Person, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Person
	if rf, ok := ret.Get(0).(func(context.Context, *personService.PersonRequest, ...grpc.CallOption) *types.Person); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *personService.PersonRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransactionPin provides a mock function with given fields: ctx, in, opts
func (_m *PersonServiceClient) SetTransactionPin(ctx context.Context, in *personService.TransactionPinRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *personService.TransactionPinRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *personService.TransactionPinRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}