package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

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
	"ms.api/protos/pb/verifyService"
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

	tokens := make([]*protoTypes.DeviceToken, len(device.Tokens))

	for k, v := range device.Tokens {
		tokens[k] = &protoTypes.DeviceToken{
			Type:  string(v.Type),
			Value: v.Value,
		}
	}

	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboardingService.CreatePhoneRequest{
			PhoneNumber: phone,
			Device: &protoTypes.Device{
				Identifier: device.Identifier,
				Brand:      device.Brand,
				Os:         device.Os,
				Tokens:     tokens,
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

	// Validate request
	if err := datevalidator.ValidateDob(person.Dob); err != nil {
		return nil, err
	}
	if err := r.validateAddress(address); err != nil {
		return nil, err
	}

	// Build person bio data payload
	payload := onboardingService.UpdatePersonRequest{
		PersonId:         personId.PersonId,
		Address:          &onboardingService.InputAddress{},
		FirstName:        strings.TrimSpace(person.FirstName),
		LastName:         strings.TrimSpace(person.LastName),
		Dob:              person.Dob,
		CountryResidence: person.CountryResidence,
	}
	if address.Country != nil {
		payload.Address.Country = *address.Country
	}
	if address.Street != nil {
		payload.Address.Street = *address.Street
	}
	if address.City != nil {
		payload.Address.City = *address.City
	}
	if address.Postcode != nil {
		payload.Address.Postcode = *address.Postcode
	}
	if address.Country2 != nil {
		payload.Address.Country2 = *address.Country2
	}
	if address.State != nil {
		payload.Address.State = *address.State
	}
	if address.County != nil {
		payload.Address.County = *address.County
	}
	if person.Bvn != nil {
		payload.Bvn = *person.Bvn
	}
	// Call onboarding service
	response, err := r.onBoardingService.UpdatePersonBioData(context.Background(), &payload)
	if err != nil {
		return nil, err
	}

	// Build response
	identities := make([]*types.Identity, len(response.Identities))
	emails := make([]*types.Email, len(response.Emails))
	phones := make([]*types.Phone, len(response.Phones))
	addresses := make([]*types.Address, len(response.Addresses))
	nationality := make([]*string, len(response.Nationality))
	for index, id := range response.Identities {
		identities[index] = &types.Identity{
			ID:             id.Id,
			Nickname:       &id.Nickname,
			Active:         &id.Active,
			Authentication: &id.Authentication,
		}
	}
	for index, email := range response.Emails {
		emails[index] = &types.Email{
			Value:    email.Value,
			Verified: email.Verified,
		}
	}
	for index, phone := range response.Phones {
		phones[index] = &types.Phone{
			Value:    phone.Number,
			Verified: phone.Verified,
		}
	}
	for index, addr := range response.Addresses {
		if addr == nil {
			continue
		}

		addresses[index] = &types.Address{
			Primary:  &addr.Primary,
			Street:   &addr.Street,
			City:     &addr.City,
			County:   &addr.County,
			State:    &addr.State,
			Postcode: &addr.Postcode,
			Country:  &types.Country{CountryName: addr.Country},
		}
	}
	for index, next := range response.Nationality {
		nationality[index] = &next
	}
	return &types.Person{
		ID:               response.Id,
		Title:            &response.Title,
		FirstName:        response.FirstName,
		LastName:         response.LastName,
		MiddleName:       &response.MiddleName,
		Phones:           phones,
		Emails:           emails,
		Dob:              response.Dob,
		CountryResidence: &response.CountryResidence,
		Nationality:      nationality,
		Addresses:        addresses,
		Identities:       identities,
		Ts:               response.Ts,
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
	tokens := make([]*protoTypes.DeviceToken, len(credentials.Device.Tokens))

	for k, v := range credentials.Device.Tokens {
		tokens[k] = &protoTypes.DeviceToken{
			Type:  string(v.Type),
			Value: v.Value,
		}
	}
	// TODO: change authService.LoginRequest{}.Tokens to slice datatype
	req := &authService.LoginRequest{
		Email:    newEmail,
		Passcode: credentials.Passcode,
		Device: &protoTypes.Device{
			Os:         credentials.Device.Os,
			Brand:      credentials.Device.Brand,
			Identifier: credentials.Device.Identifier,
			Tokens:     tokens,
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

func (r *mutationResolver) LoginWithToken(ctx context.Context, token string, authType types.AuthType) (*types.AuthResponse, error) {
	req := &authService.LoginWithTokenRequest{Token: token, AuthType: authType.String()}
	resp, err := r.authService.LoginWithToken(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.LoginWithToken() failed: %v", err))
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
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	response, err := r.onBoardingService.UpdateValidationStatus(ctx, &onboardingService.UpdateValidationStatusRequest{
		Validation: validation,
		Status:     string(status),
		Message:    message,
		Owner:      claims.PersonId,
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
		Token:   &response.Token,
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

func (r *mutationResolver) ResubmitReports(ctx context.Context, reports []*types.ReportInput) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	request := onboardingService.ResubmitReportRequest{
		PersonId: claims.PersonId,
		Reports:  make([]*onboardingService.ReportInput, len(reports)),
	}
	for index, report := range reports {
		request.Reports[index] = &onboardingService.ReportInput{
			Id: report.ID,
		}
	}
	response, err := r.onBoardingService.ResubmitReport(ctx, &request)
	if err != nil {
		r.logger.Error("call onBoardingService.ResubmitReport()", zap.Error(err))
		return nil, err
	}

	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) CreatePayment(ctx context.Context, payment types.PaymentInput, password string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	p, err := validator.ValidatePayment(&payment)
	if err != nil {
		r.logger.Error("validate payment request details", zap.Error(err))
		return nil, err
	}

	// Build request
	request := &paymentService.CreatePaymentRequest{
		IdentityId:     claims.IdentityId,
		TransactionPin: password,
		IdempotencyKey: p.IdempotencyKey,
		Owner:          p.Owner,
		Tags:           p.Tags,
		Beneficiary: &paymentService.BeneficiaryInput{
			Account: p.Beneficiary.Account,
			Amount:  p.Beneficiary.Amount,
		},
		FundingSource: p.FundingSource,
		FundingAmount: p.FundingAmount,
	}
	if p.Charge != nil {
		request.Charge = *p.Charge
	}
	if p.Reference != nil {
		request.Reference = *p.Reference
	}
	if p.Image != nil {
		request.Image = *p.Image
	}
	if p.Notes != nil {
		request.Notes = *p.Notes
	}
	if p.Quote != nil {
		request.Quote = *p.Quote
	}
	if p.Beneficiary.Currency != nil {
		request.Beneficiary.Currency = *p.Beneficiary.Currency
	}
	if p.Currency != nil {
		request.Currency = *p.Currency
	}
	// Make call
	response, err := r.paymentService.CreatePayment(ctx, request)
	if err != nil {
		r.logger.Error("error calling paymentService.CreatePayment()", zap.Error(err))
		return nil, err
	}

	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) ValidateBvn(ctx context.Context, bvn string, phone string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	response, err := r.accountService.ValidateBVN(ctx, &accountService.ValidateBVNRequest{
		PersonId: claims.PersonId,
		Bvn:      bvn,
		Phone:    phone,
	})
	if err != nil {
		r.logger.Error("error calling accountService.ValidateBVN()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) RequestOtp(ctx context.Context, typeArg types.DeliveryMode, target string, expireTime *int64) (*types.Response, error) {
	// Build request
	requestOTPRequest := verifyService.RequestOTPRequest{
		DeliveryMode: typeArg.String(),
		Target:       target,
	}
	if expireTime != nil {
		requestOTPRequest.ExpireTime = *expireTime
	}
	response, err := r.verifyService.RequestOTP(ctx, &requestOTPRequest)
	if err != nil {
		r.logger.Error("verifyService request OTP", zap.Error(err))
		return nil, err
	}

	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, target string, token string) (*types.Response, error) {
	// Build request
	response, err := r.verifyService.VerifyOTP(ctx, &verifyService.VerifyOTPRequest{
		Target: target,
		Token:  token,
	})
	if err != nil {
		r.logger.Error("verifyService verify OTP", zap.Error(err))
		return nil, err
	}

	return &types.Response{
		Message: response.Message,
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) ValidateEmail(ctx context.Context, email string, device types.DeviceInput) (*types.Response, error) {
	newEmail, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return nil, err
	}
	resp, err := r.authService.ValidateEmail(ctx,
		&authService.ValidateEmailRequest{
			Email: newEmail,
			Device: &protoTypes.Device{
				Identifier: device.Identifier,
				Brand:      device.Brand,
				Os:         device.Os,
			},
		},
	)
	if err != nil {
		r.logger.Error("error calling authService.ValidateEmail()", zap.Error(err))
		return nil, err
	}
	code := int64(resp.Code)
	return &types.Response{
		Message: resp.Message,
		Success: resp.Success,
		Code:    &code,
	}, nil
}

func (r *mutationResolver) ValidateUser(ctx context.Context, user types.ValidateUserInput) (*types.Response, error) {
	newEmail, err := emailvalidator.Validate(user.Email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", user.Email))
		return nil, err
	}
	resp, err := r.authService.ValidateUser(ctx,
		&authService.ValidateUserRequest{
			Email:         newEmail,
			FirstName:     strings.TrimSpace(user.FirstName),
			LastName:      strings.TrimSpace(user.LastName),
			DOB:           user.Dob,
			AccountNumber: user.AccountNumber,
			SortCode:      user.SortCode,
			Device: &protoTypes.Device{
				Identifier: user.Device.Identifier,
				Brand:      user.Device.Brand,
				Os:         user.Device.Os,
			},
		},
	)
	if err != nil {
		r.logger.Error("error calling authService.ValidateUser()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: resp.Message,
		Success: resp.Success,
	}, nil
}

func (r *mutationResolver) RequestTransactionPasscodeReset(ctx context.Context, email string) (*types.Response, error) {
	res, err := r.identityService.RequestResetTransactionPassword(ctx, &identityService.RequestResetTransactionPasswordRequest{
		Email: email,
	})
	if err != nil {
		r.logger.Error("identity service request reset transaction password", zap.Error(err))
		return nil, err
	}

	return &types.Response{
		Message: res.Message,
		Success: res.Success,
	}, nil
}

func (r *mutationResolver) ResetTransactionPasscode(ctx context.Context, email string, currentPasscode string, newPasscode string) (*types.Response, error) {
	res, err := r.identityService.ResetTransactionPassword(ctx, &identityService.ResetTransactionPasswordRequest{
		Email:           email, // TODO: Update GraphQL
		CurrentPasscode: currentPasscode,
		NewPasscode:     newPasscode,
	})
	if err != nil {
		r.logger.Error("identity service reset transaction password", zap.Error(err))
		return nil, err
	}

	return &types.Response{
		Message: res.Message,
		Success: res.Success,
	}, nil
}

func (r *mutationResolver) SetDevicePreference(ctx context.Context, typeArg types.DevicePreferenceType, status bool) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	res, err := r.identityService.SetDevicePreference(ctx, &identityService.SetDevicePreferenceRequest{
		IdentityId:           claims.IdentityId,
		DeviceId:             claims.DeviceId,
		DevicePreferenceType: string(typeArg),
		Status:               status,
	})
	if err != nil {
		r.logger.Error("identity service set device preference", zap.Error(err))
		return nil, err
	}

	return &types.Response{
		Message: res.Message,
		Success: res.Success,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
