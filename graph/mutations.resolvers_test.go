package graph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	terror "github.com/roava/zebra/errors"
	"github.com/roava/zebra/middleware"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	errorvalues "ms.api/libs/errors"
	devicevalidator "ms.api/libs/validator/device"
	emailvalidator "ms.api/libs/validator/email"
	phonenumbervalidator "ms.api/libs/validator/phonenumbervalidator"
	"ms.api/mocks"
	accountPb "ms.api/protos/pb/account"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/messaging"
	"ms.api/protos/pb/onboarding"
	"ms.api/protos/pb/payment"
	"ms.api/protos/pb/pricing"
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
		Title:     types.CustomerTitleMr,
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
		name string
		args struct {
			email       string
			deviceInput types.DeviceInput
		}
		testType int
	}{
		{
			name: "Test check email found successful",
			args: struct {
				email       string
				deviceInput types.DeviceInput
			}{
				email: "f@mail.com",
				deviceInput: types.DeviceInput{
					Identifier: "identifier",
					Os:         "IOS",
					Brand:      "iPhoneX",
					Tokens: []*types.DeviceTokenInput{
						{
							Type:  types.DeviceTokenTypesFirebase,
							Value: "firebase-token-string",
						},
					},
					Preferences: []*types.DevicePreferencesInput{
						{
							Type:  types.DevicePreferencesTypesPush,
							Value: false,
						},
					},
				},
			},
			testType: success,
		},

		{
			name: "Test error check email not found",
			args: struct {
				email       string
				deviceInput types.DeviceInput
			}{
				email: "noEmail@mail.com",
				deviceInput: types.DeviceInput{
					Identifier: "xxxxxxxxxx",
					Os:         "Android",
					Brand:      "Samsung",
					Tokens: []*types.DeviceTokenInput{
						{
							Type:  types.DeviceTokenTypesFirebase,
							Value: "sample-firebase-token-string",
						},
					},
					Preferences: []*types.DevicePreferencesInput{
						{
							Type:  types.DevicePreferencesTypesPush,
							Value: false,
						},
					},
				},
			},
			testType: emailNotFound,
		},
		{
			name: "Test invalid email error",
			args: struct {
				email       string
				deviceInput types.DeviceInput
			}{
				email: "invalidEmail",
				deviceInput: types.DeviceInput{
					Identifier: "abczzabczz",
					Os:         "Android",
					Brand:      "Nokia1",
					Tokens: []*types.DeviceTokenInput{
						{
							Type:  types.DeviceTokenTypesFirebase,
							Value: "sample-firebase-token-string",
						},
					},
					Preferences: []*types.DevicePreferencesInput{
						{
							Type:  types.DevicePreferencesTypesPush,
							Value: false,
						},
					},
				},
			},
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
			helper := helpersfactory{}
			switch test.testType {
			case success:
				tokens := make([]*pbTypes.DeviceTokenInput, len(test.args.deviceInput.Tokens))
				for index, deviceToken := range test.args.deviceInput.Tokens {
					tokens[index] = &pbTypes.DeviceTokenInput{
						Type:  helper.GetProtoDeviceTokenType(deviceToken.Type),
						Value: deviceToken.Value,
					}
				}

				preferences := make([]*pbTypes.DevicePreferencesInput, len(test.args.deviceInput.Preferences))
				for index, devicePreference := range test.args.deviceInput.Preferences {
					preferences[index] = &pbTypes.DevicePreferencesInput{
						Type:  helper.GetProtoDevicePreferencesType(devicePreference.Type),
						Value: devicePreference.Value,
					}
				}

				request := &customer.CheckCustomerEmailRequest{
					Email: test.args.email,
					Device: &pbTypes.DeviceInput{
						Identifier:  test.args.deviceInput.Identifier,
						Os:          test.args.deviceInput.Os,
						Brand:       test.args.deviceInput.Brand,
						Tokens:      tokens,
						Preferences: preferences,
					},
				}

				customerServiceClient.EXPECT().CheckCustomerEmail(context.Background(), request).Return(
					&pbTypes.DefaultResponse{Success: true, Code: 1}, nil,
				)

				resp, err := resolver.CheckCustomerEmail(context.Background(), test.args.email, test.args.deviceInput)
				assert.NoError(t, err)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    1,
				}, resp)

			case emailNotFound:
				tokens := make([]*pbTypes.DeviceTokenInput, len(test.args.deviceInput.Tokens))
				for index, deviceToken := range test.args.deviceInput.Tokens {
					tokens[index] = &pbTypes.DeviceTokenInput{
						Type:  helper.GetProtoDeviceTokenType(deviceToken.Type),
						Value: deviceToken.Value,
					}
				}

				preferences := make([]*pbTypes.DevicePreferencesInput, len(test.args.deviceInput.Preferences))
				for index, devicePreference := range test.args.deviceInput.Preferences {
					preferences[index] = &pbTypes.DevicePreferencesInput{
						Type:  helper.GetProtoDevicePreferencesType(devicePreference.Type),
						Value: devicePreference.Value,
					}
				}

				request := &customer.CheckCustomerEmailRequest{
					Email: test.args.email,
					Device: &pbTypes.DeviceInput{
						Identifier:  test.args.deviceInput.Identifier,
						Os:          test.args.deviceInput.Os,
						Brand:       test.args.deviceInput.Brand,
						Tokens:      tokens,
						Preferences: preferences,
					},
				}

				customerServiceClient.EXPECT().CheckCustomerEmail(context.Background(), request).Return(
					nil, errorvalues.Format(500, errors.New("")),
				)

				resp, err := resolver.CheckCustomerEmail(context.Background(), test.args.email, test.args.deviceInput)
				assert.Error(t, err)
				assert.Nil(t, resp)

			case invalidEmail:
				resp, err := resolver.CheckCustomerEmail(context.Background(), test.args.email, test.args.deviceInput)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_CheckCustomerData(t *testing.T) {
	const (
		success = iota
		invalidEmail
		invalidDob
		customerNotFound
	)

	tests := []struct {
		name              string
		customerDataInput types.CheckCustomerDataInput
		testType          int
	}{
		{
			name: "Test checkCustomerData successfully",
			customerDataInput: types.CheckCustomerDataInput{
				Email:            "test@mail.com",
				FirstName:        "first_name",
				LastName:         "last_name",
				Dob:              "01-02-2001",
				AccountNumber:    "00011133",
				SortCode:         "040695",
				DeviceIdentifier: "zzz-ccxx-aaa",
			},
			testType: success,
		},

		{
			name: "Test error check customer data with invalid email",
			customerDataInput: types.CheckCustomerDataInput{
				Email:            "invalidEmail",
				FirstName:        "first_name",
				LastName:         "last_name",
				Dob:              "01-02-2001",
				AccountNumber:    "00011133",
				SortCode:         "040695",
				DeviceIdentifier: "zzz-ccxx-aaa",
			},
			testType: invalidEmail,
		},

		{
			name: "Test error check customer data with invalid Dob",
			customerDataInput: types.CheckCustomerDataInput{
				Email:            "invalidEmail",
				FirstName:        "first_name",
				LastName:         "last_name",
				Dob:              "2001/01/02",
				AccountNumber:    "00011133",
				SortCode:         "040695",
				DeviceIdentifier: "zzz-ccxx-aaa",
			},
			testType: invalidDob,
		},

		{
			name: "Test error check customer data not found",
			customerDataInput: types.CheckCustomerDataInput{
				Email:            "noEmail@mail.com",
				FirstName:        "first_name",
				LastName:         "last_name",
				Dob:              "01-10-1998",
				AccountNumber:    "00011133",
				SortCode:         "040695",
				DeviceIdentifier: "zzz-ccxx-aaa",
			},
			testType: customerNotFound,
		},
	}

	for _, test := range tests {

		controller := gomock.NewController(t)
		defer controller.Finish()
		customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
		resolverOpts := &ResolverOpts{
			CustomerService: customerServiceClient,
		}
		resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case success:
				request := &customer.CheckCustomerDataRequest{
					Email:            test.customerDataInput.Email,
					FirstName:        test.customerDataInput.FirstName,
					LastName:         test.customerDataInput.LastName,
					Dob:              test.customerDataInput.Dob,
					AccountNumber:    test.customerDataInput.AccountNumber,
					SortCode:         test.customerDataInput.SortCode,
					DeviceIdentifier: test.customerDataInput.DeviceIdentifier,
				}

				customerServiceClient.EXPECT().CheckCustomerData(context.Background(), request).Return(
					&pbTypes.DefaultResponse{
						Success: true,
						Code:    1,
					}, nil)

				resp, err := resolver.CheckCustomerData(context.Background(), test.customerDataInput)
				assert.NoError(t, err)
				assert.Equal(t, &types.Response{Success: true, Code: int64(1)}, resp)

			case invalidEmail:
				resp, err := resolver.CheckCustomerData(context.Background(), test.customerDataInput)
				assert.Error(t, err)
				assert.Nil(t, resp)

			case invalidDob:
				resp, err := resolver.CheckCustomerData(context.Background(), test.customerDataInput)
				assert.Error(t, err)
				assert.Nil(t, resp)

			case customerNotFound:
				request := &customer.CheckCustomerDataRequest{
					Email:            test.customerDataInput.Email,
					FirstName:        test.customerDataInput.FirstName,
					LastName:         test.customerDataInput.LastName,
					Dob:              test.customerDataInput.Dob,
					AccountNumber:    test.customerDataInput.AccountNumber,
					SortCode:         test.customerDataInput.SortCode,
					DeviceIdentifier: test.customerDataInput.DeviceIdentifier,
				}

				customerServiceClient.EXPECT().CheckCustomerData(context.Background(), request).Return(nil, errors.New(""))

				resp, err := resolver.CheckCustomerData(context.Background(), test.customerDataInput)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
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
					Title:     pbTypes.Customer_MR,
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
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusInternalServerError,
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
	const (
		success = iota
		errorUnauthenticated
		errorSettingTransactionPassword
	)

	var tests = []struct {
		name     string
		arg      string
		testType int
	}{
		{
			name:     "Test success",
			arg:      "password",
			testType: success,
		},
		{
			name:     "Test error unauthenticated user",
			arg:      "password",
			testType: errorUnauthenticated,
		},
		{
			name:     "Test error setting transaction password",
			arg:      "password",
			testType: errorSettingTransactionPassword,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			validCtx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})
			if err != nil {
				assert.NoError(t, err)
				t.Fail()
			}
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
			switch testCase.testType {
			case success:
				customerServiceClient.EXPECT().SetTransactionPassword(validCtx, &customer.SetTransactionPasswordRequest{
					Password: testCase.arg,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SetTransactionPassword(validCtx, testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.SetTransactionPassword(context.Background(), testCase.arg)
				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthenticationError, err.(*terror.Terror).Code())
				assert.Nil(t, resp)
			case errorSettingTransactionPassword:
				customerServiceClient.EXPECT().SetTransactionPassword(validCtx, &customer.SetTransactionPasswordRequest{
					Password: testCase.arg,
				}).Return(nil, errors.New(""))

				resp, err := resolver.SetTransactionPassword(validCtx, testCase.arg)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_ForgotTransactionPassword(t *testing.T) {
	const (
		success = iota
		errorUnauthenticated
		errorSettingTransactionPassword
	)

	var tests = []struct {
		name     string
		arg      string
		testType int
	}{
		{
			name:     "Test success",
			arg:      "password",
			testType: success,
		},
		{
			name:     "Test error unauthenticated user",
			arg:      "password",
			testType: errorUnauthenticated,
		},
		{
			name:     "Test error setting transaction password",
			arg:      "password",
			testType: errorSettingTransactionPassword,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			validCtx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})
			if err != nil {
				assert.NoError(t, err)
				t.Fail()
			}
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
			switch testCase.testType {
			case success:
				customerServiceClient.EXPECT().ForgotTransactionPassword(validCtx, &customer.ForgotTransactionPasswordRequest{
					NewPassword: testCase.arg,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.ForgotTransactionPassword(validCtx, testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.ForgotTransactionPassword(context.Background(), testCase.arg)
				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthenticationError, err.(*terror.Terror).Code())
				assert.Nil(t, resp)
			case errorSettingTransactionPassword:
				customerServiceClient.EXPECT().ForgotTransactionPassword(validCtx, &customer.ForgotTransactionPasswordRequest{
					NewPassword: testCase.arg,
				}).Return(nil, errors.New(""))

				resp, err := resolver.ForgotTransactionPassword(validCtx, testCase.arg)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_ResetTransactionPassword(t *testing.T) {
	const (
		success = iota
		errorUnauthenticated
		errorResettingTransactionPassword
	)

	type arg struct {
		otpToken                   string
		email                      string
		newTransactionPassword     string
		currentTransactionPassword string
	}
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				otpToken:                   "token",
				email:                      "email@roava.app",
				newTransactionPassword:     "newTransactionPassword",
				currentTransactionPassword: "oldTransactionPassword",
			},
			testType: success,
		},
		{
			name:     "Test error unauthenticated customer",
			testType: errorUnauthenticated,
		},
		{
			name: "Test error resetting transaction password",
			arg: arg{
				otpToken:                   "token",
				email:                      "email@roava.app",
				newTransactionPassword:     "newTransactionPassword",
				currentTransactionPassword: "oldTransactionPassword",
			},
			testType: errorResettingTransactionPassword,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			validCtx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})
			if err != nil {
				assert.NoError(t, err)
				t.Fail()
			}
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
			switch testCase.testType {
			case success:
				customerServiceClient.EXPECT().ResetTransactionPassword(validCtx, &customer.ResetTransactionPasswordRequest{
					OtpToken:        testCase.arg.otpToken,
					Email:           testCase.arg.email,
					NewPassword:     testCase.arg.newTransactionPassword,
					CurrentPassword: testCase.arg.currentTransactionPassword,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.ResetTransactionPassword(validCtx, testCase.arg.otpToken, testCase.arg.email,
					testCase.arg.newTransactionPassword, testCase.arg.currentTransactionPassword)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.ResetTransactionPassword(context.Background(), testCase.arg.otpToken, testCase.arg.email,
					testCase.arg.newTransactionPassword, testCase.arg.currentTransactionPassword)
				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthenticationError, err.(*terror.Terror).Code())
				assert.Nil(t, resp)
			case errorResettingTransactionPassword:
				customerServiceClient.EXPECT().ResetTransactionPassword(validCtx, &customer.ResetTransactionPasswordRequest{
					OtpToken:        testCase.arg.otpToken,
					Email:           testCase.arg.email,
					NewPassword:     testCase.arg.newTransactionPassword,
					CurrentPassword: testCase.arg.currentTransactionPassword,
				}).Return(nil, errors.New(""))

				resp, err := resolver.ResetTransactionPassword(validCtx, testCase.arg.otpToken, testCase.arg.email,
					testCase.arg.newTransactionPassword, testCase.arg.currentTransactionPassword)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
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
	const (
		success = iota
		errorUnauthenticated
		errorCheckingBvn
	)

	type arg struct {
		bvn   string
		phone string
	}

	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				bvn:   "22241890998",
				phone: "08060223673",
			},
			testType: success,
		},
		{
			name: "Test error unauthenticated user",
			arg: arg{
				bvn:   "22241890998",
				phone: "08060223673",
			},
			testType: errorUnauthenticated,
		},
		{
			name: "Test error setting transaction password",
			arg: arg{
				bvn:   "22241890998",
				phone: "08060223673",
			},
			testType: errorCheckingBvn,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			validCtx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})
			if err != nil {
				assert.NoError(t, err)
				t.Fail()
			}
			controller := gomock.NewController(t)
			defer controller.Finish()
			customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
			resolverOpts := &ResolverOpts{
				CustomerService: customerServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
			switch testCase.testType {
			case success:
				customerServiceClient.EXPECT().CheckBVN(validCtx, &customer.CheckBVNRequest{
					Bvn:   testCase.arg.bvn,
					Phone: testCase.arg.phone,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.CheckBvn(validCtx, testCase.arg.bvn, testCase.arg.phone)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.CheckBvn(context.Background(), testCase.arg.bvn, testCase.arg.phone)
				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthenticationError, err.(*terror.Terror).Code())
				assert.Nil(t, resp)
			case errorCheckingBvn:
				customerServiceClient.EXPECT().CheckBVN(validCtx, &customer.CheckBVNRequest{
					Bvn:   testCase.arg.bvn,
					Phone: testCase.arg.phone,
				}).Return(nil, errors.New(""))

				resp, err := resolver.CheckBvn(validCtx, testCase.arg.bvn, testCase.arg.phone)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_CreateAccount(t *testing.T) {
	const (
		success = iota
		errorAuthentication
		errorCreatingAccount
	)

	var tests = []struct {
		name     string
		arg      types.AccountInput
		testType int
	}{
		{
			name: "Test name",
			arg: types.AccountInput{
				ProductID: "produtId",
			},
			testType: success,
		},
		{
			name:     "Test error unauthenticated",
			testType: errorAuthentication,
		},
		{
			name: "Test error calling CreateAccount",
			arg: types.AccountInput{
				ProductID: "produtId",
			},
			testType: errorCreatingAccount,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			accountServiceClient := mocks.NewMockAccountServiceClient(controller)
			resolverOpts := &ResolverOpts{
				AccountService: accountServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{
					Client:   models.APP,
					ID:       "123456",
					Email:    "email@roava.app",
					DeviceID: "12345"})

				request := accountPb.CreateAccountRequest{
					ProductId: testCase.arg.ProductID,
				}
				accountServiceClient.EXPECT().CreateAccount(ctx, &request).Return(&pbTypes.Account{Id: "accountId"}, nil)

				resp, err := resolver.CreateAccount(ctx, testCase.arg)

				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorAuthentication:
				resp, err := resolver.CreateAccount(context.Background(), testCase.arg)

				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthenticationError, err.(*terror.Terror).Code())
				assert.Nil(t, resp)
			case errorCreatingAccount:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{
					Client:   models.APP,
					ID:       "123456",
					Email:    "email@roava.app",
					DeviceID: "12345"})

				request := accountPb.CreateAccountRequest{
					ProductId: testCase.arg.ProductID,
				}
				accountServiceClient.EXPECT().CreateAccount(ctx, &request).Return(nil, errors.New(""))

				resp, err := resolver.CreateAccount(ctx, testCase.arg)

				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_CreateVaultAccount(t *testing.T) {
	const (
		failOnAuthenticationError = iota
		failOnGRPCError
		success
	)
	testCases := []struct {
		name     string
		testType int
	}{
		{
			name:     "should fail on authentication error",
			testType: failOnAuthenticationError,
		},
		{
			name:     "should fail on gRPC error",
			testType: failOnGRPCError,
		},
		{
			name:     "success",
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	accountServiceClient := mocks.NewMockAccountServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failOnAuthenticationError:
				resp, err := resolver.CreateVaultAccount(context.Background(), types.VaultAccountInput{}, "")
				assert.Error(t, err)
				assert.Nil(t, resp)
				switch newTerror := err.(type) {
				case *terror.Terror:
					assert.Equal(t, errorvalues.InvalidAuthenticationError, newTerror.Code())
				default:
					t.Error("Should return an error of type InvalidAuthenticationError")
					t.Fail()
				}
			case failOnGRPCError:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				accountServiceClient.EXPECT().CreateVaultAccount(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resp, err := resolver.CreateVaultAccount(ctx, types.VaultAccountInput{}, "")
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				accountServiceClient.EXPECT().CreateVaultAccount(gomock.Any(), gomock.Any()).Return(&pbTypes.Account{}, nil).Times(1)
				input := &types.VaultAccountInput{
					Name: func() *string {
						str := "test"
						return &str
					}(),
				}
				resp, err := resolver.CreateVaultAccount(ctx, *input, "")
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_CreateBeneficiary(t *testing.T) {
	const (
		failOnAuthenticationError = iota
		failOnGRPCError
		success
	)
	testCases := []struct {
		name     string
		testType int
	}{
		{
			name:     "should fail on authentication error",
			testType: failOnAuthenticationError,
		},
		{
			name:     "should fail on gRPC error",
			testType: failOnGRPCError,
		},
		{
			name:     "success",
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	paymentServiceClient := mocks.NewMockPaymentServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failOnAuthenticationError:
				resp, err := resolver.CreateBeneficiary(context.Background(), types.BeneficiaryInput{}, "")
				assert.Error(t, err)
				assert.Nil(t, resp)
				switch newTerror := err.(type) {
				case *terror.Terror:
					assert.Equal(t, errorvalues.InvalidAuthenticationError, newTerror.Code())
				default:
					t.Error("Should return an error of type InvalidAuthenticationError")
					t.Fail()
				}
			case failOnGRPCError:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().CreateBeneficiary(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resp, err := resolver.CreateBeneficiary(ctx, types.BeneficiaryInput{}, "")
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().CreateBeneficiary(gomock.Any(), gomock.Any()).Return(&pbTypes.Beneficiary{}, nil).Times(1)
				input := &types.BeneficiaryInput{}
				resp, err := resolver.CreateBeneficiary(ctx, *input, "")
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_CreateBeneficiariesByPhone(t *testing.T) {
	const (
		failOnAuthenticationError = iota
		failOnGRPCError
		success
	)
	testCases := []struct {
		name     string
		testType int
	}{
		{
			name:     "should fail on authentication error",
			testType: failOnAuthenticationError,
		},
		{
			name:     "should fail on gRPC error",
			testType: failOnGRPCError,
		},
		{
			name:     "success",
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	paymentServiceClient := mocks.NewMockPaymentServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failOnAuthenticationError:
				resp, err := resolver.CreateBeneficiariesByPhone(context.Background(), []*types.BeneficiaryByPhoneInput{}, "")
				assert.Error(t, err)
				assert.Nil(t, resp)
				switch newTerror := err.(type) {
				case *terror.Terror:
					assert.Equal(t, errorvalues.InvalidAuthenticationError, newTerror.Code())
				default:
					t.Error("Should return an error of type InvalidAuthenticationError")
					t.Fail()
				}
			case failOnGRPCError:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().CreateBeneficiariesByPhone(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resp, err := resolver.CreateBeneficiariesByPhone(ctx, []*types.BeneficiaryByPhoneInput{}, "")
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().CreateBeneficiariesByPhone(gomock.Any(), gomock.Any()).Return(&pbTypes.DefaultResponse{}, nil).Times(1)
				input := []*types.BeneficiaryByPhoneInput{}
				resp, err := resolver.CreateBeneficiariesByPhone(ctx, input, "")
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_AddBeneficiaryAccount(t *testing.T) {
	const (
		failOnAuthenticationError = iota
		failOnGRPCError
		success
	)
	testCases := []struct {
		name     string
		testType int
	}{
		{
			name:     "should fail on authentication error",
			testType: failOnAuthenticationError,
		},
		{
			name:     "should fail on gRPC error",
			testType: failOnGRPCError,
		},
		{
			name:     "success",
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	paymentServiceClient := mocks.NewMockPaymentServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failOnAuthenticationError:
				resp, err := resolver.AddBeneficiaryAccount(context.Background(), "", types.BeneficiaryAccountInput{}, "")
				assert.Error(t, err)
				assert.Nil(t, resp)
				switch newTerror := err.(type) {
				case *terror.Terror:
					assert.Equal(t, errorvalues.InvalidAuthenticationError, newTerror.Code())
				default:
					t.Error("Should return an error of type InvalidAuthenticationError")
					t.Fail()
				}
			case failOnGRPCError:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().AddBeneficiaryAccount(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resp, err := resolver.AddBeneficiaryAccount(ctx, "", types.BeneficiaryAccountInput{}, "")
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().AddBeneficiaryAccount(gomock.Any(), gomock.Any()).Return(&pbTypes.BeneficiaryAccount{}, nil).Times(1)
				input := &types.BeneficiaryAccountInput{}
				resp, err := resolver.AddBeneficiaryAccount(ctx, "", *input, "")
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_DeleteBeneficiaryAccount(t *testing.T) {
	const (
		failOnAuthenticationError = iota
		failOnGRPCError
		success
	)
	testCases := []struct {
		name     string
		testType int
	}{
		{
			name:     "should fail on authentication error",
			testType: failOnAuthenticationError,
		},
		{
			name:     "should fail on gRPC error",
			testType: failOnGRPCError,
		},
		{
			name:     "success",
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	paymentServiceClient := mocks.NewMockPaymentServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failOnAuthenticationError:
				resp, err := resolver.DeleteBeneficiaryAccount(context.Background(), "", "", "")
				assert.Error(t, err)
				assert.Nil(t, resp)
				switch newTerror := err.(type) {
				case *terror.Terror:
					assert.Equal(t, errorvalues.InvalidAuthenticationError, newTerror.Code())
				default:
					t.Error("Should return an error of type InvalidAuthenticationError")
					t.Fail()
				}
			case failOnGRPCError:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().DeleteBeneficiaryAccount(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resp, err := resolver.DeleteBeneficiaryAccount(ctx, "", "", "")
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().DeleteBeneficiaryAccount(gomock.Any(), gomock.Any()).Return(&pbTypes.DefaultResponse{}, nil).Times(1)
				resp, err := resolver.DeleteBeneficiaryAccount(ctx, "", "", "")
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_CreateTransfer(t *testing.T) {
	const (
		failOnAuthenticationError = iota
		failOnGRPCError
		success
	)
	testCases := []struct {
		name     string
		testType int
	}{
		{
			name:     "should fail on authentication error",
			testType: failOnAuthenticationError,
		},
		{
			name:     "should fail on gRPC error",
			testType: failOnGRPCError,
		},
		{
			name:     "success",
			testType: success,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	paymentServiceClient := mocks.NewMockPaymentServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failOnAuthenticationError:
				resp, err := resolver.CreateTransfer(context.Background(), types.TransactionInput{}, "")
				assert.Error(t, err)
				assert.Nil(t, resp)
				switch newTerror := err.(type) {
				case *terror.Terror:
					assert.Equal(t, errorvalues.InvalidAuthenticationError, newTerror.Code())
				default:
					t.Error("Should return an error of type InvalidAuthenticationError")
					t.Fail()
				}
			case failOnGRPCError:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().CreateTransfer(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)
				resp, err := resolver.CreateTransfer(ctx, types.TransactionInput{}, "")
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.False(t, resp.Success)
				assert.NotNil(t, resp.Message)
				assert.Equal(t, int64(http.StatusInternalServerError), resp.Code)
			case success:
				ctx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})
				assert.NoError(t, err)
				paymentServiceClient.EXPECT().CreateTransfer(gomock.Any(), gomock.Any()).Return(&pbTypes.DefaultResponse{Success: true, Code: http.StatusOK}, nil).Times(1)
				reference := "some-reference"
				exchangeRateId := "some-exchange-rate-id"
				resp, err := resolver.CreateTransfer(ctx, types.TransactionInput{Reference: &reference, ExchangeRateID: &exchangeRateId}, "")
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.True(t, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			}
		})
	}
}

func TestMutationResolver_SendNotification(t *testing.T) {
	const (
		success = iota
		successSMS
		successPush
		errorUnauthenticated
		errorCallingRPC

		templateId        = "templateId"
		emailDeliveryMode = types.DeliveryMode("EMAIL")
		smsDeliveryMode   = types.DeliveryMode("SMS")
		pushDeliveryMode  = types.DeliveryMode("PUSH")
	)

	type arg struct {
		deliveryMode types.DeliveryMode
		content      string
		templateID   string
	}
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test success",
			arg: arg{
				deliveryMode: emailDeliveryMode,
				content:      "Something to send in success!",
				templateID:   templateId,
			},
			testType: success,
		},
		{
			name: "Test success SMS",
			arg: arg{
				deliveryMode: smsDeliveryMode,
				content:      "Something to send in sms success",
				templateID:   templateId,
			},
			testType: successSMS,
		},
		{
			name: "Test success Push",
			arg: arg{
				deliveryMode: pushDeliveryMode,
				content:      "Something to send in push success",
				templateID:   templateId,
			},
			testType: successPush,
		},
		{
			name:     "Test error unathenticated user",
			testType: errorUnauthenticated,
		},
		{
			name: "Test error requesting resubmit",
			arg: arg{
				deliveryMode: pushDeliveryMode,
				content:      "Something to send in push success",
				templateID:   templateId,
			},
			testType: errorCallingRPC,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			messagingServiceClient := mocks.NewMockMessagingServiceClient(controller)
			resolverOpts := &ResolverOpts{
				MessagingService: messagingServiceClient,
			}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				messagingServiceClient.EXPECT().SendNotification(ctx, &messaging.SendNotificationRequest{
					Type:       messaging.SendNotificationRequest_DeliveryMode(0),
					Content:    testCase.arg.content,
					TemplateId: testCase.arg.templateID,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SendNotification(ctx, testCase.arg.deliveryMode, testCase.arg.content, testCase.arg.templateID)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successSMS:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				messagingServiceClient.EXPECT().SendNotification(ctx, &messaging.SendNotificationRequest{
					Type:       messaging.SendNotificationRequest_DeliveryMode(1),
					Content:    testCase.arg.content,
					TemplateId: testCase.arg.templateID,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SendNotification(ctx, testCase.arg.deliveryMode, testCase.arg.content, testCase.arg.templateID)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case successPush:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				messagingServiceClient.EXPECT().SendNotification(ctx, &messaging.SendNotificationRequest{
					Type:       messaging.SendNotificationRequest_DeliveryMode(2),
					Content:    testCase.arg.content,
					TemplateId: testCase.arg.templateID,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.SendNotification(ctx, testCase.arg.deliveryMode, testCase.arg.content, testCase.arg.templateID)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.SendNotification(context.Background(), testCase.arg.deliveryMode, testCase.arg.content, testCase.arg.templateID)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusInternalServerError,
				}, resp)
			case errorCallingRPC:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "customerID"})

				messagingServiceClient.EXPECT().SendNotification(ctx, &messaging.SendNotificationRequest{
					Type:       messaging.SendNotificationRequest_DeliveryMode(2),
					Content:    testCase.arg.content,
					TemplateId: testCase.arg.templateID,
				}).Return(nil, errors.New(""))

				resp, err := resolver.SendNotification(ctx, testCase.arg.deliveryMode, testCase.arg.content, testCase.arg.templateID)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
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
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusInternalServerError,
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
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusInternalServerError,
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
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusInternalServerError,
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

func TestMutationResolver_DeactivateCredential(t *testing.T) {
	const (
		successLogin = iota
		successPin
		errorUnauthenticated
		errorDeactivatingCredential
	)

	var tests = []struct {
		name     string
		arg      types.IdentityCredentialsTypes
		testType int
	}{
		{
			name:     "Test success deactivate login password",
			arg:      types.IdentityCredentialsTypesLogin,
			testType: successLogin,
		},
		{
			name:     "Test success deactivate transaction password",
			arg:      types.IdentityCredentialsTypesPin,
			testType: successPin,
		},
		{
			name:     "Test error unauthenticated user",
			arg:      types.IdentityCredentialsTypesLogin,
			testType: errorUnauthenticated,
		},
		{
			name:     "Test error deactivating credential",
			arg:      types.IdentityCredentialsTypesLogin,
			testType: errorDeactivatingCredential,
		},
	}

	validCtx, err := middleware.PutClaimsOnContext(
		context.Background(),
		&models.JWTClaims{},
	)
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}

	mockResponse := pbTypes.DefaultResponse{
		Success: true,
		Code:    http.StatusOK,
	}

	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case successLogin:
				request := customer.DeactivateCredentialRequest{
					CredentialType: pbTypes.IdentityCredentials_LOGIN,
				}

				customerServiceClient.EXPECT().
					DeactivateCredential(validCtx, &request).
					Return(&mockResponse, nil)

				resp, err := resolver.DeactivateCredential(validCtx, testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, true, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			case successPin:
				request := customer.DeactivateCredentialRequest{
					CredentialType: pbTypes.IdentityCredentials_PIN,
				}

				customerServiceClient.EXPECT().
					DeactivateCredential(validCtx, &request).
					Return(&mockResponse, nil)

				resp, err := resolver.DeactivateCredential(validCtx, testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, true, resp.Success)
				assert.Equal(t, int64(http.StatusOK), resp.Code)
			case errorUnauthenticated:
				resp, err := resolver.DeactivateCredential(context.Background(), testCase.arg)
				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthenticationError, err.(*terror.Terror).Code())
				assert.Nil(t, resp)
			case errorDeactivatingCredential:
				request := customer.DeactivateCredentialRequest{
					CredentialType: pbTypes.IdentityCredentials_LOGIN,
				}

				customerServiceClient.EXPECT().
					DeactivateCredential(validCtx, &request).
					Return(nil, errors.New(""))

				resp, err := resolver.DeactivateCredential(validCtx, testCase.arg)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_UpdateCustomerDetails(t *testing.T) {
	const (
		success = iota
		errorUnauthenticated
		errorUpdatingCustomerDetails
	)
	type arg struct {
		CustomerDetails     types.CustomerDetailsUpdateInput
		TransactionPassword string
	}
	var (
		firstName = "firstName"
		lastName  = "lastName"
		phone     = "07000000000"
		email     = "janedoe@gmail.com"
		state     = "state"
		city      = "city"
	)
	var tests = []struct {
		name     string
		arg      arg
		testType int
	}{
		{
			name: "Test Success",
			arg: arg{
				CustomerDetails: types.CustomerDetailsUpdateInput{
					FirstName: &firstName,
					LastName:  &lastName,
					Phone:     &phone,
					Email:     &email,
					Address: &types.AddressInput{
						CountryID: "countryId",
						State:     &state,
						City:      &city,
						Street:    "street",
						Postcode:  "postCode",
						Cordinates: &types.CordinatesInput{
							Latitude:  1.2333,
							Longitude: 1.4555,
						},
					},
				},
			},
			testType: success,
		},
		{
			name:     "Test error unauthenticated customer",
			testType: errorUnauthenticated,
		},
		{
			name: "Test error updating customer details",
			arg: arg{
				CustomerDetails: types.CustomerDetailsUpdateInput{
					FirstName: &firstName,
					LastName:  &lastName,
					Phone:     &phone,
					Email:     &email,
					Address: &types.AddressInput{
						CountryID: "countryId",
						State:     &state,
						City:      &city,
						Street:    "street",
						Postcode:  "postCode",
						Cordinates: &types.CordinatesInput{
							Latitude:  1.2333,
							Longitude: 1.4555,
						},
					},
				},
			},
			testType: errorUpdatingCustomerDetails,
		},
	}

	validCtx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				customerServiceClient.EXPECT().CustomerDetailsUpdate(validCtx, &customer.CustomerDetailsUpdateRequest{
					FirstName: *testCase.arg.CustomerDetails.FirstName,
					LastName:  *testCase.arg.CustomerDetails.LastName,
					Phone:     *testCase.arg.CustomerDetails.Phone,
					Email:     *testCase.arg.CustomerDetails.Email,
					Address: &customer.AddressInput{
						CountryId: testCase.arg.CustomerDetails.Address.CountryID,
						State:     *testCase.arg.CustomerDetails.Address.State,
						City:      *testCase.arg.CustomerDetails.Address.City,
						Street:    testCase.arg.CustomerDetails.Address.Street,
						Postcode:  testCase.arg.CustomerDetails.Address.Postcode,
						Cordinates: &customer.CordinatesInput{
							Latitude:  float32(testCase.arg.CustomerDetails.Address.Cordinates.Latitude),
							Longitude: float32(testCase.arg.CustomerDetails.Address.Cordinates.Longitude),
						},
					},
					TransactionPassword: testCase.arg.TransactionPassword,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateCustomerDetails(validCtx, testCase.arg.CustomerDetails, testCase.arg.TransactionPassword)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				resp, err := resolver.UpdateCustomerDetails(context.Background(), testCase.arg.CustomerDetails, testCase.arg.TransactionPassword)
				assert.Error(t, err)
				assert.IsType(t, &terror.Terror{}, err)
				assert.Equal(t, errorvalues.InvalidAuthenticationError, err.(*terror.Terror).Code())
				assert.Nil(t, resp)
			case errorUpdatingCustomerDetails:
				customerServiceClient.EXPECT().CustomerDetailsUpdate(validCtx, &customer.CustomerDetailsUpdateRequest{
					FirstName: *testCase.arg.CustomerDetails.FirstName,
					LastName:  *testCase.arg.CustomerDetails.LastName,
					Phone:     *testCase.arg.CustomerDetails.Phone,
					Email:     *testCase.arg.CustomerDetails.Email,
					Address: &customer.AddressInput{
						CountryId: testCase.arg.CustomerDetails.Address.CountryID,
						State:     *testCase.arg.CustomerDetails.Address.State,
						City:      *testCase.arg.CustomerDetails.Address.City,
						Street:    testCase.arg.CustomerDetails.Address.Street,
						Postcode:  testCase.arg.CustomerDetails.Address.Postcode,
						Cordinates: &customer.CordinatesInput{
							Latitude:  float32(testCase.arg.CustomerDetails.Address.Cordinates.Latitude),
							Longitude: float32(testCase.arg.CustomerDetails.Address.Cordinates.Longitude),
						},
					},
					TransactionPassword: testCase.arg.TransactionPassword,
				}).Return(nil, errors.New(""))
				resp, err := resolver.UpdateCustomerDetails(validCtx, testCase.arg.CustomerDetails, testCase.arg.TransactionPassword)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_UpdateFx(t *testing.T) {
	const (
		success = iota
		errorCreatingFX
	)
	var FX = types.UpdateFXInput{
		BaseCurrencyID: "baseCurrencyId",
		CurrencyID:     "currencyId",
		BuyPrice:       0.1,
	}
	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test Success",
			testType: success,
		},
		{
			name:     "Test error creating FX",
			testType: errorCreatingFX,
		},
	}

	validCtx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	pricingServiceClient := mocks.NewMockPricingServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PricingService: pricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				pricingServiceClient.EXPECT().UpdateFX(validCtx, &pricing.UpdateFXRequest{
					BaseCurrencyId: FX.BaseCurrencyID,
					CurrencyId:     FX.CurrencyID,
					BuyPrice:       float32(FX.BuyPrice),
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateFx(validCtx, FX)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorCreatingFX:
				pricingServiceClient.EXPECT().UpdateFX(validCtx, &pricing.UpdateFXRequest{
					BaseCurrencyId: FX.BaseCurrencyID,
					CurrencyId:     FX.CurrencyID,
					BuyPrice:       float32(FX.BuyPrice),
				}).Return(nil, errors.New(""))
				resp, err := resolver.UpdateFx(validCtx, FX)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_UpdateFees(t *testing.T) {
	const (
		success = iota
		errorCreatingFees
	)
	var (
		lower      = 1000.0
		upper      = 2000.0
		amount     = 100.0
		percentage = 0.3
		fees       = []*types.UpdateFeesInput{
			{
				TransactionTypeID: "transactionTypeId",
				Type:              types.FeeTypesFixed,
				Boundaries: []*types.BoundaryFee{
					{
						Lower:      &lower,
						Upper:      &upper,
						Amount:     &amount,
						Percentage: &percentage,
					},
				},
			},
			{
				TransactionTypeID: "transactionTypeId",
				Type:              types.FeeTypesVariable,
				Boundaries: []*types.BoundaryFee{
					{
						Lower:      &lower,
						Upper:      &upper,
						Amount:     &amount,
						Percentage: &percentage,
					},
				},
			},
		}
		feesRequest []*pricing.UpdateFeesRequest
	)
	for _, fee := range fees {
		var feeRequest pricing.UpdateFeesRequest
		feeRequest.TransactionTypeId = fee.TransactionTypeID
		var boundaryRequest pbTypes.FeeBoundaries
		feeRequest.Type = pbTypes.Fee_FIXED
		for _, reqBoundary := range fee.Boundaries {
			boundaryRequest.Lower = float32(*reqBoundary.Lower)
			boundaryRequest.Upper = float32(*reqBoundary.Upper)
			boundaryRequest.Amount = float32(*reqBoundary.Amount)
			boundaryRequest.Percentage = float32(*reqBoundary.Percentage)
			feeRequest.Boundaries = append(feeRequest.Boundaries, &boundaryRequest)
		}
		feesRequest = append(feesRequest, &feeRequest)
	}

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test Success",
			testType: success,
		},
		{
			name:     "Test error creating FX",
			testType: errorCreatingFees,
		},
	}

	validCtx, err := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{})
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	pricingServiceClient := mocks.NewMockPricingServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PricingService: pricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				pricingServiceClient.EXPECT().UpdateFees(validCtx, &pricing.UpdateFeesRequests{
					Fees: feesRequest,
				}).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)
				resp, err := resolver.UpdateFees(validCtx, fees)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorCreatingFees:
				pricingServiceClient.EXPECT().UpdateFees(validCtx, &pricing.UpdateFeesRequests{
					Fees: feesRequest,
				}).Return(nil, errors.New(""))
				resp, err := resolver.UpdateFees(validCtx, fees)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}

}

func TestMutationResolver_WithdrawVaultAccount(t *testing.T) {
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

	mockClaims := models.JWTClaims{ID: "customerID"}

	mockRequest := payment.WithdrawVaultAccountRequest{
		SourceAccountId:     "SourceAccountId",
		TargetAccountId:     "TargetAccountId",
		TransactionPassword: "TransactionPassword",
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	paymentsServiceClient := mocks.NewMockPaymentServiceClient(controller)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentsServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &mockClaims)

				paymentsServiceClient.EXPECT().
					WithdrawVaultAccount(ctx, &mockRequest).
					Return(
						&pbTypes.DefaultResponse{
							Success: true,
							Code:    http.StatusOK,
						},
						nil,
					)

				resp, err := resolver.WithdrawVaultAccount(ctx, mockRequest.SourceAccountId, mockRequest.TargetAccountId, mockRequest.TransactionPassword)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				ctx := context.Background()
				resp, err := resolver.WithdrawVaultAccount(ctx, mockRequest.SourceAccountId, mockRequest.TargetAccountId, mockRequest.TransactionPassword)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusInternalServerError,
				}, resp)
			case errorCallingRPC:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &mockClaims)

				paymentsServiceClient.EXPECT().
					WithdrawVaultAccount(ctx, &mockRequest).
					Return(nil, errors.New("mock error"))

				resp, err := resolver.WithdrawVaultAccount(ctx, mockRequest.SourceAccountId, mockRequest.TargetAccountId, mockRequest.TransactionPassword)
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Contains(t, err.Error(), "mock error")
			}
		})
	}
}

func TestMutationResolver_CheckCustomerDetails(t *testing.T) {
	const (
		success = iota
		errorCheckingCustomerDetails
	)
	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "test Success",
			testType: success,
		},
		{
			name:     "Test error checking customer details",
			testType: errorCheckingCustomerDetails,
		},
	}
	customerDetails := types.CheckCustomerDetailsInput{
		Password:    "qr1234e",
		PhoneNumber: "uva",
		Dob:         "01-01-1900",
	}
	typeArg := types.ActionTypeDeviceUpdate
	request := &customer.CheckCustomerDetailsRequest{
		Password:    "qr1234e",
		PhoneNumber: "uva",
		Dob:         "01-01-1900",
		ActionType:  customer.CheckCustomerDetailsRequest_DEVICE_UPDATE,
	}
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				customerServiceClient.EXPECT().CheckCustomerDetails(context.Background(), request).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)
				resp, err := resolver.CheckCustomerDetails(context.Background(), customerDetails, typeArg)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorCheckingCustomerDetails:
				customerServiceClient.EXPECT().CheckCustomerDetails(context.Background(), request).Return(nil, errors.New(""))
				resp, err := resolver.CheckCustomerDetails(context.Background(), customerDetails, typeArg)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_UpdateDevice(t *testing.T) {
	const (
		success = iota
		errorUpdatingDevice
	)
	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "test Success",
			testType: success,
		},
		{
			name:     "Test error updating device",
			testType: errorUpdatingDevice,
		},
	}
	phoneNumber := "phoneNumber"
	device := types.DeviceInput{
		Identifier: "identifier",
		Os:         "os",
		Brand:      "brand",
		Tokens: []*types.DeviceTokenInput{
			{
				Type:  types.DeviceTokenTypesFirebase,
				Value: "hjhfwifwr83283r9nvow9r8r731nvpo1391_=38238r",
			},
		},
		Preferences: []*types.DevicePreferencesInput{
			{
				Type:  types.DevicePreferencesTypesPush,
				Value: true,
			},
		},
	}
	request := &customer.DeviceInputRequest{
		PhoneNumber: phoneNumber,
		Device: &pbTypes.DeviceInput{
			Identifier: device.Identifier,
			Os:         device.Os,
			Brand:      device.Brand,
			Tokens: []*pbTypes.DeviceTokenInput{
				{
					Type:  pbTypes.DeviceToken_FIREBASE,
					Value: "hjhfwifwr83283r9nvow9r8r731nvpo1391_=38238r",
				},
			},
			Preferences: []*pbTypes.DevicePreferencesInput{
				{
					Type:  pbTypes.DevicePreferences_PUSH,
					Value: true,
				},
			},
		},
	}
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				customerServiceClient.EXPECT().UpdateDevice(context.Background(), request).Return(&pbTypes.DefaultResponse{
					Success: true,
					Code:    http.StatusOK,
				}, nil)

				resp, err := resolver.UpdateDevice(context.Background(), phoneNumber, device)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUpdatingDevice:
				customerServiceClient.EXPECT().UpdateDevice(context.Background(), request).Return(nil, errors.New(""))
				resp, err := resolver.UpdateDevice(context.Background(), phoneNumber, device)
				assert.Error(t, err)
				assert.Nil(t, resp)
			}
		})
	}
}

func TestMutationResolver_StaffUpdateCustomerDetails(t *testing.T) {
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

	mockClaims := models.JWTClaims{ID: "customerID"}

	mockRequestModel := customer.StaffCustomerDetailsUpdateRequest{
		CustomerID: "mockRequestCustomerID",
		FirstName:  "mockRequestFirstName",
		LastName:   "mockRequestLastName",
		Email:      "mockRequestEmail",
		Address: &customer.AddressInput{
			CountryId: "mockRequestAddressCountryID",
			State:     "mockRequestAddressState",
			City:      "mockRequestAddressCity",
			Street:    "mockRequestAddressStreet",
			Postcode:  "mockRequestAddressPostCode",
			Cordinates: &customer.CordinatesInput{
				Latitude:  1,
				Longitude: 2,
			},
		},
	}
	mockRequest := types.StaffCustomerDetailsUpdateInput{
		CustomerID: mockRequestModel.CustomerID,
		FirstName:  &mockRequestModel.FirstName,
		LastName:   &mockRequestModel.LastName,
		Email:      &mockRequestModel.Email,
		Address: &types.AddressInput{
			CountryID: mockRequestModel.Address.CountryId,
			State:     &mockRequestModel.Address.State,
			City:      &mockRequestModel.Address.City,
			Street:    mockRequestModel.Address.Street,
			Postcode:  mockRequestModel.Address.Postcode,
			Cordinates: &types.CordinatesInput{
				Latitude:  float64(mockRequestModel.Address.Cordinates.Latitude),
				Longitude: float64(mockRequestModel.Address.Cordinates.Longitude),
			},
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &mockClaims)

				customerServiceClient.EXPECT().
					StaffCustomerDetailsUpdate(ctx, &mockRequestModel).
					Return(
						&pbTypes.DefaultResponse{
							Success: true,
							Code:    http.StatusOK,
						},
						nil,
					)

				resp, err := resolver.StaffUpdateCustomerDetails(ctx, mockRequest)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Success: true,
					Code:    http.StatusOK,
				}, resp)
			case errorUnauthenticated:
				ctx := context.Background()
				resp, err := resolver.StaffUpdateCustomerDetails(ctx, mockRequest)
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, &types.Response{
					Message: &authFailedMessage,
					Success: false,
					Code:    http.StatusInternalServerError,
				}, resp)
			case errorCallingRPC:
				ctx, _ := middleware.PutClaimsOnContext(context.Background(), &mockClaims)

				customerServiceClient.EXPECT().
					StaffCustomerDetailsUpdate(ctx, &mockRequestModel).
					Return(nil, errors.New("mock error"))

				resp, err := resolver.StaffUpdateCustomerDetails(ctx, mockRequest)
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Contains(t, err.Error(), "mock error")
			}
		})
	}
}

func Test_mutationResolver_CloseAccount(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx               context.Context
		accountCloseInput types.AccountCloseInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Response
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.CloseAccount(tt.args.ctx, tt.args.accountCloseInput)
			if !tt.wantErr(t, err, fmt.Sprintf("CloseAccount(%v, %v)", tt.args.ctx, tt.args.accountCloseInput)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CloseAccount(%v, %v)", tt.args.ctx, tt.args.accountCloseInput)
		})
	}
}
