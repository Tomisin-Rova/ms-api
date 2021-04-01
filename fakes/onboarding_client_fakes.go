package fakes

import (
	"context"

	"ms.api/protos/pb/types"

	"google.golang.org/grpc"
	"ms.api/protos/pb/onboardingService"
)

var _ onboardingService.OnBoardingServiceClient = &FakeOnBoardingClient{}

type FakeOnBoardingClient struct {
	resp            *types.Response
	createPhoneResp *onboardingService.CreatePhoneResponse
	otpResp         *onboardingService.OtpVerificationResponse
	err             error
}

func NewFakeOnBoardingClient(resp *types.Response,
	cResp *onboardingService.CreatePhoneResponse,
	otpResp *onboardingService.OtpVerificationResponse,
	err error) *FakeOnBoardingClient {
	return &FakeOnBoardingClient{resp: resp, err: err, createPhoneResp: cResp, otpResp: otpResp}
}

func (f *FakeOnBoardingClient) CreatePhone(ctx context.Context,
	req *onboardingService.CreatePhoneRequest, opts ...grpc.CallOption) (*onboardingService.CreatePhoneResponse, error) {
	return f.createPhoneResp, f.err
}

func (f *FakeOnBoardingClient) UpdatePersonBiodata(ctx context.Context, in *onboardingService.UpdatePersonRequest, opts ...grpc.CallOption) (*types.Person, error) {
	return &types.Person{}, f.err
}

func (f *FakeOnBoardingClient) CreatePerson(ctx context.Context, req *onboardingService.CreatePersonRequest,
	opts ...grpc.CallOption) (*onboardingService.CreatePersonResponse, error) {
	return &onboardingService.CreatePersonResponse{
		JwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
		Message:  "success",
	}, f.err
}

func (f *FakeOnBoardingClient) AddReasonsForUsingRoava(ctx context.Context, req *onboardingService.RoavaReasonsRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return f.resp, f.err
}

func (f *FakeOnBoardingClient) CheckEmailExistence(ctx context.Context, req *onboardingService.CheckEmailExistenceRequest,
	opts ...grpc.CallOption) (*onboardingService.CheckEmailExistenceResponse, error) {
	return &onboardingService.CheckEmailExistenceResponse{Message: "", Exists: false}, f.err
}

func (f *FakeOnBoardingClient) VerifySmsOtp(ctx context.Context, req *onboardingService.OtpVerificationRequest,
	opts ...grpc.CallOption) (*onboardingService.OtpVerificationResponse, error) {
	return f.otpResp, f.err
}

func (f *FakeOnBoardingClient) ResendOtp(ctx context.Context, req *onboardingService.ResendOtpRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{Message: ""}, f.err
}

func (f *FakeOnBoardingClient) CreateOnfidoApplicant(ctx context.Context, req *onboardingService.CreateOnfidoApplicantRequest,
	opts ...grpc.CallOption) (*onboardingService.CreateOnfidoApplicantResponse, error) {
	return nil, f.err
}

func (f *FakeOnBoardingClient) GetOnfidoSDKToken(ctx context.Context, in *onboardingService.GetOnfidoSDKTokenRequest, opts ...grpc.CallOption) (*onboardingService.GetOnfidoSDKTokenResponse, error) {
	return nil, f.err
}

func (f *FakeOnBoardingClient) SubmitCheck(ctx context.Context, req *onboardingService.SubmitCheckRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{Message: "sucess"}, f.err
}

func (f *FakeOnBoardingClient) VerifyEmailMagicLInk(ctx context.Context, req *onboardingService.VerifyEmailMagicLInkRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{Message: "success"}, f.err
}

func (f *FakeOnBoardingClient) ResendEmailMagicLInk(ctx context.Context, req *onboardingService.ResendEmailMagicLInkRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{Message: "success"}, f.err
}

func (f *FakeOnBoardingClient) FetchCountries(ctx context.Context, req *onboardingService.FetchCountriesRequest,
	opts ...grpc.CallOption) (*onboardingService.FetchCountriesResponse, error) {
	return &onboardingService.FetchCountriesResponse{}, f.err
}

func (f *FakeOnBoardingClient) AcceptTermsAndConditions(ctx context.Context, req *onboardingService.TermsAndConditionsRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{}, f.err
}

func (f *FakeOnBoardingClient) FetchReasons(ctx context.Context, req *onboardingService.FetchReasonsRequest,
	opts ...grpc.CallOption) (*onboardingService.ReasonResponse, error) {
	return &onboardingService.ReasonResponse{}, f.err
}

func (f *FakeOnBoardingClient) UpdateFirebaseToken(ctx context.Context, req *onboardingService.UpdateFirebaseTokenRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{}, f.err
}

func (f *FakeOnBoardingClient) AddressLookup(ctx context.Context, req *onboardingService.AddressLookupRequest,
	opts ...grpc.CallOption) (*onboardingService.AddressLookupResponse, error) {
	return &onboardingService.AddressLookupResponse{}, f.err
}

func (f *FakeOnBoardingClient) SubmitApplication(ctx context.Context, in *onboardingService.SubmitApplicationRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{}, f.err
}

func (f *FakeOnBoardingClient) SubmitProof(ctx context.Context, in *onboardingService.SubmitProofRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{}, f.err
}

func (f *FakeOnBoardingClient) UpdateValidationStatus(ctx context.Context, in *onboardingService.UpdateValidationStatusRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return &types.Response{}, f.err
}
