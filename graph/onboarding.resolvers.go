package graph

import (
	"context"
	"errors"
	onboarding "ms.api/protos/pb/onboardingService"
	verify "ms.api/protos/pb/verifyService"
	"ms.api/types"
)


func (r *mutationResolver) CreatePhone(ctx context.Context, input types.CreatePhoneInput) (*types.Result, error) {
	if input.Phone == "" || len(input.Phone) < 6 {
		return nil, errors.New("invalid phone number")
	}
	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboarding.CreatePhoneRequest{PhoneNumber: input.Phone, Device: &onboarding.Device{Os: input.Device.Os}})
	if err != nil {
		r.logger.Infof("onBoardingService.createPhone() failed: %v", err)
		return nil, err
	}
	return &types.Result{Message: result.Message, Success: true}, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, phone string, code string) (*types.Result, error) {
	resp, err := r.verifyService.VerifySmsOtp(context.Background(), &verify.OtpVerificationRequest{
		Phone: phone, Code: code,
	})
	if err != nil {
		r.logger.Infof("verifyService.verifySmsOtp() failed: %v", err)
		return nil, err
	}
	return &types.Result{Success: resp.Match, Message: resp.Message}, nil
}
