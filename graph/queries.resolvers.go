package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"ms.api/graph/generated"
	"ms.api/protos/pb/customer"
	apiTypes "ms.api/types"
)

func (r *queryResolver) CheckEmail(ctx context.Context, email string) (bool, error) {
	return false, errors.New("not implemented")
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
	return &apiTypes.Questionary{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Questionaries(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.QuestionaryStatuses, typeArg []apiTypes.QuestionaryTypes) (*apiTypes.QuestionaryConnection, error) {
	return &apiTypes.QuestionaryConnection{
		Nodes: []*apiTypes.Questionary{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
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
	panic(fmt.Errorf(panicMsg))
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*apiTypes.Customer, error) {
	return &apiTypes.Customer{
		ID: "n/a",
	}, errors.New("not implemented")
}

func (r *queryResolver) Customers(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.CustomerStatuses) (*apiTypes.CustomerConnection, error) {
	return &apiTypes.CustomerConnection{
		Nodes: []*apiTypes.Customer{
			{
				ID: "n/a",
			},
		},
	}, errors.New("not implemented")
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
