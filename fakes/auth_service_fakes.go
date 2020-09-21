package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/authService"
)

type FakeAuthClient struct {
	loginResp *authService.LoginResponse
	resp *authService.ValidateTokenResponse
	rResp *authService.RefreshTokenResponse
	err error
}

func NewFakeAuthClient(resp *authService.ValidateTokenResponse,
	loginResp *authService.LoginResponse, rResp *authService.RefreshTokenResponse, err error) *FakeAuthClient {
	return &FakeAuthClient{resp: resp, err: err, loginResp: loginResp, rResp: rResp}
}

func (f *FakeAuthClient) Login(ctx context.Context, req *authService.LoginRequest,
	opts...grpc.CallOption) (*authService.LoginResponse, error) {
	return f.loginResp, f.err
}

func (f *FakeAuthClient) ValidateToken(ctx context.Context, req *authService.ValidateTokenRequest,
	opts...grpc.CallOption) (*authService.ValidateTokenResponse, error) {
	return f.resp, f.err
}

func (f *FakeAuthClient) RefreshToken(ctx context.Context, req *authService.RefreshTokenRequest,
	opts... grpc.CallOption) (*authService.RefreshTokenResponse, error) {
	return f.rResp, f.err
}
