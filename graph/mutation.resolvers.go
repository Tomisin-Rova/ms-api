package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"ms.api/graph/generated"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/onboardingService"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *mutationResolver) CreatePhone(ctx context.Context, phone string, device types.DeviceInput) (*types.Response, error) {
	if err := phonenumbervalidator.ValidatePhoneNumber(phone); err != nil {
		r.logger.Error("failed to validate phone number",
			zap.Error(err),
			zap.String("phone", phone),
		)
		return nil, err
	}
	// TODO: change onboardingService.CreatePhoneRequest{}.Tokens to slice datatype
	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboardingService.CreatePhoneRequest{PhoneNumber: phone,
			Device: &protoTypes.Device{Os: device.Os, Brand: device.Brand,
				DeviceId: device.Identifier, DeviceToken: ""}})
	if err != nil {
		r.logger.Info(fmt.Sprintf("OnBoardingService.createPhone() failed: %v", err))
		return nil, err
	}
	return &types.Response{Message: result.Message, Success: true, Token: &result.Token}, nil
}

func (r *mutationResolver) ConfirmPhone(ctx context.Context, token string, code string) (*types.Response, error) {
	resp, err := r.onBoardingService.VerifySmsOtp(context.Background(), &onboardingService.OtpVerificationRequest{
		Token: token, Code: code,
	})
	if err != nil {
		r.logger.Info(fmt.Sprintf("onboardingService.verifySmsOtp() failed: %v", err))
		return nil, err
	}
	return &types.Response{Success: resp.Match, Message: resp.Message}, nil
}

func (r *mutationResolver) Signup(ctx context.Context, token string, email string, passcode *string) (*types.AuthResponse, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}

	result, err := r.onBoardingService.CreatePerson(ctx, &onboardingService.CreatePersonRequest{
		Email:    email,
		Passcode: *passcode,
		Token:    token,
	})
	if err != nil {
		r.logger.Error("error calling onboardingService.CreatePerson.", zap.Error(err))
		return nil, err
	}

	return &types.AuthResponse{
		Message: "successful",
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    result.JwtToken,
			Refresh: result.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) Registration(ctx context.Context, personid string, person types.PersonInput, address types.AddressInput) (*types.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) IntendedActivities(ctx context.Context, activities []string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	req := &onboardingService.RoavaReasonsRequest{
		PersonId: claims.PersonId,
		Reasons:  activities,
	}

	resp, err := r.onBoardingService.AddReasonsForUsingRoava(ctx, req)
	if err != nil {
		r.logger.Error("Adding reasons for using roava failed", zap.Error(err))
		return nil, fmt.Errorf("an error ocurred saving the reasons of use, pls try again")
	}

	r.logger.Info("Adding reasons for using roava succeed")

	return &types.Response{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) CreateApplication(ctx context.Context, applicant types.ApplicantInput) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) VerifyEmail(ctx context.Context, email string, token string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResendOtp(ctx context.Context, phone string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResendEmailMagicLInk(ctx context.Context, email string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, credentials types.AuthInput, biometric *bool) (*types.AuthResponse, error) {
	if err := emailvalidator.Validate(credentials.Email); err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", credentials.Email))
		return nil, err
	}
	// TODO: change authService.LoginRequest{}.Tokens to slice datatype
	bio := biometric != nil && *biometric
	req := &authService.LoginRequest{
		Email:     credentials.Email,
		Passcode:  credentials.Passcode,
		Biometric: bio,
		Device: &protoTypes.Device{
			Os:       credentials.Device.Os,
			Brand:    credentials.Device.Brand,
			DeviceId: credentials.Device.Identifier,
		}}
	resp, err := r.authService.Login(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.Login() failed: %v", err))
		return nil, err
	}
	return &types.AuthResponse{
		Message: "Successful!",
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    resp.AccessToken,
			Refresh: resp.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, token string) (*types.AuthResponse, error) {
	req := &authService.RefreshTokenRequest{RefreshToken: token}
	resp, err := r.authService.RefreshToken(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.RefreshToken() failed: %v", err))
		return nil, err
	}
	return &types.AuthResponse{
		Message: "Successful!",
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    resp.Token,
			Refresh: resp.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) UpdateDeviceToken(ctx context.Context, token string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetBiometricAuth(ctx context.Context, activate *bool) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	setActive := activate != nil && *activate
	if setActive {
		resp, err := r.authService.ActivateBioLogin(ctx, &authService.ActivateBioLoginRequest{
			IdentityId: claims.IdentityId,
			DeviceId:   claims.DeviceId,
		})
		if err != nil {
			r.logger.Info("authService.ActivateBioLogin() failed", zap.Error(err))
			return nil, err
		}
		return &types.Response{Message: resp.Message, Success: true, Token: &resp.BiometricPasscode}, nil
	} else {
		resp, err := r.authService.DeactivateBioLogin(ctx, &authService.DeactivateBioLoginRequest{
			IdentityId: claims.IdentityId,
			DeviceId:   claims.DeviceId,
		})
		if err != nil {
			r.logger.Info("authService.DeactivateBioLogin() failed", zap.Error(err))
			return nil, err
		}
		return &types.Response{Message: resp.Message, Success: true}, nil
	}
}

func (r *mutationResolver) ResetPasscode(ctx context.Context, credentials *types.AuthInput, token string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SubmitApplication(ctx context.Context) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AcceptTerms(ctx context.Context, documents []*string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
