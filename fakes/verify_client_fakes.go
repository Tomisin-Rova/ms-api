package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/verifyService"
)

type FakeVerifyClient struct {
	resp        *verifyService.OtpVerificationResponse
	successResp *verifyService.SuccessResponse
	err         error
}

func NewFakeVerifyClient(res *verifyService.OtpVerificationResponse,
	successResp *verifyService.SuccessResponse, err error) *FakeVerifyClient {
	return &FakeVerifyClient{resp: res, err: err}
}

func (f *FakeVerifyClient) ValidateEmail(ctx context.Context, req *verifyService.ValidateEmailRequest,
	opts ...grpc.CallOption) (*verifyService.SuccessResponse, error) {
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
