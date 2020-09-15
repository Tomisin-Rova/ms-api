package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/onboardingService"
)

type FakeOnBoardingClient struct {
	resp    *onboardingService.SuccessResponse
	msgResp *onboardingService.MsgResponse
	err     error
}

func NewFakeOnBoardingClient(resp *onboardingService.SuccessResponse, msg *onboardingService.MsgResponse,
	err error) *FakeOnBoardingClient {
	return &FakeOnBoardingClient{resp: resp, err: err, msgResp: msg}
}

func (f *FakeOnBoardingClient) CreatePhone(ctx context.Context,
	req *onboardingService.CreatePhoneRequest, opts ...grpc.CallOption) (*onboardingService.MsgResponse, error) {
	return f.msgResp, f.err
}

func (f *FakeOnBoardingClient) CreatePasscode(ctx context.Context,
	req *onboardingService.CreatePasscodeRequest, opts ...grpc.CallOption) (*onboardingService.MsgResponse, error) {
	return f.msgResp, f.err
}

func (f *FakeOnBoardingClient) UpdatePersonBiodata(ctx context.Context,
	req *onboardingService.UpdatePersonRequest, opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	return f.resp, f.err
}

func (f *FakeOnBoardingClient) AddReasonsForUsingRoava(ctx context.Context, req *onboardingService.RoavaReasonsRequest,
	opts ...grpc.CallOption) (*onboardingService.SuccessResponse, error) {
	return f.resp, f.err
}
