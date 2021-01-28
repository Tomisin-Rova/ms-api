package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/types"
	"ms.api/protos/pb/verifyService"
)

type FakeVerifyClient struct {
	resp        *verifyService.OtpVerificationResponse
	successResp *types.Response
	err         error
}

func NewFakeVerifyClient(res *verifyService.OtpVerificationResponse,
	successResp *types.Response, err error) *FakeVerifyClient {
	return &FakeVerifyClient{resp: res, err: err}
}

func (f *FakeVerifyClient) ValidateEmail(ctx context.Context, req *verifyService.ValidateEmailRequest,
	opts ...grpc.CallOption) (*types.Response, error) {
	return f.successResp, f.err
}

func (f *FakeVerifyClient) VerifySmsOtp(ctx context.Context, req *verifyService.OtpVerificationRequest,
	opt ...grpc.CallOption) (*verifyService.OtpVerificationResponse, error) {
	return f.resp, f.err
}

func (f *FakeVerifyClient) VerifyEmailOtp(ctx context.Context, req *verifyService.OtpVerificationByEmailRequest,
	opt ...grpc.CallOption) (*verifyService.OtpVerificationResponse, error) {
	return f.resp, f.err
}

func (f *FakeVerifyClient) ResendOtp(ctx context.Context, req *verifyService.ResendOtpRequest,
	opt ...grpc.CallOption) (*types.Response, error) {
	return f.successResp, f.err
}
