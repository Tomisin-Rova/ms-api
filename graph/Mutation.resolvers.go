package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"ms.api/graph/generated"
	rerrors "ms.api/libs/errors"
	"ms.api/libs/validator/datevalidator"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *mutationResolver) ResetPasscode(ctx context.Context, email string, newPasscode string, verificationToken string) (*types.Result, error) {
	result, err := r.authService.ResetPassword(ctx, &authService.PasswordResetRequest{
		Email:             email,
		NewPassword:       newPasscode,
		VerificationToken: verificationToken,
	})
	if err != nil {
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.Result{
		Success: true,
		Message: result.Message,
	}, err
}

func (r *mutationResolver) ConfirmPasscodeResetDetails(ctx context.Context, email string, dob string, address types.InputAddress) (*types.Result, error) {
	result, err := r.authService.ConfirmPasswordResetDetails(ctx, &authService.PasswordResetUserDetails{
		Email: email,
		Dob:   dob,
		Address: &authService.Address{
			Country:  address.Country,
			Street:   address.Street,
			City:     address.City,
			Postcode: address.Postcode,
		},
	})
	if err != nil {
		r.logger.Infof("authService.ConfirmPasswordResetDetails() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.Result{
		Success: true,
		Message: result.Message,
	}, nil
}

func (r *mutationResolver) ConfirmPasscodeResetOtp(ctx context.Context, email string, otp string) (*types.Result, error) {
	result, err := r.authService.ConfirmPasswordResetOtp(ctx, &authService.PasswordResetOtpRequest{
		Email: email,
		Code:  otp,
	})
	if err != nil {
		r.logger.Infof("authService.ConfirmPasswordResetOtp() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.Result{
		Success: true,
		Message: result.Message,
	}, nil
}

func (r *mutationResolver) UpdatePersonBiodata(ctx context.Context, input *types.UpdateBioDataInput) (*types.Result, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	if err := datevalidator.ValidateDob(input.Dob); err != nil {
		return nil, err
	}
	if err := r.validateAddress(input.Address); err != nil {
		return nil, err
	}
	if input.Address.Postcode == "" {
		input.Address.Postcode = "NA"
	}

	payload := onboardingService.UpdatePersonRequest{
		PersonId: personId,
		Address: &onboardingService.Address{
			Postcode: input.Address.Postcode, Street: input.Address.Street,
			City: input.Address.City, Country: input.Address.Country,
		},
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Dob:       input.Dob,
	}
	res, err := r.onBoardingService.UpdatePersonBiodata(context.Background(), &payload)
	if err != nil {
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) AddReasonsForUsingRoava(ctx context.Context, personID string, reasonValues []*string) (*types.Result, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	reasons := make([]string, 0, len(reasonValues))
	for _, reason := range reasonValues {
		reasons = append(reasons, *reason)
	}
	payload := onboardingService.RoavaReasonsRequest{
		PersonId: personId,
		Reasons:  reasons,
	}
	res, err := r.onBoardingService.AddReasonsForUsingRoava(context.Background(), &payload)
	if err != nil {
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) CreatePhone(ctx context.Context, input types.CreatePhoneInput) (*types.CreatePhoneResult, error) {
	if err := phonenumbervalidator.ValidatePhoneNumber(input.Phone); err != nil {
		r.logger.WithField("phone", input.Phone).WithError(err).
			Error("failed to validate phone number")
		return nil, err
	}
	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboardingService.CreatePhoneRequest{PhoneNumber: input.Phone,
			Device: &onboardingService.Device{Os: input.Device.Os, Brand: input.Device.Brand,
				DeviceId: input.Device.DeviceID, DeviceToken: input.Device.DeviceToken}})
	if err != nil {
		r.logger.Infof("onBoardingService.createPhone() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.CreatePhoneResult{Message: result.Message, Success: true, Token: result.Token}, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, phone string, code string) (*types.Result, error) {
	resp, err := r.onBoardingService.VerifySmsOtp(context.Background(), &onboardingService.OtpVerificationRequest{
		Phone: phone, Code: code,
	})
	if err != nil {
		r.logger.Infof("onboardingService.verifySmsOtp() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.Result{Success: resp.Match, Message: resp.Message}, nil
}

func (r *mutationResolver) CreateEmail(ctx context.Context, input *types.CreateEmailInput) (*types.AuthResult, error) {
	if err := emailvalidator.Validate(input.Email); err != nil {
		return nil, err
	}

	resp, err := r.onBoardingService.CreateEmail(ctx, &onboardingService.CreateEmailRequest{
		Email:    input.Email,
		Token:    input.Token,
		Passcode: input.Passcode,
	})
	if err != nil {
		r.logger.Infof("onBoardingService.createEmail() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}
	tokens, err := r.authService.GenerateToken(ctx, &authService.GenerateTokenRequest{PersonId: resp.PersonId})
	if err != nil {
		r.logger.Infof("authService.generateToken() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.AuthResult{Token: tokens.Token, RefreshToken: tokens.RefreshToken}, nil
}

func (r *mutationResolver) AuthenticateCustomer(ctx context.Context, email string, passcode string) (*types.AuthResult, error) {
	if err := emailvalidator.Validate(email); err != nil {
		r.logger.WithField("email", email).Info("invalid email supplied")
		return nil, errors.New("invalid email address")
	}
	req := &authService.LoginRequest{Email: email, Passcode: passcode}
	resp, err := r.authService.Login(ctx, req)
	if err != nil {
		r.logger.Infof("authService.Login() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.AuthResult{
		Token: resp.AccessToken, RefreshToken: resp.RefreshToken,
		Person: &types.APIPerson{
			FirstName:               resp.Person.FirstName,
			LastName:                resp.Person.LastName,
			Email:                   resp.Person.Email,
			IsEmailActive:           resp.Person.IsEmailActive,
			IsBiometricLoginEnabled: resp.Person.IsBiometricLoginEnabled,
			IsTransactionPinEnabled: resp.Person.IsTransactionPinEnabled,
			RegistrationCheckPoint:  resp.Person.RegistrationCheckPoint,
		},
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string) (*types.AuthResult, error) {
	req := &authService.RefreshTokenRequest{RefreshToken: refreshToken}
	resp, err := r.authService.RefreshToken(ctx, req)
	if err != nil {
		r.logger.Infof("authService.RefreshToken() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.AuthResult{Token: resp.Token, RefreshToken: resp.RefreshToken}, nil
}

func (r *mutationResolver) ResendOtp(ctx context.Context, phone string) (*types.Result, error) {
	if err := phonenumbervalidator.ValidatePhoneNumber(phone); err != nil {
		r.logger.WithField("phone", phone).WithError(err).
			Error("failed to validate phone number")
		return nil, err
	}
	resp, err := r.onBoardingService.ResendOtp(ctx, &onboardingService.ResendOtpRequest{Phone: phone})
	if err != nil {
		r.logger.Infof("onboardingService.ResendOtp() failed: %v", err)
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) CheckEmailExistence(ctx context.Context, email string) (*types.CheckEmailExistenceResult, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}
	resp, err := r.onBoardingService.CheckEmailExistence(ctx, &onboardingService.CheckEmailExistenceRequest{Email: email})
	if err != nil {
		r.logger.WithError(err).Info("onboardingService.checkEmailExistence() failed")
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.CheckEmailExistenceResult{Message: resp.Message, Exists: resp.Exists}, nil
}

func (r *mutationResolver) ActivateBioLogin(ctx context.Context, deviceID string) (*types.ActivateBioLoginResponse, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.authService.ActivateBioLogin(ctx, &authService.ActivateBioLoginRequest{
		PersonId: personId,
		DeviceId: deviceID,
	})

	if err != nil {
		r.logger.WithError(err).Info("authService.ActivateBioLogin() failed")
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.ActivateBioLoginResponse{BiometricPasscode: resp.BiometricPasscode, Message: resp.Message}, nil
}

func (r *mutationResolver) BioLoginRequest(ctx context.Context, input types.BioLoginInput) (*types.AuthResult, error) {
	if err := emailvalidator.Validate(input.Email); err != nil {
		return nil, err
	}
	resp, err := r.authService.BioLogin(ctx, &authService.BioLoginRequest{
		Email:             input.Email,
		BiometricPasscode: input.BiometricPasscode,
		DeviceId:          input.DeviceID,
	})

	if err != nil {
		r.logger.WithError(err).Info("authService.BioLogin() failed")
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.AuthResult{
		Token:        resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		Person: &types.APIPerson{
			FirstName:               resp.Person.FirstName,
			LastName:                resp.Person.LastName,
			Email:                   resp.Person.Email,
			IsEmailActive:           resp.Person.IsEmailActive,
			IsBiometricLoginEnabled: resp.Person.IsBiometricLoginEnabled,
			IsTransactionPinEnabled: resp.Person.IsTransactionPinEnabled,
			RegistrationCheckPoint:  resp.Person.RegistrationCheckPoint,
		},
	}, nil
}

func (r *mutationResolver) DeactivateBioLogin(ctx context.Context, input types.DeactivateBioLoginInput) (*types.Result, error) {
	if err := emailvalidator.Validate(input.Email); err != nil {
		return nil, err
	}
	resp, err := r.authService.DeactivateBioLogin(ctx, &authService.DeactivateBioLoginRequest{
		Email:    input.Email,
		DeviceId: input.DeviceID,
	})
	if err != nil {
		r.logger.WithError(err).Info("authService.DeactivateBioLogin() failed")
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.Result{
		Message: resp.Message,
	}, nil
}

func (r *mutationResolver) VerifyEmailMagicLInk(ctx context.Context, email string, verificationToken string) (*types.Result, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}

	resp, err := r.onBoardingService.VerifyEmailMagicLInk(ctx, &onboardingService.VerifyEmailMagicLInkRequest{
		Email:             email,
		VerificationToken: verificationToken, // The service will validate the token. ( it's an independent logic )
	})

	if err != nil {
		r.logger.WithError(err).Info("onBoardingService.VerifyEmailMagicLInk() failed")
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.Result{
		Success: true,
		Message: resp.Message,
	}, nil
}

func (r *mutationResolver) ResendEmailMagicLInk(ctx context.Context, email string) (*types.Result, error) {
	// Validate email
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}

	resp, err := r.onBoardingService.ResendEmailMagicLInk(ctx, &onboardingService.ResendEmailMagicLInkRequest{Email: email})
	if err != nil {
		r.logger.WithError(err).Info("onBoardingService.ResendEmailMagicLInk() failed")
		return nil, rerrors.NewFromGrpc(err)
	}

	return &types.Result{
		Success: true,
		Message: resp.Message,
	}, nil
}

func (r *mutationResolver) SubmitApplication(ctx context.Context) (*types.Result, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.onBoardingService.SubmitCheck(ctx, &onboardingService.SubmitCheckRequest{
		PersonId: personId,
	})
	if err != nil {
		r.logger.WithError(err).Error("submitCheck() failed")
		return nil, rerrors.NewFromGrpc(err)
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
