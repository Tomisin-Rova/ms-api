package graph

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"ms.api/libs/mapper"
	"ms.api/protos/pb/accountService"
	"ms.api/protos/pb/pricingService"

	"ms.api/protos/pb/paymentService"

	"ms.api/mocks"
	cddService "ms.api/protos/pb/cddService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/personService"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"

	"github.com/golang/mock/gomock"
	coreErrors "github.com/roava/zebra/errors"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	maxAddresses = 100
)

func genMockAddresses() []*protoTypes.AddressLookup {
	var addressRes []*protoTypes.AddressLookup
	for i := 0; i < maxAddresses; i++ {
		gs := fmt.Sprintf("Street # {%d}", i)
		addressRes = append(addressRes, &protoTypes.AddressLookup{
			Street:    gs,
			Latitude:  "52.5859730797",
			Longitude: "1.3491603533",
		})
	}
	return addressRes
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
				text:  String("Baker"),
				first: Int64(10),
			},
			testCaseType: testFirstParam,
		},
		{
			name: "Test last param (4 elements)",
			args: args{
				text: String("Baker"),
				last: Int64(4),
			},
			testCaseType: testLastParam,
		},
		{
			name: "Test after param (2 elements)",
			args: args{
				text:  String("Baker"),
				first: Int64(2),
				after: String("Y3Vyc29yOjI="),
			},
			testCaseType: testAfterParam,
		},
		{
			name: "Test before param (2 elements)",
			args: args{
				text:   String("Baker"),
				first:  Int64(2),
				before: String("Y3Vyc29yOjI="),
			},
			testCaseType: testBeforeParam,
		},
		{
			name: "Test hasNextPage",
			args: args{
				text:  String("Baker"),
				first: Int64(2),
			},
			testCaseType: testHasNextPage,
		},
		{
			name: "Test hasNextPage false",
			args: args{
				text:  String("Baker"),
				first: Int64(maxAddresses + 1),
			},
			testCaseType: testHasNextPageFalse,
		},
		{
			name: "Test without pagination params",
			args: args{
				text: String("Baker"),
			},
			testCaseType: testWithoutPaginationParams,
		},
		{
			name: "Test unexpected error",
			args: args{
				text: String("Baker"),
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
				assert.Equal(t, *res.TotalCount, int64(maxAddresses))
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
				assert.Equal(t, *res.TotalCount, int64(maxAddresses))
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
				assert.Equal(t, *res.TotalCount, int64(maxAddresses))
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
				assert.Equal(t, *res.TotalCount, int64(maxAddresses))
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

func Test_queryResolver_Person(t *testing.T) {
	const (
		personNotFound = iota
		personFound
	)

	tests := []struct {
		name     string
		testType int
	}{
		{
			name:     "Test person not found",
			testType: personNotFound,
		},
		{
			name:     "Test person  found",
			testType: personFound,
		},
	}

	personServiceClient := new(mocks.PersonServiceClient)
	cddServiceClient := new(mocks.CddServiceClient)

	resolver := NewResolver(&ResolverOpts{
		personService: personServiceClient,
		cddClient:     cddServiceClient,
	}, zaptest.NewLogger(t)).Query()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.testType {
			case personNotFound:
				personServiceClient.On("Person", context.Background(), &personService.PersonRequest{
					Id: "01eyx7ew2gt0en7e613tkyt1x4",
				}).Return(nil, errors.New(""))

				response, err := resolver.Person(context.Background(), "01eyx7ew2gt0en7e613tkyt1x4")
				assert.Error(t, err)
				assert.Nil(t, response)
			case personFound:
				personServiceClient.On("Person", context.Background(), &personService.PersonRequest{
					Id: "01eyx7ew2gt0en7e613tkyt1xc",
				}).Return(&protoTypes.Person{
					Id:        "01eyx7ew2gt0en7e613tkyt1xc",
					Title:     "Title",
					FirstName: "FirstName",
					LastName:  "LastName",
				}, nil)
				cddServiceClient.On("GetCDDByOwner", context.Background(), &cddService.GetCDDByOwnerRequest{
					PersonId: "01eyx7ew2gt0en7e613tkyt1xc",
				}).Return(&protoTypes.Cdd{
					Id: "cddId",
				}, nil)

				response, err := resolver.Person(context.Background(), "01eyx7ew2gt0en7e613tkyt1xc")
				assert.NoError(t, err)
				assert.NotNil(t, response)
			}

		})
	}
}

func TestQueryResolver_Me(t *testing.T) {
	const (
		success            = iota
		successNotCDDFound = iota
		errorNotAuthenticatedUser
		errorGettingPerson
		errorGettingCDD
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test query me successfully",
			testType: success,
		},
		{
			name:     "Test query me successfully with no cdd",
			testType: successNotCDDFound,
		},
		{
			name:     "Test error not authenticated user",
			testType: errorNotAuthenticatedUser,
		},
		{
			name:     "Test error getting person",
			testType: errorGettingPerson,
		},
		{
			name:     "Test error getting cdd",
			testType: errorGettingCDD,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			personServiceClient := new(mocks.PersonServiceClient)
			cddServiceClient := new(mocks.CddServiceClient)

			resolver := NewResolver(&ResolverOpts{
				cddClient:     cddServiceClient,
				personService: personServiceClient,
			}, zaptest.NewLogger(t))
			ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey,
				models.Claims{
					PersonId:   "personId",
					IdentityId: "identityId",
					DeviceId:   "deviceId",
				})
			switch testCase.testType {
			case success:
				personServiceClient.On("Person", ctx, &personService.PersonRequest{
					Id: "personId",
				}).Return(&protoTypes.Person{
					Id: "personId",
				}, nil)
				cddServiceClient.On("GetCDDByOwner", ctx, &cddService.GetCDDByOwnerRequest{
					PersonId: "personId",
				}).Return(&protoTypes.Cdd{
					Id: "cddId",
					Owner: &protoTypes.Person{
						Id: "ownerId",
					},
					Validations: []*protoTypes.Validation{
						{
							ValidationType: models.SCREEN,
							Data: &anypb.Any{
								TypeUrl: models.SCREEN,
								Value:   []byte("{}"),
							},
							Organisation: &protoTypes.Organisation{},
						},
					},
				}, nil)

				me, err := resolver.Query().Me(ctx)
				assert.NoError(t, err)
				assert.NotNil(t, me)
				assert.NotEmpty(t, me)
			case successNotCDDFound:
				personServiceClient.On("Person", ctx, &personService.PersonRequest{
					Id: "personId",
				}).Return(&protoTypes.Person{
					Id: "personId",
				}, nil)
				cddServiceClient.On("GetCDDByOwner", ctx, &cddService.GetCDDByOwnerRequest{
					PersonId: "personId",
				}).Return(nil, errors.New("rpc error: code = Unknown desc = {\"error\":{\"code\":1105,\"type\":\"CddNotFound\",\"message\":\"cdd record not found for this customer\",\"detail\":\"cdd record not found for this customer: cdd not found by id\"}}"))

				me, err := resolver.Query().Me(ctx)
				assert.NoError(t, err)
				assert.NotNil(t, me)
				assert.NotEmpty(t, me)
			case errorNotAuthenticatedUser:
				me, err := resolver.Query().Me(context.Background())
				assert.Error(t, err)
				assert.Nil(t, me)
			case errorGettingPerson:
				personServiceClient.On("Person", ctx, &personService.PersonRequest{
					Id: "personId",
				}).Return(nil, errors.New(""))

				me, err := resolver.Query().Me(ctx)
				assert.Error(t, err)
				assert.Nil(t, me)
			case errorGettingCDD:
				personServiceClient.On("Person", ctx, &personService.PersonRequest{
					Id: "personId",
				}).Return(&protoTypes.Person{
					Id: "personId",
				}, nil)
				cddServiceClient.On("GetCDDByOwner", ctx, &cddService.GetCDDByOwnerRequest{
					PersonId: "personId",
				}).Return(nil, errors.New(""))

				me, err := resolver.Query().Me(ctx)
				assert.Error(t, err)
				assert.Nil(t, me)
			}

			personServiceClient.AssertExpectations(t)
			cddServiceClient.AssertExpectations(t)
		})
	}
}

func Test_queryResolver_People(t *testing.T) {
	const (
		peopleNotFound = iota
		peopleFound
	)

	tests := []struct {
		name     string
		testType int
	}{
		{
			name:     "Test people not found",
			testType: peopleNotFound,
		},
		{
			name:     "Test people found",
			testType: peopleFound,
		},
	}

	personServiceClient := new(mocks.PersonServiceClient)

	resolver := NewResolver(&ResolverOpts{
		personService: personServiceClient,
	}, zaptest.NewLogger(t)).Query()

	var first *int64
	var after *string
	var last *int64
	var before *string

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.testType {
			case peopleNotFound:
				var onboarded bool
				personServiceClient.On("People", context.Background(), &personService.PeopleRequest{
					First: 100, After: "", Last: 0, Before: "", Keywords: "John Smith", Onboarded: "NOT_ONBOARDED",
				}).Return(nil, errors.New(""))
				kw := "John Smith"
				response, err := resolver.People(context.Background(), &kw, first, after, last, before, &onboarded)
				assert.NotNil(t, err)
				assert.Nil(t, response)
			case peopleFound:
				personServiceClient.On("People", context.Background(), &personService.PeopleRequest{
					First: 100, After: "", Last: 0, Before: "", Keywords: "Luke", Onboarded: "NOT_ONBOARDED",
				}).Return(&protoTypes.Persons{}, nil)
				kw := "Luke"
				var onboarded bool
				response, err := resolver.People(context.Background(), &kw, first, after, last, before, &onboarded)
				assert.Nil(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

func TestQueryResolver_Cdds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOwner := "personId"
	firstCdd := protoTypes.Cdd{
		Id: "id1",
		Owner: &protoTypes.Person{
			Id: mockOwner,
		},
		Validations: []*protoTypes.Validation{
			{ValidationType: "CHECK", Data: &anypb.Any{Value: []byte("{\"id\": \"checkId\"}")}, Organisation: &protoTypes.Organisation{Id: "orgId"}, Applicant: mockOwner},
			{ValidationType: "SCREEN", Data: &anypb.Any{Value: []byte("{\"id\": \"screenId\"}")}, Organisation: &protoTypes.Organisation{Id: "orgId"}, Applicant: mockOwner},
		},
	}
	secondCdd := protoTypes.Cdd{
		Id: "id2",
		Owner: &protoTypes.Person{
			Id: mockOwner,
		},
		Validations: []*protoTypes.Validation{
			{ValidationType: "CHECK", Data: &anypb.Any{Value: []byte("{id: \"checkId\"}")}, Organisation: &protoTypes.Organisation{Id: "orgId"}, Applicant: mockOwner},
			{ValidationType: "SCREEN", Data: &anypb.Any{Value: []byte("{id: \"screenId\"}")}, Organisation: &protoTypes.Organisation{Id: "orgId"}, Applicant: mockOwner},
		},
	}
	mockCdds := &protoTypes.Cdds{
		Results: []*protoTypes.Cdd{
			&firstCdd,
			&secondCdd,
		},
	}

	mockStore := mocks.NewMockDataStore(ctrl)

	var first *int64
	var after *string
	var last *int64
	var before *string

	t.Run("CDDS_No_Data", func(t *testing.T) {
		cddServiceClient := new(mocks.CddServiceClient)
		resolver := NewResolver(&ResolverOpts{
			cddClient: cddServiceClient,
			DataStore: mockStore,
		}, zaptest.NewLogger(t)).Query()
		kw := "John Smith"
		cddServiceClient.On("CDDS", context.Background(), &cddService.CDDSRequest{
			Page:     1,
			PerPage:  100,
			Keywords: kw,
			First:    100,
		}).Return(nil, errors.New("No Data"))

		response, err := resolver.Cdds(context.Background(), &kw, nil, first, after, last, before)
		assert.NotNil(t, err, "should return an error if no data is found")
		assert.Nil(t, response, "should return empty response if no data is found")
	})

	t.Run("CDDS_With_Data", func(t *testing.T) {
		cddServiceClient := new(mocks.CddServiceClient)
		resolver := NewResolver(&ResolverOpts{
			cddClient: cddServiceClient,
			DataStore: mockStore,
		}, zaptest.NewLogger(t)).Query()
		kw := "John Smith"
		cddServiceClient.On("CDDS", context.Background(), &cddService.CDDSRequest{
			Page:     1,
			PerPage:  100,
			Keywords: kw,
			First:    100,
		}).Return(mockCdds, nil)
		response, err := resolver.Cdds(context.Background(), &kw, nil, first, after, last, before)
		assert.Nil(t, err, "should not return an error if data is found")
		assert.NotNil(t, response, "should return a valid response if data is found")
	})
}

func TestQueryResolver_Payee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPayee := &protoTypes.Payee{
		Id:     "ID",
		Owner:  "owner",
		Name:   "Name",
		Avatar: "Avatar",
		Ts:     time.Now().Unix(),
		Accounts: []*protoTypes.PayeeAccount{
			{
				Id:            "accountId",
				Iban:          "iban",
				AccountNumber: "accountnumber",
			},
		},
	}

	mockPerson := &protoTypes.Person{
		Id:         "ID",
		Title:      "Title",
		FirstName:  "FirstName",
		LastName:   "LastName",
		MiddleName: "MiddleName",
		Dob:        "Dob",
		Ts:         time.Now().Unix(),
	}

	mockIdentity := &models.Identity{
		ID:     "identityId",
		Owner:  "personId",
		Active: true,
		Credentials: models.Credentials{
			Identifier: "hashuser@email.com",
			Password:   "hashpasscode",
			Pin:        "transactionPin8",
		},
	}

	paymentClient, personClient, preloader := &mocks.PaymentServiceClient{}, &mocks.PersonServiceClient{}, &mocks.Preloader{}
	mockStore := mocks.NewMockDataStore(ctrl)

	ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey,
		models.Claims{
			PersonId:   "personId",
			IdentityId: "identityId",
			DeviceId:   "deviceId",
		})

	paymentClient.On("GetPayee", ctx, &paymentService.GetPayeeRequest{
		PayeeId:    "payeeId",
		IdentityId: "identityId",
	}).Return(mockPayee, nil)

	personClient.On("Person", ctx, &personService.PersonRequest{Id: "personId"}).
		Return(mockPerson, nil)

	preloader.On("GetPreloads", ctx).Return([]string{"owner", "owner.owner"})

	mockStore.EXPECT().GetIdentityById("identityId").Return(mockIdentity, nil)

	resolverOpts := &ResolverOpts{
		paymentService: paymentClient,
		personService:  personClient,
		preloader:      preloader,
		DataStore:      mockStore,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()

	payee, err := resolver.Payee(ctx, "payeeId")
	assert.Nil(t, err)
	assert.NotNil(t, payee)
	assert.Equal(t, payee.ID, mockPayee.Id)
	assert.Equal(t, payee.Name, mockPayee.Name)
	assert.Equal(t, mockPayee.Avatar, *payee.Avatar)
	assert.Equal(t, mockPayee.Accounts[0].Id, payee.Accounts[0].ID)
	assert.Equal(t, payee.Owner.ID, mockIdentity.ID)
}

func TestQueryResolver_Product(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProduct := &protoTypes.Product{
		Id:             "ID",
		Identification: "identification",
		Scheme:         "scheme",
		Details: &protoTypes.ProductDetails{
			OverdraftSetting: &protoTypes.OverdraftSetting{
				AllowTechnicalOverdraft: true,
				InterestSettings: &protoTypes.InterestSettings{
					DaysInYear: "123",
					RateTiers: []*protoTypes.RateTiers{{
						EndingBalance: 12,
					}},
				},
			},
		},
	}

	accountServiceClient := &mocks.AccountServiceClient{}

	ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey,
		models.Claims{
			PersonId:   "personId",
			IdentityId: "identityId",
			DeviceId:   "deviceId",
		})

	resolverOpts := &ResolverOpts{
		accountService: accountServiceClient,
		mapper:         &mapper.GQLMapper{},
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	accountServiceClient.On("GetProduct", ctx, &accountService.GetProductRequest{Id: "productId"}).
		Return(mockProduct, nil)
	product, err := resolver.Product(ctx, "productId")
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.ID, mockProduct.Id)
	assert.Equal(t, *product.Identification, mockProduct.Identification)
	assert.Equal(t, *product.Details.OverdraftSetting.InterestSettings.DaysInYear, mockProduct.Details.OverdraftSetting.InterestSettings.DaysInYear)
}

func TestQueryResolver_MeStaff(t *testing.T) {
	const (
		success = iota
		errorNotAuthenticatedUser
		errorGettingStaff
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test query me staff successfully",
			testType: success,
		},
		{
			name:     "Test error not authenticated user",
			testType: errorNotAuthenticatedUser,
		},
		{
			name:     "Test error getting staff",
			testType: errorGettingStaff,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			personServiceClient := new(mocks.PersonServiceClient)

			resolver := NewResolver(&ResolverOpts{
				personService: personServiceClient,
			}, zaptest.NewLogger(t))
			ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey,
				models.Claims{
					PersonId:   "personId",
					IdentityId: "identityId",
					DeviceId:   "deviceId",
				})
			switch testCase.testType {
			case success:
				personServiceClient.On("GetStaffById", ctx, &personService.StaffRequest{
					Id: "personId",
				}).Return(&protoTypes.Staff{
					Id: "personId",
				}, nil)

				me, err := resolver.Query().MeStaff(ctx)
				assert.NoError(t, err)
				assert.NotNil(t, me)
				assert.NotEmpty(t, me)
			case errorNotAuthenticatedUser:
				me, err := resolver.Query().MeStaff(context.Background())
				assert.Error(t, err)
				assert.Nil(t, me)
			}

			personServiceClient.AssertExpectations(t)
		})
	}
}

func TestQueryResolver_Accounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccounts := []*protoTypes.Account{{
		Id:      "ID",
		Product: "product",
		AccountData: &protoTypes.AccountData{
			Balances: &protoTypes.Balances{
				TotalBalance: 100,
			},
		},
		Transactions: []*protoTypes.Transaction{
			{
				Id: "trans1",
				TransactionData: &protoTypes.TransactionData{
					Amount: 500,
				},
			},
			{
				Id: "trans2",
				TransactionData: &protoTypes.TransactionData{
					Amount: 400,
				},
			},
		},
	}}

	accountServiceClient, preloader := &mocks.AccountServiceClient{}, &mocks.Preloader{}

	ctx := context.WithValue(context.Background(), middlewares.AuthenticatedUserContextKey,
		models.Claims{
			PersonId:   "personId",
			IdentityId: "identityId",
			DeviceId:   "deviceId",
		})

	accountServiceClient.On("GetAccounts", mock.Anything, &accountService.GetAccountsRequest{IdentityId: "identityId"}).
		Return(&accountService.GetAccountsResponse{Accounts: mockAccounts}, nil)
	preloader.On("GetPreloads", ctx).Return([]string{"nodes.transactions"})
	var argMap = map[string]interface{}{}
	preloader.On("GetArgMap", ctx, "Transactions").Return(argMap)

	resolverOpts := &ResolverOpts{
		accountService: accountServiceClient,
		mapper:         &mapper.GQLMapper{},
		preloader:      preloader,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	accounts, err := resolver.Accounts(ctx, Int64(1), String(""), Int64(10), String(""))
	assert.Nil(t, err)
	assert.NotNil(t, accounts)
	assert.Equal(t, accounts.Nodes[0].ID, mockAccounts[0].Id)
	assert.Equal(t, accounts.Nodes[0].AccountData.Balances.TotalBalance, Int64(int64(mockAccounts[0].AccountData.Balances.TotalBalance)))
	assert.Equal(t, accounts.Nodes[0].Transactions.Nodes[0].ID, mockAccounts[0].Transactions[0].Id)
	assert.Equal(t, accounts.Nodes[0].Transactions.Nodes[0].TransactionData.Amount, &mockAccounts[0].Transactions[0].TransactionData.Amount)
}

func TestQueryResolver_TransferFees(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	PricingServiceClient := new(mocks.PricingServiceClient)

	mockFee := &protoTypes.TransferFees{
		Currency:     "GBP",
		BaseCurrency: "NGN",
		Fees: []*protoTypes.Fee{{
			LowerBoundary: 12,
			UpperBoundary: 2,
			Fee:           50,
		}},
		Ts: 12,
	}
	mockReq := &pricingService.TransferFeesRequest{
		Currency:     "GBP",
		BaseCurrency: "NGN",
	}
	PricingServiceClient.On("GetTransferFees", mock.Anything, mockReq).Return(mockFee, nil)

	resolverOpts := &ResolverOpts{
		pricingService: PricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	fees, err := resolver.TransferFees(context.Background(), mockReq.Currency, mockReq.BaseCurrency)
	assert.Nil(t, err)
	assert.NotNil(t, fees)
	assert.Equal(t, mockFee.Currency, fees.Currency)
	assert.Equal(t, mockFee.BaseCurrency, fees.BaseCurrency)
	assert.Equal(t, mockFee.Fees[0].Fee, float32(fees.Fees[0].Fee))
}

func TestQueryResolver_Fx(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	PricingServiceClient := new(mocks.PricingServiceClient)

	mockFx := &protoTypes.Fx{
		Currency:     "GBP",
		BaseCurrency: "NGN",
		BuyRate:      123,
		SellRate:     2313,
		Ts:           12,
	}
	mockReq := &pricingService.FxRequest{
		Currency:     "GBP",
		BaseCurrency: "NGN",
	}
	PricingServiceClient.On("GetFxRates", mock.Anything, mockReq).Return(mockFx, nil)

	resolverOpts := &ResolverOpts{
		pricingService: PricingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Query()
	fx, err := resolver.Fx(context.Background(), mockReq.Currency, mockReq.BaseCurrency)
	assert.Nil(t, err)
	assert.NotNil(t, fx)
	assert.Equal(t, mockFx.Currency, fx.Currency)
	assert.Equal(t, mockFx.BaseCurrency, fx.BaseCurrency)
	assert.Equal(t, mockFx.BuyRate, float32(fx.BuyRate))
}
