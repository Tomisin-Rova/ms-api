// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	onboardingService "ms.api/protos/pb/onboardingService"

	types "ms.api/protos/pb/types"
)

// OnBoardingServiceClient is an autogenerated mock type for the OnBoardingServiceClient type
type OnBoardingServiceClient struct {
	mock.Mock
}

// AcceptTermsAndConditions provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) AcceptTermsAndConditions(ctx context.Context, in *onboardingService.TermsAndConditionsRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.TermsAndConditionsRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
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
func (_m *OnBoardingServiceClient) AddReasonsForUsingRoava(ctx context.Context, in *onboardingService.RoavaReasonsRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.RoavaReasonsRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
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

// AddressLookup provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) AddressLookup(ctx context.Context, in *onboardingService.AddressLookupRequest, opts ...grpc.CallOption) (*onboardingService.AddressLookupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.AddressLookupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.AddressLookupRequest, ...grpc.CallOption) *onboardingService.AddressLookupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.AddressLookupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.AddressLookupRequest, ...grpc.CallOption) error); ok {
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

// CreateOnfidoApplicant provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) CreateOnfidoApplicant(ctx context.Context, in *onboardingService.CreateOnfidoApplicantRequest, opts ...grpc.CallOption) (*onboardingService.CreateOnfidoApplicantResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.CreateOnfidoApplicantResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.CreateOnfidoApplicantRequest, ...grpc.CallOption) *onboardingService.CreateOnfidoApplicantResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.CreateOnfidoApplicantResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.CreateOnfidoApplicantRequest, ...grpc.CallOption) error); ok {
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
func (_m *OnBoardingServiceClient) FetchReasons(ctx context.Context, in *onboardingService.FetchReasonsRequest, opts ...grpc.CallOption) (*onboardingService.ReasonResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.ReasonResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.FetchReasonsRequest, ...grpc.CallOption) *onboardingService.ReasonResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.ReasonResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.FetchReasonsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOnfidoSDKToken provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) GetOnfidoSDKToken(ctx context.Context, in *onboardingService.GetOnfidoSDKTokenRequest, opts ...grpc.CallOption) (*onboardingService.GetOnfidoSDKTokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *onboardingService.GetOnfidoSDKTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.GetOnfidoSDKTokenRequest, ...grpc.CallOption) *onboardingService.GetOnfidoSDKTokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*onboardingService.GetOnfidoSDKTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.GetOnfidoSDKTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResendEmailMagicLInk provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) ResendEmailMagicLInk(ctx context.Context, in *onboardingService.ResendEmailMagicLInkRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.ResendEmailMagicLInkRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
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
func (_m *OnBoardingServiceClient) ResendOtp(ctx context.Context, in *onboardingService.ResendOtpRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.ResendOtpRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
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

// Resubmit provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) Resubmit(ctx context.Context, in *onboardingService.ResubmitRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.ResubmitRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.ResubmitRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResubmitReport provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) ResubmitReport(ctx context.Context, in *onboardingService.ResubmitReportRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.ResubmitReportRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.ResubmitReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitApplication provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) SubmitApplication(ctx context.Context, in *onboardingService.SubmitApplicationRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.SubmitApplicationRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.SubmitApplicationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitCheck provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) SubmitCheck(ctx context.Context, in *onboardingService.SubmitCheckRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.SubmitCheckRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
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

// SubmitProof provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) SubmitProof(ctx context.Context, in *onboardingService.SubmitProofRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.SubmitProofRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.SubmitProofRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFirebaseToken provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) UpdateFirebaseToken(ctx context.Context, in *onboardingService.UpdateFirebaseTokenRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.UpdateFirebaseTokenRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
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
func (_m *OnBoardingServiceClient) UpdatePersonBiodata(ctx context.Context, in *onboardingService.UpdatePersonRequest, opts ...grpc.CallOption) (*types.Person, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Person
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.UpdatePersonRequest, ...grpc.CallOption) *types.Person); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Person)
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

// UpdateValidationStatus provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) UpdateValidationStatus(ctx context.Context, in *onboardingService.UpdateValidationStatusRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.UpdateValidationStatusRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *onboardingService.UpdateValidationStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyEmailMagicLInk provides a mock function with given fields: ctx, in, opts
func (_m *OnBoardingServiceClient) VerifyEmailMagicLInk(ctx context.Context, in *onboardingService.VerifyEmailMagicLInkRequest, opts ...grpc.CallOption) (*types.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Response
	if rf, ok := ret.Get(0).(func(context.Context, *onboardingService.VerifyEmailMagicLInkRequest, ...grpc.CallOption) *types.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Response)
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
