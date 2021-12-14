package graph

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ms.api/mocks"
	"ms.api/protos/pb/customer"
	pbTypes "ms.api/protos/pb/types"
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
			name:     "Test error conte",
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
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	resp, err := resolver.CheckEmail(context.Background(), "")

	assert.Error(t, err)
	assert.Equal(t, resp, false)
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
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	resp, err := resolver.Questionary(context.Background(), "")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Questionaries(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"
	keywords := "keywords"

	resp, err := resolver.Questionaries(context.Background(), &keywords, &first, &after, &last, &before, []types.QuestionaryStatuses{}, []types.QuestionaryTypes{})
	assert.Error(t, err)
	assert.NotNil(t, resp)
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
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	resp, err := resolver.Customer(context.Background(), "")
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func Test_queryResolver_Customers(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	first := int64(10)
	after := "after"
	last := int64(10)
	before := "before"
	keywords := "keywords"

	resp, err := resolver.Customers(context.Background(), &keywords, &first, &after, &last, &before, []types.CustomerStatuses{})
	assert.Error(t, err)
	assert.NotNil(t, resp)
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
