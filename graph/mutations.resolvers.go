package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"ms.api/graph/generated"
	errorvalues "ms.api/libs/errors"
	"ms.api/libs/validator/datevalidator"
	devicevalidator "ms.api/libs/validator/device"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/onboarding"
	pbTypes "ms.api/protos/pb/types"
	"ms.api/protos/pb/verification"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

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

func (r *mutationResolver) CheckCustomerEmail(ctx context.Context, email string, device types.DeviceInput) (*types.Response, error) {
	_, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return nil, err
	}

	resp, err := r.CustomerService.CheckEmail(ctx, &customer.CheckEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}

	// TODO: Temporal implementation, refactor to use correct RPC call when implemented
	return &types.Response{
		Success: resp.Success,
		Code:    0,
	}, nil
}

func (r *mutationResolver) CheckCustomerData(ctx context.Context, customerData types.CheckCustomerDataInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

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

func (r *mutationResolver) SetTransactionPassword(ctx context.Context, password string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) ResetTransactionPassword(ctx context.Context, otpToken string, email string, newTransactionPassword string, currentTransactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

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

func (r *mutationResolver) CheckBvn(ctx context.Context, bvn string, phone string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, account types.AccountInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateVaultAccount(ctx context.Context, account types.VaultAccountInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateBeneficiary(ctx context.Context, beneficiary types.BeneficiaryInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) AddBeneficiaryAccount(ctx context.Context, beneficiaryID string, account types.BeneficiaryAccountInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) DeleteBeneficaryAccount(ctx context.Context, beneficiaryID string, accountID string, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateTransfer(ctx context.Context, transfer types.TransactionInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
