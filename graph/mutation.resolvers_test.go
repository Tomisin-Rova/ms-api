package graph

import (
	"errors"
	"testing"

	"ms.api/mocks"
	"ms.api/protos/pb/accountService"
	"ms.api/protos/pb/authService"
	identitySvc "ms.api/protos/pb/identityService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/paymentService"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/protos/pb/verifyService"
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

const (
	emailFound = iota
	gbpAccountFound
	emailNotFound
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
					Tokens: []*types.DeviceTokenInput{{
						Type:  "firebase",
						Value: "AHRFRR",
					}},
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
					Tokens: []*types.DeviceTokenInput{{
						Type:  "firebase",
						Value: "AHRFRR",
					}},
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
					Tokens: []*types.DeviceTokenInput{{
						Type:  "firebase",
						Value: "AHRFRR",
					}},
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
				tokens := make([]*protoTypes.DeviceToken, len(testCase.args.device.Tokens))

				for k, v := range testCase.args.device.Tokens {
					tokens[k] = &protoTypes.DeviceToken{
						Type:  string(v.Type),
						Value: v.Value,
					}
				}
				onBoardingServiceClient.On("CreatePhone", mock.Anything, &onboardingService.CreatePhoneRequest{
					PhoneNumber: testCase.args.phone,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
						Tokens:     tokens,
					},
				}).Return(nil, errors.New(""))

				response, err := mutationResolver.CreatePhone(context.Background(), testCase.args.phone, testCase.args.device)

				assert.Error(t, err)
				assert.Nil(t, response)
			case success:
				tokens := make([]*protoTypes.DeviceToken, len(testCase.args.device.Tokens))

				for k, v := range testCase.args.device.Tokens {
					tokens[k] = &protoTypes.DeviceToken{
						Type:  string(v.Type),
						Value: v.Value,
					}
				}
				onBoardingServiceClient.On("CreatePhone", mock.Anything, &onboardingService.CreatePhoneRequest{
					PhoneNumber: testCase.args.phone,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
						Tokens:     tokens,
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
		errorEmptyCurrency
		errorOnboardingSvcSubmitApplication
	)
	accountName, accountNumber, passcode, currency := "accountName", "12345678", "passcode", "GBP"
	payeeInput := types.PayeeInput{
		Name: "test name",
		Accounts: []*types.PayeeAccountInput{{
			Name:          &accountName,
			AccountNumber: &accountNumber,
			Currency:      &currency,
		}},
	}
	mockReq := &paymentService.CreatePayeeRequest{
		IdentityId:     identityId,
		TransactionPin: passcode,
		Name:           payeeInput.Name,
		AccountName:    *payeeInput.Accounts[0].Name,
		AccountNumber:  *payeeInput.Accounts[0].AccountNumber,
		Currency:       currency,
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
			name:     "Test error empty currency field",
			testType: errorEmptyCurrency,
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
			case errorEmptyCurrency:
				payeeInput.Accounts[0].Currency = nil
				response, err := mutationResolver.CreatePayee(context.Background(), payeeInput, passcode)

				assert.Error(t, err)
				assert.Equal(t, 7012, err.(*coreErrors.Terror).Code())
				assert.Nil(t, response)
				assert.Empty(t, response)
			case errorOnboardingSvcSubmitApplication:
				payeeInput.Accounts[0].Currency = &currency
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

func TestMutationResolver_CreateAccount(t *testing.T) {
	const (
		success = iota
		errorNotAuthenticatedUser
		errorCreatingAccount
	)

	var tests = []struct {
		name     string
		arg      types.ProductInput
		testType int
	}{
		{
			name: "Test create account successfully",
			arg: types.ProductInput{
				ID: "productId",
			},
			testType: success,
		},
		{
			name: "Test error not authenticated user",
			arg: types.ProductInput{
				ID: "productId",
			},
			testType: errorNotAuthenticatedUser,
		},
		{
			name: "Test error creating account",
			arg: types.ProductInput{
				ID: "productId",
			},
			testType: errorCreatingAccount,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			accountServiceClient := new(mocks.AccountServiceClient)

			resolver := NewResolver(&ResolverOpts{
				accountService: accountServiceClient,
			}, zaptest.NewLogger(t))
			ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey,
				models.Claims{
					PersonId:   "personId",
					IdentityId: "identityId",
					DeviceId:   "deviceId",
				})
			switch testCase.testType {
			case success:
				accountServiceClient.On("CreateAccount", ctx, &accountService.CreateAccountRequest{
					IdentityId: "identityId",
					Product: &protoTypes.ProductInput{
						Id: testCase.arg.ID,
					},
				}).Return(&protoTypes.Response{Success: true}, nil)

				response, err := resolver.Mutation().CreateAccount(ctx, testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response)
			case errorNotAuthenticatedUser:
				response, err := resolver.Mutation().CreateAccount(context.Background(), testCase.arg)
				assert.Error(t, err)
				assert.Nil(t, response)
			case errorCreatingAccount:
				accountServiceClient.On("CreateAccount", ctx, &accountService.CreateAccountRequest{
					IdentityId: "identityId",
					Product: &protoTypes.ProductInput{
						Id: testCase.arg.ID,
					},
				}).Return(nil, errors.New(""))

				response, err := resolver.Mutation().CreateAccount(ctx, testCase.arg)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_ValidateBvn(t *testing.T) {
	const (
		success = iota
		errorValidatingBVN
	)

	mockReq := &accountService.ValidateBVNRequest{
		PersonId: personId,
		Bvn:      "123412412342",
		Phone:    "01204201242",
	}

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test validate bvn successfully",
			testType: success,
		},
		{
			name:     "Test account service error",
			testType: errorValidatingBVN,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			accountServiceClient := new(mocks.AccountServiceClient)
			resolverOpts := &ResolverOpts{accountService: accountServiceClient}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t))
			ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey,
				models.Claims{PersonId: personId})

			switch testCase.testType {
			case success:
				accountServiceClient.On("ValidateBVN", ctx, mockReq).
					Return(&protoTypes.Response{Success: true}, nil)

				response, err := resolver.Mutation().ValidateBvn(ctx, mockReq.Bvn, mockReq.Phone)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response)
			case errorValidatingBVN:
				accountServiceClient.On("ValidateBVN", ctx, mockReq).
					Return(nil, errors.New(""))

				response, err := resolver.Mutation().ValidateBvn(ctx, mockReq.Bvn, mockReq.Phone)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_ResubmitReports(t *testing.T) {
	const (
		success = iota
		errorUnauthenticatedUser
		errorResubmittingReports
	)

	var tests = []struct {
		name     string
		arg      []*types.ReportInput
		testType int
	}{
		{
			name: "Test resubmit reports successfully",
			arg: []*types.ReportInput{
				{
					ID: "123",
				},
			},
			testType: success,
		},
		{
			name: "Test error unauthenticated user",
			arg: []*types.ReportInput{
				{
					ID: "123",
				},
			},
			testType: errorUnauthenticatedUser,
		},
		{
			name: "Test error submitting reports",
			arg: []*types.ReportInput{
				{
					ID: "123",
				},
			},
			testType: errorResubmittingReports,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			onboardingServiceClient := new(mocks.OnBoardingServiceClient)
			resolverOpts := &ResolverOpts{OnBoardingService: onboardingServiceClient}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t))

			switch testCase.testType {
			case success:
				request := onboardingService.ResubmitReportRequest{
					PersonId: personId,
				}
				for _, report := range testCase.arg {
					request.Reports = append(request.Reports, &onboardingService.ReportInput{
						Id: report.ID,
					})
				}
				onboardingServiceClient.On("ResubmitReport", validUserCtx, &request).Return(&protoTypes.Response{
					Success: true,
				}, nil)

				response, err := resolver.Mutation().ResubmitReports(validUserCtx, testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, true, response.Success)
			case errorUnauthenticatedUser:
				response, err := resolver.Mutation().ResubmitReports(context.Background(), testCase.arg)
				assert.Error(t, err)
				assert.IsType(t, &coreErrors.Terror{}, err)
				assert.Equal(t, 7012, err.(*coreErrors.Terror).Code())
				assert.Nil(t, response)
			case errorResubmittingReports:
				request := onboardingService.ResubmitReportRequest{
					PersonId: personId,
				}
				for _, report := range testCase.arg {
					request.Reports = append(request.Reports, &onboardingService.ReportInput{
						Id: report.ID,
					})
				}
				onboardingServiceClient.On("ResubmitReport", validUserCtx, &request).Return(nil, errors.New(""))

				response, err := resolver.Mutation().ResubmitReports(validUserCtx, testCase.arg)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_CreatePayment(t *testing.T) {
	const (
		success = iota
	)

	var tests = []struct {
		name     string
		arg      *paymentService.CreatePaymentRequest
		testType int
	}{
		{
			name: "Test successful funding of Vault account",
			arg: &paymentService.CreatePaymentRequest{
				IdentityId:     "identityId",
				TransactionPin: "123456789",
				IdempotencyKey: "1234567",
				Owner:          "Princewill Chiaka",
				Charge:         50,
				Reference:      "Test payment to Vault account",
				Beneficiary: &paymentService.BeneficiaryInput{
					Account:  "1234567",
					Currency: "GBP",
					Amount:   100000,
				},
				FundingSource: "1234567",
				Currency:      "GBP",
				FundingAmount: 100000,
			},
			testType: success,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			serviceClient := new(mocks.PaymentServiceClient)
			resolverOpts := &ResolverOpts{paymentService: serviceClient}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t))

			switch testCase.testType {
			case success:
				serviceClient.On("CreatePayment", validUserCtx, testCase.arg).Return(&protoTypes.Response{
					Success: true,
				}, nil)

				response, err := resolver.paymentService.CreatePayment(validUserCtx, testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, true, response.Success)
			}
		})
	}
}

func TestMutationResolver_RequestOtp(t *testing.T) {
	const (
		success = iota
		errorRequestingOTP
	)

	var expireTime int64 = 20

	var tests = []struct {
		name string
		arg  struct {
			typeArg    types.DeliveryMode
			target     string
			expireTime *int64
		}
		testType int
	}{
		{
			name: "Test request OTP successfully",
			arg: struct {
				typeArg    types.DeliveryMode
				target     string
				expireTime *int64
			}{typeArg: types.DeliveryModeEmail, target: "test@email.com", expireTime: &expireTime},
			testType: success,
		},
		{
			name: "Test error calling RequestOTP",
			arg: struct {
				typeArg    types.DeliveryMode
				target     string
				expireTime *int64
			}{typeArg: types.DeliveryModeEmail, target: "test@email.com", expireTime: nil},
			testType: errorRequestingOTP,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			verifyServiceMock := new(mocks.VerifyServiceClient)
			resolverOpts := &ResolverOpts{verifyService: verifyServiceMock}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t))

			switch testCase.testType {
			case success:
				verifyServiceMock.On("RequestOTP", context.Background(), &verifyService.RequestOTPRequest{
					DeliveryMode: testCase.arg.typeArg.String(),
					Target:       testCase.arg.target,
					ExpireTime:   expireTime,
				}).Return(&protoTypes.Response{}, nil)

				response, err := resolver.Mutation().RequestOtp(context.Background(), testCase.arg.typeArg, testCase.arg.target,
					testCase.arg.expireTime)
				assert.NoError(t, err)
				assert.NotNil(t, response)
			case errorRequestingOTP:
				verifyServiceMock.On("RequestOTP", context.Background(), &verifyService.RequestOTPRequest{
					DeliveryMode: testCase.arg.typeArg.String(),
					Target:       testCase.arg.target,
					ExpireTime:   0,
				}).Return(nil, errors.New(""))

				response, err := resolver.Mutation().RequestOtp(context.Background(), testCase.arg.typeArg, testCase.arg.target,
					testCase.arg.expireTime)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_VerifyOtp(t *testing.T) {
	const (
		success = iota
		errorVerifyingOTP
	)

	var tests = []struct {
		name string
		arg  struct {
			target string
			token  string
		}
		testType int
	}{
		{
			name: "Test verify OTP successfully",
			arg: struct {
				target string
				token  string
			}{target: "target", token: "12345"},
			testType: success,
		},
		{
			name: "Test error calling VerifyOTP",
			arg: struct {
				target string
				token  string
			}{target: "target", token: "12345"},
			testType: errorVerifyingOTP,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			verifyServiceMock := new(mocks.VerifyServiceClient)
			resolverOpts := &ResolverOpts{verifyService: verifyServiceMock}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t))

			switch testCase.testType {
			case success:
				verifyServiceMock.On("VerifyOTP", context.Background(), &verifyService.VerifyOTPRequest{
					Target: testCase.arg.target,
					Token:  testCase.arg.token,
				}).Return(&protoTypes.Response{}, nil)

				response, err := resolver.Mutation().VerifyOtp(context.Background(), testCase.arg.target, testCase.arg.token)
				assert.NoError(t, err)
				assert.NotNil(t, response)
			case errorVerifyingOTP:
				verifyServiceMock.On("VerifyOTP", context.Background(), &verifyService.VerifyOTPRequest{
					Target: testCase.arg.target,
					Token:  testCase.arg.token,
				}).Return(nil, errors.New(""))

				response, err := resolver.Mutation().VerifyOtp(context.Background(), testCase.arg.target, testCase.arg.token)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestMutationResolver_ValidateEmail(t *testing.T) {
	const (
		success = iota
		errorValidateEmail
		emailFoundOnly
	)

	var tests = []struct {
		name string
		args struct {
			email  string
			device types.DeviceInput
		}
		testType int
	}{
		{
			name: "Test validate email successfully",
			args: struct {
				email  string
				device types.DeviceInput
			}{
				email: "johnsmith@gmail.com",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: success,
		},
		{
			name: "Test error calling ValidateEmail",
			args: struct {
				email  string
				device types.DeviceInput
			}{
				email: "johnsmith@gmail.com",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: errorValidateEmail,
		},
		{
			name: "Test email found only",
			args: struct {
				email  string
				device types.DeviceInput
			}{
				email: "johnsmith@gmail.com",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: emailFoundOnly,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			authServiceMock := new(mocks.AuthServiceClient)
			accountServiceMock := new(mocks.AccountServiceClient)
			resolverOpts := &ResolverOpts{AuthService: authServiceMock, accountService: accountServiceMock}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t))

			switch testCase.testType {
			case success:
				authServiceMock.On("ValidateEmail", context.Background(), &authService.ValidateEmailRequest{
					Email: testCase.args.email,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(&protoTypes.Response{
					Success: true,
					Message: "successful",
					Code:    gbpAccountFound,
				}, nil)
				response, err := resolver.Mutation().ValidateEmail(context.Background(), testCase.args.email, testCase.args.device)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Message, "successful")
				assert.Equal(t, response.Success, true)
				assert.Equal(t, *response.Code, int64(gbpAccountFound))
			case errorValidateEmail:
				authServiceMock.On("ValidateEmail", context.Background(), &authService.ValidateEmailRequest{
					Email: testCase.args.email,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(&protoTypes.Response{
					Success: false,
					Code:    emailNotFound,
				}, nil)
				response, err := resolver.Mutation().ValidateEmail(context.Background(), testCase.args.email, testCase.args.device)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Success, false)
			case emailFoundOnly:
				authServiceMock.On("ValidateEmail", context.Background(), &authService.ValidateEmailRequest{
					Email: testCase.args.email,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(&protoTypes.Response{
					Success: true,
					Code:    emailFound,
				}, nil)
				response, err := resolver.Mutation().ValidateEmail(context.Background(), testCase.args.email, testCase.args.device)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Success, true)
				assert.Equal(t, *response.Code, int64(emailFound))
			}
		})
	}
}

func TestMutationResolver_ValidateUser(t *testing.T) {
	const (
		success = iota
		errorValidateEmail
	)

	var tests = []struct {
		name string
		args struct {
			email  string
			device types.DeviceInput
		}
		testType int
	}{
		{
			name: "Test validate email successfully",
			args: struct {
				email  string
				device types.DeviceInput
			}{
				email: "johnsmith@gmail.com",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: success,
		},
		{
			name: "Test error calling ValidateEmail",
			args: struct {
				email  string
				device types.DeviceInput
			}{
				email: "johnsmith@gmail.com",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: errorValidateEmail,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			authServiceMock := new(mocks.AuthServiceClient)
			accountServiceMock := new(mocks.AccountServiceClient)
			resolverOpts := &ResolverOpts{AuthService: authServiceMock, accountService: accountServiceMock}
			resolver := NewResolver(resolverOpts, zaptest.NewLogger(t))

			switch testCase.testType {
			case success:
				authServiceMock.On("ValidateUser", context.Background(), &authService.ValidateUserRequest{
					Email: testCase.args.email,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(&protoTypes.Response{
					Success: true,
					Message: "successful",
				}, nil)
				response, err := resolver.Mutation().ValidateUser(context.Background(), types.ValidateUserInput{
					Email: testCase.args.email,
					Device: &types.DeviceInput{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				})
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Message, "successful")
				assert.Equal(t, response.Success, true)
			case errorValidateEmail:
				authServiceMock.On("ValidateUser", context.Background(), &authService.ValidateUserRequest{
					Email: testCase.args.email,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(&protoTypes.Response{
					Success: false,
					Message: "",
				}, nil)
				response, err := resolver.Mutation().ValidateUser(context.Background(), types.ValidateUserInput{
					Email: testCase.args.email,
					Device: &types.DeviceInput{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				})
				assert.NoError(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

func TestRequestTransactionPasscodeReset(t *testing.T) {
	const (
		success = iota
		rpcError
	)

	tests := []struct {
		name     string
		email    string
		testType int
	}{
		{
			name:     "Test that RequestTransactionPasscodeReset call is successful.",
			email:    "princewill@example.com",
			testType: success,
		},
		{
			name:     "Test that RequestTransactionPasscodeReset call fails and is handled properly.",
			email:    "princewill@example.com",
			testType: rpcError,
		},
	}

	for _, testCase := range tests {
		identityService := new(mocks.IdentityServiceClient)
		resolver := NewResolver(&ResolverOpts{identityService: identityService}, zaptest.NewLogger(t))

		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				identityService.On("RequestResetTransactionPassword", validUserCtx, &identitySvc.RequestResetTransactionPasswordRequest{
					Email: testCase.email,
				}).Return(&protoTypes.Response{
					Message: "OTP Sent successfully.",
					Success: true,
				}, nil)

				response, err := resolver.Mutation().RequestTransactionPasscodeReset(validUserCtx, testCase.email)
				assert.NoError(t, err)
				assert.NotNil(t, response)

			case rpcError:
				identityService.On("RequestResetTransactionPassword", validUserCtx, &identitySvc.RequestResetTransactionPasswordRequest{
					Email: testCase.email,
				}).Return(nil, errors.New(""))

				response, err := resolver.Mutation().RequestTransactionPasscodeReset(validUserCtx, testCase.email)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestResetTransactionPasscode(t *testing.T) {
	const (
		success = iota
		rpcError
	)

	tests := []struct {
		name     string
		arg      *identitySvc.ResetTransactionPasswordRequest
		testType int
	}{
		{
			name: "Test that ResetTransactionPasscode call is successful.",
			arg: &identitySvc.ResetTransactionPasswordRequest{
				Email:           "princewill@example.com",
				Token:           "1234567",
				CurrentPasscode: "secret",
				NewPasscode:     "new-secret",
			},
			testType: success,
		},
	}

	for _, testCase := range tests {
		identityService := new(mocks.IdentityServiceClient)
		resolver := NewResolver(&ResolverOpts{identityService: identityService}, zaptest.NewLogger(t))

		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				identityService.On("ResetTransactionPassword", validUserCtx, testCase.arg).Return(&protoTypes.Response{
					Message: "Password reset successful.",
					Success: true,
				}, nil)

				response, err := resolver.Mutation().ResetTransactionPasscode(validUserCtx, testCase.arg.Email, testCase.arg.Token, testCase.arg.CurrentPasscode, testCase.arg.NewPasscode)
				assert.NoError(t, err)
				assert.NotNil(t, response)

			case rpcError:
				identityService.On("ResetTransactionPassword", validUserCtx, testCase.arg).Return(nil, errors.New(""))

				response, err := resolver.Mutation().ResetTransactionPasscode(validUserCtx, testCase.arg.Email, testCase.arg.Token, testCase.arg.CurrentPasscode, testCase.arg.NewPasscode)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}
