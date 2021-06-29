// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	context "context"

	authService "ms.api/protos/pb/authService"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "ms.api/protos/pb/types"
)

// AuthServiceClient is an autogenerated mock type for the AuthServiceClient type
type AuthServiceClient struct {
	mock.Mock
}

// ConfirmPasswordResetDetails provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) ConfirmPasswordResetDetails(ctx context.Context, in *authService.PasswordResetUserDetails, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *authService.PasswordResetUserDetails, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.PasswordResetUserDetails, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfirmPasswordResetOtp provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) ConfirmPasswordResetOtp(ctx context.Context, in *authService.PasswordResetOtpRequest, opts ...grpc.CallOption) (*authService.PasswordResetOtpResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *authService.PasswordResetOtpResponse
	if rf, ok := ret.Get(0).(func(context.Context, *authService.PasswordResetOtpRequest, ...grpc.CallOption) *authService.PasswordResetOtpResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authService.PasswordResetOtpResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.PasswordResetOtpRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPerson provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) GetPerson(ctx context.Context, in *authService.GetPersonRequest, opts ...grpc.CallOption) (*authService.GetPersonResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *authService.GetPersonResponse
	if rf, ok := ret.Get(0).(func(context.Context, *authService.GetPersonRequest, ...grpc.CallOption) *authService.GetPersonResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authService.GetPersonResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.GetPersonRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) Login(ctx context.Context, in *authService.LoginRequest, opts ...grpc.CallOption) (*authService.AuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *authService.AuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *authService.LoginRequest, ...grpc.CallOption) *authService.AuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authService.AuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginWithToken provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) LoginWithToken(ctx context.Context, in *authService.LoginWithTokenRequest, opts ...grpc.CallOption) (*authService.AuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *authService.AuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *authService.LoginWithTokenRequest, ...grpc.CallOption) *authService.AuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authService.AuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.LoginWithTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshToken provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) RefreshToken(ctx context.Context, in *authService.RefreshTokenRequest, opts ...grpc.CallOption) (*authService.RefreshTokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *authService.RefreshTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *authService.RefreshTokenRequest, ...grpc.CallOption) *authService.RefreshTokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authService.RefreshTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.RefreshTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) ResetPassword(ctx context.Context, in *authService.PasswordResetRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *authService.PasswordResetRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.PasswordResetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateEmail provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) ValidateEmail(ctx context.Context, in *authService.ValidateEmailRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *authService.ValidateEmailRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.ValidateEmailRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) ValidateToken(ctx context.Context, in *authService.ValidateTokenRequest, opts ...grpc.CallOption) (*authService.ValidateTokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *authService.ValidateTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *authService.ValidateTokenRequest, ...grpc.CallOption) *authService.ValidateTokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authService.ValidateTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.ValidateTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateUser provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) ValidateUser(ctx context.Context, in *authService.ValidateUserRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *authService.ValidateUserRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *authService.ValidateUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
