package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/verifyService"

	"ms.api/graph/generated"
	"ms.api/types"
)

func (r *mutationResolver) SubmitKYCApplication(ctx context.Context, applicationID string) (*types.Result, error) {
	if _, err := r.kycClient.StartApplicationCDD(ctx, &kycService.ApplicationIdRequest{
		ApplicationId: applicationID,
	}); err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: "Successfully started CDD check, you'll be notified once completed.",
	}, nil
}
func (r *mutationResolver) CreatePasscode(ctx context.Context, userID string, passcode string) (*types.Result, error) {
	payload := onboardingService.CreatePasscodeRequest{
		PersonId: userID,
		Passcode: passcode,
	}
	res, err := r.onBoardingService.CreatePasscode(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) UpdatePersonBiodata(ctx context.Context, input *types.UpdateBioDataInput) (*types.Result, error) {
	payload := onboardingService.UpdatePersonRequest{
		PersonId:  input.PersonID,
		Address:   input.Address,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Dob:       input.Dob,
	}
	res, err := r.onBoardingService.UpdatePersonBiodata(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) AddReasonsForUsingRoava(ctx context.Context, personID string, reasons string) (*types.Result, error) {
	payload := onboardingService.RoavaReasonsRequest{
		PersonId: personID,
		Reasons:  reasons,
	}
	res, err := r.onBoardingService.AddReasonsForUsingRoava(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) CreatePhone(ctx context.Context, input types.CreatePhoneInput) (*types.CreatePhoneResult, error) {
	if input.Phone == "" || len(input.Phone) < 6 {
		return nil, errors.New("invalid phone number")
	}
	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboardingService.CreatePhoneRequest{PhoneNumber: input.Phone, Device: &onboardingService.Device{Os: input.Device.Os}})
	if err != nil {
		r.logger.Infof("onBoardingService.createPhone() failed: %v", err)
		return nil, err
	}
	return &types.CreatePhoneResult{Message: result.Message, Success: true, EmailToken: result.EmailToken}, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, phone string, code string) (*types.Result, error) {
	resp, err := r.verifyService.VerifySmsOtp(context.Background(), &verifyService.OtpVerificationRequest{
		Phone: phone, Code: code,
	})
	if err != nil {
		r.logger.Infof("verifyService.verifySmsOtp() failed: %v", err)
		return nil, err
	}
	return &types.Result{Success: resp.Match, Message: resp.Message}, nil
}

func (r *mutationResolver) CreateEmail(ctx context.Context, input *types.CreateEmailInput) (*types.Result, error) {
	resp, err := r.onBoardingService.CreateEmail(ctx, &onboardingService.CreateEmailRequest{
		Value:      input.Value,
		EmailToken: input.EmailToken,
	})
	if err != nil {
		r.logger.Infof("onBoardingService.createEmail() failed: %v", err)
		return nil, err
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
