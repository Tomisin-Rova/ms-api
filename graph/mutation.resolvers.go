package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"ms.api/graph/generated"
	onboarding "ms.api/protos/pb/onboardingService"
	verify "ms.api/protos/pb/verifyService"
	"ms.api/types"
)

func (r *mutationResolver) SubmitLiveVideo(ctx context.Context, id string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PingKYCService(ctx context.Context, message string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
