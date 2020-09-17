package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"ms.api/graph/generated"
	emailvalidator "ms.api/libs/email"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/verifyService"
	"ms.api/server/http/middlewares"
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
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	payload := onboardingService.UpdatePersonRequest{
		PersonId:  personId,
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
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	payload := onboardingService.RoavaReasonsRequest{
		PersonId: personId,
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

func (r *mutationResolver) AuthenticateCustomer(ctx context.Context, email string, passcode string) (*types.AuthResult, error) {
	if err := emailvalidator.Validate(email); err != nil {
		r.logger.WithField("email", email).Info("invalid email supplied")
		return nil, errors.New("invalid email address")
	}
	req := &authService.LoginRequest{Email: email, Passcode: passcode}
	resp, err := r.authService.Login(ctx, req)
	if err != nil {
		r.logger.Info("authService.Login() failed: %v", err)
		return nil, err
	}
	return &types.AuthResult{
		Token: resp.Token, RefreshToken: resp.RefreshToken,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	ErrUnAuthenticated = errors.New("user not authenticated")
)
