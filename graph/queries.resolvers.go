package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"net/http"

	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"ms.api/graph/generated"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/protos/pb/account"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/onboarding"
	"ms.api/protos/pb/payment"
	protoTypes "ms.api/protos/pb/types"
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
		return false, nil
	}

	return resp.Success, nil
}

func (r *queryResolver) Addresses(ctx context.Context, first *int64, after *string, last *int64, before *string, postcode *string) (*apiTypes.AddressConnection, error) {
	// Build request
	var request customer.GetAddressesRequest
	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}
	if postcode != nil {
		request.Postcode = *postcode
	}

	// Execute RPC call
	response, err := r.CustomerService.GetAddresses(ctx, &request)
	if err != nil {
		return nil, err
	}

	// Build response
	nodes := make([]*apiTypes.Address, len(response.Nodes))
	for index, node := range response.Nodes {
		address := apiTypes.Address{
			Country: func() *apiTypes.Country {
				if node.Country == nil {
					return nil
				}

				return &apiTypes.Country{
					ID:         node.Country.Id,
					CodeAlpha2: node.Country.CodeAlpha2,
					CodeAlpha3: node.Country.CodeAlpha3,
					Name:       node.Country.Name,
				}
			}(),
			State:    &node.State,
			City:     &node.City,
			Street:   node.Street,
			Postcode: node.Postcode,
			Cordinates: func() *apiTypes.Cordinates {
				if node.Coordinates == nil {
					return nil
				}

				return &apiTypes.Cordinates{
					Latitude:  float64(node.Coordinates.Latitude),
					Longitude: float64(node.Coordinates.Longitude),
				}
			}(),
		}

		nodes[index] = &address
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     response.PaginationInfo.HasNextPage,
		HasPreviousPage: response.PaginationInfo.HasPreviousPage,
		StartCursor:     &response.PaginationInfo.StartCursor,
		EndCursor:       &response.PaginationInfo.EndCursor,
	}

	result := &apiTypes.AddressConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(response.TotalCount),
	}

	return result, nil
}

func (r *queryResolver) Countries(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string) (*apiTypes.CountryConnection, error) {
	// Build request
	var request customer.GetCountriesRequest
	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}
	if keywords != nil {
		request.Keywords = *keywords
	}

	// Execute RPC call
	response, err := r.CustomerService.GetCountries(ctx, &request)
	if err != nil {
		return nil, err
	}

	// Build response
	nodes := make([]*apiTypes.Country, len(response.Nodes))
	for index, node := range response.Nodes {
		address := apiTypes.Country{
			ID:         node.Id,
			CodeAlpha2: node.CodeAlpha2,
			CodeAlpha3: node.CodeAlpha3,
			Name:       node.Name,
		}

		nodes[index] = &address
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     response.PaginationInfo.HasNextPage,
		HasPreviousPage: response.PaginationInfo.HasPreviousPage,
		StartCursor:     &response.PaginationInfo.StartCursor,
		EndCursor:       &response.PaginationInfo.EndCursor,
	}

	result := &apiTypes.CountryConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(response.TotalCount),
	}

	return result, nil
}

func (r *queryResolver) OnfidoSDKToken(ctx context.Context) (*apiTypes.TokenResponse, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return &apiTypes.TokenResponse{Message: &authFailedMessage, Success: false, Code: http.StatusUnauthorized}, err
	}

	// Execute RPC call
	response, err := r.OnBoardingService.GetOnfidoSDKToken(ctx, &onboarding.GetOnfidoSDKTokenRequest{})
	if err != nil {
		return nil, err
	}

	return &apiTypes.TokenResponse{
		Success: true,
		Code:    http.StatusOK,
		Token:   response.Token,
	}, nil
}

func (r *queryResolver) Cdd(ctx context.Context, filter apiTypes.CommonQueryFilterInput) (*apiTypes.Cdd, error) {
	request := onboarding.GetCDDRequest{Last: false}
	if filter.ID != nil {
		request.Id = *filter.ID
	}
	if filter.Last != nil {
		request.Last = *filter.Last
	}
	if filter.CustomerID != nil {
		request.CustomerId = *filter.CustomerID
	}

	resp, err := r.OnBoardingService.GetCDD(ctx, &request)
	if err != nil {
		return nil, err
	}

	helpers := helpersfactory{}
	cdd := &apiTypes.Cdd{
		ID:       resp.Id,
		Customer: helpers.makeCustomerFromProto(resp.Customer),
		Amls:     helpers.makeAMLsFromProto(resp.Amls),
		Kycs:     helpers.makeKYCsFromProto(resp.Kycs),
		Poas:     helpers.makePOAsFromProto(resp.Poas),
		Status:   helpers.MapProtoCDDStatuses(resp.Status),
		StatusTs: resp.StatusTs.AsTime().Unix(),
		Ts:       resp.Ts.AsTime().Unix(),
	}

	return cdd, nil
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
	// Build request
	var request customer.GetContentsRequest
	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}

	res, err := r.CustomerService.GetContents(ctx, &request)
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
	result, err := r.AccountService.GetProduct(ctx, &account.GetProductRequest{Id: id})
	if err != nil {
		return nil, err
	}

	helpers := &helpersfactory{}

	return helpers.MakeProductFromProto(result), nil
}

func (r *queryResolver) Products(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.ProductStatuses, typeArg *apiTypes.ProductTypes) (*apiTypes.ProductConnection, error) {
	helper := helpersfactory{}
	productStatuses := make([]protoTypes.Product_ProductStatuses, len(statuses))

	if len(statuses) > 0 {
		for index, state := range statuses {
			productStatuses[index] = helper.GetProtoProductStatuses(state)
		}
	}

	// Build request
	request := account.GetProductsRequest{}

	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}

	if typeArg != nil {
		request.Type = helper.GetProtoProductTypes(*typeArg)
	}

	if len(statuses) > 0 {
		request.Statuses = productStatuses
	}

	resp, err := r.AccountService.GetProducts(ctx, &request)

	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Product, len(resp.Nodes))
	for index, node := range resp.Nodes {
		nodes[index] = helper.MakeProductFromProto(node)
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     resp.PaginationInfo.HasNextPage,
		HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
		StartCursor:     &resp.PaginationInfo.StartCursor,
		EndCursor:       &resp.PaginationInfo.EndCursor,
	}

	return &apiTypes.ProductConnection{
		Nodes: nodes, PageInfo: &pageInfo,
		TotalCount: int64(resp.TotalCount)}, nil
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
	request := account.GetAccountRequest{Id: id}
	account, err := r.AccountService.GetAccount(ctx, &request)
	if err != nil {
		return nil, err
	}

	helpers := helpersfactory{}
	return helpers.MakeAccountFromProto(account), nil
}

func (r *queryResolver) Accounts(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.AccountStatuses, types []apiTypes.ProductTypes) (*apiTypes.AccountConnection, error) {
	helpers := helpersfactory{}
	request := account.GetAccountsRequest{
		Statuses:     make([]protoTypes.Account_AccountStatuses, 0),
		ProductTypes: make([]protoTypes.Product_ProductTypes, 0),
	}

	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}
	if len(statuses) > 0 {
		for _, status := range statuses {
			request.Statuses = append(request.Statuses, helpers.MapAccountStatuses(status))
		}
	}
	if len(types) > 0 {
		for _, productType := range types {
			request.ProductTypes = append(request.ProductTypes, helpers.GetProtoProductTypes(productType))
		}
	}

	resp, err := r.AccountService.GetAccounts(ctx, &request)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Account, len(resp.Nodes))
	for i, account := range resp.Nodes {
		nodes[i] = helpers.MakeAccountFromProto(account)
	}

	pageInfo := &apiTypes.PageInfo{}
	if resp.PaginationInfo != nil {
		pageInfo.StartCursor = &resp.PaginationInfo.StartCursor
		pageInfo.EndCursor = &resp.PaginationInfo.EndCursor
		pageInfo.HasNextPage = resp.PaginationInfo.HasNextPage
		pageInfo.HasPreviousPage = resp.PaginationInfo.HasPreviousPage
	}

	return &apiTypes.AccountConnection{
		Nodes:      nodes,
		PageInfo:   pageInfo,
		TotalCount: int64(resp.TotalCount),
	}, nil
}

func (r *queryResolver) Transaction(ctx context.Context, id string) (*apiTypes.Transaction, error) {
	result, err := r.PaymentService.GetTransaction(ctx, &payment.GetTransactionRequest{Id: id})
	if err != nil {
		return nil, err
	}

	helpers := &helpersfactory{}

	return helpers.MakeTransactionFromProto(result), nil
}

func (r *queryResolver) Transactions(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.TransactionStatuses, accountIds []string, beneficiaryIds []string) (*apiTypes.TransactionConnection, error) {
	helper := helpersfactory{}
	transactionStatuses := make([]protoTypes.Transaction_TransactionStatuses, len(statuses))

	if len(statuses) > 0 {
		for index, state := range statuses {
			transactionStatuses[index] = helper.GetProtoTransactionStatuses(state)
		}
	}

	// Build request
	request := payment.GetTransactionsRequest{}

	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}
	if len(accountIds) > 0 {
		request.AccountIds = accountIds
	}
	if len(beneficiaryIds) > 0 {
		request.BeneficiaryIds = beneficiaryIds
	}
	if len(statuses) > 0 {
		request.Statuses = transactionStatuses
	}

	resp, err := r.PaymentService.GetTransactions(ctx, &request)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Transaction, len(resp.Nodes))
	for index, node := range resp.Nodes {
		nodes[index] = helper.MakeTransactionFromProto(node)
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     resp.PaginationInfo.HasNextPage,
		HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
		StartCursor:     &resp.PaginationInfo.StartCursor,
		EndCursor:       &resp.PaginationInfo.EndCursor,
	}

	return &apiTypes.TransactionConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(resp.TotalCount),
	}, nil
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

	// Build response
	response := apiTypes.Questionary{
		ID:        resp.Id,
		Type:      apiTypes.QuestionaryTypes(resp.Type),
		Questions: questions,
		Status:    apiTypes.QuestionaryStatuses(resp.Status),
		StatusTs:  resp.StatusTs.AsTime().Unix(),
		Ts:        resp.Ts.AsTime().Unix(),
	}
	switch resp.Type {
	case protoTypes.Questionary_REASONS:
		response.Type = apiTypes.QuestionaryTypesReasons
	}
	switch resp.Status {
	case protoTypes.Questionary_ACTIVE:
		response.Status = apiTypes.QuestionaryStatusesActive
	case protoTypes.Questionary_INACTIVE:
		response.Status = apiTypes.QuestionaryStatusesInactive
	}

	return &response, nil
}

func (r *queryResolver) Questionaries(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.QuestionaryStatuses, typeArg []apiTypes.QuestionaryTypes) (*apiTypes.QuestionaryConnection, error) {
	helper := helpersfactory{}
	questionaryStatuses := make([]protoTypes.Questionary_QuestionaryStatuses, 0)
	questionaryTypes := make([]protoTypes.Questionary_QuestionaryTypes, 0)

	if len(statuses) > 0 {
		for _, state := range statuses {
			questionaryStatuses = append(questionaryStatuses, helper.MapQuestionaryStatus(state))
		}
	}

	if len(typeArg) > 0 {
		for _, arg := range typeArg {
			questionaryTypes = append(questionaryTypes, helper.MapQuestionaryType(arg))
		}
	}

	// Build request
	request := customer.GetQuestionariesRequest{
		Statuses: questionaryStatuses,
		Types:    questionaryTypes,
	}
	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}
	if keywords != nil {
		request.Keywords = *keywords
	}

	resp, err := r.CustomerService.GetQuestionaries(ctx, &request)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Questionary, 0)
	for _, node := range resp.Nodes {
		questions := make([]*apiTypes.QuestionaryQuestion, 0)
		for _, q := range node.Questions {
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

		content := apiTypes.Questionary{
			ID:        node.Id,
			Type:      helper.MapProtoQuestionaryType(node.Type),
			Questions: questions,
			Status:    helper.MapProtoQuesionaryStatus(node.Status),
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
	claims, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return apiTypes.Staff{}, err
	}

	addresses := make([]*apiTypes.Address, 0)
	phones := make([]*apiTypes.Phone, 0)

	resp, err := r.CustomerService.Me(ctx, &customer.MeRequest{})
	if err != nil {
		return apiTypes.Staff{}, err
	}

	helpers := &helpersfactory{}

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
			Status:    helpers.MapProtoStaffStatuses(staff.Status),
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
			Status:   helpers.MapProtoCustomerStatuses(appCustomer.Status),
			StatusTs: appCustomer.StatusTs.AsTime().Unix(),
			Ts:       appCustomer.Ts.AsTime().Unix(),
		}, nil
	}

	return apiTypes.Customer{}, errors.New("unknown error occurred")
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*apiTypes.Customer, error) {
	result, err := r.CustomerService.GetCustomer(ctx, &customer.GetCustomerRequest{Id: id})
	if err != nil {
		return &apiTypes.Customer{}, err
	}

	helpers := &helpersfactory{}

	return helpers.makeCustomerFromProto(result), nil
}

func (r *queryResolver) Customers(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.CustomerStatuses) (*apiTypes.CustomerConnection, error) {
	helper := helpersfactory{}
	customerStatuses := make([]protoTypes.Customer_CustomerStatuses, 0)

	if len(statuses) > 0 {
		for _, state := range statuses {
			customerStatuses = append(customerStatuses, helper.GetProtoCustomerStatuses(state))
		}
	}

	// Build request
	request := customer.GetCustomersRequest{
		Statuses: customerStatuses,
	}
	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}
	if keywords != nil {
		request.Keywords = *keywords
	}

	resp, err := r.CustomerService.GetCustomers(ctx, &request)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Customer, 0)
	for _, node := range resp.Nodes {
		customer_ := helper.makeCustomerFromProto(node)
		nodes = append(nodes, customer_)
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
	helpers := helpersfactory{}
	cddStatuses := make([]protoTypes.CDD_CDDStatuses, len(statuses))

	for i, state := range statuses {
		cddStatuses[i] = helpers.MapCDDStatusesFromModel(state)
	}

	// Build request
	request := onboarding.GetCDDsRequest{
		Statuses: cddStatuses,
	}

	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}

	resp, err := r.OnBoardingService.GetCDDs(ctx, &request)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Cdd, len(resp.Nodes))

	for i, node := range resp.Nodes {
		cdd := &apiTypes.Cdd{
			ID:       node.Id,
			Customer: helpers.makeCustomerFromProto(node.Customer),
			Amls:     helpers.makeAMLsFromProto(node.Amls),
			Kycs:     helpers.makeKYCsFromProto(node.Kycs),
			Poas:     helpers.makePOAsFromProto(node.Poas),
			Status:   helpers.MapProtoCDDStatuses(node.Status),
			StatusTs: node.StatusTs.AsTime().Unix(),
			Ts:       node.Ts.AsTime().Unix(),
		}

		nodes[i] = cdd
	}

	return &apiTypes.CDDConnection{
		Nodes: nodes,
		PageInfo: &apiTypes.PageInfo{
			HasNextPage:     resp.PaginationInfo.HasNextPage,
			HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
			StartCursor:     &resp.PaginationInfo.StartCursor,
			EndCursor:       &resp.PaginationInfo.EndCursor,
		},
		TotalCount: int64(resp.TotalCount),
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
