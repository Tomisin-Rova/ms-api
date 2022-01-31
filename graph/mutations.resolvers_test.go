package graph

import (
	"context"
	"errors"
	"net/http"
	"testing"

	terror "github.com/roava/zebra/errors"

	"github.com/golang/mock/gomock"
	"github.com/roava/zebra/middleware"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	errorvalues "ms.api/libs/errors"
	devicevalidator "ms.api/libs/validator/device"
	emailvalidator "ms.api/libs/validator/email"
	phonenumbervalidator "ms.api/libs/validator/phonenumbervalidator"
	"ms.api/mocks"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/onboarding"
	pbTypes "ms.api/protos/pb/types"
	"ms.api/protos/pb/verification"
	"ms.api/types"
)

var (
	state       = "lagos"
	city        = "lagos"
	answer      = "My lifestyle"
	mockAddress = types.AddressInput{
		CountryID: "111xcc",
		State:     &state,
		City:      &city,
		Street:    "vi",
		Postcode:  "23401",
		Cordinates: &types.CordinatesInput{
			Latitude:  3.15669,
			Longitude: 3.99244,
		},
	}
	mockRegisterReq = types.CustomerDetailsInput{
		FirstName: "roava",
		LastName:  "app",
		Dob:       "18-05-1994",
		Address:   &mockAddress,
	}
)

func TestMutationResolver_RequestOtp(t *testing.T) {
	const (
		success = iota
		successWithExpireTime
		successSMS
		successPUSH
		errorCallingRPC
	)

	expireTime := int64(120)

	type arg struct {
		typeArg             types.DeliveryMode
		target              string
		expireTimeInSeconds *int64
	}
	var tests = []struct {
		name string
		arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				typeArg:             types.DeliveryModeEmail,
				target:              "email@roava.app",
				expireTimeInSeconds: nil,
			},
			testType: success,
		},
		{
			name: "Test success with expire time",
			arg: arg{
				typeArg:             types.DeliveryModeEmail,
				target:              "email@roava.app",
				expireTimeInSeconds: &expireTime,
			},
			testType: successWithExpireTime,
		},
		{
			name: "Test success SMS",
			arg: arg{
				typeArg:             types.DeliveryModeSms,
				target:              "1234567891",
				expireTimeInSeconds: nil,
			},
			testType: successSMS,
		},
		{
			name: "Test success",
			arg: arg{
				typeArg:             types.DeliveryModePush,
				target:              "1234567891",
				expireTimeInSeconds: nil,
			},
			testType: successPUSH,
		},
		{
			name: "Test error calling rpc function",
			arg: arg{
				typeArg:             types.DeliveryModeEmail,
				target:              "email@roava.app",
				expireTimeInSeconds: nil,
			},
			testType: errorCallingRPC,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			verificationServiceClient := mocks.NewMockVerificationServiceClient(controller)
			resolverOpts := &ResolverOpts{
				VerificationService: verificationServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				verificationServiceClient.EXPECT().RequestOTP(context.Background(), &verification.RequestOTPRequest{
					Type:                verification.RequestOTPRequest_EMAIL,
					Target:              testCase.arg.target,
					ExpireTimeInSeconds: 60,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				response, err := resolver.RequestOtp(context.Background(), testCase.arg.typeArg, testCase.arg.target, testCase.arg.expireTimeInSeconds)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, &types.Response{
					Success: response.Success,
					Code:    int64(response.Code),
				}, response)
			case successWithExpireTime:
				verificationServiceClient.EXPECT().RequestOTP(context.Background(), &verification.RequestOTPRequest{
					Type:                verification.RequestOTPRequest_EMAIL,
					Target:              testCase.arg.target,
					ExpireTimeInSeconds: int32(*testCase.arg.expireTimeInSeconds),
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				response, err := resolver.RequestOtp(context.Background(), testCase.arg.typeArg, testCase.arg.target, testCase.arg.expireTimeInSeconds)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, &types.Response{
					Success: response.Success,
					Code:    int64(response.Code),
				}, response)
			case successSMS:
				verificationServiceClient.EXPECT().RequestOTP(context.Background(), &verification.RequestOTPRequest{
					Type:                verification.RequestOTPRequest_SMS,
					Target:              testCase.arg.target,
					ExpireTimeInSeconds: 60,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				response, err := resolver.RequestOtp(context.Background(), testCase.arg.typeArg, testCase.arg.target, testCase.arg.expireTimeInSeconds)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, &types.Response{
					Success: response.Success,
					Code:    int64(response.Code),
				}, response)
			case successPUSH:
				verificationServiceClient.EXPECT().RequestOTP(context.Background(), &verification.RequestOTPRequest{
					Type:                verification.RequestOTPRequest_PUSH,
					Target:              testCase.arg.target,
					ExpireTimeInSeconds: 60,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				response, err := resolver.RequestOtp(context.Background(), testCase.arg.typeArg, testCase.arg.target, testCase.arg.expireTimeInSeconds)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, &types.Response{
					Success: response.Success,
					Code:    int64(response.Code),
				}, response)
			case errorCallingRPC:
				verificationServiceClient.EXPECT().RequestOTP(context.Background(), &verification.RequestOTPRequest{
					Type:                verification.RequestOTPRequest_EMAIL,
					Target:              testCase.arg.target,
					ExpireTimeInSeconds: 60,
				}).Return(nil, errors.New(""))

				response, err := resolver.RequestOtp(context.Background(), testCase.arg.typeArg, testCase.arg.target, testCase.arg.expireTimeInSeconds)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_VerifyOtp(t *testing.T) {
	const (
		success = iota
		errorCallingRPC
	)

	type arg struct {
		target   string
		otpToken string
	}
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				target:   "email@roava.app",
				otpToken: "12345",
			},
			testType: success,
		},
		{
			name: "Test error calling rpc function",
			arg: arg{
				target:   "email@roava.app",
				otpToken: "12345",
			},
			testType: errorCallingRPC,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			verificationServiceClient := mocks.NewMockVerificationServiceClient(controller)
			resolverOpts := &ResolverOpts{
				VerificationService: verificationServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
			switch testCase.testType {
			case success:
				verificationServiceClient.EXPECT().VerifyOTP(context.Background(), &verification.VerifyOTPRequest{
					Target:   testCase.arg.target,
					OtpToken: testCase.arg.otpToken,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				response, err := resolver.VerifyOtp(context.Background(), testCase.arg.target, testCase.arg.otpToken)
				assert.NoError(t, err)
				assert.NotNil(t, resolver)
				assert.Equal(t, &types.Response{
					Success: response.Success,
					Code:    int64(response.Code),
				}, response)
			case errorCallingRPC:
				verificationServiceClient.EXPECT().VerifyOTP(context.Background(), &verification.VerifyOTPRequest{
					Target:   testCase.arg.target,
					OtpToken: testCase.arg.otpToken,
				}).Return(nil, errors.New(""))

				response, err := resolver.VerifyOtp(context.Background(), testCase.arg.target, testCase.arg.otpToken)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_Signup(t *testing.T) {
	const (
		failInvalidPhone = iota
		failInvalidEmail
		failInvalidDevice
		failAuthServiceError
		successWithNilDeviceTokensAndPreferences
		success
	)

	defaultPayload := types.CustomerInput{
		Phone:         "+1122233334444",
		Email:         "EMAIL@user.org",
		LoginPassword: "123456",
		Device: &types.DeviceInput{
			Identifier: "device-identifier",
			Os:         "device-os",
			Brand:      "device-brand",
			Tokens: []*types.DeviceTokenInput{
				{
					Type:  types.DeviceTokenTypesFirebase,
					Value: "firebase-token",
				},
			},
			Preferences: []*types.DevicePreferencesInput{
				{
					Type:  types.DevicePreferencesTypesPush,
					Value: true,
				},
				{
					Type:  types.DevicePreferencesTypesBiometrics,
					Value: true,
				},
			},
		},
	}

	authServiceResponse := &auth.TokenPairResponse{
		AuthToken:    "auth-token",
		RefreshToken: "refresh-token",
	}

	testCases := []struct {
		name     string
		input    types.CustomerInput
		testType int
	}{
		{
			name:     "Fail if invalid phone",
			input:    defaultPayload,
			testType: failInvalidPhone,
		},
		{
			name:     "Fail if invalid e-mail",
			input:    defaultPayload,
			testType: failInvalidEmail,
		},
		{
			name:     "Fail if invalid device input",
			input:    defaultPayload,
			testType: failInvalidDevice,
		},
		{
			name:     "Fail on auth service error",
			input:    defaultPayload,
			testType: failAuthServiceError,
		},
		{
			name: "Success request despite empty device tokens and preferences",
			input: types.CustomerInput{
				Phone:         "+1122233334444",
				Email:         "EMAIL@user.org",
				LoginPassword: "123456",
				Device: &types.DeviceInput{
					Identifier: "device-identifier",
					Os:         "device-os",
					Brand:      "device-brand",
				},
			},
			testType: successWithNilDeviceTokensAndPreferences,
		},
		{
			name:     "Success request",
			input:    defaultPayload,
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	authServiceClient := mocks.NewMockAuthServiceClient(controller)
	phoneNumberValidator := mocks.NewMockPhoneNumberValidator(controller)
	emailValidator := mocks.NewMockEmailValidator(controller)
	deviceValidator := mocks.NewMockDeviceValidator(controller)
	resolverOpts := &ResolverOpts{
		AuthService:     authServiceClient,
		PhoneValidator:  phoneNumberValidator,
		EmailValidator:  emailValidator,
		DeviceValidator: deviceValidator,
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failInvalidPhone:
				phoneNumberValidator.EXPECT().ValidatePhoneNumber(testCase.input.Phone).Return(errors.New("")).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Signup(context.Background(), testCase.input)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, phonenumbervalidator.ErrInvalidPhoneNumber.Message(), *resp.Message)
				assert.Equal(t, int64(http.StatusBadRequest), resp.Code)
			case failInvalidEmail:
				phoneNumberValidator.EXPECT().ValidatePhoneNumber(testCase.input.Phone).Return(nil).Times(1)
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("", errors.New("")).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Signup(context.Background(), testCase.input)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, emailvalidator.ErrInvalidEmail.Message(), *resp.Message)
				assert.Equal(t, int64(http.StatusBadRequest), resp.Code)
			case failInvalidDevice:
				phoneNumberValidator.EXPECT().ValidatePhoneNumber(defaultPayload.Phone).Return(nil).Times(1)
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("email@user.org", nil).Times(1)
				deviceValidator.EXPECT().Validate(testCase.input.Device).Return(errors.New("")).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Signup(context.Background(), testCase.input)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, devicevalidator.ErrInvalidDevice.Message(), *resp.Message)
				assert.Equal(t, int64(http.StatusBadRequest), resp.Code)
			case failAuthServiceError:
				phoneNumberValidator.EXPECT().ValidatePhoneNumber(testCase.input.Phone).Return(nil).Times(1)
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("email@user.org", nil).Times(1)
				deviceValidator.EXPECT().Validate(testCase.input.Device).Return(nil).Times(1)
				authServiceClient.EXPECT().Signup(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Signup(context.Background(), testCase.input)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, errorvalues.Message(errorvalues.InternalErr), *resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case successWithNilDeviceTokensAndPreferences:
				phoneNumberValidator.EXPECT().ValidatePhoneNumber(testCase.input.Phone).Return(nil).Times(1)
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("email@user.org", nil).Times(1)
				deviceValidator.EXPECT().Validate(testCase.input.Device).Return(nil).Times(1)
				authServiceClient.EXPECT().Signup(gomock.Any(), gomock.Any()).Return(authServiceResponse, nil).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Signup(context.Background(), testCase.input)
				assert.Nil(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, "Success", *resp.Message)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			case success:
				phoneNumberValidator.EXPECT().ValidatePhoneNumber(testCase.input.Phone).Return(nil).Times(1)
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("email@user.org", nil).Times(1)
				deviceValidator.EXPECT().Validate(testCase.input.Device).Return(nil).Times(1)
				authServiceClient.EXPECT().Signup(gomock.Any(), gomock.Any()).Return(authServiceResponse, nil).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Signup(context.Background(), testCase.input)
				assert.Nil(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, "Success", *resp.Message)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_ResetLoginPassword(t *testing.T) {
	const (
		success = iota
		errorUnauthenticatedUser
		errorCallingResetLoginPassword
	)

	type arg struct {
		otpToken      string
		email         string
		loginPassword string
	}
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				otpToken:      "validOTP",
				email:         "email@roava.app",
				loginPassword: "newLoginPassword",
			},
			testType: success,
		},
		{
			name:     "Test error unauthenticated user",
			testType: errorUnauthenticatedUser,
		},
		{
			name: "Test error calling RPC",
			arg: arg{
				otpToken:      "validOTP",
				email:         "email@roava.app",
				loginPassword: "newLoginPassword",
			},
			testType: errorCallingResetLoginPassword,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerService := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerService,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				customerService.EXPECT().ResetLoginPassword(ctx, &customer.ResetLoginPasswordRequest{
					OtpToken:      "validOTP",
					Email:         "email@roava.app",
					LoginPassword: "newLoginPassword",
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				response, err := resolver.ResetLoginPassword(ctx, testCase.arg.otpToken, testCase.arg.email, testCase.arg.loginPassword)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, response)
			case errorUnauthenticatedUser:
				response, err := resolver.ResetLoginPassword(context.Background(), testCase.arg.otpToken, testCase.arg.email, testCase.arg.loginPassword)
				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthentication, err.(*terror.Terror).Code())
				assert.Nil(t, response)
			case errorCallingResetLoginPassword:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				customerService.EXPECT().ResetLoginPassword(ctx, &customer.ResetLoginPasswordRequest{
					OtpToken:      "validOTP",
					Email:         "email@roava.app",
					LoginPassword: "newLoginPassword",
				}).Return(nil, errors.New(""))

				response, err := resolver.ResetLoginPassword(ctx, testCase.arg.otpToken, testCase.arg.email, testCase.arg.loginPassword)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_CheckCustomerEmail(t *testing.T) {
	const (
		success = iota
		emailNotFound
		invalidEmail
	)

	tests := []struct {
		name     string
		arg      string
		testType int
	}{
		{
			name:     "Test check email found successful",
			arg:      "f@mail.com",
			testType: success,
		},

		{
			name:     "Test error check email not found",
			arg:      "f@mail.com",
			testType: emailNotFound,
		},
		{
			name:     "Test invalid email error",
			arg:      "invalidEmail",
			testType: invalidEmail,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch test.testType {
			case success:
				customerServiceClient.EXPECT().CheckEmail(context.Background(),
					&customer.CheckEmailRequest{Email: test.arg},
				).Return(&pbTypes.DefaultResponse{Success: true}, nil)

				resp, err := resolver.CheckCustomerEmail(context.Background(), test.arg, types.DeviceInput{})
				assert.NoError(t, err)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    0,
				}, resp)

			case emailNotFound:
				customerServiceClient.EXPECT().CheckEmail(context.Background(),
					&customer.CheckEmailRequest{Email: test.arg},
				).Return(&pbTypes.DefaultResponse{Success: false}, errors.New("not found"))

				resp, err := resolver.CheckCustomerEmail(context.Background(), test.arg, types.DeviceInput{})
				assert.Error(t, err)
				assert.Nil(t, resp)

			case invalidEmail:
				resp, err := resolver.CheckCustomerEmail(context.Background(), test.arg, types.DeviceInput{})
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_CheckCustomerData(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CheckCustomerData(context.Background(), types.CheckCustomerDataInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_Register(t *testing.T) {
	const (
		success = iota
		invalidDob
		errUserAuthentication
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test register success",
			testType: success,
		},
		{
			name:     "Test invalid Dob",
			testType: invalidDob,
		},
		{
			name:     "Test error authenticating user",
			testType: errUserAuthentication,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{Client: models.APP, ID: "123456", Email: "f@roava.app", DeviceID: "129594533fsdd"})

				customerServiceClient.EXPECT().Register(ctx, &customer.RegisterRequest{
					FirstName: "roava",
					LastName:  "app",
					Dob:       "18-05-1994",
					Address: &customer.AddressInput{
						CountryId: "111xcc",
						State:     state,
						City:      city,
						Street:    "vi",
						Postcode:  "23401",
						Cordinates: &customer.CordinatesInput{
							Latitude:  3.15669,
							Longitude: 3.99244,
						},
					},
				}).Return(nil, nil).Times(1)
				resp, err := resolver.Register(ctx, mockRegisterReq)

				assert.Nil(t, err)
				assert.Equal(t, resp.Code, int64(200))

			case invalidDob:

				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{Client: models.APP, ID: "123456", Email: "f@roava.app", DeviceID: "129594533fsdd"})
				mockRegisterReq.Dob = "1994-10-02"

				_, err := resolver.Register(ctx, mockRegisterReq)

				assert.Error(t, err)

			case errUserAuthentication:
				_, err := resolver.Register(context.Background(), mockRegisterReq)

				assert.Error(t, err)
			}

		})
	}

}

func TestMutationResolver_SubmitCdd(t *testing.T) {
	const (
		success = iota
		successNoKYC
		successNoAML
		successNoPOA
		errorUnauthenticatedUser
		errorSubmitCDD
	)

	ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})

	type arg struct {
		ctx      context.Context
		cddInput types.CDDInput
	}
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success all validations",
			arg: arg{
				ctx: ctx,
				cddInput: types.CDDInput{
					Kyc: &types.KYCInput{
						ReportTypes: []types.KYCTypes{types.KYCTypesDocument, types.KYCTypesFacialVideo},
					},
					Aml: true,
					Poa: &types.POAInput{
						Data: "base64Image",
					},
				},
			},
			testType: success,
		},
		{
			name: "Test success no KYC validation",
			arg: arg{
				ctx: ctx,
				cddInput: types.CDDInput{
					Kyc: nil,
					Aml: true,
					Poa: &types.POAInput{
						Data: "base64Image",
					},
				},
			},
			testType: successNoKYC,
		},
		{
			name: "Test success no aml validation",
			arg: arg{
				ctx: ctx,
				cddInput: types.CDDInput{
					Kyc: &types.KYCInput{
						ReportTypes: []types.KYCTypes{types.KYCTypesDocument, types.KYCTypesFacialVideo},
					},
					Aml: false,
					Poa: &types.POAInput{
						Data: "base64Image",
					},
				},
			},
			testType: successNoAML,
		},
		{
			name: "Test success no POA validation",
			arg: arg{
				ctx: ctx,
				cddInput: types.CDDInput{
					Kyc: &types.KYCInput{
						ReportTypes: []types.KYCTypes{types.KYCTypesDocument, types.KYCTypesFacialVideo},
					},
					Aml: true,
					Poa: nil,
				},
			},
			testType: successNoPOA,
		},
		{
			name: "Test error unauthenticated user",
			arg: arg{
				ctx: context.Background(),
			},
			testType: errorUnauthenticatedUser,
		},
		{
			name: "Test error submitting CDD",
			arg: arg{
				ctx: ctx,
				cddInput: types.CDDInput{
					Kyc: &types.KYCInput{
						ReportTypes: []types.KYCTypes{types.KYCTypesDocument, types.KYCTypesFacialVideo},
					},
					Aml: true,
					Poa: &types.POAInput{
						Data: "base64Image",
					},
				},
			},
			testType: errorSubmitCDD,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			onboardingServiceClient := mocks.NewMockOnboardingServiceClient(controller)
			resolverOpts := &ResolverOpts{
				OnboardingService: onboardingServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				onboardingServiceClient.EXPECT().SubmitCDD(testCase.arg.ctx, &onboarding.SubmitCDDRequest{
					Kyc: &onboarding.KYCInput{
						ReportTypes: []onboarding.KYCInput_ReportTypes{onboarding.KYCInput_DOCUMENT, onboarding.KYCInput_FACIAL_VIDEO},
					},
					Aml: true,
					Poa: &onboarding.POAInput{
						Data: testCase.arg.cddInput.Poa.Data,
					},
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SubmitCdd(testCase.arg.ctx, testCase.arg.cddInput)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successNoKYC:
				onboardingServiceClient.EXPECT().SubmitCDD(testCase.arg.ctx, &onboarding.SubmitCDDRequest{
					Kyc: nil,
					Aml: true,
					Poa: &onboarding.POAInput{
						Data: testCase.arg.cddInput.Poa.Data,
					},
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SubmitCdd(testCase.arg.ctx, testCase.arg.cddInput)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successNoAML:
				onboardingServiceClient.EXPECT().SubmitCDD(testCase.arg.ctx, &onboarding.SubmitCDDRequest{
					Kyc: &onboarding.KYCInput{
						ReportTypes: []onboarding.KYCInput_ReportTypes{onboarding.KYCInput_DOCUMENT, onboarding.KYCInput_FACIAL_VIDEO},
					},
					Aml: false,
					Poa: &onboarding.POAInput{
						Data: testCase.arg.cddInput.Poa.Data,
					},
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SubmitCdd(testCase.arg.ctx, testCase.arg.cddInput)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successNoPOA:
				onboardingServiceClient.EXPECT().SubmitCDD(testCase.arg.ctx, &onboarding.SubmitCDDRequest{
					Kyc: &onboarding.KYCInput{
						ReportTypes: []onboarding.KYCInput_ReportTypes{onboarding.KYCInput_DOCUMENT, onboarding.KYCInput_FACIAL_VIDEO},
					},
					Aml: true,
					Poa: nil,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SubmitCdd(testCase.arg.ctx, testCase.arg.cddInput)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticatedUser:
				resp, err := resolver.SubmitCdd(testCase.arg.ctx, testCase.arg.cddInput)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusUnauthorized,
				}, resp)
			case errorSubmitCDD:
				onboardingServiceClient.EXPECT().SubmitCDD(testCase.arg.ctx, &onboarding.SubmitCDDRequest{
					Kyc: &onboarding.KYCInput{
						ReportTypes: []onboarding.KYCInput_ReportTypes{onboarding.KYCInput_DOCUMENT, onboarding.KYCInput_FACIAL_VIDEO},
					},
					Aml: true,
					Poa: &onboarding.POAInput{
						Data: testCase.arg.cddInput.Poa.Data,
					},
				}).Return(nil, errors.New(""))

				resp, err := resolver.SubmitCdd(testCase.arg.ctx, testCase.arg.cddInput)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_AnswerQuestionary(t *testing.T) {
	const (
		success = iota
		errUserAuthentication
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test answer questionary successful",
			testType: success,
		},
		{
			name:     "Test error while authenticating user",
			testType: errUserAuthentication,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(
					context.Background(),
					&models.JWTClaims{
						Client:   models.APP,
						ID:       "123456",
						Email:    "f@roava.app",
						DeviceID: "129594533fsdd"},
				)

				customerServiceClient.EXPECT().AnswerQuestionary(ctx,
					&customer.AnswerQuestionaryRequest{
						Id: "questionaire_id",
						Answers: []*customer.AnswerInput{
							{
								Id:                "question_id",
								Answer:            "My lifestyle",
								PredefinedAnswers: []string{"a", "b", "c"},
							},
						},
					},
				).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    200,
				}, nil)

				req := types.QuestionaryAnswerInput{
					ID: "questionaire_id",
					Answers: []*types.AnswerInput{
						{
							ID:                "question_id",
							Answer:            &answer,
							PredefinedAnswers: []string{"a", "b", "c"},
						},
					},
				}

				resp, err := resolver.AnswerQuestionary(ctx, req)

				assert.NoError(t, err)
				assert.Equal(t, resp.Code, int64(200))

			case errUserAuthentication:
				ctx := context.Background()

				req := types.QuestionaryAnswerInput{
					ID: "questionaire_id",
					Answers: []*types.AnswerInput{
						{
							ID:                "question_id",
							Answer:            &answer,
							PredefinedAnswers: []string{"a", "b", "c"},
						},
					},
				}

				resp, err := resolver.AnswerQuestionary(ctx, req)

				assert.Error(t, err)
				assert.Equal(t, resp.Code, int64(500))
			}
		})
	}

}

func TestMutationResolver_SetTransactionPassword(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.SetTransactionPassword(context.Background(), "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_ResetTransactionPassword(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.ResetTransactionPassword(context.Background(), "", "", "", "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_Login(t *testing.T) {
	const (
		failInvalidEmail = iota
		failAuthServiceError
		success
	)

	defaultPayload := types.AuthInput{
		Email:            "EMAIL@user.org",
		Password:         "123456",
		DeviceIdentifier: "device-identifier",
	}

	authServiceResponse := &auth.TokenPairResponse{
		AuthToken:    "auth-token",
		RefreshToken: "refresh-token",
	}

	testCases := []struct {
		name     string
		input    types.AuthInput
		testType int
	}{

		{
			name:     "Fail if invalid e-mail",
			input:    defaultPayload,
			testType: failInvalidEmail,
		},
		{
			name:     "Fail on auth service error",
			input:    defaultPayload,
			testType: failAuthServiceError,
		},

		{
			name:     "Success request",
			input:    defaultPayload,
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	authServiceClient := mocks.NewMockAuthServiceClient(controller)
	emailValidator := mocks.NewMockEmailValidator(controller)
	resolverOpts := &ResolverOpts{
		AuthService:    authServiceClient,
		EmailValidator: emailValidator,
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failInvalidEmail:
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("", errors.New("")).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Login(context.Background(), testCase.input)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, emailvalidator.ErrInvalidEmail.Message(), *resp.Message)
				assert.Equal(t, int64(http.StatusBadRequest), resp.Code)
			case failAuthServiceError:
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("email@user.org", nil).Times(1)
				authServiceClient.EXPECT().Login(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Login(context.Background(), testCase.input)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, errorvalues.Message(errorvalues.InternalErr), *resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				emailValidator.EXPECT().Validate(testCase.input.Email).Return("email@user.org", nil).Times(1)
				authServiceClient.EXPECT().Login(gomock.Any(), gomock.Any()).Return(authServiceResponse, nil).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.Login(context.Background(), testCase.input)
				assert.Nil(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, "Success", *resp.Message)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}
func TestMutationResolver_RefreshToken(t *testing.T) {
	const (
		success = iota
		errorRefreshingToken
	)

	var tests = []struct {
		name     string
		arg      string
		testType int
	}{
		{
			name:     "Test refresh token success",
			arg:      "sample-jwt-token",
			testType: success,
		},

		{
			name:     "Test error refresh token",
			arg:      "",
			testType: errorRefreshingToken,
		},
	}

	for _, testCase := range tests {
		controller := gomock.NewController(t)
		defer controller.Finish()
		authServiceClient := mocks.NewMockAuthServiceClient(controller)
		resolverOpts := &ResolverOpts{
			AuthService: authServiceClient,
		}

		resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				ctx := context.Background()
				authServiceClient.EXPECT().RefreshToken(ctx, &auth.RefreshTokenRequest{Token: testCase.arg}).
					Return(
						&auth.TokenPairResponse{
							AuthToken:    "valid-auth-token",
							RefreshToken: "refreshed-token",
						}, nil).Times(1)

				resp, err := resolver.RefreshToken(ctx, testCase.arg)
				assert.NotNil(t, resp)
				assert.NoError(t, err)

			case errorRefreshingToken:
				authServiceClient.EXPECT().RefreshToken(context.Background(), &auth.RefreshTokenRequest{Token: testCase.arg}).
					Return(nil, errors.New("")).Times(1)

				resp, err := resolver.RefreshToken(context.Background(), testCase.arg)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})

	}
}

func TestMutationResolver_SetDeviceToken(t *testing.T) {
	const (
		success = iota
		errUserAuthentication
	)

	var tests = []struct {
		name     string
		args     []*types.DeviceTokenInput
		testType int
	}{
		{
			name: "Test set device token successful",
			args: []*types.DeviceTokenInput{
				{
					Type:  types.DeviceTokenTypesFirebase,
					Value: "hjhfwifwr83283r9nvow9r8r731nvpo1391_=38238r",
				},
			},
			testType: success,
		},
		{
			name:     "Test set device token error authenticating user",
			args:     []*types.DeviceTokenInput{},
			testType: errUserAuthentication,
		},
	}

	for _, testCase := range tests {
		controller := gomock.NewController(t)
		defer controller.Finish()
		customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
		resolverOpts := &ResolverOpts{
			CustomerService: customerServiceClient,
		}

		resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

		switch testCase.testType {

		case success:
			ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{
				Client:   models.APP,
				ID:       "123456",
				Email:    "f@roava.app",
				DeviceID: "129594533fsdd"})

			customerServiceClient.EXPECT().SetDeviceToken(ctx,
				&customer.SetDeviceTokenRequest{
					Tokens: []*pbTypes.DeviceTokenInput{
						{
							Type:  pbTypes.DeviceToken_FIREBASE,
							Value: "hjhfwifwr83283r9nvow9r8r731nvpo1391_=38238r",
						},
					},
				},
			).Return(&pbTypes.DefaultResponse{
				Success: true,
				Code:    200,
			}, nil)

			resp, err := resolver.SetDeviceToken(ctx, testCase.args)
			assert.NoError(t, err)
			assert.Equal(t, resp.Code, int64(200))

		case errUserAuthentication:
			ctx := context.Background()

			resp, err := resolver.SetDeviceToken(ctx, testCase.args)
			assert.Error(t, err)
			assert.Equal(t, resp.Code, int64(500))
		}
	}
}

func TestMutationResolver_SetDevicePreferences(t *testing.T) {
	const (
		success = iota
		errUserAuthentication
	)

	var tests = []struct {
		name     string
		args     []*types.DevicePreferencesInput
		testType int
	}{
		{
			name: "Test set device preferences successful",
			args: []*types.DevicePreferencesInput{
				{
					Type:  types.DevicePreferencesTypesPush,
					Value: true,
				},
			},
			testType: success,
		},
		{
			name: "Test set device preferences error authenticating user",
			args: []*types.DevicePreferencesInput{
				{
					Type:  types.DevicePreferencesTypesPush,
					Value: true,
				},
			},
			testType: errUserAuthentication,
		},
	}

	for _, testCase := range tests {
		controller := gomock.NewController(t)
		defer controller.Finish()
		customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
		resolverOpts := &ResolverOpts{
			CustomerService: customerServiceClient,
		}

		resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

		switch testCase.testType {

		case success:
			ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{
				Client:   models.APP,
				ID:       "123456",
				Email:    "f@roava.app",
				DeviceID: "129594533fsdd"})

			customerServiceClient.EXPECT().SetDevicePreferences(ctx,
				&customer.SetDevicePreferencesRequest{
					Preferences: []*pbTypes.DevicePreferencesInput{
						{
							Type:  pbTypes.DevicePreferences_PUSH,
							Value: true,
						},
					},
				},
			).Return(&pbTypes.Device{Id: "deviceId"}, nil)

			resp, err := resolver.SetDevicePreferences(ctx, testCase.args)
			assert.NoError(t, err)
			assert.NotNil(t, resp)

		case errUserAuthentication:
			ctx := context.Background()

			resp, err := resolver.SetDevicePreferences(ctx, testCase.args)
			assert.Error(t, err)
			assert.Equal(t, resp.Success, false)
		}
	}
}

func TestMutationResolver_CheckBvn(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CheckBvn(context.Background(), "", "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	accountServiceClient := mocks.NewMockAccountServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateAccount(context.Background(), types.AccountInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateVaultAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	accountServiceClient := mocks.NewMockAccountServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateVaultAccount(context.Background(), types.VaultAccountInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateBeneficiary(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateBeneficiary(context.Background(), types.BeneficiaryInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_AddBeneficiaryAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.AddBeneficiaryAccount(context.Background(), "", types.BeneficiaryAccountInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_DeleteBeneficaryAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.DeleteBeneficaryAccount(context.Background(), "", "", "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateTransfer(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	paymentServiceClient := mocks.NewMockPaymentServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateTransfer(context.Background(), types.TransactionInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_RequestResubmit(t *testing.T) {
	const (
		success = iota
		errorUnauthenticated
		errorCallingRPC
	)

	type arg struct {
		customerID string
		reportIds  []string
		message    *string
	}
	message := "Message"
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				customerID: "customerId",
				reportIds:  []string{"reportId1", "reportId2"},
				message:    &message,
			},
			testType: success,
		},
		{
			name:     "Test error unathenticated user",
			testType: errorUnauthenticated,
		},
		{
			name: "Test error requesting resubmit",
			arg: arg{
				customerID: "customerId",
				reportIds:  []string{"reportId1", "reportId2"},
				message:    &message,
			},
			testType: errorCallingRPC,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			onboardingServiceClient := mocks.NewMockOnboardingServiceClient(controller)
			resolverOpts := &ResolverOpts{
				OnboardingService: onboardingServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().RequestResubmit(ctx, &onboarding.RequestResubmitRequest{
					CustomerId: testCase.arg.customerID,
					ReportIds:  testCase.arg.reportIds,
					Message:    *testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.RequestResubmit(ctx, testCase.arg.customerID, testCase.arg.reportIds,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.RequestResubmit(context.Background(), testCase.arg.customerID, testCase.arg.reportIds,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusUnauthorized,
				}, resp)
			case errorCallingRPC:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().RequestResubmit(ctx, &onboarding.RequestResubmitRequest{
					CustomerId: testCase.arg.customerID,
					ReportIds:  testCase.arg.reportIds,
					Message:    *testCase.arg.message,
				}).Return(nil, errors.New(""))

				resp, err := resolver.RequestResubmit(ctx, testCase.arg.customerID, testCase.arg.reportIds,
					testCase.arg.message)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_StaffLogin(t *testing.T) {
	const (
		failAuthServiceError = iota
		success
	)

	authServiceResponse := &auth.TokenPairResponse{
		AuthToken:    "auth-token",
		RefreshToken: "refresh-token",
	}

	testCases := []struct {
		name     string
		input    string
		testType int
	}{

		{
			name:     "Fail on auth service error",
			input:    "some-token",
			testType: failAuthServiceError,
		},

		{
			name:     "Success request",
			input:    "some-token",
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	authServiceClient := mocks.NewMockAuthServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failAuthServiceError:
				authServiceClient.EXPECT().StaffLogin(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.StaffLogin(context.Background(), testCase.input, types.AuthTypeGoogle)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, errorvalues.Message(errorvalues.InternalErr), *resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				authServiceClient.EXPECT().StaffLogin(gomock.Any(), gomock.Any()).Return(authServiceResponse, nil).Times(1)
				resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
				resp, err := resolver.StaffLogin(context.Background(), testCase.input, types.AuthTypeGoogle)
				assert.Nil(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, "Success", *resp.Message)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_UpdateKYCStatus(t *testing.T) {
	const (
		success = iota
		successManualReview
		successApproved
		successDeclined
		errorUnauthenticated
		errorCallRPC
	)

	type arg struct {
		id      string
		status  types.KYCStatuses
		message string
	}
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				id:      "kycId",
				status:  types.KYCStatusesPending,
				message: "Message",
			},
			testType: success,
		},
		{
			name: "Test success manual review",
			arg: arg{
				id:      "kycId",
				status:  types.KYCStatusesManualReview,
				message: "Message",
			},
			testType: successManualReview,
		},
		{
			name: "Test success approved",
			arg: arg{
				id:      "kycId",
				status:  types.KYCStatusesApproved,
				message: "Message",
			},
			testType: successApproved,
		},
		{
			name: "Test success declined",
			arg: arg{
				id:      "kycId",
				status:  types.KYCStatusesDeclined,
				message: "Message",
			},
			testType: successDeclined,
		},
		{
			name:     "Test error customer not authenticated",
			testType: errorUnauthenticated,
		},
		{
			name: "Test error calling updating kyc status",
			arg: arg{
				id:      "kycId",
				status:  types.KYCStatusesPending,
				message: "Message",
			},
			testType: errorCallRPC,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			onboardingServiceClient := mocks.NewMockOnboardingServiceClient(controller)
			resolverOpts := &ResolverOpts{
				OnboardingService: onboardingServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateKYCStatus(ctx, &onboarding.UpdateKYCStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.KYC_PENDING,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateKYCStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successManualReview:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateKYCStatus(ctx, &onboarding.UpdateKYCStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.KYC_MANUAL_REVIEW,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateKYCStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successApproved:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateKYCStatus(ctx, &onboarding.UpdateKYCStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.KYC_APPROVED,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateKYCStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successDeclined:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateKYCStatus(ctx, &onboarding.UpdateKYCStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.KYC_DECLINED,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateKYCStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.UpdateKYCStatus(context.Background(), testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusUnauthorized,
				}, resp)
			case errorCallRPC:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateKYCStatus(ctx, &onboarding.UpdateKYCStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.KYC_PENDING,
					Message: testCase.arg.message,
				}).Return(nil, errors.New(""))

				resp, err := resolver.UpdateKYCStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}

}

func TestMutationResolver_UpdateAMLStatus(t *testing.T) {
	const (
		success = iota
		successManualReview
		successApproved
		successDeclined
		errorUnauthenticated
		errorCallRPC
	)

	type arg struct {
		id      string
		status  types.AMLStatuses
		message string
	}
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				id:      "amlId",
				status:  types.AMLStatusesPending,
				message: "Message",
			},
			testType: success,
		},
		{
			name: "Test success manual review",
			arg: arg{
				id:      "amlId",
				status:  types.AMLStatusesManualReview,
				message: "Message",
			},
			testType: successManualReview,
		},
		{
			name: "Test success approved",
			arg: arg{
				id:      "amlId",
				status:  types.AMLStatusesApproved,
				message: "Message",
			},
			testType: successApproved,
		},
		{
			name: "Test success declined",
			arg: arg{
				id:      "amlId",
				status:  types.AMLStatusesDeclined,
				message: "Message",
			},
			testType: successDeclined,
		},
		{
			name:     "Test error customer not authenticated",
			testType: errorUnauthenticated,
		},
		{
			name: "Test error calling updating kyc status",
			arg: arg{
				id:      "amlId",
				status:  types.AMLStatusesPending,
				message: "Message",
			},
			testType: errorCallRPC,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			onboardingServiceClient := mocks.NewMockOnboardingServiceClient(controller)
			resolverOpts := &ResolverOpts{
				OnboardingService: onboardingServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateAMLStatus(ctx, &onboarding.UpdateAMLStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.AML_PENDING,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateAMLStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successManualReview:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateAMLStatus(ctx, &onboarding.UpdateAMLStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.AML_MANUAL_REVIEW,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateAMLStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successApproved:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateAMLStatus(ctx, &onboarding.UpdateAMLStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.AML_APPROVED,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateAMLStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successDeclined:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateAMLStatus(ctx, &onboarding.UpdateAMLStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.AML_DECLINED,
					Message: testCase.arg.message,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateAMLStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.UpdateAMLStatus(context.Background(), testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusUnauthorized,
				}, resp)
			case errorCallRPC:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				onboardingServiceClient.EXPECT().UpdateAMLStatus(ctx, &onboarding.UpdateAMLStatusRequest{
					Id:      testCase.arg.id,
					Status:  pbTypes.AML_PENDING,
					Message: testCase.arg.message,
				}).Return(nil, errors.New(""))

				resp, err := resolver.UpdateAMLStatus(ctx, testCase.arg.id, testCase.arg.status,
					testCase.arg.message)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}
