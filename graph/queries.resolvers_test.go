package graph

import (
	"context"
	"errors"
	"testing"

	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ms.api/mocks"
	"ms.api/protos/pb/customer"
	pbTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

var (
	mockExpectedContents = &customer.GetContentsResponse{
		Nodes: []*pbTypes.Content{
			{
				Id:   "01xc2",
				Type: pbTypes.Content_PRIVACY_NOTICE,
				Link: "http://sample-1-link",
				Ts:   timestamppb.Now(),
			},
			{
				Id:   "21xc2",
				Type: pbTypes.Content_GENERAL_TC,
				Link: "http://sample-2-link",
				Ts:   timestamppb.Now(),
			},
			{
				Id:   "2rvc2",
				Type: pbTypes.Content_GENERAL_TC,
				Link: "http://sample-3-link",
				Ts:   timestamppb.Now(),
			},
		},

		PaginationInfo: &pbTypes.PaginationInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
			StartCursor:     "start_cursor",
			EndCursor:       "end_cursor",
		},

		TotalCount: 3,
	}
)

func Test_queryResolver_Content(t *testing.T) {
	const (
		success = iota
		contentNotFound
	)

	tests := []struct {
		name     string
		testType int
		arg      string
	}{
		{
			name:     "Test content found successfully with given contentId ",
			testType: success,
			arg:      "1",
		}, {
			name:     "Test error content",
			testType: contentNotFound,
			arg:      "wrongcontentId",
		},
	}

	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case success:
				mockExpectedContent := &pbTypes.Content{
					Id:   test.arg,
					Type: pbTypes.Content_GENERAL_TC,
					Link: "http://sample-link",
					Ts:   timestamppb.Now(),
				}
				serviceReq := &customer.GetContentRequest{Id: test.arg}
				customerServiceClient.On("GetContent",
					context.Background(),
					serviceReq,
				).Return(mockExpectedContent, nil)

				response, err := resolver.Content(context.Background(), test.arg)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, mockExpectedContent.Id, response.ID)

			case contentNotFound:
				serviceReq := &customer.GetContentRequest{Id: test.arg}
				customerServiceClient.On("GetContent",
					context.Background(),
					serviceReq,
				).Return(nil, errors.New("contentNotfound"))

				response, err := resolver.Content(context.Background(), test.arg)
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func Test_queryResolver_Contents(t *testing.T) {
	const (
		firstFiveContents = iota
		lastFiveContents
	)
	zero := int64(0)
	firstFive := int64(5)
	lastFive := int64(5)

	tests := []struct {
		name     string
		testType int
		arg      struct {
			first  *int64
			after  string
			last   *int64
			before string
		}
	}{
		{
			name:     "Test first 5 contents found successfully ",
			testType: firstFiveContents,
			arg: struct {
				first  *int64
				after  string
				last   *int64
				before string
			}{first: &firstFive, after: "", last: &zero, before: ""},
		}, {
			name:     "Test last 5 contents found",
			testType: lastFiveContents,
			arg: struct {
				first  *int64
				after  string
				last   *int64
				before string
			}{first: &zero, after: "", last: &lastFive, before: ""},
		},
	}

	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case firstFiveContents:
				serviceReq := &customer.GetContentsRequest{First: int32(*test.arg.first), After: test.arg.after, Last: int32(*test.arg.last), Before: test.arg.before}
				customerServiceClient.On("GetContents",
					context.Background(),
					serviceReq,
				).Return(mockExpectedContents, nil)

				response, err := resolver.Contents(context.Background(), test.arg.first, &test.arg.after, test.arg.last, &test.arg.before)
				assert.NoError(t, err)
				assert.NotNil(t, response)

			case lastFiveContents:
				serviceReq := &customer.GetContentsRequest{First: int32(*test.arg.first), After: test.arg.after, Last: int32(*test.arg.last), Before: test.arg.before}
				customerServiceClient.On("GetContents",
					context.Background(),
					serviceReq,
				).Return(mockExpectedContents, nil)

				response, err := resolver.Contents(context.Background(), test.arg.first, &test.arg.after, test.arg.last, &test.arg.before)
				assert.NoError(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

func Test_queryResolver_CheckEmail(t *testing.T) {
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

	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case success:
				customerServiceClient.On("CheckEmail",
					context.Background(),
					&customer.CheckEmailRequest{Email: test.arg},
				).Return(&pbTypes.DefaultResponse{Success: true}, nil)

				resp, err := resolver.CheckEmail(context.Background(), test.arg)
				assert.NoError(t, err)
				assert.Equal(t, resp, true)

			case emailNotFound:
				customerServiceClient.On("CheckEmail",
					context.Background(),
					customer.CheckEmailRequest{Email: test.arg},
				).Return(&pbTypes.DefaultResponse{Success: false}, errors.New("not found"))

				_, err := resolver.CheckEmail(context.Background(), test.arg)
				assert.NoError(t, err)

			case invalidEmail:
				customerServiceClient.On("CheckEmail",
					context.Background(),
					customer.CheckEmailRequest{Email: test.arg},
				).Return(&pbTypes.DefaultResponse{Success: false}, errors.New("invalid email"))

				resp, err := resolver.CheckEmail(context.Background(), test.arg)
				assert.Error(t, err)
				assert.Equal(t, resp, false)
			}
		})
	}
}

func Test_queryResolver_OnfidoSDKToken(t *testing.T) {
	onboardingServiceClient := new(mocks.OnboardingServiceClient)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	resp, err := resolver.OnfidoSDKToken(context.Background())

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Cdd(t *testing.T) {
	onboardingServiceClient := new(mocks.OnboardingServiceClient)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	resp, err := resolver.Cdd(context.Background(), types.CommonQueryFilterInput{})

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Me(t *testing.T) {
	const (
		me_staff_success = iota
		me_customer_success
		me_auth_err
	)

	var tests = []struct {
		name     string
		arg      models.JWTClaims
		testType int
	}{
		{
			name: "Test ME staff successful",
			arg: models.JWTClaims{
				Client:   models.DASHBOARD,
				ID:       "123456",
				Email:    "f@roava.app",
				DeviceID: "129594533fs434kd",
			},
			testType: me_staff_success,
		},

		{
			name: "Test ME customer successful",
			arg: models.JWTClaims{
				Client:   models.APP,
				ID:       "84773442",
				Email:    "sample@roava.app",
				DeviceID: "hfewuhdfff8424",
			},
			testType: me_customer_success,
		},

		{
			name:     "Test error ME authentication",
			arg:      models.JWTClaims{},
			testType: me_auth_err,
		},
	}

	for _, testCase := range tests {

		customerServiceClient := new(mocks.CustomerServiceClient)
		resolverOpts := &ResolverOpts{
			CustomerService: customerServiceClient,
		}

		resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

		switch testCase.testType {
		case me_staff_success:

			ctx := context.WithValue(context.Background(),
				middlewares.AuthenticatedUserContextKey, testCase.arg)

			customerServiceClient.On("Me", ctx, &customer.MeRequest{}).
				Return(&customer.MeResponse{
					Data: &customer.MeResponse_Staff{
						Staff: &pbTypes.Staff{
							Id:       "staffId",
							Name:     "staff name",
							LastName: "staff lastname",
							Dob:      "dd/mm/yyyy",
							Addresses: []*pbTypes.Address{
								{
									Primary: true,
									Country: &pbTypes.Country{
										Id:         "countryId",
										CodeAlpha2: "code_alpha_2",
										CodeAlpha3: "code_alpha_3",
										Name:       "country name",
									},
									State:    "state",
									City:     "city",
									Street:   "street",
									Postcode: "12345",
									Coordinates: &pbTypes.Coordinates{
										Latitude:  3.299434,
										Longitude: 1.443499,
									},
								},
							},
						},
					},
				}, nil)

			resp, err := resolver.Me(ctx)
			assert.NoError(t, err)
			assert.NotNil(t, resp)

		case me_customer_success:
			ctx := context.WithValue(context.Background(),
				middlewares.AuthenticatedUserContextKey, testCase.arg)

			customerServiceClient.On("Me", ctx, &customer.MeRequest{}).
				Return(&customer.MeResponse{
					Data: &customer.MeResponse_Customer{
						Customer: &pbTypes.Customer{
							Id:        "id",
							FirstName: "firstname",
							LastName:  "lastname",
							Dob:       "mm-dd-yyyt",
							Bvn:       "1200488434",
							Addresses: []*pbTypes.Address{
								{
									Primary: true,
									Country: &pbTypes.Country{
										Id:         "country_id",
										CodeAlpha2: "code_alpha_2",
										CodeAlpha3: "code_alpha_3",
										Name:       "country_name",
									},
									State:    "state",
									City:     "city",
									Street:   "street",
									Postcode: "3723",
									Coordinates: &pbTypes.Coordinates{
										Latitude:  3.97434,
										Longitude: 2.94873,
									},
								},
							},
							Phones: []*pbTypes.Phone{
								{
									Primary:  true,
									Number:   "234059999594",
									Verified: true,
								},
							},
							Email: &pbTypes.Email{
								Address:  "example@mail.com",
								Verified: true,
							},
							Status:   pbTypes.Customer_SIGNEDUP,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
					},
				}, nil)

			resp, err := resolver.Me(ctx)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		case me_auth_err:
			ctx := context.WithValue(context.Background(),
				middlewares.AuthenticatedUserContextKey, testCase.arg)

			customerServiceClient.On("Me", ctx, &customer.MeRequest{}).Return(&customer.MeResponse{}, errors.New("auth problem"))

			_, err := resolver.Me(ctx)
			assert.Error(t, err)
		}

	}
}

func Test_queryResolver_Product(t *testing.T) {
	accountServiceClient := new(mocks.AccountServiceClient)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	resp, err := resolver.Product(context.Background(), "")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Products(t *testing.T) {
	accountServiceClient := new(mocks.AccountServiceClient)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"

	resp, err := resolver.Products(context.Background(), &first, &after, &last, &before, []types.ProductStatuses{})

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Banks(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"

	resp, err := resolver.Banks(context.Background(), &first, &after, &last, &before)

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Account(t *testing.T) {
	accountServiceClient := new(mocks.AccountServiceClient)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	resp, err := resolver.Account(context.Background(), "")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Transactions(t *testing.T) {
	paymentServiceClient := new(mocks.PaymentServiceClient)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"

	resp, err := resolver.Transactions(context.Background(), &first, &after, &last, &before, []types.AccountStatuses{}, []string{}, []string{})
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Beneficiary(t *testing.T) {
	paymentServiceClient := new(mocks.PaymentServiceClient)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	resp, err := resolver.Beneficiary(context.Background(), "")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Beneficiaries(t *testing.T) {
	paymentServiceClient := new(mocks.PaymentServiceClient)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"
	keyword := "search_keyworkd"

	resp, err := resolver.Beneficiaries(context.Background(), &keyword, &first, &after, &last, &before, []types.BeneficiaryStatuses{})
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_TransactionTypes(t *testing.T) {
	paymentServiceClient := new(mocks.PaymentServiceClient)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"

	resp, err := resolver.TransactionTypes(context.Background(), &first, &after, &last, &before, []types.TransactionTypeStatuses{})
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Questionary(t *testing.T) {
	const (
		success = iota
		questionaryNotFound
	)
	tests := []struct {
		name     string
		arg      string
		testType int
	}{
		{
			name:     "Test questionary found successfully with a given questionaryId",
			arg:      "1",
			testType: success,
		},

		{
			name:     "Test error questionary not found with an invalidId",
			arg:      "invalidId",
			testType: questionaryNotFound,
		},
	}

	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case success:
				customerServiceClient.On("GetQuestionary",
					context.Background(),
					&customer.GetQuestionaryRequest{Id: test.arg},
				).Return(&pbTypes.Questionary{
					Id:   test.arg,
					Type: pbTypes.Questionary_REASONS,
					Questions: []*pbTypes.QuestionaryQuestion{
						{
							Id:    "questionId",
							Value: "Question text",
							PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
								{
									Id:    "predefinedAnswer Id 1",
									Value: "predefinedAnswer Value 1",
								},
								{
									Id:    "predefinedAnswer Id 2",
									Value: "predefinedAnswer Value 2",
								},
							},
							Required:        true,
							MultipleOptions: true,
						},
					},
					Status:   pbTypes.Questionary_ACTIVE,
					StatusTs: timestamppb.Now(),
					Ts:       timestamppb.Now(),
				}, nil)

				resp, err := resolver.Questionary(context.Background(), test.arg)
				assert.NoError(t, err)
				assert.NotNil(t, resp)

			case questionaryNotFound:
				customerServiceClient.On("GetQuestionary",
					context.Background(),
					&customer.GetQuestionaryRequest{Id: test.arg},
				).Return(&pbTypes.Questionary{}, errors.New("questionary not found"))

				resp, err := resolver.Questionary(context.Background(), test.arg)
				assert.Error(t, err)
				assert.Equal(t, resp, &types.Questionary{})
			}
		})
	}
}

func Test_queryResolver_Questionaries(t *testing.T) {
	const (
		first_ten_questionaries = iota
		last_ten_questionaries
		first_ten_active_questionaries
	)

	tests := []struct {
		name string
		args struct {
			keywords string
			first    int64
			after    string
			last     int64
			before   string
			statuses []types.QuestionaryStatuses
			types    []types.QuestionaryTypes
		}
		testType int
	}{
		{
			name: "Test first ten questionaries successfully",
			args: struct {
				keywords string
				first    int64
				after    string
				last     int64
				before   string
				statuses []types.QuestionaryStatuses
				types    []types.QuestionaryTypes
			}{
				keywords: "",
				first:    int64(10),
				after:    "",
				last:     0,
				before:   "",
				statuses: []types.QuestionaryStatuses{types.QuestionaryStatusesActive, types.QuestionaryStatusesInactive},
				types:    []types.QuestionaryTypes{types.QuestionaryTypesReasons},
			},
		},
		{
			name: "Test last ten questionaries successfully",
			args: struct {
				keywords string
				first    int64
				after    string
				last     int64
				before   string
				statuses []types.QuestionaryStatuses
				types    []types.QuestionaryTypes
			}{
				keywords: "",
				first:    0,
				after:    "",
				last:     int64(10),
				before:   "",
				statuses: []types.QuestionaryStatuses{types.QuestionaryStatusesActive, types.QuestionaryStatusesInactive},
				types:    []types.QuestionaryTypes{types.QuestionaryTypesReasons},
			},
		},

		{
			name: "Test first ten active questionaries successfully",
			args: struct {
				keywords string
				first    int64
				after    string
				last     int64
				before   string
				statuses []types.QuestionaryStatuses
				types    []types.QuestionaryTypes
			}{
				keywords: "",
				first:    int64(10),
				after:    "",
				last:     0,
				before:   "",
				statuses: []types.QuestionaryStatuses{types.QuestionaryStatusesActive},
				types:    []types.QuestionaryTypes{types.QuestionaryTypesReasons},
			},
		},
	}

	for _, test := range tests {
		customerServiceClient := new(mocks.CustomerServiceClient)
		resolverOpts := &ResolverOpts{
			CustomerService: customerServiceClient,
		}
		resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case first_ten_questionaries:
				helpers := helpersfactory{}
				// convert statuses to Questionary_QuestionaryStatuses
				statuses := make([]pbTypes.Questionary_QuestionaryStatuses, 0)
				if len(test.args.statuses) > 0 {
					for _, state := range test.args.statuses {
						statuses = append(statuses, pbTypes.Questionary_QuestionaryStatuses(helpers.GetQuestionaryStatusIndex(state)))
					}
				}

				// convert types to Questionary_QuestionaryTypes
				questionaryTypes := make([]pbTypes.Questionary_QuestionaryTypes, 0)
				if len(test.args.types) > 0 {
					for _, qstType := range test.args.types {
						questionaryTypes = append(questionaryTypes, pbTypes.Questionary_QuestionaryTypes(helpers.GetQuestionaryTypesIndex(qstType)))
					}
				}

				customerServiceClient.On("GetQuestionaries", context.Background(),
					&customer.GetQuestionariesRequest{
						Keywords: test.args.keywords,
						First:    int32(test.args.first),
						After:    test.args.after,
						Last:     int32(test.args.last),
						Before:   test.args.before,
						Statuses: statuses,
						Types:    questionaryTypes,
					}).Return(&customer.GetQuestionariesResponse{
					Nodes: []*pbTypes.Questionary{
						{
							Id:   "1",
							Type: pbTypes.Questionary_REASONS,
							Questions: []*pbTypes.QuestionaryQuestion{
								{
									Id:    "1",
									Value: "Do you have criminal record",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},

								{
									Id:    "2",
									Value: "Do you have an existing foreign account",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},
							},
							Status:   pbTypes.Questionary_ACTIVE,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
						{
							Id:   "2",
							Type: pbTypes.Questionary_REASONS,
							Questions: []*pbTypes.QuestionaryQuestion{
								{
									Id:    "1",
									Value: "Do you have medical record",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},

								{
									Id:    "2",
									Value: "Would you want to own foreign account",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},
							},
							Status:   pbTypes.Questionary_ACTIVE,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
					},

					PaginationInfo: &pbTypes.PaginationInfo{
						HasNextPage:     false,
						HasPreviousPage: false,
						EndCursor:       "end_cursor",
						StartCursor:     "start_cursor",
					},

					TotalCount: 2,
				}, nil)

				resp, err := resolver.Questionaries(context.Background(), &test.args.keywords, &test.args.first, &test.args.after, &test.args.last, &test.args.before, test.args.statuses, test.args.types)

				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, resp.TotalCount, int64(2))

			case last_ten_questionaries:
				// convert statuses to Questionary_QuestionaryStatuses
				statuses := make([]pbTypes.Questionary_QuestionaryStatuses, 0)
				if len(test.args.statuses) > 0 {
					for _, state := range test.args.statuses {
						statuses = append(statuses, pbTypes.Questionary_QuestionaryStatuses(pbTypes.Questionary_QuestionaryStatuses_value[string(state)]))
					}
				}

				// convert types to Questionary_QuestionaryTypes
				questionaryTypes := make([]pbTypes.Questionary_QuestionaryTypes, 0)
				if len(test.args.types) > 0 {
					for _, qstType := range test.args.types {
						questionaryTypes = append(questionaryTypes, pbTypes.Questionary_QuestionaryTypes(pbTypes.Questionary_QuestionaryTypes_value[string(qstType)]))
					}
				}

				customerServiceClient.On("GetQuestionaries", context.Background(),
					&customer.GetQuestionariesRequest{
						Keywords: test.args.keywords,
						First:    int32(test.args.first),
						After:    test.args.after,
						Last:     int32(test.args.last),
						Before:   test.args.before,
						Statuses: statuses,
						Types:    questionaryTypes,
					}).Return(&customer.GetQuestionariesResponse{
					Nodes: []*pbTypes.Questionary{
						{
							Id:   "2",
							Type: pbTypes.Questionary_REASONS,
							Questions: []*pbTypes.QuestionaryQuestion{
								{
									Id:    "1",
									Value: "Do you have medical record",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},

								{
									Id:    "2",
									Value: "Would you want to own foreign account",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},
							},
							Status:   pbTypes.Questionary_ACTIVE,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
						{
							Id:   "1",
							Type: pbTypes.Questionary_REASONS,
							Questions: []*pbTypes.QuestionaryQuestion{
								{
									Id:    "1",
									Value: "Do you have criminal record",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},

								{
									Id:    "2",
									Value: "Do you have an existing foreign account",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},
							},
							Status:   pbTypes.Questionary_ACTIVE,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
					},

					PaginationInfo: &pbTypes.PaginationInfo{
						HasNextPage:     false,
						HasPreviousPage: false,
						EndCursor:       "end_cursor",
						StartCursor:     "start_cursor",
					},

					TotalCount: 2,
				}, nil)

				resp, err := resolver.Questionaries(context.Background(), &test.args.keywords, &test.args.first, &test.args.after, &test.args.last, &test.args.before, test.args.statuses, test.args.types)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, resp.TotalCount, int64(2))

			case first_ten_active_questionaries:
				// convert statuses to Questionary_QuestionaryStatuses
				statuses := make([]pbTypes.Questionary_QuestionaryStatuses, 0)
				if len(test.args.statuses) > 0 {
					for _, state := range test.args.statuses {
						statuses = append(statuses, pbTypes.Questionary_QuestionaryStatuses(pbTypes.Questionary_QuestionaryStatuses_value[string(state)]))
					}
				}

				// convert types to Questionary_QuestionaryTypes
				questionaryTypes := make([]pbTypes.Questionary_QuestionaryTypes, 0)
				if len(test.args.types) > 0 {
					for _, qstType := range test.args.types {
						questionaryTypes = append(questionaryTypes, pbTypes.Questionary_QuestionaryTypes(pbTypes.Questionary_QuestionaryTypes_value[string(qstType)]))
					}
				}

				customerServiceClient.On("GetQuestionaries", context.Background(),
					&customer.GetQuestionariesRequest{
						Keywords: test.args.keywords,
						First:    int32(test.args.first),
						After:    test.args.after,
						Last:     int32(test.args.last),
						Before:   test.args.before,
						Statuses: statuses,
						Types:    questionaryTypes,
					}).Return(&customer.GetQuestionariesResponse{
					Nodes: []*pbTypes.Questionary{
						{
							Id:   "2",
							Type: pbTypes.Questionary_REASONS,
							Questions: []*pbTypes.QuestionaryQuestion{
								{
									Id:    "1",
									Value: "Do you have medical record",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},

								{
									Id:    "2",
									Value: "Would you want to own foreign account",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},
							},
							Status:   pbTypes.Questionary_ACTIVE,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
						{
							Id:   "1",
							Type: pbTypes.Questionary_REASONS,
							Questions: []*pbTypes.QuestionaryQuestion{
								{
									Id:    "1",
									Value: "Do you have criminal record",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},

								{
									Id:    "2",
									Value: "Do you have an existing foreign account",
									PredefinedAnswers: []*pbTypes.QuestionaryPredefinedAnswer{
										{
											Id:    "predefinedAnswer Id 1",
											Value: "predefinedAnswer Value 1",
										},
										{
											Id:    "predefinedAnswer Id 2",
											Value: "predefinedAnswer Value 2",
										},
									},
									Required:        true,
									MultipleOptions: true,
								},
							},
							Status:   pbTypes.Questionary_ACTIVE,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
					},

					PaginationInfo: &pbTypes.PaginationInfo{
						HasNextPage:     false,
						HasPreviousPage: false,
						EndCursor:       "end_cursor",
						StartCursor:     "start_cursor",
					},

					TotalCount: 2,
				}, nil)

				resp, err := resolver.Questionaries(context.Background(), &test.args.keywords, &test.args.first, &test.args.after, &test.args.last, &test.args.before, test.args.statuses, test.args.types)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, resp.TotalCount, int64(2))
			}

		})
	}
}

func Test_queryResolver_Currency(t *testing.T) {
	pricingServiceClient := new(mocks.PricingServiceClient)
	resolverOpts := &ResolverOpts{
		PricingService: pricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	resp, err := resolver.Currency(context.Background(), "")
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Currencies(t *testing.T) {
	pricingServiceClient := new(mocks.PricingServiceClient)
	resolverOpts := &ResolverOpts{
		PricingService: pricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"
	keywords := "keywords"

	resp, err := resolver.Currencies(context.Background(), &keywords, &first, &after, &last, &before)
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Fees(t *testing.T) {
	pricingServiceClient := new(mocks.PricingServiceClient)
	resolverOpts := &ResolverOpts{
		PricingService: pricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	resp, err := resolver.Fees(context.Background(), "")
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_ExchangeRate(t *testing.T) {
	pricingServiceClient := new(mocks.PricingServiceClient)
	resolverOpts := &ResolverOpts{
		PricingService: pricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	resp, err := resolver.ExchangeRate(context.Background(), "")
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Customer(t *testing.T) {

	const (
		success = iota
		customerNotFound
	)
	tests := []struct {
		name     string
		arg      string
		testType int
	}{
		{
			name:     "Test customer found successfully with a given customerId",
			arg:      "1",
			testType: success,
		},

		{
			name:     "Test error customer not found with an invalidId",
			arg:      "invalidId",
			testType: customerNotFound,
		},
	}

	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case success:
				customerServiceClient.On("GetCustomer",
					context.Background(),
					&customer.GetCustomerRequest{Id: test.arg},
				).Return(&pbTypes.Customer{
					Id:        test.arg,
					FirstName: "firstname",
					LastName:  "lastname",
					Dob:       "mm-dd-yyyt",
					Bvn:       "1200488434",
					Addresses: []*pbTypes.Address{
						{
							Primary: true,
							Country: &pbTypes.Country{
								Id:         "country_id",
								CodeAlpha2: "code_alpha_2",
								CodeAlpha3: "code_alpha_3",
								Name:       "country_name",
							},
						},
					},
					Phones: []*pbTypes.Phone{
						{
							Primary:  true,
							Number:   "234059999594",
							Verified: true,
						},
					},
					Email: &pbTypes.Email{
						Address:  "example@mail.com",
						Verified: true,
					},
					Status:   pbTypes.Customer_SIGNEDUP,
					StatusTs: timestamppb.Now(),
					Ts:       timestamppb.Now(),
				}, nil)

				resp, err := resolver.Customer(context.Background(), test.arg)

				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, resp.ID, test.arg)

			case customerNotFound:
				customerServiceClient.On("GetCustomer",
					context.Background(),
					&customer.GetCustomerRequest{Id: test.arg},
				).Return(&pbTypes.Customer{}, errors.New("customer not found"))

				resp, err := resolver.Customer(context.Background(), test.arg)
				assert.Error(t, err)
				assert.Equal(t, resp, &types.Customer{})
			}
		})
	}
}

func Test_queryResolver_Customers(t *testing.T) {
	const (
		first_ten_customers = iota
		last_ten_customers
	)

	tests := []struct {
		name string
		args struct {
			keywords string
			first    int64
			after    string
			last     int64
			before   string
			statuses []types.CustomerStatuses
		}
		testType int
	}{
		{
			name: "Test first ten questionaries successfully",
			args: struct {
				keywords string
				first    int64
				after    string
				last     int64
				before   string
				statuses []types.CustomerStatuses
			}{
				keywords: "",
				first:    int64(10),
				after:    "",
				last:     0,
				before:   "",
				statuses: []types.CustomerStatuses{
					types.CustomerStatusesSignedup,
					types.CustomerStatusesOnboarded,
					types.CustomerStatusesVerified,
					types.CustomerStatusesExited,
					types.CustomerStatusesRejected,
				},
			},
		},
		{
			name: "Test last ten questionaries successfully",
			args: struct {
				keywords string
				first    int64
				after    string
				last     int64
				before   string
				statuses []types.CustomerStatuses
			}{
				keywords: "",
				first:    0,
				after:    "",
				last:     int64(10),
				before:   "",
				statuses: []types.CustomerStatuses{
					types.CustomerStatusesSignedup,
					types.CustomerStatusesOnboarded,
					types.CustomerStatusesVerified,
					types.CustomerStatusesExited,
					types.CustomerStatusesRejected,
				},
			},
		},
	}

	for _, test := range tests {
		customerServiceClient := new(mocks.CustomerServiceClient)
		resolverOpts := &ResolverOpts{
			CustomerService: customerServiceClient,
		}
		resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
		helpers := helpersfactory{}

		t.Run(test.name, func(t *testing.T) {
			switch test.testType {
			case first_ten_customers:

				// convert statuses to Customer_CustomerStatuses
				statuses := make([]pbTypes.Customer_CustomerStatuses, 0)
				if len(test.args.statuses) > 0 {
					for _, state := range test.args.statuses {
						statuses = append(statuses, pbTypes.Customer_CustomerStatuses(helpers.GetCustomerStatusIndex(state)))
					}
				}

				customerServiceClient.On("GetCustomers", context.Background(),
					&customer.GetCustomersRequest{
						Keywords: test.args.keywords,
						First:    int32(test.args.first),
						After:    test.args.after,
						Last:     int32(test.args.last),
						Before:   test.args.before,
						Statuses: statuses,
					}).Return(&customer.GetCustomersResponse{
					Nodes: []*pbTypes.Customer{
						{
							Id:        "1",
							FirstName: "firstname",
							LastName:  "lastname",
							Dob:       "mm-dd-yyyt",
							Bvn:       "1200488434",
							Addresses: []*pbTypes.Address{
								{
									Primary: true,
									Country: &pbTypes.Country{
										Id:         "country_id",
										CodeAlpha2: "code_alpha_2",
										CodeAlpha3: "code_alpha_3",
										Name:       "country_name",
									},
								},
							},
							Phones: []*pbTypes.Phone{
								{
									Primary:  true,
									Number:   "234059999594",
									Verified: true,
								},
							},
							Email: &pbTypes.Email{
								Address:  "example@mail.com",
								Verified: true,
							},
							Status:   pbTypes.Customer_SIGNEDUP,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},

						{
							Id:        "2",
							FirstName: "firstname_2",
							LastName:  "lastname_2",
							Dob:       "mm-dd-yyyy",
							Bvn:       "1200488434",
							Addresses: []*pbTypes.Address{
								{
									Primary: true,
									Country: &pbTypes.Country{
										Id:         "country_id",
										CodeAlpha2: "code_alpha_2",
										CodeAlpha3: "code_alpha_3",
										Name:       "country_name",
									},
								},
							},
							Phones: []*pbTypes.Phone{
								{
									Primary:  true,
									Number:   "2349599997294",
									Verified: true,
								},
							},
							Email: &pbTypes.Email{
								Address:  "example2@mail.com",
								Verified: true,
							},
							Status:   pbTypes.Customer_REGISTERED,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
					},

					PaginationInfo: &pbTypes.PaginationInfo{
						HasNextPage:     false,
						HasPreviousPage: false,
						StartCursor:     "start_cursor",
						EndCursor:       "end_cursor",
					},

					TotalCount: 2,
				}, nil)

				resp, err := resolver.Customers(context.Background(), &test.args.keywords, &test.args.first, &test.args.after, &test.args.last, &test.args.before, test.args.statuses)

				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, resp.TotalCount, int64(2))

			case last_ten_customers:
				// convert statuses to Customer_CustomerStatuses
				statuses := make([]pbTypes.Customer_CustomerStatuses, 0)
				if len(test.args.statuses) > 0 {
					for _, state := range test.args.statuses {
						statuses = append(statuses, pbTypes.Customer_CustomerStatuses(helpers.GetCustomerStatusIndex(state)))
					}
				}

				customerServiceClient.On("GetCustomers", context.Background(),
					&customer.GetCustomersRequest{
						Keywords: test.args.keywords,
						First:    int32(test.args.first),
						After:    test.args.after,
						Last:     int32(test.args.last),
						Before:   test.args.before,
						Statuses: statuses,
					}).Return(&customer.GetCustomersResponse{
					Nodes: []*pbTypes.Customer{

						{
							Id:        "2",
							FirstName: "firstname_2",
							LastName:  "lastname_2",
							Dob:       "mm-dd-yyyy",
							Bvn:       "1200488434",
							Addresses: []*pbTypes.Address{
								{
									Primary: true,
									Country: &pbTypes.Country{
										Id:         "country_id",
										CodeAlpha2: "code_alpha_2",
										CodeAlpha3: "code_alpha_3",
										Name:       "country_name",
									},
								},
							},
							Phones: []*pbTypes.Phone{
								{
									Primary:  true,
									Number:   "2349599997294",
									Verified: true,
								},
							},
							Email: &pbTypes.Email{
								Address:  "example2@mail.com",
								Verified: true,
							},
							Status:   pbTypes.Customer_REGISTERED,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},

						{
							Id:        "1",
							FirstName: "firstname",
							LastName:  "lastname",
							Dob:       "mm-dd-yyyt",
							Bvn:       "1200488434",
							Addresses: []*pbTypes.Address{
								{
									Primary: true,
									Country: &pbTypes.Country{
										Id:         "country_id",
										CodeAlpha2: "code_alpha_2",
										CodeAlpha3: "code_alpha_3",
										Name:       "country_name",
									},
								},
							},
							Phones: []*pbTypes.Phone{
								{
									Primary:  true,
									Number:   "234059999594",
									Verified: true,
								},
							},
							Email: &pbTypes.Email{
								Address:  "example@mail.com",
								Verified: true,
							},
							Status:   pbTypes.Customer_SIGNEDUP,
							StatusTs: timestamppb.Now(),
							Ts:       timestamppb.Now(),
						},
					},

					PaginationInfo: &pbTypes.PaginationInfo{
						HasNextPage:     false,
						HasPreviousPage: false,
						StartCursor:     "start_cursor",
						EndCursor:       "end_cursor",
					},

					TotalCount: 2,
				}, nil)

				resp, err := resolver.Customers(context.Background(), &test.args.keywords, &test.args.first, &test.args.after, &test.args.last, &test.args.before, test.args.statuses)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, resp.Nodes[0].ID, "2")
			}
		})
	}
}

func Test_queryResolver_Cdds(t *testing.T) {
	onboardingServiceClient := new(mocks.OnboardingServiceClient)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"

	resp, err := resolver.Cdds(context.Background(), &first, &after, &last, &before, []types.CDDStatuses{})
	assert.Error(t, err)
	assert.NotNil(t, resp)
}
