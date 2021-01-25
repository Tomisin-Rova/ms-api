package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"ms.api/graph/generated"
	"ms.api/libs/validator/datevalidator"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/payeeService"
	"ms.api/protos/pb/paymentService"
	"ms.api/protos/pb/personService"
	"ms.api/protos/pb/productService"
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
		return nil, err
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
		r.logger.Info(fmt.Sprintf("authService.ConfirmPasswordResetDetails() failed: %v", err))
		return nil, err
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
		r.logger.Info(fmt.Sprintf("authService.ConfirmPasswordResetOtp() failed: %v", err))
		return nil, err
	}

	return &types.Result{
		Success: true,
		Message: result.Message,
	}, nil
}

func (r *mutationResolver) UpdatePersonBiodata(ctx context.Context, input *types.UpdateBioDataInput) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
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
		PersonId: claims.PersonId,
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
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) AddReasonsForUsingRoava(ctx context.Context, reasonValues []*string) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	reasons := make([]string, 0, len(reasonValues))
	for _, reason := range reasonValues {
		reasons = append(reasons, *reason)
	}
	payload := onboardingService.RoavaReasonsRequest{
		PersonId: claims.PersonId,
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

func (r *mutationResolver) CreatePhone(ctx context.Context, phone string, device types.DeviceInput) (*types.Response, error) {
	if err := phonenumbervalidator.ValidatePhoneNumber(phone); err != nil {
		r.logger.Error("failed to validate phone number",
			zap.Error(err),
			zap.String("phone", phone),
		)
		return nil, err
	}
	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboardingService.CreatePhoneRequest{PhoneNumber: phone,
			Device: &onboardingService.Device{Os: device.Os, Brand: device.Brand,
				DeviceId: device.Identifier, DeviceToken: device.Tokens.Firebase}})
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

func (r *mutationResolver) Signup(ctx context.Context, token string, email string, passcode string) (*types.AuthResponse, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}

	result, err := r.onBoardingService.CreatePerson(ctx, &onboardingService.CreatePersonRequest{
		Email:    email,
		Passcode: passcode,
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
			Jwt:     result.JwtToken,
			Refresh: result.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, credentials types.AuthInput, biometric *bool) (*types.AuthResponse, error) {
	if err := emailvalidator.Validate(credentials.Email); err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", credentials.Email))
		return nil, err
	}
	bio := biometric != nil && *biometric
	req := &authService.LoginRequest{
		Email:     credentials.Email,
		Passcode:  credentials.Passcode,
		Biometric: bio,
		Device: &authService.Device{
			Os:          credentials.Device.Os,
			Brand:       credentials.Device.Brand,
			DeviceToken: credentials.Device.Tokens.Firebase,
			DeviceId:    credentials.Device.Identifier,
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
			Jwt:     resp.AccessToken,
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
			Jwt:     resp.Token,
			Refresh: resp.RefreshToken,
		},
	}, nil
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
		resp, err := r.authService.ActivateBioLogin(ctx, &authService.ActivateBioLoginRequest{
			IdentityId: claims.IdentityId,
			DeviceId:   claims.DeviceId,
		})
		if err != nil {
			r.logger.Info("authService.ActivateBioLogin() failed", zap.Error(err))
			return nil, err
		}
		return &types.Response{Message: resp.Message, Success: true, Token: &resp.BiometricPasscode}, nil
	}
}

func (r *mutationResolver) ResendOtp(ctx context.Context, phone string) (*types.Result, error) {
	if err := phonenumbervalidator.ValidatePhoneNumber(phone); err != nil {
		r.logger.Error("failed to validate phone number", zap.Error(err), zap.String("phone", phone))
		return nil, err
	}
	resp, err := r.onBoardingService.ResendOtp(ctx, &onboardingService.ResendOtpRequest{Phone: phone})
	if err != nil {
		r.logger.Info(fmt.Sprintf("onboardingService.ResendOtp() failed: %v", err))
		return nil, err
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) VerifyEmail(ctx context.Context, email string, token string) (*types.Result, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}

	resp, err := r.onBoardingService.VerifyEmailMagicLInk(ctx, &onboardingService.VerifyEmailMagicLInkRequest{
		Email:             email,
		VerificationToken: token,
	})
	if err != nil {
		r.logger.Error("error calling onboardingService.VerifyEmailMagicLInk.", zap.Error(err))
		return nil, err
	}

	return &types.Result{
		Success: true,
		Message: resp.Message,
	}, nil
}

func (r *mutationResolver) ResendEmailMagicLInk(ctx context.Context, email string) (*types.Result, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}

	resp, err := r.onBoardingService.ResendEmailMagicLInk(ctx, &onboardingService.ResendEmailMagicLInkRequest{Email: email})
	if err != nil {
		r.logger.Info("OnBoardingService.ResendEmailMagicLInk() failed", zap.Error(err))
		return nil, err
	}

	return &types.Result{
		Success: true,
		Message: resp.Message,
	}, nil
}

func (r *mutationResolver) SubmitApplication(ctx context.Context) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.onBoardingService.SubmitCheck(ctx, &onboardingService.SubmitCheckRequest{
		PersonId: claims.PersonId,
	})
	if err != nil {
		r.logger.Error("submitCheck() failed", zap.Error(err))
		return nil, err
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) AcceptTermsAndConditions(ctx context.Context) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.onBoardingService.AcceptTermsAndConditions(ctx, &onboardingService.TermsAndConditionsRequest{
		PersonId: claims.PersonId,
	})
	if err != nil {
		r.logger.Error("AcceptTermsAndConditions() failed", zap.Error(err))
		return nil, err
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) CreateApplication(ctx context.Context) (*types.CreateApplicationResponse, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.onBoardingService.CreateApplication(ctx,
		&onboardingService.CreateApplicationRequest{PersonId: claims.PersonId})
	if err != nil {
		r.logger.Error("onBoardingService.createApplication() failed", zap.Error(err))
		return nil, err
	}
	return &types.CreateApplicationResponse{Token: resp.Token}, nil
}

func (r *mutationResolver) CreateCurrencyAccount(ctx context.Context, currencyCode string) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.productService.CreateAccount(ctx, &productService.CreateAccountRequest{
		PersonId: claims.PersonId,
		Currency: currencyCode,
	})
	if err != nil {
		r.logger.Error("productService.createAccount() failed", zap.Error(err))
		return nil, err
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) UpdateFirebaseToken(ctx context.Context, token string) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.onBoardingService.UpdateFirebaseToken(ctx,
		&onboardingService.UpdateFirebaseTokenRequest{PersonId: claims.PersonId, Token: token})
	if err != nil {
		r.logger.Error("onBoardingService.UpdateFirebaseToken() failed", zap.Error(err))
		return nil, err
	}
	return &types.Result{Message: resp.Message}, nil
}

func (r *mutationResolver) CreateTransactionPin(ctx context.Context, pin string) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	if len(pin) != 4 {
		r.logger.Error("invalid pin length")
		return nil, errors.New("he transaction pin has to contain only 4 digits")
	}
	resp, err := r.personService.CreateTransactionPin(ctx,
		&personService.CreateTransactionPinRequest{PersonId: claims.PersonId, Pin: pin})
	if err != nil {
		r.logger.Error("personService.CreateTransactionPin() failed", zap.Error(err))
		return nil, err
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) MakeTransfer(ctx context.Context, input *types.MakeTransferInput) (*types.Result, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	if input.Amount <= 0 {
		return nil, errors.New("invalid transaction amount")
	}
	if len(input.ToAccountNumber) < 10 || len(input.FromAccountNumber) < 10 {
		return nil, errors.New("invalid account details")
	}
	if len(input.Notes) == 0 {
		return nil, errors.New("a transaction note is required")
	}
	resp, err := r.paymentService.MakeTransfer(ctx, &paymentService.TransferRequest{
		AccountNumber:            input.FromAccountNumber,
		BeneficiaryAccountNumber: input.ToAccountNumber,
		Amount:                   input.Amount,
		Notes:                    input.Notes,
		PersonId:                 claims.PersonId,
		TransactionPin:           input.TransactionPin,
	})
	if err != nil {
		r.logger.Error("paymentService.MakeTransfer() failed", zap.Error(err))
		return nil, err
	}
	return &types.Result{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) CreatePayee(ctx context.Context, input types.CreatePayeeInput) (*types.CreatePayeeResult, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	req := &payeeService.CreatePayeeRequest{
		PersonId:       claims.PersonId,
		TransactionPin: input.TransactionPin,
		Name:           input.Name,
		Country:        input.Country,
		AccountNumber:  input.AccountNumber,
	}
	switch input.Country {
	case "UK":
		if input.SortCode == nil {
			return nil, errors.New("Sort code required for a UK account")
		}
		req.SortCode = *input.SortCode
	case "Nigeria":
		if input.BankCode == nil || input.BankName == nil {
			return nil, errors.New("Bank code and name required for a Nigerian account")
		}
		req.BankCode = *input.BankCode
		req.BankName = *input.BankName
	case "US":
		if input.RoutingNumber == nil || input.RoutingType == nil {
			return nil, errors.New("Routing number / type required for a US account")
		}
		req.RoutingNumber = *input.RoutingNumber
		req.RoutingType = *input.RoutingType
	}
	resp, err := r.PayeeService.CreatePayee(ctx, req)

	if err != nil {
		r.logger.Error("payeeService.CreateBeneficiary() failed", zap.Error(err))
		return nil, err
	}
	var account *types.PayeeAccount
	if len(resp.Payee.Accounts) > 0 {
		account = &types.PayeeAccount{
			AccountNumber: resp.Payee.Accounts[0].AccountNumber,
			RoutingNumber: &resp.Payee.Accounts[0].RoutingNumber,
			Bic:           &resp.Payee.Accounts[0].BIC,
			Iban:          &resp.Payee.Accounts[0].IBAN,
			Country:       resp.Payee.Accounts[0].Country,
			BankName:      &resp.Payee.Accounts[0].BankName,
			BankCode:      &resp.Payee.Accounts[0].BankCode,
			RoutingType:   &resp.Payee.Accounts[0].RoutingType,
			SortCode:      &resp.Payee.Accounts[0].SortCode,
		}
	}
	return &types.CreatePayeeResult{
		Success: true,
		Message: resp.Message,
		Beneficiary: &types.Beneficiary{
			PayeeID: resp.Payee.PayeeId,
			Owner:   resp.Payee.Owner,
			Name:    resp.Payee.Name,
			Accounts: []*types.PayeeAccount{
				account,
			},
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
