package graph

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"ms.api/mocks"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/onboarding"
	pbTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

var (
	state       = "lagos"
	city        = "lagos"
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
		Dob:       "18/05/1994",
		Address:   &mockAddress,
	}
)

func TestMutationResolver_RequestOtp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	verificationServiceClient := mocks.NewMockVerificationServiceClient(controller)
	resolverOpts := &ResolverOpts{
		VerificationService: verificationServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	expire := int64(3600)

	resp, err := resolver.RequestOtp(context.Background(), types.DeliveryMode(""), "", &expire)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_VerifyOtp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	verificationServiceClient := mocks.NewMockVerificationServiceClient(controller)
	resolverOpts := &ResolverOpts{
		VerificationService: verificationServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.VerifyOtp(context.Background(), "", "")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_Signup(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	authServiceClient := mocks.NewMockAuthServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	resp, err := resolver.Signup(context.Background(), types.CustomerInput{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_ResetLoginPassword(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.ResetLoginPassword(context.Background(), "", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CheckCustomerEmail(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	customerServiceClient := mocks.NewMockCustomerServiceClient(controller)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}

	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CheckCustomerEmail(context.Background(), "", types.DeviceInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
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
				ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.JWTClaims{Client: models.APP, ID: "123456", Email: "f@roava.app", DeviceID: "129594533fsdd"})

				customerServiceClient.EXPECT().Register(ctx, &customer.RegisterRequest{
					FirstName: "roava",
					LastName:  "app",
					Dob:       "18/05/1994",
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

				ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.JWTClaims{Client: models.APP, ID: "123456", Email: "f@roava.app", DeviceID: "129594533fsdd"})
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
				ctx: context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.JWTClaims{}),
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
				ctx: context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.JWTClaims{}),
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
				ctx: context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.JWTClaims{}),
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
				ctx: context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.JWTClaims{}),
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
				ctx: context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.JWTClaims{}),
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
						ReportTypes: []pbTypes.Reports_KYCTypes{pbTypes.Reports_DOCUMENT, pbTypes.Reports_FACIAL_VIDEO},
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
						ReportTypes: []pbTypes.Reports_KYCTypes{pbTypes.Reports_DOCUMENT, pbTypes.Reports_FACIAL_VIDEO},
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
						ReportTypes: []pbTypes.Reports_KYCTypes{pbTypes.Reports_DOCUMENT, pbTypes.Reports_FACIAL_VIDEO},
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
					Code:    http.StatusUnauthorized,
				}, resp)
			case errorSubmitCDD:
				onboardingServiceClient.EXPECT().SubmitCDD(testCase.arg.ctx, &onboarding.SubmitCDDRequest{
					Kyc: &onboarding.KYCInput{
						ReportTypes: []pbTypes.Reports_KYCTypes{pbTypes.Reports_DOCUMENT, pbTypes.Reports_FACIAL_VIDEO},
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
				ctx := context.WithValue(context.Background(),
					middlewares.AuthenticatedUserContextKey,
					models.JWTClaims{
						Client:   models.APP,
						ID:       "123456",
						Email:    "f@roava.app",
						DeviceID: "129594533fsdd"})

				customerServiceClient.EXPECT().AnswerQuestionary(ctx,
					&customer.AnswerQuestionaryRequest{
						Id: "questionaire_id",
						Answers: []*customer.AnswerInput{
							{
								Id:     "question_id",
								Answer: "My lifestyle",
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
							ID:     "question_id",
							Answer: "My lifestyle",
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
							ID:     "question_id",
							Answer: "My lifestyle",
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
	controller := gomock.NewController(t)
	defer controller.Finish()
	authServiceClient := mocks.NewMockAuthServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.Login(context.Background(), types.AuthInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_RefreshToken(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	authServiceClient := mocks.NewMockAuthServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.RefreshToken(context.Background(), "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
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
			ctx := context.WithValue(context.Background(),
				middlewares.AuthenticatedUserContextKey,
				models.JWTClaims{
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
			ctx := context.WithValue(context.Background(),
				middlewares.AuthenticatedUserContextKey,
				models.JWTClaims{
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
	controller := gomock.NewController(t)
	defer controller.Finish()
	onboardingServiceClient := mocks.NewMockOnboardingServiceClient(controller)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}

	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	message := ""

	resp, err := resolver.RequestResubmit(context.Background(), "", []string{}, &message)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_StaffLogin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	authServiceClient := mocks.NewMockAuthServiceClient(controller)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.StaffLogin(context.Background(), "", types.AuthTypeGoogle)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_UpdateKYCStatus(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	onboardingServiceClient := mocks.NewMockOnboardingServiceClient(controller)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.UpdateKYCStatus(context.Background(), "", types.KYCStatusesApproved, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_UpdateAMLStatus(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	onboardingServiceClient := mocks.NewMockOnboardingServiceClient(controller)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.UpdateAMLStatus(context.Background(), "", types.AMLStatusesPending, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
