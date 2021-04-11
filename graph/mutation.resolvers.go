package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"ms.api/graph/generated"
	"ms.api/libs/validator"
	"ms.api/libs/validator/datevalidator"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/accountService"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/identityService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/paymentService"
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

	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboardingService.CreatePhoneRequest{
			PhoneNumber: phone,
			Device: &protoTypes.Device{
				Identifier: device.Identifier,
				Brand:      device.Brand,
				Os:         device.Os,
			},
		},
	)
	if err != nil {
		r.logger.Error("error calling onBoardingService.CreatePhone()", zap.Error(err))
		return nil, err
	}
	return &types.Response{Message: result.Message, Success: true, Token: &result.Token}, nil
}

func (r *mutationResolver) ConfirmPhone(ctx context.Context, token string, code string) (*types.Response, error) {
	resp, err := r.onBoardingService.VerifySmsOtp(context.Background(), &onboardingService.OtpVerificationRequest{
		Token: token, Code: code,
	})
	if err != nil {
		r.logger.Error("error calling onBoardingService.ConfirmPhone()", zap.Error(err))
		return nil, err
	}
	return &types.Response{Success: resp.Match, Message: resp.Message}, nil
}

func (r *mutationResolver) Signup(ctx context.Context, token string, email string, passcode string) (*types.AuthResponse, error) {
	newEmail, err := emailvalidator.Validate(email)
	if err != nil {
		return nil, err
	}
	// Validate passCode
	if err := validator.IsValidPassCode(passcode); err != nil {
		return nil, err
	}

	result, err := r.onBoardingService.CreatePerson(ctx, &onboardingService.CreatePersonRequest{
		Email:    newEmail,
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
			Auth:    result.JwtToken,
			Refresh: result.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) Register(ctx context.Context, person types.PersonInput, address types.AddressInput) (*types.Person, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	if err := datevalidator.ValidateDob(person.Dob); err != nil {
		return nil, err
	}
	if err := r.validateAddress(address); err != nil {
		return nil, err
	}
	postCode, bvn := "", ""
	if address.Postcode != nil {
		postCode = *address.Postcode
	}
	if person.Bvn != nil {
		bvn = *person.Bvn
	}

	payload := onboardingService.UpdatePersonRequest{
		PersonId: personId.PersonId,
		Address: &onboardingService.InputAddress{
			Postcode: postCode, Street: *address.Street,
			City: *address.City, Country: *address.Country,
		},
		FirstName:        person.FirstName,
		LastName:         person.LastName,
		Dob:              person.Dob,
		CountryResidence: person.CountryResidence,
		Bvn:              bvn,
	}
	res, err := r.onBoardingService.UpdatePersonBiodata(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	identities := make([]*types.Identity, 0)
	emails := make([]*types.Email, 0)
	phones := make([]*types.Phone, 0)
	addresses := make([]*types.Address, 0)
	for _, id := range res.Identities {
		identities = append(identities, &types.Identity{
			ID:             id.Id,
			Nickname:       &id.Nickname,
			Active:         &id.Active,
			Authentication: &id.Authentication,
		})
	}
	for _, email := range res.Emails {
		emails = append(emails, &types.Email{
			Value:    email.Value,
			Verified: email.Verified,
		})
	}
	for _, phone := range res.Phones {
		phones = append(phones, &types.Phone{
			Value:    phone.Number,
			Verified: phone.Verified,
		})
	}
	for _, addr := range res.Addresses {
		addresses = append(addresses, &types.Address{
			Street:   &addr.Street,
			State:    &addr.State,
			Postcode: &addr.Postcode,
			Country:  &types.Country{CountryName: addr.Country},
		})
	}
	nationality := make([]*string, 0)
	for _, next := range res.Nationality {
		nationality = append(nationality, &next)
	}
	return &types.Person{
		ID:               res.Id,
		Title:            &res.Title,
		FirstName:        res.FirstName,
		LastName:         res.LastName,
		MiddleName:       &res.MiddleName,
		Phones:           phones,
		Emails:           emails,
		Dob:              res.Dob,
		CountryResidence: &res.CountryResidence,
		Nationality:      nationality,
		Addresses:        addresses,
		Identities:       identities,
		Ts:               int64(res.Ts),
	}, nil
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

func (r *mutationResolver) CreateApplication(ctx context.Context) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	req := &onboardingService.CreateOnfidoApplicantRequest{
		PersonId: claims.PersonId,
	}
	resp, err := r.onBoardingService.CreateOnfidoApplicant(ctx, req)
	if err != nil {
		r.logger.Error("create applicant request failed", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: "successful",
		Success: true,
		Token:   &resp.Token,
	}, nil
}

func (r *mutationResolver) VerifyEmail(ctx context.Context, email string, code string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResendOtp(ctx context.Context, phone string) (*types.Response, error) {
	if err := phonenumbervalidator.ValidatePhoneNumber(phone); err != nil {
		r.logger.Error("failed to validate phone number",
			zap.Error(err),
			zap.String("phone", phone),
		)
		return nil, err
	}

	result, err := r.onBoardingService.ResendOtp(ctx,
		&onboardingService.ResendOtpRequest{
			Phone: phone,
		},
	)
	if err != nil {
		r.logger.Error("error calling onBoardingService.ResendOtp()", zap.Error(err))
		return nil, err
	}
	return &types.Response{Message: result.Message, Success: true, Token: &result.Token}, nil
}

func (r *mutationResolver) ResendEmailMagicLInk(ctx context.Context, email string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, credentials types.AuthInput) (*types.AuthResponse, error) {
	newEmail, err := emailvalidator.Validate(credentials.Email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", credentials.Email))
		return nil, err
	}
	// TODO: change authService.LoginRequest{}.Tokens to slice datatype
	req := &authService.LoginRequest{
		Email:    newEmail,
		Passcode: credentials.Passcode,
		Device: &protoTypes.Device{
			Os:         credentials.Device.Os,
			Brand:      credentials.Device.Brand,
			Identifier: credentials.Device.Identifier,
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

func (r *mutationResolver) UpdateDeviceToken(ctx context.Context, token []*types.DeviceTokenInput) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	// Get tokens
	var requestTokens []*identityService.DeviceTokens
	for _, token := range token {
		if token != nil {
			requestTokens = append(requestTokens, &identityService.DeviceTokens{
				Type:  string(token.Type),
				Value: token.Value,
			})
		}
	}
	response, err := r.identityService.UpdateDeviceTokens(ctx, &identityService.UpdateDeviceTokensRequest{
		DeviceId:   claims.DeviceId,
		IdentityId: claims.IdentityId,
		Tokens:     requestTokens,
	})
	if err != nil {
		r.logger.Error("error calling identityService.UpdateDeviceTokens()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: "successful",
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) ResetPasscode(ctx context.Context, token string, email string, passcode string) (*types.Response, error) {
	newEmail, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return nil, err
	}
	// Validate passCode
	if err := validator.IsValidPassCode(passcode); err != nil {
		return nil, err
	}

	req := &authService.PasswordResetRequest{
		Email:             newEmail,
		NewPassword:       passcode,
		VerificationToken: token,
	}
	resp, err := r.authService.ResetPassword(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.ResetPasscode() failed: %v", err))
		return nil, err
	}
	return &types.Response{
		Message: resp.Message,
		Success: true,
	}, nil
}

func (r *mutationResolver) RequestPasscodeReset(ctx context.Context, email string, device types.DeviceInput) (*types.Response, error) {
	newEmail, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return nil, err
	}
	req := &authService.PasswordResetUserDetails{
		Email: newEmail,
		Device: &protoTypes.Device{
			Os:         device.Os,
			Brand:      device.Brand,
			Identifier: device.Identifier,
		}}
	resp, err := r.authService.ConfirmPasswordResetDetails(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.RequestPasscodeReset() failed: %v", err))
		return nil, err
	}
	return &types.Response{
		Message: resp.Message,
		Success: true,
	}, nil
}

func (r *mutationResolver) ConfirmPasscodeResetOtp(ctx context.Context, email string, otp string) (*types.Response, error) {
	newEmail, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return nil, err
	}
	req := &authService.PasswordResetOtpRequest{
		Email: newEmail,
		Code:  otp,
	}
	resp, err := r.authService.ConfirmPasswordResetOtp(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.ConfirmPasscodeResetOtp() failed: %v", err))
		return nil, err
	}
	return &types.Response{
		Token:   &resp.VerificationToken,
		Message: resp.Message,
		Success: true,
	}, nil
}

func (r *mutationResolver) SubmitApplication(ctx context.Context) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	// Get tokens
	response, err := r.onBoardingService.SubmitApplication(ctx, &onboardingService.SubmitApplicationRequest{
		PersonId: claims.PersonId,
	})
	if err != nil {
		r.logger.Error("onBoardingService.SubmitApplication()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) AcceptTerms(ctx context.Context, documents []*string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateValidationStatus(ctx context.Context, validation string, status types.State, message string) (*types.Response, error) {
	response, err := r.onBoardingService.UpdateValidationStatus(ctx, &onboardingService.UpdateValidationStatusRequest{
		Validation: validation,
		Status:     string(status),
		Message:    message,
	})
	if err != nil {
		r.logger.Error("onBoardingService.UpdateValidationStatus()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) SubmitProof(ctx context.Context, proof types.SubmitProofInput) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	response, err := r.onBoardingService.SubmitProof(ctx, &onboardingService.SubmitProofRequest{
		Owner:  claims.PersonId,
		Type:   proof.Type.String(),
		Data:   proof.Data,
		Status: proof.Status.String(),
	})
	if err != nil {
		r.logger.Error("onBoardingService.SubmitProof()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) CreateTransactionPassword(ctx context.Context, password string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	if err := validator.ValidateTransactionPassword(password); err != nil {
		r.logger.Error("error validating password", zap.Error(err))
		return nil, err
	}
	response, err := r.identityService.CreateTransactionPassword(ctx, &identityService.CreateTransactionPasswordRequest{
		PersonId: claims.PersonId,
		Pin:      password,
	})
	if err != nil {
		r.logger.Error("error calling identityService.createTransactionPin()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, product types.ProductInput) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	// Build request
	request := accountService.CreateAccountRequest{
		IdentityId: claims.IdentityId,
		Product: &protoTypes.ProductInput{
			Id: product.ID,
		},
	}
	if product.Identification != nil {
		request.Product.Identification = *product.Identification
	}
	if product.Scheme != nil {
		request.Product.Scheme = *product.Scheme
	}
	// Execute create account
	response, err := r.accountService.CreateAccount(ctx, &request)
	if err != nil {
		r.logger.Error("error calling accountService.CreateAccount()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) CreatePayee(ctx context.Context, payee types.PayeeInput, password string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	payeeAccount, err := validator.ValidatePayeeAccount(payee.Accounts[0])
	if err != nil {
		r.logger.Error("validating payee account details", zap.Error(err))
		return nil, err
	}
	avatar := ""
	if payee.Avatar != nil {
		avatar = *payee.Avatar
	}
	response, err := r.paymentService.CreatePayee(ctx, &paymentService.CreatePayeeRequest{
		IdentityId:     claims.IdentityId,
		TransactionPin: password,
		Name:           payee.Name,
		Avatar:         avatar,
		AccountName:    payeeAccount.Name,
		AccountNumber:  payeeAccount.AccountNumber,
		SortCode:       payeeAccount.SortCode,
		BankCode:       payeeAccount.BankCode,
		Iban:           payeeAccount.Iban,
		SwiftBic:       payeeAccount.SwiftBic,
		RoutingNumber:  payeeAccount.RoutingNumber,
		PhoneNumber:    payeeAccount.PhoneNumber,
		Currency:       payeeAccount.Currency,
	})
	if err != nil {
		r.logger.Error("error calling paymentService.CreatePayee()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) UpdatePayee(ctx context.Context, payee string, payeeInput *types.PayeeInput, password string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	avatar := ""
	if payeeInput.Avatar != nil {
		avatar = *payeeInput.Avatar
	}
	response, err := r.paymentService.UpdatePayee(ctx, &paymentService.UpdatePayeeRequest{
		IdentityId:     claims.IdentityId,
		TransactionPin: password,
		Name:           payeeInput.Name,
		Avatar:         avatar,
		PayeeId:        payee,
	})
	if err != nil {
		r.logger.Error("error calling paymentService.UpdatePayee()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) AddPayeeAccount(ctx context.Context, payee string, payeeAccount types.PayeeAccountInput) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	p, err := validator.ValidatePayeeAccount(&payeeAccount)
	if err != nil {
		r.logger.Error("validating payee account details", zap.Error(err))
		return nil, err
	}
	response, err := r.paymentService.AddPayeeAccount(ctx, &paymentService.AddPayeeAccountRequest{
		IdentityId:    claims.IdentityId,
		PayeeId:       payee,
		AccountName:   p.Name,
		AccountNumber: p.AccountNumber,
		SortCode:      p.SortCode,
		BankCode:      p.BankCode,
		Iban:          p.Iban,
		SwiftBic:      p.SwiftBic,
		RoutingNumber: p.RoutingNumber,
		PhoneNumber:   p.PhoneNumber,
		Currency:      p.Currency,
	})
	if err != nil {
		r.logger.Error("error calling paymentService.AddPayeeAccount()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) DeletePayeeAccount(ctx context.Context, payee string, payeeAccount string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	response, err := r.paymentService.DeletePayeeAccount(ctx, &paymentService.DeletePayeeAccountRequest{
		IdentityId:     claims.IdentityId,
		PayeeId:        payee,
		PayeeAccountId: payeeAccount,
	})
	if err != nil {
		r.logger.Error("error calling paymentService.DeletePayeeAccount()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) Resubmit(ctx context.Context, reports []*types.ReportInput, message *string) (*types.Response, error) {
	reportIds := make([]string, 0)
	for _, reportId := range reports {
		reportIds = append(reportIds, reportId.ID)
	}
	response, err := r.onBoardingService.Resubmit(ctx, &onboardingService.ResubmitRequest{
		ReportIds: reportIds,
		Message:   *message,
	})
	if err != nil {
		r.logger.Error("onBoardingService.Resubmit()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) CreatePayment(ctx context.Context, payment types.PaymentInput) (*types.Response, error) {
	_, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	response, err := r.paymentService.CreatePayment(ctx, &paymentService.CreatePaymentRequest{
		IdempotencyKey: payment.IdempotencyKey,
		Owner:          payment.Owner,
		Beneficiary: &paymentService.Beneficiary{
			Account:  payment.Beneficiary.Account,
			Currency: *payment.Beneficiary.Currency,
			Amount:   *payment.Beneficiary.Amount,
		},
		Charge:        *payment.Charge,
		Reference:     *payment.Reference,
		Status:        string(*payment.Status),
		Image:         *payment.Image,
		Notes:         *payment.Notes,
		Tags:          payment.Tags,
		FundingSource: payment.FundingSource,
		Currency:      *payment.Currency,
		FundingAmount: payment.FundingAmount,
		Quote:         *payment.Quote,
	})
	if err != nil {
		r.logger.Error("error calling paymentService.CreatePayment()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		// Success: response.Success,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
