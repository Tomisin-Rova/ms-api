package graph

import (
	"errors"
	"ms.api/protos/pb/paymentService"
	"testing"

	"ms.api/mocks"
	identitySvc "ms.api/protos/pb/identityService"
	"ms.api/protos/pb/onboardingService"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"
	"ms.api/types"

	coreErrors "github.com/roava/zebra/errors"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"golang.org/x/net/context"
)

var (
	personId     = "p123456"
	identityId   = "i123456"
	deviceId     = "d123456"
	validUserCtx = context.WithValue(
		context.Background(), middlewares.AuthenticatedUserContextKey, models.Claims{
			PersonId:   personId,
			IdentityId: identityId,
			DeviceId:   deviceId,
		})
)

func TestMutationResolver_CreatePhone(t *testing.T) {
	const (
		success = iota
		errorInvalidPhone
		errorOnboardingSvcCreatePhone
	)

	var tests = []struct {
		name string
		args struct {
			phone  string
			device types.DeviceInput
		}
		testType int
	}{
		{
			name: "Test create phone successfully",
			args: struct {
				phone  string
				device types.DeviceInput
			}{
				phone: "5522552255",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: success,
		},
		{
			name: "Test invalid phone",
			args: struct {
				phone  string
				device types.DeviceInput
			}{
				phone: "invalidPhone",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: errorInvalidPhone,
		},
		{
			name: "Test error calling OnboardingService.CreatePhone()",
			args: struct {
				phone  string
				device types.DeviceInput
			}{
				phone: "5522552255",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: errorOnboardingSvcCreatePhone,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			onBoardingServiceClient := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onBoardingServiceClient,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case errorInvalidPhone:
				response, err := mutationResolver.CreatePhone(context.Background(), testCase.args.phone, testCase.args.device)

				assert.Error(t, err)
				assert.Equal(t, 7010, err.(*coreErrors.Terror).Code())
				assert.Nil(t, response)
			case errorOnboardingSvcCreatePhone:
				onBoardingServiceClient.On("CreatePhone", mock.Anything, &onboardingService.CreatePhoneRequest{
					PhoneNumber: testCase.args.phone,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(nil, errors.New(""))

				response, err := mutationResolver.CreatePhone(context.Background(), testCase.args.phone, testCase.args.device)

				assert.Error(t, err)
				assert.Nil(t, response)
			case success:
				onBoardingServiceClient.On("CreatePhone", mock.Anything, &onboardingService.CreatePhoneRequest{
					PhoneNumber: testCase.args.phone,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(&onboardingService.CreatePhoneResponse{}, nil)

				response, err := mutationResolver.CreatePhone(context.Background(), testCase.args.phone, testCase.args.device)

				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, true, response.Success)
			}

			onBoardingServiceClient.AssertExpectations(t)
		})
	}
}

func TestMutationResolver_ConfirmPhone(t *testing.T) {
	const (
		success = iota
		errorOnboardingSvcVerifySmsOtp
	)

	var tests = []struct {
		name string
		args struct {
			token string
			code  string
		}
		testType int
	}{
		{
			name: "Test confirm phone correctly",
			args: struct {
				token string
				code  string
			}{
				token: "testToken",
				code:  "123456",
			},
			testType: success,
		},
		{
			name: "Test error calling OnboardingService..VerifySmsOtp()",
			args: struct {
				token string
				code  string
			}{
				token: "testToken",
				code:  "123456",
			},
			testType: errorOnboardingSvcVerifySmsOtp,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			onBoardingServiceClient := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onBoardingServiceClient,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case errorOnboardingSvcVerifySmsOtp:
				onBoardingServiceClient.On("VerifySmsOtp", mock.Anything, &onboardingService.OtpVerificationRequest{
					Token: testCase.args.token, Code: testCase.args.code,
				}).Return(nil, errors.New(""))

				response, err := mutationResolver.ConfirmPhone(context.Background(), testCase.args.token, testCase.args.code)

				assert.Error(t, err)
				assert.Nil(t, response)
			case success:
				onBoardingServiceClient.On("VerifySmsOtp", mock.Anything, &onboardingService.OtpVerificationRequest{
					Token: testCase.args.token, Code: testCase.args.code,
				}).Return(&onboardingService.OtpVerificationResponse{}, nil)

				response, err := mutationResolver.ConfirmPhone(context.Background(), testCase.args.token, testCase.args.code)

				assert.NoError(t, err)
				assert.NotNil(t, response)
			}

			onBoardingServiceClient.AssertExpectations(t)
		})
	}
}

func TestMutationResolver_Signup(t *testing.T) {
	const (
		success = iota
		errorInvalidEmail
		errorOnboardingSvcCreatePerson
	)

	var tests = []struct {
		name string
		args struct {
			token    string
			email    string
			passcode string
		}
		testType int
	}{
		{
			name: "Test sign up successfully",
			args: struct {
				token    string
				email    string
				passcode string
			}{
				token:    "123456",
				email:    "test@email.com",
				passcode: "123456",
			},
			testType: success,
		},
		{
			name: "Test error invalid email",
			args: struct {
				token    string
				email    string
				passcode string
			}{
				token:    "123456",
				email:    "invalidEmail",
				passcode: "123456",
			},
			testType: errorInvalidEmail,
		},
		{
			name: "Test error calling OnboardingService.CreatePerson()",
			args: struct {
				token    string
				email    string
				passcode string
			}{
				token:    "123456",
				email:    "test@email.com",
				passcode: "123456",
			},
			testType: errorOnboardingSvcCreatePerson,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			onboardingServiceClient := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onboardingServiceClient,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case success:
				onboardingServiceClient.On("CreatePerson", context.Background(), &onboardingService.CreatePersonRequest{
					Email:    testCase.args.email,
					Passcode: testCase.args.passcode,
					Token:    testCase.args.token,
				}).Return(&onboardingService.CreatePersonResponse{}, nil)
				response, err := mutationResolver.Signup(context.Background(), testCase.args.token, testCase.args.email,
					testCase.args.passcode)

				assert.NoError(t, err)
				assert.NotNil(t, response)
			case errorInvalidEmail:
				response, err := mutationResolver.Signup(context.Background(), testCase.args.token, testCase.args.email,
					testCase.args.passcode)

				assert.Error(t, err)
				assert.Equal(t, 1100, err.(*coreErrors.Terror).Code())
				assert.Nil(t, response)
			case errorOnboardingSvcCreatePerson:
				onboardingServiceClient.On("CreatePerson", context.Background(), &onboardingService.CreatePersonRequest{
					Email:    testCase.args.email,
					Passcode: testCase.args.passcode,
					Token:    testCase.args.token,
				}).Return(nil, errors.New(""))
				response, err := mutationResolver.Signup(context.Background(), testCase.args.token, testCase.args.email,
					testCase.args.passcode)

				assert.Error(t, err)
				assert.Nil(t, response)
			}

			onboardingServiceClient.AssertExpectations(t)
		})
	}
}

func TestMutationResolver_UpdateDeviceToken(t *testing.T) {
	const (
		success = iota
		errorInvalidUser
		errorIdentitySvcUpdateDeviceTokens
	)

	var tests = []struct {
		name                 string
		args                 []*types.DeviceTokenInput
		expectedDeviceTokens []*identitySvc.DeviceTokens
		testType             int
	}{
		{
			name: "Test update device successfully",
			args: []*types.DeviceTokenInput{
				{
					Type:  types.DeviceTokenTypeFirebase,
					Value: "123456",
				},
				{
					Type:  types.DeviceTokenTypeBiometric,
					Value: "123456",
				},
			},
			expectedDeviceTokens: []*identitySvc.DeviceTokens{
				{
					Type:  string(types.DeviceTokenTypeFirebase),
					Value: "123456",
				},
				{
					Type:  string(types.DeviceTokenTypeBiometric),
					Value: "123456",
				},
			},
			testType: success,
		},
		{
			name: "Test error invalid user provided on jwtToken",
			args: []*types.DeviceTokenInput{
				{
					Type:  types.DeviceTokenTypeFirebase,
					Value: "123456",
				},
				{
					Type:  types.DeviceTokenTypeBiometric,
					Value: "123456",
				},
			},
			expectedDeviceTokens: []*identitySvc.DeviceTokens{
				{
					Type:  string(types.DeviceTokenTypeFirebase),
					Value: "123456",
				},
				{
					Type:  string(types.DeviceTokenTypeBiometric),
					Value: "123456",
				},
			},
			testType: errorInvalidUser,
		},
		{
			name: "Test error calling identityService.UpdateDeviceTokens()",
			args: []*types.DeviceTokenInput{
				{
					Type:  types.DeviceTokenTypeFirebase,
					Value: "123456",
				},
				{
					Type:  types.DeviceTokenTypeBiometric,
					Value: "123456",
				},
			},
			expectedDeviceTokens: []*identitySvc.DeviceTokens{
				{
					Type:  string(types.DeviceTokenTypeFirebase),
					Value: "123456",
				},
				{
					Type:  string(types.DeviceTokenTypeBiometric),
					Value: "123456",
				},
			},
			testType: errorIdentitySvcUpdateDeviceTokens,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			identityService := new(mocks.IdentityServiceClient)

			resolver := NewResolver(&ResolverOpts{
				identityService: identityService,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case success:
				identityService.On("UpdateDeviceTokens", validUserCtx, &identitySvc.UpdateDeviceTokensRequest{
					DeviceId:   deviceId,
					IdentityId: identityId,
					Tokens:     testCase.expectedDeviceTokens,
				}).Return(&protoTypes.Response{
					Success: true,
				}, nil)

				response, err := mutationResolver.UpdateDeviceToken(validUserCtx, testCase.args)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, &types.Response{
					Message: "successful",
					Success: true,
				}, response)
			case errorInvalidUser:
				response, err := mutationResolver.UpdateDeviceToken(context.Background(), testCase.args)
				assert.Error(t, err)
				assert.Nil(t, response)
				assert.Equal(t, 7012, err.(*coreErrors.Terror).Code())
			case errorIdentitySvcUpdateDeviceTokens:
				identityService.On("UpdateDeviceTokens", validUserCtx, &identitySvc.UpdateDeviceTokensRequest{
					DeviceId:   deviceId,
					IdentityId: identityId,
					Tokens:     testCase.expectedDeviceTokens,
				}).Return(nil, errors.New(""))

				response, err := mutationResolver.UpdateDeviceToken(validUserCtx, testCase.args)
				assert.Error(t, err)
				assert.Nil(t, response)
			}

			identityService.AssertExpectations(t)
		})
	}
}

func TestMutationResolver_SubmitApplication(t *testing.T) {
	const (
		success = iota
		errorInvalidUser
		errorOnboardingSvcSubmitApplication
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test submit application successfully",
			testType: success,
		},
		{
			name:     "Test error invalid user context",
			testType: errorInvalidUser,
		},
		{
			name:     "Test error calling onBoardingService.SubmitApplication",
			testType: errorOnboardingSvcSubmitApplication,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			onboardingServiceMock := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onboardingServiceMock,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case success:
				onboardingServiceMock.On("SubmitApplication", validUserCtx, &onboardingService.SubmitApplicationRequest{
					PersonId: personId,
				}).Return(&protoTypes.Response{
					Message: "success",
					Success: true,
				}, nil)

				response, err := mutationResolver.SubmitApplication(validUserCtx)

				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response)
			case errorInvalidUser:
				response, err := mutationResolver.SubmitApplication(context.Background())

				assert.Error(t, err)
				assert.Equal(t, 7012, err.(*coreErrors.Terror).Code())
				assert.Nil(t, response)
				assert.Empty(t, response)
			case errorOnboardingSvcSubmitApplication:
				onboardingServiceMock.On("SubmitApplication", validUserCtx, &onboardingService.SubmitApplicationRequest{
					PersonId: personId,
				}).Return(nil, errors.New(""))

				response, err := mutationResolver.SubmitApplication(validUserCtx)

				assert.Error(t, err)
				assert.Nil(t, response)
				assert.Empty(t, response)
			}

			onboardingServiceMock.AssertExpectations(t)
		})
	}
}

func TestMutationResolver_CreatePayee(t *testing.T) {
	const (
		success = iota
		errorInvalidUser
		errorOnboardingSvcSubmitApplication
	)
	accountName, accountNumber, passcode := "accountName", "1023413534", "passcode"
	payeeInput := types.PayeeInput{
		Name: "test name",
		Accounts: []*types.PayeeAccountInput{{
			Name:          &accountName,
			AccountNumber: &accountNumber,
		}},
	}
	mockReq := &paymentService.CreatePayeeRequest{
		IdentityId:     identityId,
		TransactionPin: passcode,
		Name:           payeeInput.Name,
		AccountName:    *payeeInput.Accounts[0].Name,
		AccountNumber:  *payeeInput.Accounts[0].AccountNumber,
	}
	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test submit application successfully",
			testType: success,
		},
		{
			name:     "Test error invalid user context",
			testType: errorInvalidUser,
		},
		{
			name:     "Test error calling onBoardingService.SubmitApplication",
			testType: errorOnboardingSvcSubmitApplication,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			paymentServiceMock := new(mocks.PaymentServiceClient)

			resolver := NewResolver(&ResolverOpts{
				paymentService: paymentServiceMock,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case success:
				paymentServiceMock.On("CreatePayee", validUserCtx, mockReq).Return(&protoTypes.Response{
					Message: "success",
					Success: true,
				}, nil)

				response, err := mutationResolver.CreatePayee(validUserCtx, payeeInput, passcode)

				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response)
			case errorInvalidUser:
				response, err := mutationResolver.CreatePayee(context.Background(), payeeInput, passcode)

				assert.Error(t, err)
				assert.Equal(t, 7012, err.(*coreErrors.Terror).Code())
				assert.Nil(t, response)
				assert.Empty(t, response)
			case errorOnboardingSvcSubmitApplication:
				paymentServiceMock.On("CreatePayee", validUserCtx, mockReq).Return(nil, errors.New(""))
				response, err := mutationResolver.CreatePayee(validUserCtx, payeeInput, passcode)

				assert.Error(t, err)
				assert.Nil(t, response)
				assert.Empty(t, response)
			}

			paymentServiceMock.AssertExpectations(t)
		})
	}
}
