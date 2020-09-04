package fakes

import (
	"context"
	"google.golang.org/grpc"
	verify "ms.api/protos/pb/verifyService"
)

type FakeVerifyClient struct {
	resp *verify.OtpVerificationResponse
	successResp *verify.SuccessResponse
	err error
}

func NewFakeVerifyClient(res *verify.OtpVerificationResponse, successResp *verify.SuccessResponse, err error) *FakeVerifyClient {
	return &FakeVerifyClient{resp: res, err: err}
}

func (f *FakeVerifyClient) ValidateEmail(ctx context.Context, req *verify.ValidateEmailRequest,
	opts... grpc.CallOption) (*verify.SuccessResponse, error) {
	return f.successResp, f.err
}

func (f *FakeVerifyClient) VerifySmsOtp(ctx context.Context, req *verify.OtpVerificationRequest,
	opt... grpc.CallOption) (*verify.OtpVerificationResponse, error) {
	return f.resp, f.err
}

func (f *FakeVerifyClient) VerifyEmailOtp(ctx context.Context, req *verify.OtpVerificationByEmailRequest,
	opt... grpc.CallOption) (*verify.OtpVerificationResponse, error) {
	return f.resp, f.err
}
