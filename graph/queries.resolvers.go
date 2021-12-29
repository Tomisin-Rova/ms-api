package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"ms.api/graph/generated"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"
	apiTypes "ms.api/types"
)

func (r *queryResolver) CheckEmail(ctx context.Context, email string) (bool, error) {
	_, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return false, err
	}

	resp, err := r.CustomerService.CheckEmail(ctx, &customer.CheckEmailRequest{Email: email})
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

func (r *queryResolver) Addresses(ctx context.Context, first *int64, after *string, last *int64, before *string, postcode *string) (*apiTypes.AddressConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Countries(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string) (*apiTypes.CountryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OnfidoSDKToken(ctx context.Context) (*apiTypes.TokenResponse, error) {
	msg := "Not implemented"
	return &apiTypes.TokenResponse{
		Message: &msg,
		Code:    int64(500),
		Success: false,
	}, errors.New("not implemented")
}

func (r *queryResolver) Cdd(ctx context.Context, filter apiTypes.CommonQueryFilterInput) (*apiTypes.Cdd, error) {
	return &apiTypes.Cdd{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Content(ctx context.Context, id string) (*apiTypes.Content, error) {
	res, err := r.CustomerService.GetContent(ctx, &customer.GetContentRequest{Id: id})
	if err != nil {
		r.logger.Info("error fetching content", zap.String("id", id))
		return nil, err
	}
	result := &apiTypes.Content{
		ID:   res.Id,
		Type: apiTypes.ContentType(res.Type.String()),
		Link: &res.Link,
		Ts:   res.Ts.Seconds,
	}

	return result, nil
}

func (r *queryResolver) Contents(ctx context.Context, first *int64, after *string, last *int64, before *string) (*apiTypes.ContentConnection, error) {
	req := &customer.GetContentsRequest{First: int32(*first), After: *after, Last: int32(*last), Before: *before}

	res, err := r.CustomerService.GetContents(ctx, req)
	if err != nil {
		r.logger.Info("error fetching contents")
		return nil, err
	}

	nodes := make([]*apiTypes.Content, 0)
	for _, node := range res.Nodes {
		content := apiTypes.Content{
			ID:   node.Id,
			Type: apiTypes.ContentType(node.Type.String()),
			Link: &node.Link,
			Ts:   node.Ts.AsTime().Unix(),
		}

		nodes = append(nodes, &content)
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     res.PaginationInfo.HasNextPage,
		HasPreviousPage: res.PaginationInfo.HasPreviousPage,
		StartCursor:     &res.PaginationInfo.StartCursor,
		EndCursor:       &res.PaginationInfo.EndCursor,
	}

	result := &apiTypes.ContentConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(res.TotalCount),
	}

	return result, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*apiTypes.Product, error) {
	return &apiTypes.Product{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Products(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.ProductStatuses) (*apiTypes.ProductConnection, error) {
	return &apiTypes.ProductConnection{
		Nodes: []*apiTypes.Product{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
}

func (r *queryResolver) Banks(ctx context.Context, first *int64, after *string, last *int64, before *string) (*apiTypes.BankConnection, error) {
	return &apiTypes.BankConnection{
		Nodes: []*apiTypes.Bank{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
}

func (r *queryResolver) Account(ctx context.Context, id string) (*apiTypes.Account, error) {
	return &apiTypes.Account{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Accounts(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.AccountStatuses, types []apiTypes.ProductTypes) (*apiTypes.AccountConnection, error) {
	panic(panicMsg)
}

func (r *queryResolver) Transaction(ctx context.Context, id string) (*apiTypes.Transaction, error) {
	return &apiTypes.Transaction{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Transactions(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.AccountStatuses, accountIds []string, beneficiaryIds []string) (*apiTypes.TransactionConnection, error) {
	return &apiTypes.TransactionConnection{
		Nodes: []*apiTypes.Transaction{
			{ID: "n/a"},
		},
	}, errors.New("not implemented")
}

func (r *queryResolver) Beneficiary(ctx context.Context, id string) (*apiTypes.Beneficiary, error) {
	return &apiTypes.Beneficiary{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Beneficiaries(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.BeneficiaryStatuses) (*apiTypes.BeneficiaryConnection, error) {
	return &apiTypes.BeneficiaryConnection{
		Nodes: []*apiTypes.Beneficiary{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
}

func (r *queryResolver) TransactionTypes(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.TransactionTypeStatuses) (*apiTypes.TransactionTypeConnection, error) {
	return &apiTypes.TransactionTypeConnection{
		Nodes: []*apiTypes.TransactionType{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
}

func (r *queryResolver) Questionary(ctx context.Context, id string) (*apiTypes.Questionary, error) {
	resp, err := r.CustomerService.GetQuestionary(ctx, &customer.GetQuestionaryRequest{Id: id})
	if err != nil {
		return &apiTypes.Questionary{}, err
	}

	questions := make([]*apiTypes.QuestionaryQuestion, 0)
	for _, q := range resp.Questions {

		predefinedAnswers := make([]*apiTypes.QuestionaryPredefinedAnswer, 0)
		for _, pa := range q.PredefinedAnswers {
			predefinedAnswers = append(predefinedAnswers, &apiTypes.QuestionaryPredefinedAnswer{
				ID:    pa.Id,
				Value: pa.Value,
			})
		}

		question := &apiTypes.QuestionaryQuestion{
			ID:                q.Id,
			Value:             q.Value,
			PredefinedAnswers: predefinedAnswers,
			Required:          q.Required,
			MultipleOptions:   q.MultipleOptions,
		}
		questions = append(questions, question)
	}

	return &apiTypes.Questionary{
		ID:        resp.Id,
		Type:      apiTypes.QuestionaryTypes(resp.Type),
		Questions: questions,
		Status:    apiTypes.QuestionaryStatuses(resp.Status),
		StatusTs:  resp.StatusTs.AsTime().Unix(),
		Ts:        resp.Ts.AsTime().Unix(),
	}, nil
}

func (r *queryResolver) Questionaries(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.QuestionaryStatuses, typeArg []apiTypes.QuestionaryTypes) (*apiTypes.QuestionaryConnection, error) {
	helper := helpersfactory{}
	questionaryStatuses := make([]types.Questionary_QuestionaryStatuses, 0)
	questionaryTypes := make([]types.Questionary_QuestionaryTypes, 0)

	if len(statuses) > 0 {
		for _, state := range statuses {
			questionaryStatuses = append(questionaryStatuses, types.Questionary_QuestionaryStatuses(helper.GetQuestionaryStatusIndex(state)))
		}
	}

	if len(typeArg) > 0 {
		for _, arg := range typeArg {
			questionaryTypes = append(questionaryTypes, types.Questionary_QuestionaryTypes(helper.GetQuestionaryTypesIndex(arg)))
		}
	}

	customerQuestionariesReq := customer.GetQuestionariesRequest{
		Keywords: *keywords,
		First:    int32(*first),
		After:    *after,
		Last:     int32(*last),
		Statuses: questionaryStatuses,
		Types:    questionaryTypes,
	}

	resp, err := r.CustomerService.GetQuestionaries(ctx, &customerQuestionariesReq)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Questionary, 0)
	for _, node := range resp.Nodes {
		questions := make([]*apiTypes.QuestionaryQuestion, 0)
		for _, q := range node.Questions {
			question := &apiTypes.QuestionaryQuestion{
				ID:    q.Id,
				Value: q.Value,
			}
			questions = append(questions, question)
		}

		content := apiTypes.Questionary{
			ID:        node.Id,
			Type:      apiTypes.QuestionaryTypes(node.Type),
			Questions: questions,
			Status:    apiTypes.QuestionaryStatuses(node.Status),
			StatusTs:  node.StatusTs.AsTime().Unix(),
			Ts:        node.Ts.AsTime().Unix(),
		}

		nodes = append(nodes, &content)
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     resp.PaginationInfo.HasNextPage,
		HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
		StartCursor:     &resp.PaginationInfo.StartCursor,
		EndCursor:       &resp.PaginationInfo.EndCursor,
	}

	return &apiTypes.QuestionaryConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(resp.TotalCount),
	}, nil
}

func (r *queryResolver) Currency(ctx context.Context, id string) (*apiTypes.Currency, error) {
	return &apiTypes.Currency{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Currencies(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string) (*apiTypes.CurrencyConnection, error) {
	return &apiTypes.CurrencyConnection{
		Nodes: []*apiTypes.Currency{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
}

func (r *queryResolver) Fees(ctx context.Context, transactionTypeID string) ([]*apiTypes.Fee, error) {
	return []*apiTypes.Fee{
		{ID: "n/a"},
	}, errors.New("not implemented")
}

func (r *queryResolver) ExchangeRate(ctx context.Context, transactionTypeID string) (*apiTypes.ExchangeRate, error) {
	return &apiTypes.ExchangeRate{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Me(ctx context.Context) (apiTypes.MeResult, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return apiTypes.Staff{}, err
	}

	addresses := make([]*apiTypes.Address, 0)
	phones := make([]*apiTypes.Phone, 0)

	resp, err := r.CustomerService.Me(ctx, &customer.MeRequest{})
	if err != nil {
		return apiTypes.Staff{}, err
	}

	switch claims.Client {

	case models.DASHBOARD:
		staff := resp.Data.(*customer.MeResponse_Staff).Staff

		for _, addr := range staff.Addresses {
			address := apiTypes.Address{
				Primary: addr.Primary,
				Country: &apiTypes.Country{
					ID:         addr.Country.Id,
					CodeAlpha2: addr.Country.CodeAlpha2,
					CodeAlpha3: addr.Country.CodeAlpha3,
					Name:       addr.Country.Name,
				},
				State:    &addr.State,
				City:     &addr.City,
				Street:   addr.Street,
				Postcode: addr.Postcode,
				Cordinates: &apiTypes.Cordinates{
					Latitude:  float64(addr.Coordinates.Latitude),
					Longitude: float64(addr.Coordinates.Longitude),
				},
			}
			addresses = append(addresses, &address)
		}

		for _, phone := range staff.Phones {
			phone := apiTypes.Phone{
				Primary:  phone.Primary,
				Number:   phone.Number,
				Verified: phone.Verified,
			}

			phones = append(phones, &phone)
		}

		return apiTypes.Staff{
			ID:        staff.Id,
			Name:      staff.Name,
			LastName:  staff.LastName,
			Dob:       &staff.Dob,
			Addresses: addresses,
			Phones:    phones,
			Email:     staff.Email,
			Status:    apiTypes.StaffStatuses(staff.Status),
			StatusTs:  staff.StatusTs.AsTime().Unix(),
			Ts:        staff.Ts.AsTime().Unix(),
		}, nil

	case models.APP:
		appCustomer := resp.Data.(*customer.MeResponse_Customer).Customer

		for _, addr := range appCustomer.Addresses {
			address := apiTypes.Address{
				Primary: addr.Primary,
				Country: &apiTypes.Country{
					ID:         addr.Country.Id,
					CodeAlpha2: addr.Country.CodeAlpha2,
					CodeAlpha3: addr.Country.CodeAlpha3,
					Name:       addr.Country.Name,
				},
				State:    &addr.State,
				City:     &addr.City,
				Street:   addr.Street,
				Postcode: addr.Postcode,
				Cordinates: &apiTypes.Cordinates{
					Latitude:  float64(addr.Coordinates.Latitude),
					Longitude: float64(addr.Coordinates.Longitude),
				},
			}
			addresses = append(addresses, &address)
		}

		for _, phone := range appCustomer.Phones {
			phone := apiTypes.Phone{
				Primary:  phone.Primary,
				Number:   phone.Number,
				Verified: phone.Verified,
			}

			phones = append(phones, &phone)
		}

		return apiTypes.Customer{
			ID:        appCustomer.Id,
			FirstName: appCustomer.FirstName,
			LastName:  appCustomer.LastName,
			Dob:       appCustomer.Dob,
			Bvn:       &appCustomer.Bvn,
			Addresses: addresses,
			Phones:    phones,
			Email: &apiTypes.Email{
				Address:  appCustomer.Email.Address,
				Verified: appCustomer.Email.Verified,
			},
			Status:   apiTypes.CustomerStatuses(appCustomer.Status),
			StatusTs: appCustomer.StatusTs.AsTime().Unix(),
			Ts:       appCustomer.Ts.AsTime().Unix(),
		}, nil
	}

	return apiTypes.Customer{}, errors.New("unknown error occurred")
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*apiTypes.Customer, error) {
	result, err := r.CustomerService.GetCustomer(ctx, &customer.GetCustomerRequest{Id: id})
	if err != nil {
		return &apiTypes.Customer{}, errors.New("not implemented")
	}

	addresses := make([]*apiTypes.Address, 0)
	for _, addr := range result.Addresses {
		address := apiTypes.Address{
			Primary: addr.Primary,
			Country: &apiTypes.Country{
				ID:         addr.Country.Id,
				CodeAlpha2: addr.Country.CodeAlpha2,
				CodeAlpha3: addr.Country.CodeAlpha3,
				Name:       addr.Country.Name,
			},
		}
		addresses = append(addresses, &address)
	}

	phones := make([]*apiTypes.Phone, 0)
	for _, phone := range result.Phones {
		phone := apiTypes.Phone{
			Primary:  phone.Primary,
			Number:   phone.Number,
			Verified: phone.Verified,
		}

		phones = append(phones, &phone)
	}

	helpers := &helpersfactory{}

	return &apiTypes.Customer{
		ID:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Dob:       result.Dob,
		Bvn:       &result.Bvn,
		Addresses: addresses,
		Phones:    phones,
		Email: &apiTypes.Email{
			Address:  result.Email.Address,
			Verified: result.Email.Verified,
		},
		Status:   apiTypes.CustomerStatuses(helpers.GetCustomer_CustomerStatusIndex(result.Status)),
		StatusTs: result.StatusTs.AsTime().Unix(),
		Ts:       result.Ts.AsTime().Unix(),
	}, nil
}

func (r *queryResolver) Customers(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.CustomerStatuses) (*apiTypes.CustomerConnection, error) {
	helper := helpersfactory{}
	customerStatuses := make([]types.Customer_CustomerStatuses, 0)

	if len(statuses) > 0 {
		for _, state := range statuses {
			customerStatuses = append(customerStatuses, types.Customer_CustomerStatuses(helper.GetCustomerStatusIndex(state)))
		}
	}

	customerQuestionariesReq := customer.GetCustomersRequest{
		Keywords: *keywords,
		First:    int32(*first),
		After:    *after,
		Last:     int32(*last),
		Statuses: customerStatuses,
	}

	resp, err := r.CustomerService.GetCustomers(ctx, &customerQuestionariesReq)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Customer, 0)
	for _, node := range resp.Nodes {

		addresses := make([]*apiTypes.Address, 0)
		for _, addr := range node.Addresses {
			address := apiTypes.Address{
				Primary: addr.Primary,
				Country: &apiTypes.Country{
					ID:         addr.Country.Id,
					CodeAlpha2: addr.Country.CodeAlpha2,
					CodeAlpha3: addr.Country.CodeAlpha3,
					Name:       addr.Country.Name,
				},
			}
			addresses = append(addresses, &address)
		}

		phones := make([]*apiTypes.Phone, 0)
		for _, phone := range node.Phones {
			phone := apiTypes.Phone{
				Primary:  phone.Primary,
				Number:   phone.Number,
				Verified: phone.Verified,
			}

			phones = append(phones, &phone)
		}

		customer_ := apiTypes.Customer{
			ID:        node.Id,
			FirstName: node.FirstName,
			LastName:  node.LastName,
			Dob:       node.Dob,
			Bvn:       &node.Bvn,
			Addresses: addresses,
			Phones:    phones,
			Email: &apiTypes.Email{
				Address:  node.Email.Address,
				Verified: node.Email.Verified,
			},
			Status:   apiTypes.CustomerStatuses(helper.GetCustomer_CustomerStatusIndex(node.Status)),
			StatusTs: node.StatusTs.AsTime().Unix(),
			Ts:       node.Ts.AsTime().Unix(),
		}

		nodes = append(nodes, &customer_)
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     resp.PaginationInfo.HasNextPage,
		HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
		StartCursor:     &resp.PaginationInfo.StartCursor,
		EndCursor:       &resp.PaginationInfo.EndCursor,
	}

	return &apiTypes.CustomerConnection{
		Nodes: nodes, PageInfo: &pageInfo,
		TotalCount: int64(resp.TotalCount)}, nil
}

func (r *queryResolver) Cdds(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.CDDStatuses) (*apiTypes.CDDConnection, error) {
	return &apiTypes.CDDConnection{
		Nodes: []*apiTypes.Cdd{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const (
	panicMsg = "not implemented"
)
