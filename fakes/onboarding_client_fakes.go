package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/onboardingService"
)

type FakeOnBoardingClient struct {
	resp            *onboardingService.SuccessResponse
	createPhoneResp *onboardingService.CreatePhoneResponse
	otpResp         *onboardingService.OtpVerificationResponse
	err             error
}

func NewFakeOnBoardingClient(resp *onboardingService.SuccessResponse,
	cResp *onboardingService.CreatePhoneResponse,
	otpResp *onboardingService.OtpVerificationResponse,
	err error) *FakeOnBoardingClient {
	return &FakeOnBoardingClient{resp: resp, err: err, createPhoneResp: cResp, otpResp: otpResp}
}

func (f *FakeOnBoardingClient) CreatePhone(ctx context.Context,
	req *onboardingService.CreatePhoneRequest, opts ...grpc.CallOption) (*onboardingService.CreatePhoneResponse, error) {
	return f.createPhoneResp, f.err
}

func (f *FakeOnBoardingClient) UpdatePersonBiodata(ctx context.Context,
	req *onboardingService.UpdatePersonRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	return f.resp, f.err
}

func (f *FakeOnBoardingClient) CreateEmail(ctx context.Context, req *onboardingService.CreateEmailRequest,
	opts ...grpc.CallOption) (*onboardingService.CreateEmailResponse, error) {
	return &onboardingService.CreateEmailResponse{PersonId: "personId"}, f.err
}

func (f *FakeOnBoardingClient) AddReasonsForUsingRoava(ctx context.Context, req *onboardingService.RoavaReasonsRequest,
	opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
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

func (f *FakeOnBoardingClient) VerifyEmailOtp(ctx context.Context, req *onboardingService.OtpVerificationByEmailRequest,
	opts ...grpc.CallOption) (*onboardingService.OtpVerificationResponse, error) {
	return f.otpResp, f.err
}

func (f *FakeOnBoardingClient) ResendOtp(ctx context.Context, req *onboardingService.ResendOtpRequest,
	opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	return &onboardingService.SuccessResponse{Message: ""}, f.err
}

func (f *FakeOnBoardingClient) CreateApplication(ctx context.Context, req *onboardingService.CreateApplicationRequest,
	opts ...grpc.CallOption) (onboardingService.OnBoardingService_CreateApplicationClient, error) {
	return nil, f.err
}

func (f *FakeOnBoardingClient) SubmitCheck(ctx context.Context, req *onboardingService.SubmitCheckRequest,
	opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	return &onboardingService.SuccessResponse{Message: "sucess"}, f.err
}
