package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/onboardingService"
)

type FakeOnBoardingClient struct {
	resp            *onboardingService.SuccessResponse
	createPhoneResp *onboardingService.CreatePhoneResponse
	err             error
}

func NewFakeOnBoardingClient(resp *onboardingService.SuccessResponse, cResp *onboardingService.CreatePhoneResponse,
	err error) *FakeOnBoardingClient {
	return &FakeOnBoardingClient{resp: resp, err: err, createPhoneResp: cResp}
}

func (f *FakeOnBoardingClient) CreatePhone(ctx context.Context,
	req *onboardingService.CreatePhoneRequest, opts ...grpc.CallOption) (*onboardingService.CreatePhoneResponse, error) {
	return f.createPhoneResp, f.err
}

func (f *FakeOnBoardingClient) CreatePasscode(ctx context.Context,
	req *onboardingService.CreatePasscodeRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	return f.resp, f.err
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
