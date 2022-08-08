package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ms.api/graph/generated"
	errorvalues "ms.api/libs/errors"
	"ms.api/libs/validator/datevalidator"
	devicevalidator "ms.api/libs/validator/device"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	accountPb "ms.api/protos/pb/account"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/messaging"
	"ms.api/protos/pb/onboarding"
	"ms.api/protos/pb/payment"
	"ms.api/protos/pb/pricing"
	pbTypes "ms.api/protos/pb/types"
	"ms.api/protos/pb/verification"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

// RequestOtp is the resolver for the requestOTP field.
func (r *mutationResolver) RequestOtp(ctx context.Context, typeArg types.DeliveryMode, target string, expireTimeInSeconds *int64) (*types.Response, error) {
	const defaultExpirationTime = 60
	// Build request
	request := verification.RequestOTPRequest{
		Target:              target,
		ExpireTimeInSeconds: defaultExpirationTime,
	}
	if expireTimeInSeconds != nil {
		request.ExpireTimeInSeconds = int32(*expireTimeInSeconds)
	}
	switch typeArg {
	case types.DeliveryModeEmail:
		request.Type = verification.RequestOTPRequest_EMAIL
	case types.DeliveryModeSms:
		request.Type = verification.RequestOTPRequest_SMS
	case types.DeliveryModePush:
		request.Type = verification.RequestOTPRequest_PUSH
	}

	// Execute RPC call
	response, err := r.VerificationService.RequestOTP(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// VerifyOtp is the resolver for the verifyOTP field.
func (r *mutationResolver) VerifyOtp(ctx context.Context, target string, otpToken string) (*types.Response, error) {
	// Build request
	request := verification.VerifyOTPRequest{
		Target:   target,
		OtpToken: otpToken,
	}
	// Execute RPC call
	response, err := r.VerificationService.VerifyOTP(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, customer types.CustomerInput) (*types.AuthResponse, error) {
	err := r.phoneValidator.ValidatePhoneNumber(customer.Phone)
	if err != nil {
		invalidMsg := phonenumbervalidator.ErrInvalidPhoneNumber.Message()
		return &types.AuthResponse{
			Message: &invalidMsg,
			Success: false,
			Code:    http.StatusBadRequest,
		}, err
	}
	email, err := r.emailValidator.Validate(customer.Email)
	if err != nil {
		invalidMsg := emailvalidator.ErrInvalidEmail.Message()
		return &types.AuthResponse{
			Message: &invalidMsg,
			Success: false,
			Code:    http.StatusBadRequest,
		}, err
	}
	err = r.deviceValidator.Validate(customer.Device)
	if err != nil {
		invalidMsg := devicevalidator.ErrInvalidDevice.Message()
		return &types.AuthResponse{
			Message: &invalidMsg,
			Success: false,
			Code:    http.StatusBadRequest,
		}, err
	}
	customerInputTokens := customer.Device.Tokens
	if customerInputTokens == nil {
		customerInputTokens = []*types.DeviceTokenInput{}
	}
	deviceTokenInputs := make([]*pbTypes.DeviceTokenInput, len(customerInputTokens))
	for _, tokenInput := range customerInputTokens {
		deviceTokenInputs = append(deviceTokenInputs, &pbTypes.DeviceTokenInput{
			Value: tokenInput.Value,
			Type:  r.helper.DeviceTokenInputFromModel(tokenInput.Type),
		})
	}
	customerInputPreferences := customer.Device.Preferences
	if customerInputPreferences == nil {
		customerInputPreferences = []*types.DevicePreferencesInput{}
	}
	preferences := make([]*pbTypes.DevicePreferencesInput, len(customerInputPreferences))
	for _, preference := range customerInputPreferences {
		preferences = append(preferences, &pbTypes.DevicePreferencesInput{
			Value: preference.Value,
			Type:  r.helper.PreferenceInputFromModel(preference.Type),
		})
	}
	req := auth.SignupRequest{
		CustomerInput: &auth.CustomerInput{
			Phone:         customer.Phone,
			Email:         email,
			LoginPassword: customer.LoginPassword,
		},
		Device: &pbTypes.DeviceInput{
			Identifier:  customer.Device.Identifier,
			Os:          customer.Device.Os,
			Brand:       customer.Device.Brand,
			Tokens:      deviceTokenInputs,
			Preferences: preferences,
		},
	}
	tokens, err := r.AuthService.Signup(context.Background(), &req)
	if err != nil {
		invalidMsg := errorvalues.Message(errorvalues.InternalErr)
		return &types.AuthResponse{
			Message: &invalidMsg,
			Success: false,
			Code:    http.StatusInternalServerError,
		}, err
	}
	msg := "Success"
	return &types.AuthResponse{
		Message: &msg,
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    tokens.AuthToken,
			Refresh: &tokens.RefreshToken,
		},
		Code: int64(http.StatusOK),
	}, nil
}

// ResetLoginPassword is the resolver for the resetLoginPassword field.
func (r *mutationResolver) ResetLoginPassword(ctx context.Context, otpToken string, email string, loginPassword string) (*types.Response, error) {
	// Build request
	request := customer.ResetLoginPasswordRequest{
		OtpToken:      otpToken,
		Email:         email,
		LoginPassword: loginPassword,
	}
	// Make RPC call
	response, err := r.CustomerService.ResetLoginPassword(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// CheckCustomerEmail is the resolver for the checkCustomerEmail field.
func (r *mutationResolver) CheckCustomerEmail(ctx context.Context, email string, device types.DeviceInput) (*types.Response, error) {
	_, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return nil, err
	}

	deviceTokens := make([]*pbTypes.DeviceTokenInput, len(device.Tokens))
	for index, deviceToken := range device.Tokens {
		deviceTokens[index] = &pbTypes.DeviceTokenInput{
			Type:  r.helper.GetProtoDeviceTokenType(deviceToken.Type),
			Value: deviceToken.Value,
		}
	}

	devicePreferences := make([]*pbTypes.DevicePreferencesInput, len(device.Preferences))
	for index, devicePreference := range device.Preferences {
		devicePreferences[index] = &pbTypes.DevicePreferencesInput{
			Type:  r.helper.GetProtoDevicePreferencesType(devicePreference.Type),
			Value: devicePreference.Value,
		}
	}

	// Build request
	request := &customer.CheckCustomerEmailRequest{
		Email: email,
		Device: &pbTypes.DeviceInput{
			Identifier:  device.Identifier,
			Os:          device.Os,
			Brand:       device.Brand,
			Tokens:      deviceTokens,
			Preferences: devicePreferences,
		},
	}

	// Make RPC call
	resp, err := r.CustomerService.CheckCustomerEmail(ctx, request)
	if err != nil {
		return nil, err
	}

	return &types.Response{Success: resp.Success, Code: int64(resp.Code)}, nil
}

// CheckCustomerData is the resolver for the checkCustomerData field.
func (r *mutationResolver) CheckCustomerData(ctx context.Context, customerData types.CheckCustomerDataInput) (*types.Response, error) {
	_, err := emailvalidator.Validate(customerData.Email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", customerData.Email))
		return nil, err
	}

	if err = datevalidator.ValidateDob(customerData.Dob); err != nil {
		r.logger.Info("invalid Dob supplied", zap.String("Dob", customerData.Dob))
		return nil, err
	}

	// Build request
	request := &customer.CheckCustomerDataRequest{
		Email:            customerData.Email,
		FirstName:        customerData.FirstName,
		LastName:         customerData.LastName,
		Dob:              customerData.Dob,
		AccountNumber:    customerData.AccountNumber,
		SortCode:         customerData.SortCode,
		DeviceIdentifier: customerData.DeviceIdentifier,
	}

	// Make RPC call to customer service
	resp, err := r.CustomerService.CheckCustomerData(ctx, request)
	if err != nil {
		return nil, err
	}

	return &types.Response{Success: resp.Success, Code: int64(resp.Code)}, nil
}

// UpdateCustomerDetails is the resolver for the updateCustomerDetails field.
func (r *mutationResolver) UpdateCustomerDetails(ctx context.Context, customerDetails types.CustomerDetailsUpdateInput, transactionPassword string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	var (
		firstName, lastName, phone, email string
	)
	customerAddress := customerDetails.Address
	if customerAddress == nil {
		customerAddress = &types.AddressInput{}
	}

	var customerCountryId string
	if customerAddress.CountryID != "" {
		customerCountryId = customerAddress.CountryID
	}

	var customerState string
	if customerAddress.State != nil {
		customerState = *customerAddress.State
	}

	var customerCity string
	if customerAddress.City != nil {
		customerCity = *customerAddress.City
	}

	var customerStreet string
	if customerAddress.Street != "" {
		customerStreet = customerAddress.Street
	}
	var customerPostCode string
	if customerAddress.Postcode != "" {
		customerPostCode = customerAddress.Postcode
	}

	customerCoordinates := customerAddress.Cordinates
	if customerCoordinates == nil {
		customerCoordinates = &types.CordinatesInput{}
	}

	if customerDetails.FirstName != nil {
		firstName = *customerDetails.FirstName
	}
	if customerDetails.LastName != nil {
		lastName = *customerDetails.LastName
	}
	if customerDetails.Phone != nil {
		phone = *customerDetails.Phone
	}
	if customerDetails.Email != nil {
		email = *customerDetails.Email
	}
	// Build request
	request := customer.CustomerDetailsUpdateRequest{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		Address: &customer.AddressInput{
			CountryId: customerCountryId,
			State:     customerState,
			City:      customerCity,
			Street:    customerStreet,
			Postcode:  customerPostCode,
			Cordinates: &customer.CordinatesInput{
				Latitude:  float32(customerCoordinates.Latitude),
				Longitude: float32(customerCoordinates.Longitude),
			},
		},
		TransactionPassword: transactionPassword,
	}
	// Execute RPC call
	response, err := r.CustomerService.CustomerDetailsUpdate(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, customerDetails types.CustomerDetailsInput) (*types.Response, error) {
	var responseMessage string
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage = "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	err = datevalidator.ValidateDob(customerDetails.Dob)
	if err != nil {
		responseMessage = "Dob validation failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	customerAddress := customerDetails.Address
	if customerAddress == nil {
		customerAddress = &types.AddressInput{}
	}

	var customerState string
	if customerAddress.State != nil {
		customerState = *customerAddress.State
	}

	var customerCity string
	if customerAddress.City != nil {
		customerCity = *customerAddress.City
	}

	customerCoordinates := customerAddress.Cordinates
	if customerCoordinates == nil {
		customerCoordinates = &types.CordinatesInput{}
	}

	customerReq := &customer.RegisterRequest{
		Title:     r.helper.MapCustomerTitle(customerDetails.Title),
		FirstName: customerDetails.FirstName,
		LastName:  customerDetails.LastName,
		Dob:       customerDetails.Dob,
		Address: &customer.AddressInput{
			CountryId: customerAddress.CountryID,
			State:     customerState,
			City:      customerCity,
			Street:    customerAddress.Street,
			Postcode:  customerAddress.Postcode,
			Cordinates: &customer.CordinatesInput{
				Longitude: float32(customerCoordinates.Longitude),
				Latitude:  float32(customerCoordinates.Latitude),
			},
		},
	}

	_, err = r.CustomerService.Register(ctx, customerReq)
	if err != nil {
		responseMessage = "Failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	responseMessage = "Successful"
	return &types.Response{Message: &responseMessage, Success: true, Code: int64(200)}, nil
}

// SubmitCdd is the resolver for the submitCDD field.
func (r *mutationResolver) SubmitCdd(ctx context.Context, cdd types.CDDInput) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	// Build request
	var request onboarding.SubmitCDDRequest
	// KYC
	if cdd.Kyc != nil {
		var kycReportTypes = make([]onboarding.KYCInput_ReportTypes, len(cdd.Kyc.ReportTypes))
		for index, value := range cdd.Kyc.ReportTypes {
			switch value {
			case types.KYCTypesDocument:
				kycReportTypes[index] = onboarding.KYCInput_DOCUMENT
			case types.KYCTypesFacialVideo:
				kycReportTypes[index] = onboarding.KYCInput_FACIAL_VIDEO
			}
		}
		request.Kyc = &onboarding.KYCInput{
			ReportTypes: kycReportTypes,
		}
	}
	// AML
	request.Aml = cdd.Aml
	// POA
	if cdd.Poa != nil {
		request.Poa = &onboarding.POAInput{
			Data: cdd.Poa.Data,
		}
	}
	// Execute RPC call
	response, err := r.OnBoardingService.SubmitCDD(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// AnswerQuestionary is the resolver for the answerQuestionary field.
func (r *mutationResolver) AnswerQuestionary(ctx context.Context, questionary types.QuestionaryAnswerInput) (*types.Response, error) {
	var responseMessage string

	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage = "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	customerAnswers := make([]*customer.AnswerInput, len(questionary.Answers))
	if len(questionary.Answers) < 1 {
		// Should never happen
		responseMessage = "Questionary answers not found"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(400)}, nil
	}
	for index, ans := range questionary.Answers {
		answer := &customer.AnswerInput{
			Id:                ans.ID,
			PredefinedAnswers: ans.PredefinedAnswers,
		}
		if ans.Answer != nil {
			answer.Answer = *ans.Answer
		}
		customerAnswers[index] = answer
	}

	req := &customer.AnswerQuestionaryRequest{
		Id:      questionary.ID,
		Answers: customerAnswers,
	}

	resp, err := r.CustomerService.AnswerQuestionary(ctx, req)
	if err != nil {
		return nil, err
	}

	responseMessage = "Success"

	return &types.Response{Message: &responseMessage, Success: resp.Success, Code: int64(resp.Code)}, nil
}

// AcceptContent is the resolver for the acceptContent field.
func (r *mutationResolver) AcceptContent(ctx context.Context, contentID string) (*types.Response, error) {
	res, err := r.CustomerService.SetAcceptance(ctx, &customer.SetAcceptanceRequest{ContentId: contentID})
	if err != nil {
		return nil, err
	}

	message := "Successful"
	return &types.Response{
		Message: &message,
		Success: res.Success,
		Code:    int64(res.Code),
	}, nil
}

// SetTransactionPassword is the resolver for the setTransactionPassword field.
func (r *mutationResolver) SetTransactionPassword(ctx context.Context, password string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Build request
	request := customer.SetTransactionPasswordRequest{
		Password: password,
	}
	// Execute RPC call
	response, err := r.CustomerService.SetTransactionPassword(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// ForgotTransactionPassword is the resolver for the forgotTransactionPassword field.
func (r *mutationResolver) ForgotTransactionPassword(ctx context.Context, newTransactionPassword string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Build request
	request := customer.ForgotTransactionPasswordRequest{
		NewPassword: newTransactionPassword,
	}
	// Execute RPC call
	response, err := r.CustomerService.ForgotTransactionPassword(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// ResetTransactionPassword is the resolver for the resetTransactionPassword field.
func (r *mutationResolver) ResetTransactionPassword(ctx context.Context, otpToken string, email string, newTransactionPassword string, currentTransactionPassword string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}

	// Build request
	request := customer.ResetTransactionPasswordRequest{
		OtpToken:        otpToken,
		Email:           email,
		NewPassword:     newTransactionPassword,
		CurrentPassword: currentTransactionPassword,
	}
	// Execute RPC call
	response, err := r.CustomerService.ResetTransactionPassword(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, credentials types.AuthInput) (*types.AuthResponse, error) {
	email, err := r.emailValidator.Validate(credentials.Email)
	if err != nil {
		invalidMsg := emailvalidator.ErrInvalidEmail.Message()
		return &types.AuthResponse{
			Message: &invalidMsg,
			Success: false,
			Code:    http.StatusBadRequest,
		}, err
	}
	req := auth.LoginRequest{
		CustomerInput: &auth.CustomerInput{
			Email:         email,
			LoginPassword: credentials.Password,
		},
		Device: &pbTypes.DeviceInput{
			Identifier: credentials.DeviceIdentifier,
			Tokens:     []*pbTypes.DeviceTokenInput{},
		},
	}
	tokens, err := r.AuthService.Login(context.Background(), &req)
	if err != nil {
		invalidMsg := errorvalues.Message(errorvalues.InternalErr)
		return &types.AuthResponse{
			Message: &invalidMsg,
			Success: false,
			Code:    http.StatusInternalServerError,
		}, err
	}
	msg := "Success"
	return &types.AuthResponse{
		Message: &msg,
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    tokens.AuthToken,
			Refresh: &tokens.RefreshToken,
		},
		Code: int64(http.StatusOK),
	}, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, token string) (*types.AuthResponse, error) {
	result, err := r.AuthService.RefreshToken(ctx, &auth.RefreshTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}

	message := "Success"
	return &types.AuthResponse{
		Message: &message,
		Success: true,
		Code:    int64(http.StatusOK),
		Tokens: &types.AuthTokens{
			Auth:    result.AuthToken,
			Refresh: &result.RefreshToken,
		},
	}, nil
}

// SetDeviceToken is the resolver for the setDeviceToken field.
func (r *mutationResolver) SetDeviceToken(ctx context.Context, tokens []*types.DeviceTokenInput) (*types.Response, error) {
	var responseMessage string
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage = "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}
	token_ := make([]*pbTypes.DeviceTokenInput, 0)
	if len(tokens) < 1 {
		responseMessage = "device token empty"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(400)}, errors.New("device token empty")
	}

	for _, token := range tokens {
		token_ = append(token_, &pbTypes.DeviceTokenInput{Value: token.Value, Type: pbTypes.DeviceToken_FIREBASE})
	}

	resp, err := r.CustomerService.SetDeviceToken(ctx, &customer.SetDeviceTokenRequest{Tokens: token_})
	if err != nil {
		responseMessage = "set device token failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	responseMessage = "Successful"
	return &types.Response{Message: &responseMessage, Success: resp.Success, Code: int64(resp.Code)}, err
}

// SetDevicePreferences is the resolver for the setDevicePreferences field.
func (r *mutationResolver) SetDevicePreferences(ctx context.Context, preferences []*types.DevicePreferencesInput) (*types.Response, error) {
	var responseMessage string
	helpers := &helpersfactory{}
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage = "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	preferences_ := make([]*pbTypes.DevicePreferencesInput, 0)
	if len(preferences) < 1 {
		responseMessage = "device preferences empty"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(400)}, errors.New("device preferences empty")
	}

	for _, preference := range preferences {
		preferences_ = append(preferences_, &pbTypes.DevicePreferencesInput{Type: pbTypes.DevicePreferences_DevicePreferencesTypes(helpers.GetDeveicePreferenceTypesIndex(preference.Type)), Value: preference.Value})
	}

	resp, err := r.CustomerService.SetDevicePreferences(ctx, &customer.SetDevicePreferencesRequest{Preferences: preferences_})
	if err != nil {
		responseMessage = "set device preferences failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	if resp != nil {
		responseMessage = "Successful"
		return &types.Response{Message: &responseMessage, Success: true, Code: int64(200)}, err
	}

	responseMessage = "Unknown error occurred"
	return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
}

// CheckBvn is the resolver for the checkBVN field.
func (r *mutationResolver) CheckBvn(ctx context.Context, bvn string, phone string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Build request
	request := customer.CheckBVNRequest{
		Bvn:   bvn,
		Phone: phone,
	}
	// Execute RPC call
	response, err := r.CustomerService.CheckBVN(ctx, &request)
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, account types.AccountInput) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Build request
	request := accountPb.CreateAccountRequest{
		ProductId: account.ProductID,
	}
	// Call RPC
	_, err = r.AccountService.CreateAccount(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: true,
		Code:    http.StatusOK,
	}, nil
}

// CreateVaultAccount is the resolver for the createVaultAccount field.
func (r *mutationResolver) CreateVaultAccount(ctx context.Context, account types.VaultAccountInput, transactionPassword string) (*types.Response, error) {
	// Authenticate user
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}

	accountName := ""
	if account.Name != nil {
		accountName = *account.Name
	}

	req := accountPb.CreateVaultAccountRequest{
		VaultAccountInput: &accountPb.VaultAccountInput{
			ProductId:     account.ProductID,
			SourceAccount: account.SourceAccount,
			Amount:        float32(account.Amount),
			Name:          accountName,
		},
		TransactionPassword: transactionPassword,
	}

	_, err = r.AccountService.CreateVaultAccount(ctx, &req)
	if err != nil {
		msg := err.Error()
		return &types.Response{
			Success: false,
			Code:    int64(http.StatusInternalServerError),
			Message: &msg,
		}, err
	}

	return &types.Response{
		Success: true,
		Code:    int64(http.StatusOK),
	}, nil
}

// CreateBeneficiary is the resolver for the createBeneficiary field.
func (r *mutationResolver) CreateBeneficiary(ctx context.Context, beneficiary types.BeneficiaryInput, transactionPassword string) (*types.Response, error) {
	// Authenticate user
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}
	beneficiaryAccount := beneficiary.Account
	if beneficiaryAccount == nil {
		beneficiaryAccount = &types.BeneficiaryAccountInput{}
	}
	beneficaryAccountName := ""
	if beneficiaryAccount.Name != nil {
		beneficaryAccountName = *beneficiaryAccount.Name
	}
	req := payment.CreateBeneficiaryRequest{
		TransactionPassword: transactionPassword,
		Beneficiary: &payment.BeneficiaryInput{
			Name: beneficiary.Name,
			Account: &payment.BeneficiaryAccountInput{
				Name:          beneficaryAccountName,
				CurrencyId:    beneficiaryAccount.CurrencyID,
				AccountNumber: beneficiaryAccount.AccountNumber,
				Code:          beneficiaryAccount.Code,
			},
		},
	}
	_, err = r.PaymentService.CreateBeneficiary(ctx, &req)
	if err != nil {
		msg := err.Error()
		return &types.Response{
			Success: false,
			Code:    int64(http.StatusInternalServerError),
			Message: &msg,
		}, err
	}

	return &types.Response{
		Success: true,
		Code:    int64(http.StatusOK),
	}, nil
}

// CreateBeneficiariesByPhone is the resolver for the createBeneficiariesByPhone field.
func (r *mutationResolver) CreateBeneficiariesByPhone(ctx context.Context, beneficiaries []*types.BeneficiaryByPhoneInput, transactionPassword string) (*types.Response, error) {
	// Authenticate user
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}

	reqBeneficiaries := make([]*payment.BeneficiaryByPhoneInput, len(beneficiaries))
	for _, b := range beneficiaries {
		reqBeneficiaries = append(reqBeneficiaries, &payment.BeneficiaryByPhoneInput{
			Phone: b.Phone,
			Name:  b.Name,
		})
	}

	req := payment.CreateBeneficiariesByPhoneRequest{
		TransactionPassword: transactionPassword,
		Beneficiaries:       reqBeneficiaries,
	}

	_, err = r.PaymentService.CreateBeneficiariesByPhone(ctx, &req)
	if err != nil {
		msg := err.Error()
		return &types.Response{
			Success: false,
			Code:    int64(http.StatusInternalServerError),
			Message: &msg,
		}, err
	}

	return &types.Response{
		Success: true,
		Code:    int64(http.StatusOK),
	}, nil
}

// AddBeneficiaryAccount is the resolver for the addBeneficiaryAccount field.
func (r *mutationResolver) AddBeneficiaryAccount(ctx context.Context, beneficiaryID string, account types.BeneficiaryAccountInput, transactionPassword string) (*types.Response, error) {
	// Authenticate user
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}
	beneficaryAccountName := ""
	if account.Name != nil {
		beneficaryAccountName = *account.Name
	}
	req := payment.AddBeneficiaryAccountRequest{
		BeneficiaryId:       beneficiaryID,
		TransactionPassword: transactionPassword,
		Account: &payment.BeneficiaryAccountInput{
			Name:          beneficaryAccountName,
			CurrencyId:    account.CurrencyID,
			AccountNumber: account.AccountNumber,
			Code:          account.Code,
		},
	}
	_, err = r.PaymentService.AddBeneficiaryAccount(ctx, &req)
	if err != nil {
		msg := err.Error()
		return &types.Response{
			Success: false,
			Code:    int64(http.StatusInternalServerError),
			Message: &msg,
		}, err
	}

	return &types.Response{
		Success: true,
		Code:    int64(http.StatusOK),
	}, nil
}

// DeleteBeneficiaryAccount is the resolver for the deleteBeneficiaryAccount field.
func (r *mutationResolver) DeleteBeneficiaryAccount(ctx context.Context, beneficiaryID string, accountID string, transactionPassword string) (*types.Response, error) {
	// Authenticate user
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}

	req := payment.DeleteBeneficiaryAccountRequest{
		BeneficiaryId:       beneficiaryID,
		TransactionPassword: transactionPassword,
		AccountId:           accountID,
	}
	_, err = r.PaymentService.DeleteBeneficiaryAccount(ctx, &req)
	if err != nil {
		msg := err.Error()
		return &types.Response{
			Success: false,
			Code:    int64(http.StatusInternalServerError),
			Message: &msg,
		}, err
	}

	return &types.Response{
		Success: true,
		Code:    int64(http.StatusOK),
	}, nil
}

// CreateTransfer is the resolver for the createTransfer field.
func (r *mutationResolver) CreateTransfer(ctx context.Context, transfer types.TransactionInput, transactionPassword string) (*types.Response, error) {
	// Auhtenticate user
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}
	request := payment.CreateTransferRequest{
		Transfer: &payment.TransactionInput{
			TransactionTypeId: transfer.TransactionTypeID,
			FeeIds:            transfer.FeeIds,
			Amount:            float32(transfer.Amount),
			SourceAccountId:   transfer.SourceAccountID,
			TargetAccountId:   transfer.TargetAccountID,
			IdempotencyKey:    transfer.IdempotencyKey,
		},
		TransactionPassword: transactionPassword,
	}
	if transfer.Reference != nil {
		request.Transfer.Reference = *transfer.Reference
	}
	if transfer.ExchangeRateID != nil {
		request.Transfer.ExchangeRateId = *transfer.ExchangeRateID
	}
	resp, err := r.PaymentService.CreateTransfer(ctx, &request)
	if err != nil {
		msg := err.Error()
		return &types.Response{
			Success: false,
			Code:    int64(http.StatusInternalServerError),
			Message: &msg,
		}, err
	}

	return &types.Response{
		Success: resp.Success,
		Code:    int64(resp.Code),
	}, nil
}

// SendNotification is the resolver for the sendNotification field.
func (r *mutationResolver) SendNotification(ctx context.Context, typeArg types.DeliveryMode, content string, templateID string) (*types.Response, error) {
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	deliveryModeCode := messaging.SendNotificationRequest_DeliveryMode_value[typeArg.String()]

	request := messaging.SendNotificationRequest{
		Type:       messaging.SendNotificationRequest_DeliveryMode(deliveryModeCode),
		Content:    content,
		TemplateId: templateID,
	}

	response, err := r.MessagingService.SendNotification(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// DeactivateCredential is the resolver for the deactivateCredential field.
func (r *mutationResolver) DeactivateCredential(ctx context.Context, credentialType types.IdentityCredentialsTypes) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Build request
	request := customer.DeactivateCredentialRequest{
		CredentialType: r.helper.MapCredentialTypes(credentialType),
	}

	// Execute RPC call
	response, err := r.CustomerService.DeactivateCredential(ctx, &request)
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// WithdrawVaultAccount is the resolver for the withdrawVaultAccount field.
func (r *mutationResolver) WithdrawVaultAccount(ctx context.Context, sourceAccountID string, targetAccountID string, transactionPassword string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	// Build request
	request := payment.WithdrawVaultAccountRequest{
		SourceAccountId:     sourceAccountID,
		TargetAccountId:     targetAccountID,
		TransactionPassword: transactionPassword,
	}

	// Call RPC
	response, err := r.PaymentService.WithdrawVaultAccount(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// UpdateDevice is the resolver for the updateDevice field.
func (r *mutationResolver) UpdateDevice(ctx context.Context, phoneNumber string, otp string, device types.DeviceInput) (*types.Response, error) {
	helpers := &helpersfactory{}
	var tokens []*pbTypes.DeviceTokenInput
	for _, deviceToken := range device.Tokens {
		reqToken := &pbTypes.DeviceTokenInput{
			Type:  pbTypes.DeviceToken_FIREBASE,
			Value: deviceToken.Value,
		}
		tokens = append(tokens, reqToken)
	}

	var preferences []*pbTypes.DevicePreferencesInput
	for _, preference := range device.Preferences {
		reqPreference := &pbTypes.DevicePreferencesInput{
			Type:  pbTypes.DevicePreferences_DevicePreferencesTypes(helpers.GetDeveicePreferenceTypesIndex(preference.Type)),
			Value: preference.Value,
		}
		preferences = append(preferences, reqPreference)
	}

	request := customer.DeviceInputRequest{
		PhoneNumber: phoneNumber,
		Otp:         otp,
		Device: &pbTypes.DeviceInput{
			Identifier:  device.Identifier,
			Os:          device.Os,
			Brand:       device.Brand,
			Tokens:      tokens,
			Preferences: preferences,
		},
	}

	//	Execute RPC call
	response, err := r.CustomerService.UpdateDevice(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// CheckCustomerDetails is the resolver for the checkCustomerDetails field.
func (r *mutationResolver) CheckCustomerDetails(ctx context.Context, customerDetails types.CheckCustomerDetailsInput, typeArg types.ActionType) (*types.Response, error) {
	var request customer.CheckCustomerDetailsRequest
	switch typeArg {
	case types.ActionTypeDeviceUpdate:
		request = customer.CheckCustomerDetailsRequest{
			Password:    customerDetails.Password,
			PhoneNumber: customerDetails.PhoneNumber,
			Dob:         customerDetails.Dob,
			Email:       customerDetails.Email,
			ActionType:  customer.CheckCustomerDetailsRequest_DEVICE_UPDATE,
		}
	default:
		return &types.Response{Success: false, Code: int64(400)}, errors.New("not a valid ActionType")
	}
	// Execute RPC call
	response, err := r.CustomerService.CheckCustomerDetails(ctx, &request)
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// CloseAccount is the resolver for the closeAccount field.
func (r *mutationResolver) CloseAccount(ctx context.Context, accountCloseInput types.AccountCloseInput) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: http.StatusBadGateway}, err
	}

	//Build request
	request := accountPb.CloseAccountRequest{
		AccountId: accountCloseInput.AccountID,
		DepositAccount: &accountPb.BeneficiaryAccountInput{
			Name:          *accountCloseInput.DepositAccount.Name,
			CurrencyId:    accountCloseInput.DepositAccount.CurrencyID,
			AccountNumber: accountCloseInput.DepositAccount.AccountNumber,
			Code:          accountCloseInput.DepositAccount.Code,
		},
		TransactionPassword: accountCloseInput.TransactionPassword,
	}

	// Call RPC
	response, err := r.AccountService.CloseAccount(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    http.StatusOK,
	}, nil
}

// CreateScheduledTransfer is the resolver for the createScheduledTransfer field.
func (r *mutationResolver) CreateScheduledTransfer(ctx context.Context, scheduledTransfer types.ScheduledTransactionInput, transactionPassword string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	var reference string
	if scheduledTransfer.Reference != nil {
		reference = *scheduledTransfer.Reference
	}

	// Build request
	request := payment.CreateScheduledTransferRequest{
		Transfer: &payment.ScheduledTransactionInput{
			TransactionTypeId: scheduledTransfer.TransactionTypeID,
			Reference:         reference,
			SourceAccountId:   scheduledTransfer.SourceAccountID,
			TargetAccountId:   scheduledTransfer.TargetAccountID,
			ReferenceDate:     timestamppb.New(time.Unix(scheduledTransfer.ReferenceDate, 0)),
			RepeatType:        r.helper.MapProtoScheduledTransactionRepeatType(scheduledTransfer.RepeatType),
			Amount:            float32(scheduledTransfer.Amount),
		},
		TransactionPassword: transactionPassword,
	}

	// Call RPC
	response, err := r.PaymentService.CreateScheduledTransfer(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// SetCustomerPreferences is the resolver for the setCustomerPreferences field.
func (r *mutationResolver) SetCustomerPreferences(ctx context.Context, preferences []*types.CustomerPreferencesInput) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthenticationError, err)
	}

	// Build request
	request := customer.SetCustomerPreferencesRequest{
		Preferences: func() []*customer.CustomerPreferencesInput {
			protoPreferences := make([]*customer.CustomerPreferencesInput, len(preferences))
			for index, preference := range preferences {
				protoPreferences[index] = &customer.CustomerPreferencesInput{
					Type:  r.helper.MapProtoCustomerPreferenceType(preference.Type),
					Value: preference.Value,
				}
			}

			return protoPreferences
		}(),
	}
	// Call RPC
	response, err := r.CustomerService.SetCustomerPreferences(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// RequestResubmit is the resolver for the requestResubmit field.
func (r *mutationResolver) RequestResubmit(ctx context.Context, customerID string, reportIds []string, message *string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	// Build request
	request := onboarding.RequestResubmitRequest{
		CustomerId: customerID,
		ReportIds:  reportIds,
	}
	if message != nil {
		request.Message = *message
	}
	// Call RPC
	response, err := r.OnBoardingService.RequestResubmit(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// StaffLogin is the resolver for the staffLogin field.
func (r *mutationResolver) StaffLogin(ctx context.Context, token string, authType types.AuthType) (*types.AuthResponse, error) {
	loginType := r.helper.StaffLoginTypeFromModel(authType)
	tokens, err := r.AuthService.StaffLogin(ctx, &auth.StaffLoginRequest{Token: token, AuthType: loginType})
	if err != nil {
		invalidMsg := errorvalues.Message(errorvalues.InternalErr)
		return &types.AuthResponse{
			Message: &invalidMsg,
			Success: false,
			Code:    http.StatusInternalServerError,
		}, err
	}
	msg := "Success"
	return &types.AuthResponse{
		Message: &msg,
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    tokens.AuthToken,
			Refresh: &tokens.RefreshToken,
		},
		Code: int64(http.StatusOK),
	}, nil
}

// UpdateKYCStatus is the resolver for the updateKYCStatus field.
func (r *mutationResolver) UpdateKYCStatus(ctx context.Context, id string, status types.KYCStatuses, message string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	// Build request
	request := onboarding.UpdateKYCStatusRequest{
		Id:      id,
		Message: message,
	}
	switch status {
	case types.KYCStatusesPending:
		request.Status = pbTypes.KYC_PENDING
	case types.KYCStatusesManualReview:
		request.Status = pbTypes.KYC_MANUAL_REVIEW
	case types.KYCStatusesApproved:
		request.Status = pbTypes.KYC_APPROVED
	case types.KYCStatusesDeclined:
		request.Status = pbTypes.KYC_DECLINED

	}
	response, err := r.OnBoardingService.UpdateKYCStatus(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// UpdateAMLStatus is the resolver for the updateAMLStatus field.
func (r *mutationResolver) UpdateAMLStatus(ctx context.Context, id string, status types.AMLStatuses, message string) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	// Build request
	request := onboarding.UpdateAMLStatusRequest{
		Id:      id,
		Message: message,
	}
	switch status {
	case types.AMLStatusesPending:
		request.Status = pbTypes.AML_PENDING
	case types.AMLStatusesManualReview:
		request.Status = pbTypes.AML_MANUAL_REVIEW
	case types.AMLStatusesApproved:
		request.Status = pbTypes.AML_APPROVED
	case types.AMLStatusesDeclined:
		request.Status = pbTypes.AML_DECLINED

	}
	response, err := r.OnBoardingService.UpdateAMLStatus(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// UpdateFx is the resolver for the updateFX field.
func (r *mutationResolver) UpdateFx(ctx context.Context, exchangeRate types.UpdateFXInput) (*types.Response, error) {
	var salePrice float32
	if exchangeRate.SalePrice != nil {
		salePrice = float32(*exchangeRate.SalePrice)
	}
	request := pricing.UpdateFXRequest{
		BaseCurrencyId: exchangeRate.BaseCurrencyID,
		CurrencyId:     exchangeRate.CurrencyID,
		BuyPrice:       float32(exchangeRate.BuyPrice),
		SalePrice:      salePrice,
	}

	// Execute RPC call
	response, err := r.PricingService.UpdateFX(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// UpdateFees is the resolver for the updateFees field.
func (r *mutationResolver) UpdateFees(ctx context.Context, fees []*types.UpdateFeesInput) (*types.Response, error) {
	var feesRequest []*pricing.UpdateFeesRequest

	for _, fee := range fees {
		var feeRequest pricing.UpdateFeesRequest
		feeRequest.TransactionTypeId = fee.TransactionTypeID
		var boundaryRequest pbTypes.FeeBoundaries
		feeRequest.Type = r.helper.MapProtoFeeTypes(fee.Type)
		for _, reqBoundary := range fee.Boundaries {
			boundaryRequest.Lower = float32(*reqBoundary.Lower)
			boundaryRequest.Upper = float32(*reqBoundary.Upper)
			boundaryRequest.Amount = float32(*reqBoundary.Amount)
			boundaryRequest.Percentage = float32(*reqBoundary.Percentage)
			feeRequest.Boundaries = append(feeRequest.Boundaries, &boundaryRequest)
		}
		feesRequest = append(feesRequest, &feeRequest)
	}

	request := pricing.UpdateFeesRequests{
		Fees: feesRequest,
	}
	// Execute RPC call
	response, err := r.PricingService.UpdateFees(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// StaffUpdateCustomerDetails is the resolver for the staffUpdateCustomerDetails field.
func (r *mutationResolver) StaffUpdateCustomerDetails(ctx context.Context, customerDetails types.StaffCustomerDetailsUpdateInput) (*types.Response, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		responseMessage := "User authentication failed"
		return &types.Response{Message: &responseMessage, Success: false, Code: int64(500)}, err
	}

	var (
		firstName, lastName, email string
		customerAddress            *customer.AddressInput
	)

	if customerDetails.FirstName != nil {
		firstName = *customerDetails.FirstName
	}

	if customerDetails.LastName != nil {
		lastName = *customerDetails.LastName
	}

	if customerDetails.Email != nil {
		email = *customerDetails.Email
	}

	if customerDetails.Address != nil {
		var (
			state, city string
			coordinates *customer.CordinatesInput
		)

		if customerDetails.Address.State != nil {
			state = *customerDetails.Address.State
		}

		if customerDetails.Address.City != nil {
			city = *customerDetails.Address.City
		}

		if customerDetails.Address.Cordinates != nil {
			coordinates = &customer.CordinatesInput{
				Latitude:  float32(customerDetails.Address.Cordinates.Latitude),
				Longitude: float32(customerDetails.Address.Cordinates.Longitude),
			}
		}

		customerAddress = &customer.AddressInput{
			CountryId:  customerDetails.Address.CountryID,
			State:      state,
			City:       city,
			Street:     customerDetails.Address.Street,
			Postcode:   customerDetails.Address.Postcode,
			Cordinates: coordinates,
		}
	}

	// Build request
	request := customer.StaffCustomerDetailsUpdateRequest{
		CustomerID: customerDetails.CustomerID,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Address:    customerAddress,
	}

	// Execute RPC call
	response, err := r.CustomerService.StaffCustomerDetailsUpdate(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Success: response.Success,
		Code:    int64(response.Code),
	}, nil
}

// WithdrawVaultAccountNoSource is the resolver for the withdrawVaultAccountNoSource field.
func (r *mutationResolver) WithdrawVaultAccountNoSource(ctx context.Context, vaultAccountID string, beneficiary types.BeneficiaryAccountInput, transactionPin string) (*types.Response, error) {
	return &types.Response{Code: http.StatusNotFound, Success: false}, nil
}

// CreateFaq is the resolver for the createFAQ field.
func (r *mutationResolver) CreateFaq(ctx context.Context, faq types.CreateFAQInput) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteFaq is the resolver for the deleteFAQ field.
func (r *mutationResolver) DeleteFaq(ctx context.Context, fAQid string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateFaq is the resolver for the updateFAQ field.
func (r *mutationResolver) UpdateFaq(ctx context.Context, faq types.UpdateFAQInput) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
