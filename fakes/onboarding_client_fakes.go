package fakes

import (
	"context"
	"google.golang.org/grpc"
	onboarding "ms.api/protos/pb/onboardingService"
)

type FakeOnBoardingClient struct {
	resp *onboarding.SuccessResponse
	err error
}

func NewFakeOnBoardingClient(resp *onboarding.SuccessResponse, err error) *FakeOnBoardingClient {
	return &FakeOnBoardingClient{resp: resp, err: err}
}

func (f *FakeOnBoardingClient) CreatePhone(ctx context.Context, req *onboarding.CreatePhoneRequest, opts ...grpc.CallOption) (*onboarding.SuccessResponse, error) {
	return f.resp, f.err
}

func (f *FakeOnBoardingClient) CreateEmail(ctx context.Context, req *onboarding.CreateEmailRequest, opts ...grpc.CallOption) (*onboarding.SuccessResponse, error) {
	return f.resp, f.err
}

func (f *FakeOnBoardingClient) CreatePasscode(ctx context.Context, req *onboarding.CreatePasscodeRequest, opts ...grpc.CallOption) (*onboarding.SuccessResponse, error) {
	return f.resp, f.err
}
