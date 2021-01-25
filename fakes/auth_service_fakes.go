package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/authService"
)

type FakeAuthClient struct {
	loginResp *authService.LoginResponse
	resp      *authService.ValidateTokenResponse
	rResp     *authService.RefreshTokenResponse
	err       error
}

func (f *FakeAuthClient) ConfirmPasswordResetOtp(ctx context.Context, in *authService.PasswordResetOtpRequest, opts ...grpc.CallOption) (*authService.SuccessResponse, error) {
	return &authService.SuccessResponse{}, f.err
}

func (f *FakeAuthClient) ConfirmPasswordResetDetails(ctx context.Context, in *authService.PasswordResetUserDetails, opts ...grpc.CallOption) (*authService.SuccessResponse, error) {
	return &authService.SuccessResponse{}, nil
}

func (f *FakeAuthClient) ResetPassword(ctx context.Context, in *authService.PasswordResetRequest, opts ...grpc.CallOption) (*authService.SuccessResponse, error) {
	return &authService.SuccessResponse{}, f.err
}

func NewFakeAuthClient(resp *authService.ValidateTokenResponse,
	loginResp *authService.LoginResponse, rResp *authService.RefreshTokenResponse, err error) *FakeAuthClient {
	return &FakeAuthClient{resp: resp, err: err, loginResp: loginResp, rResp: rResp}
}

func (f *FakeAuthClient) Login(ctx context.Context, req *authService.LoginRequest,
	opts ...grpc.CallOption) (*authService.AuthResponse, error) {
	return &authService.AuthResponse{}, f.err
}

func (f *FakeAuthClient) ValidateToken(ctx context.Context, req *authService.ValidateTokenRequest,
	opts ...grpc.CallOption) (*authService.ValidateTokenResponse, error) {
	return f.resp, f.err
}

func (f *FakeAuthClient) RefreshToken(ctx context.Context, req *authService.RefreshTokenRequest,
	opts ...grpc.CallOption) (*authService.RefreshTokenResponse, error) {
	return f.rResp, f.err
}

func (f *FakeAuthClient) GetPerson(ctx context.Context, req *authService.GetPersonRequest,
	opts ...grpc.CallOption) (*authService.GetPersonResponse, error) {
	return &authService.GetPersonResponse{}, f.err
}

func (f *FakeAuthClient) ActivateBioLogin(ctx context.Context, req *authService.ActivateBioLoginRequest, opts ...grpc.CallOption) (*authService.ActivateBioLoginResponse, error) {
	return &authService.ActivateBioLoginResponse{}, f.err
}

func (f *FakeAuthClient) DeactivateBioLogin(ctx context.Context, req *authService.DeactivateBioLoginRequest, opts ...grpc.CallOption) (*authService.SuccessResponse, error) {
	return &authService.SuccessResponse{}, f.err
}
