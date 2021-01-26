// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	onboardingService "ms.api/protos/pb/onboardingService"
)

// OnBoardingServiceClient is an autogenerated mock type for the OnBoardingServiceClient type
type OnBoardingServiceClient struct {
	mock.Mock
}

// AcceptTermsAndConditions provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) AcceptTermsAndConditions(ctx context.Context, in *onboardingService.TermsAndConditionsRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.TermsAndConditionsRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.TermsAndConditionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddReasonsForUsingRoava provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) AddReasonsForUsingRoava(ctx context.Context, in *onboardingService.RoavaReasonsRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.RoavaReasonsRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.RoavaReasonsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckEmailExistence provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) CheckEmailExistence(ctx context.Context, in *onboardingService.CheckEmailExistenceRequest, opts ...grpc.CallOption) (*onboardingService.CheckEmailExistenceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.CheckEmailExistenceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.CheckEmailExistenceRequest, ...grpc.CallOption) *onboardingService.CheckEmailExistenceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.CheckEmailExistenceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.CheckEmailExistenceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateApplication provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) CreateApplication(ctx context.Context, in *onboardingService.CreateApplicationRequest, opts ...grpc.CallOption) (*onboardingService.CreateApplicationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.CreateApplicationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.CreateApplicationRequest, ...grpc.CallOption) *onboardingService.CreateApplicationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.CreateApplicationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.CreateApplicationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePerson provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) CreatePerson(ctx context.Context, in *onboardingService.CreatePersonRequest, opts ...grpc.CallOption) (*onboardingService.CreatePersonResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.CreatePersonResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.CreatePersonRequest, ...grpc.CallOption) *onboardingService.CreatePersonResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.CreatePersonResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.CreatePersonRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePhone provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) CreatePhone(ctx context.Context, in *onboardingService.CreatePhoneRequest, opts ...grpc.CallOption) (*onboardingService.CreatePhoneResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.CreatePhoneResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.CreatePhoneRequest, ...grpc.CallOption) *onboardingService.CreatePhoneResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.CreatePhoneResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.CreatePhoneRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchCountries provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) FetchCountries(ctx context.Context, in *onboardingService.FetchCountriesRequest, opts ...grpc.CallOption) (*onboardingService.FetchCountriesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.FetchCountriesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.FetchCountriesRequest, ...grpc.CallOption) *onboardingService.FetchCountriesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.FetchCountriesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.FetchCountriesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchReasons provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) FetchReasons(ctx context.Context, in *onboardingService.EmptyRequest, opts ...grpc.CallOption) (*onboardingService.ReasonResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.ReasonResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.EmptyRequest, ...grpc.CallOption) *onboardingService.ReasonResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.ReasonResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.EmptyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressesByText provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) GetAddressesByText(ctx context.Context, in *onboardingService.GetAddressesRequest, opts ...grpc.CallOption) (*onboardingService.GetAddressesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.GetAddressesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.GetAddressesRequest, ...grpc.CallOption) *onboardingService.GetAddressesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.GetAddressesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.GetAddressesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResendEmailMagicLInk provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) ResendEmailMagicLInk(ctx context.Context, in *onboardingService.ResendEmailMagicLInkRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.ResendEmailMagicLInkRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.ResendEmailMagicLInkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResendOtp provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) ResendOtp(ctx context.Context, in *onboardingService.ResendOtpRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.ResendOtpRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.ResendOtpRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitCheck provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) SubmitCheck(ctx context.Context, in *onboardingService.SubmitCheckRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.SubmitCheckRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.SubmitCheckRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFirebaseToken provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) UpdateFirebaseToken(ctx context.Context, in *onboardingService.UpdateFirebaseTokenRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.UpdateFirebaseTokenRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.UpdateFirebaseTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePersonBiodata provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) UpdatePersonBiodata(ctx context.Context, in *onboardingService.UpdatePersonRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.UpdatePersonRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.UpdatePersonRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyEmailMagicLInk provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) VerifyEmailMagicLInk(ctx context.Context, in *onboardingService.VerifyEmailMagicLInkRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.SuccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.VerifyEmailMagicLInkRequest, ...grpc.CallOption) *onboardingService.SuccessResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.SuccessResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.VerifyEmailMagicLInkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyEmailOtp provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) VerifyEmailOtp(ctx context.Context, in *onboardingService.OtpVerificationByEmailRequest, opts ...grpc.CallOption) (*onboardingService.OtpVerificationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.OtpVerificationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.OtpVerificationByEmailRequest, ...grpc.CallOption) *onboardingService.OtpVerificationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.OtpVerificationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.OtpVerificationByEmailRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifySmsOtp provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) VerifySmsOtp(ctx context.Context, in *onboardingService.OtpVerificationRequest, opts ...grpc.CallOption) (*onboardingService.OtpVerificationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.OtpVerificationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.OtpVerificationRequest, ...grpc.CallOption) *onboardingService.OtpVerificationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.OtpVerificationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.OtpVerificationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}