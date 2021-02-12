package graph

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"ms.api/mocks"
	"ms.api/protos/pb/onboardingService"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"

	coreErrors "github.com/roava/zebra/errors"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
)

const (
	maxAddresses = 100
)

func genMockAddresses() []*protoTypes.Address {
	var addressRes []*protoTypes.Address
	for i := 0; i < maxAddresses; i++ {
		gs := fmt.Sprintf("Street # {%d}", i)
		addressRes = append(addressRes, &protoTypes.Address{Street: gs})
	}
	return addressRes
}

func str(str string) *string {
	return &str
}

func i64(i int64) *int64 {
	return &i
}

func Test_queryResolver_AddressLookup(t *testing.T) {
	const (
		testUnexpectedError = iota
		testFirstParam
		testLastParam
		testAfterParam
		testBeforeParam
		testHasNextPage
		testHasNextPageFalse
		testWithoutPaginationParams
	)

	type args struct {
		text   *string
		first  *int64
		after  *string
		last   *int64
		before *string
	}
	tests := []struct {
		name         string
		args         args
		testCaseType int
	}{
		{
			name: "Test first param (10 elements)",
			args: args{
				text:  str("Baker"),
				first: i64(10),
			},
			testCaseType: testFirstParam,
		},
		{
			name: "Test last param (4 elements)",
			args: args{
				text: str("Baker"),
				last: i64(4),
			},
			testCaseType: testLastParam,
		},
		{
			name: "Test after param (2 elements)",
			args: args{
				text:  str("Baker"),
				first: i64(2),
				after: str("Y3Vyc29yOjI="),
			},
			testCaseType: testAfterParam,
		},
		{
			name: "Test before param (2 elements)",
			args: args{
				text:   str("Baker"),
				first:  i64(2),
				before: str("Y3Vyc29yOjI="),
			},
			testCaseType: testBeforeParam,
		},
		{
			name: "Test hasNextPage",
			args: args{
				text:  str("Baker"),
				first: i64(2),
			},
			testCaseType: testHasNextPage,
		},
		{
			name: "Test hasNextPage false",
			args: args{
				text:  str("Baker"),
				first: i64(maxAddresses + 1),
			},
			testCaseType: testHasNextPageFalse,
		},
		{
			name: "Test without pagination params",
			args: args{
				text: str("Baker"),
			},
			testCaseType: testWithoutPaginationParams,
		},
		{
			name: "Test unexpected error",
			args: args{
				text: str("Baker"),
			},
			testCaseType: testUnexpectedError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey, models.Claims{})
			onBoardingServiceClient := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onBoardingServiceClient,
			}, zaptest.NewLogger(t))

			switch tt.testCaseType {
			case testFirstParam:
				response := &onboardingService.AddressLookupResponse{Addresses: genMockAddresses()}
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(response, nil)

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)

				assert.NotNil(t, res)
				assert.Nil(t, err)
				assert.Equal(t, len(res.Edges), 10)
				assert.Equal(t, res.TotalCount, int64(maxAddresses))
			case testLastParam:
				response := &onboardingService.AddressLookupResponse{Addresses: genMockAddresses()}
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(response, nil)

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)

				assert.NotNil(t, res)
				assert.Nil(t, err)
				assert.Equal(t, len(res.Edges), 4)
				assert.Equal(t, res.TotalCount, int64(maxAddresses))
			case testAfterParam:
				response := &onboardingService.AddressLookupResponse{Addresses: genMockAddresses()}
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(response, nil)

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)

				assert.NotNil(t, res)
				assert.Nil(t, err)
				assert.Equal(t, len(res.Edges), 2)
				assert.Equal(t, res.TotalCount, int64(maxAddresses))
				assert.Equal(t, res.Edges[0].Cursor, "Y3Vyc29yOjM=")
				assert.Equal(t, res.Edges[1].Cursor, "Y3Vyc29yOjQ=")
			case testBeforeParam:
				response := &onboardingService.AddressLookupResponse{Addresses: genMockAddresses()}
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(response, nil)

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)

				assert.NotNil(t, res)
				assert.Nil(t, err)
				assert.Equal(t, len(res.Edges), 2)
				assert.Equal(t, res.TotalCount, int64(maxAddresses))
				assert.Equal(t, res.Edges[0].Cursor, "Y3Vyc29yOjA=")
				assert.Equal(t, res.Edges[1].Cursor, "Y3Vyc29yOjE=")
			case testHasNextPage:
				response := &onboardingService.AddressLookupResponse{Addresses: genMockAddresses()}
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(response, nil)

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)

				assert.NotNil(t, res)
				assert.Nil(t, err)
				assert.Equal(t, res.PageInfo.HasNextPage, true)
			case testHasNextPageFalse:
				response := &onboardingService.AddressLookupResponse{Addresses: genMockAddresses()}
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(response, nil)

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)

				assert.NotNil(t, res)
				assert.Nil(t, err)
				assert.Equal(t, res.PageInfo.HasNextPage, false)
			case testWithoutPaginationParams:
				response := &onboardingService.AddressLookupResponse{Addresses: genMockAddresses()}
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(response, nil)

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)
				assert.NotNil(t, res)
				assert.Nil(t, err)
				assert.Equal(t, len(res.Edges), maxAddresses)
			case testUnexpectedError:
				queryResolver := resolver.Query()
				onBoardingServiceClient.On("AddressLookup", mock.Anything, &onboardingService.AddressLookupRequest{
					Text: *tt.args.text,
				}).Return(nil, errors.New(""))

				res, err := queryResolver.AddressLookup(ctx, tt.args.text, tt.args.first, tt.args.after, tt.args.last, tt.args.before)
				assert.Nil(t, res)
				assert.NotNil(t, err)
			}

		})
	}
}

func TestQueryResolver_CheckEmail(t *testing.T) {
	const (
		success = iota
		invalidEmail
		errorOnBoardingSvcCheckEmailExistence
	)

	var tests = []struct {
		name     string
		email    string
		testType int
	}{
		{
			name:     "Test check email successfully",
			email:    "valid@mail.com",
			testType: success,
		},
		{
			name:     "Test invalid email",
			email:    "invalidEmail",
			testType: invalidEmail,
		},
		{
			name:     "Test error calling onBoardingService.CheckEmailExistence()",
			email:    "valid@mail.com",
			testType: errorOnBoardingSvcCheckEmailExistence,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			onboardingServiceClient := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onboardingServiceClient,
			}, zaptest.NewLogger(t)).Query()

			switch testCase.testType {
			case success:
				onboardingServiceClient.On("CheckEmailExistence", context.Background(), &onboardingService.CheckEmailExistenceRequest{
					Email: testCase.email,
				}).Return(&onboardingService.CheckEmailExistenceResponse{
					Exists: true,
				}, nil)

				response, err := resolver.CheckEmail(context.Background(), testCase.email)
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, true, *response)
			case invalidEmail:
				response, err := resolver.CheckEmail(context.Background(), testCase.email)
				assert.Error(t, err)
				assert.Nil(t, response)
				assert.Equal(t, 1100, err.(*coreErrors.Terror).Code())
			case errorOnBoardingSvcCheckEmailExistence:
				onboardingServiceClient.On("CheckEmailExistence", context.Background(), &onboardingService.CheckEmailExistenceRequest{
					Email: testCase.email,
				}).Return(nil, errors.New(""))

				response, err := resolver.CheckEmail(context.Background(), testCase.email)
				assert.Error(t, err)
				assert.Nil(t, response)
			}

			onboardingServiceClient.AssertExpectations(t)
		})
	}
}
