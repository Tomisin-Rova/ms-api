package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ms.api/graph/generated"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/account"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/onboarding"
	"ms.api/protos/pb/payment"
	"ms.api/protos/pb/pricing"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"
	apiTypes "ms.api/types"
)

// CheckEmail is the resolver for the checkEmail field.
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

// CheckPhoneNumber is the resolver for the checkPhoneNumber field.
func (r *queryResolver) CheckPhoneNumber(ctx context.Context, phone string) (bool, error) {
	phonevalidator := phonenumbervalidator.Validator{}
	err := phonevalidator.ValidatePhoneNumber(phone)
	if err != nil {
		r.logger.Info("invalid phone supplied", zap.String("phone", phone))
		return false, err
	}

	resp, err := r.CustomerService.CheckPhoneNumber(ctx, &customer.CheckPhoneNumberRequest{Phone: phone})
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

// Addresses is the resolver for the addresses field.
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

// Countries is the resolver for the countries field.
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
	helpers := &helpersfactory{}
	nodes := make([]*apiTypes.Country, len(response.Nodes))
	for index, node := range response.Nodes {
		address := apiTypes.Country{
			ID:         node.Id,
			CodeAlpha2: node.CodeAlpha2,
			CodeAlpha3: node.CodeAlpha3,
			Name:       node.Name,
			States:     helpers.makeStatesFromProto(node.States),
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

// OnfidoSDKToken is the resolver for the onfidoSDKToken field.
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

// Cdd is the resolver for the cdd field.
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

// Content is the resolver for the content field.
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

// Contents is the resolver for the contents field.
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

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*apiTypes.Product, error) {
	result, err := r.AccountService.GetProduct(ctx, &account.GetProductRequest{Id: id})
	if err != nil {
		return nil, err
	}

	helpers := &helpersfactory{}

	return helpers.MakeProductFromProto(result), nil
}

// Products is the resolver for the products field.
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

// Banks is the resolver for the banks field.
func (r *queryResolver) Banks(ctx context.Context, first *int64, after *string, last *int64, before *string) (*apiTypes.BankConnection, error) {
	// Build request
	var request payment.GetBanksRequest
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

	// Execute RPC call
	response, err := r.PaymentService.GetBanks(ctx, &request)
	if err != nil {
		return nil, err
	}

	// Build response
	nodes := make([]*apiTypes.Bank, len(response.Nodes))
	for index, node := range response.Nodes {
		address := apiTypes.Bank{
			ID:            node.Id,
			BankCode:      node.BankCode,
			BankName:      node.BankName,
			BankShortName: node.BankShortName,
			Active:        node.Active,
			Ts:            node.Ts.AsTime().Unix(),
		}

		nodes[index] = &address
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     response.PaginationInfo.HasNextPage,
		HasPreviousPage: response.PaginationInfo.HasPreviousPage,
		StartCursor:     &response.PaginationInfo.StartCursor,
		EndCursor:       &response.PaginationInfo.EndCursor,
	}

	result := &apiTypes.BankConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(response.TotalCount),
	}

	return result, nil
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*apiTypes.Account, error) {
	request := account.GetAccountRequest{Id: id}
	account, err := r.AccountService.GetAccount(ctx, &request)
	if err != nil {
		return nil, err
	}

	helpers := helpersfactory{}
	return helpers.MakeAccountFromProto(account), nil
}

// Accounts is the resolver for the accounts field.
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

// Transaction is the resolver for the transaction field.
func (r *queryResolver) Transaction(ctx context.Context, id string) (*apiTypes.Transaction, error) {
	result, err := r.PaymentService.GetTransaction(ctx, &payment.GetTransactionRequest{Id: id})
	if err != nil {
		return nil, err
	}

	helpers := &helpersfactory{}

	return helpers.MakeTransactionFromProto(result), nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, first *int64, after *string, last *int64, before *string, startDate *string, endDate *string, statuses []apiTypes.TransactionStatuses, accountIds []string, beneficiaryIds []string, hasBeneficiary *bool) (*apiTypes.TransactionConnection, error) {
	helper := helpersfactory{}
	transactionStatuses := make([]protoTypes.Transaction_TransactionStatuses, len(statuses))

	if len(statuses) > 0 {
		for index, state := range statuses {
			transactionStatuses[index] = helper.GetProtoTransactionStatuses(state)
		}
	}

	// Build request
	request := payment.GetTransactionsRequest{}
	const dateTemplate = "02-01-2006"
	if startDate != nil && *startDate != "" {
		formartedStartDate, err := time.Parse(dateTemplate, *startDate)
		if err != nil {
			return nil, err
		}

		request.StartDate = timestamppb.New(formartedStartDate)
	}
	if endDate != nil && *endDate != "" {
		formatedEndDate, err := time.Parse(dateTemplate, *endDate)
		if err != nil {
			return nil, err
		}
		request.EndDate = timestamppb.New(formatedEndDate)
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
	if len(accountIds) > 0 {
		request.AccountIds = accountIds
	}
	if len(beneficiaryIds) > 0 {
		request.BeneficiaryIds = beneficiaryIds
	}
	if len(statuses) > 0 {
		request.Statuses = transactionStatuses
	}
	if hasBeneficiary != nil {
		request.HasBeneficiary = *hasBeneficiary
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

// Beneficiary is the resolver for the beneficiary field.
func (r *queryResolver) Beneficiary(ctx context.Context, id string) (*apiTypes.Beneficiary, error) {
	result, err := r.PaymentService.GetBeneficiary(ctx, &payment.GetBeneficiaryRequest{Id: id})
	if err != nil {
		return nil, err
	}

	helpers := &helpersfactory{}

	return helpers.MakeBeneficiaryFromProto(result), nil
}

// Beneficiaries is the resolver for the beneficiaries field.
func (r *queryResolver) Beneficiaries(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, statuses []apiTypes.BeneficiaryStatuses, sortBy *apiTypes.BeneficiarySort) (*apiTypes.BeneficiaryConnection, error) {
	helper := helpersfactory{}
	beneficiaryStatuses := make([]protoTypes.Beneficiary_BeneficiaryStatuses, len(statuses))

	if len(statuses) > 0 {
		for index, state := range statuses {
			beneficiaryStatuses[index] = helper.GetProtoBeneficiaryStatuses(state)
		}
	}

	// Build request
	request := payment.GetBeneficiariesRequest{}

	if keywords != nil {
		request.Keywords = *keywords
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
		request.Statuses = beneficiaryStatuses
	}

	if sortBy != nil {
		sort := *sortBy
		switch sort {
		case apiTypes.BeneficiarySortName:
			request.SortBy = payment.GetBeneficiariesRequest_NAME
		case apiTypes.BeneficiarySortTs:
			request.SortBy = payment.GetBeneficiariesRequest_TS
		default:
			request.SortBy = payment.GetBeneficiariesRequest_TS
		}
	}

	resp, err := r.PaymentService.GetBeneficiaries(ctx, &request)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Beneficiary, len(resp.Nodes))
	for index, node := range resp.Nodes {
		nodes[index] = helper.MakeBeneficiaryFromProto(node)
	}

	pageInfo := &apiTypes.PageInfo{}
	if resp.PaginationInfo != nil {
		pageInfo = &apiTypes.PageInfo{
			HasNextPage:     resp.PaginationInfo.HasNextPage,
			HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
			StartCursor:     &resp.PaginationInfo.StartCursor,
			EndCursor:       &resp.PaginationInfo.EndCursor,
		}
	}

	return &apiTypes.BeneficiaryConnection{
		Nodes:      nodes,
		PageInfo:   pageInfo,
		TotalCount: int64(resp.TotalCount),
	}, nil
}

// ExistingBeneficiariesByPhone is the resolver for the existingBeneficiariesByPhone field.
func (r *queryResolver) ExistingBeneficiariesByPhone(ctx context.Context, phones []string, transactionPassword string) ([]*string, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Build request
	request := payment.GetBeneficiariesByPhoneRequest{
		Phones:              phones,
		TransactionPassword: transactionPassword,
	}

	// Execute RPC call
	response, err := r.PaymentService.GetBeneficiariesByPhone(ctx, &request)
	if err != nil {
		return nil, err
	}
	results := make([]*string, len(response.Phones))
	for i := 0; i < len(response.Phones); i++ {
		results[i] = &response.Phones[i]
	}

	return results, nil
}

// ExistingBeneficiaryByAccount is the resolver for the existingBeneficiaryByAccount field.
func (r *queryResolver) ExistingBeneficiaryByAccount(ctx context.Context, accountNumber string) (*apiTypes.BeneficiaryPreview, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Build request
	request := payment.GetBeneficiaryByAccountRequest{
		AccountNumber: accountNumber,
	}

	// Execute RPC call
	response, err := r.PaymentService.GetBeneficiaryByAccount(ctx, &request)
	if err != nil {
		return nil, err
	}

	helpers := &helpersfactory{}

	currency := helpers.MakeCurrencyFromProto(response.Currency)

	return &apiTypes.BeneficiaryPreview{
		Name:          response.Name,
		Currency:      currency,
		AccountNumber: response.AccountNumber,
		Code:          response.Code,
	}, nil
}

// LookupBeneficiary is the resolver for the lookupBeneficiary field.
func (r *queryResolver) LookupBeneficiary(ctx context.Context, accountNumber string, code string, currencyID string) (*apiTypes.BeneficiaryPreview, error) {
	// Get user claims
	_, err := middlewares.GetClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	result, err := r.PaymentService.LookupBeneficiary(ctx, &payment.LookUpBeneficiaryRequest{Code: code, AccountNumber: accountNumber, CurrencyId: currencyID})
	if err != nil {
		return nil, err
	}

	response := &apiTypes.BeneficiaryPreview{
		Name: result.Name,
		Currency: &apiTypes.Currency{
			ID:     result.Currency.Id,
			Symbol: result.Currency.Symbol,
			Code:   result.Currency.Code,
			Name:   result.Currency.Name,
		},
		AccountNumber: result.AccountNumber,
		Code:          result.Code,
	}

	return response, nil
}

// TransactionTypes is the resolver for the transactionTypes field.
func (r *queryResolver) TransactionTypes(ctx context.Context, first *int64, after *string, last *int64, before *string, statuses []apiTypes.TransactionTypeStatuses) (*apiTypes.TransactionTypeConnection, error) {
	transactionTypesStatuses := make([]protoTypes.TransactionType_TransactionTypeStatuses, len(statuses))

	if len(transactionTypesStatuses) > 0 {
		for index, state := range statuses {
			transactionTypesStatuses[index] = r.helper.GetProtoTransactionTypesStatuses(state)
		}
	}

	// Build request
	request := payment.GetTransactionTypesRequest{}

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
		request.Statuses = transactionTypesStatuses
	}

	resp, err := r.PaymentService.GetTransactionTypes(ctx, &request)

	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.TransactionType, len(resp.Nodes))
	for index, node := range resp.Nodes {
		nodes[index] = &apiTypes.TransactionType{
			ID:       node.Id,
			Name:     node.Name,
			Status:   r.helper.MapTransactionTypeStatus(node.Status),
			StatusTs: node.StatusTs.AsTime().Unix(),
			Ts:       node.Ts.AsTime().Unix(),
		}
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     resp.PaginationInfo.HasNextPage,
		HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
		StartCursor:     &resp.PaginationInfo.StartCursor,
		EndCursor:       &resp.PaginationInfo.EndCursor,
	}

	return &apiTypes.TransactionTypeConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(resp.TotalCount)}, nil
}

// Questionary is the resolver for the questionary field.
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

// Questionaries is the resolver for the questionaries field.
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

// Currency is the resolver for the currency field.
func (r *queryResolver) Currency(ctx context.Context, id string) (*apiTypes.Currency, error) {
	// Make call
	currency, err := r.PricingService.GetCurrency(ctx, &pricing.GetCurrencyRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &apiTypes.Currency{
		ID:     currency.Id,
		Symbol: currency.Symbol,
		Code:   currency.Code,
		Name:   currency.Name,
	}, nil
}

// Currencies is the resolver for the currencies field.
func (r *queryResolver) Currencies(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string) (*apiTypes.CurrencyConnection, error) {
	// Build request
	var request pricing.GetCurrenciesRequest
	if keywords != nil {
		request.Keywords = *keywords
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

	// Make call
	currencies, err := r.PricingService.GetCurrencies(ctx, &request)
	if err != nil {
		return nil, err
	}

	// Build response
	response := &apiTypes.CurrencyConnection{
		Nodes: make([]*apiTypes.Currency, len(currencies.Nodes)),
		PageInfo: &apiTypes.PageInfo{
			HasNextPage:     currencies.PaginationInfo.HasNextPage,
			HasPreviousPage: currencies.PaginationInfo.HasPreviousPage,
			StartCursor:     &currencies.PaginationInfo.StartCursor,
			EndCursor:       &currencies.PaginationInfo.EndCursor,
		},
		TotalCount: int64(currencies.TotalCount),
	}
	for index, currency := range currencies.Nodes {
		response.Nodes[index] = &apiTypes.Currency{
			ID:     currency.Id,
			Symbol: currency.Symbol,
			Code:   currency.Code,
			Name:   currency.Name,
		}
	}

	return response, nil
}

// Fees is the resolver for the fees field.
func (r *queryResolver) Fees(ctx context.Context, transactionTypeID string, sourceAccountID string, targetAccountID string) ([]*apiTypes.Fee, error) {
	resp, err := r.PricingService.GetFees(ctx, &pricing.GetFeesRequest{TransactionTypeId: transactionTypeID, SourceAccountId: sourceAccountID, TargetAccountId: targetAccountID})
	if err != nil {
		return nil, err
	}

	return r.helper.MakeFeesFromProto(resp.Fees), nil
}

// ExchangeRate is the resolver for the exchangeRate field.
func (r *queryResolver) ExchangeRate(ctx context.Context, transactionTypeID string) (*apiTypes.ExchangeRate, error) {
	resp, err := r.PricingService.GetExchangeRate(ctx, &pricing.GetExchangeRateRequest{TransactionTypeId: transactionTypeID})
	if err != nil {
		return nil, err
	}

	return r.helper.MakeExchangeRateFromProto(resp.ExchangeRate), nil
}

// Me is the resolver for the me field.
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
			Title:     r.helper.MapProtoCustomerTitle(appCustomer.Title),
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
			HasPin:     appCustomer.HasPin,
			PinBlocked: appCustomer.PinBlocked,
			Status:     helpers.MapProtoCustomerStatuses(appCustomer.Status),
			StatusTs:   appCustomer.StatusTs.AsTime().Unix(),
			Ts:         appCustomer.Ts.AsTime().Unix(),
		}, nil
	}

	return apiTypes.Customer{}, errors.New("unknown error occurred")
}

// Customer is the resolver for the customer field.
func (r *queryResolver) Customer(ctx context.Context, id string) (*apiTypes.Customer, error) {
	result, err := r.CustomerService.GetCustomer(ctx, &customer.GetCustomerRequest{Id: id})
	if err != nil {
		return &apiTypes.Customer{}, err
	}

	helpers := &helpersfactory{}

	return helpers.makeCustomerFromProto(result), nil
}

// Customers is the resolver for the customers field.
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

// Cdds is the resolver for the cdds field.
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

// StaffAuditLogs is the resolver for the staffAuditLogs field.
func (r *queryResolver) StaffAuditLogs(ctx context.Context, first *int64, after *string, last *int64, before *string, types []apiTypes.StaffAuditLogType) (*apiTypes.StaffAuditLogConnection, error) {
	// Build request
	var request customer.GetStaffAuditLogsRequest

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

	if len(types) > 0 {
		pTypes := make([]protoTypes.StaffAuditLog_StaffAuditLogTypes, len(types))
		for i, t := range types {
			pTypes[i] = r.helper.MapProtoStaffAuditLogType(t)
		}
		request.Types = pTypes
	}

	// Make call
	resp, err := r.CustomerService.GetStaffAuditLogs(ctx, &request)
	if err != nil {
		return nil, err
	}

	// Build response
	nodes := make([]*apiTypes.StaffAuditLog, len(resp.Nodes))
	helper := helpersfactory{}

	for i, node := range resp.Nodes {
		nodes[i] = helper.MakeStaffAuditLogFromProto(node)
	}

	pageInfo := apiTypes.PageInfo{
		HasNextPage:     resp.PaginationInfo.HasNextPage,
		HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
		StartCursor:     &resp.PaginationInfo.StartCursor,
		EndCursor:       &resp.PaginationInfo.EndCursor,
	}

	return &apiTypes.StaffAuditLogConnection{
		Nodes: nodes, PageInfo: &pageInfo,
		TotalCount: int64(resp.TotalCount)}, nil
}

// Statement is the resolver for the statement field.
func (r *queryResolver) Statement(ctx context.Context, accountID string, startDate string, endDate string, transactionPassword string) (*apiTypes.StatementResponse, error) {
	const dateTemplate = "02-01-2006"

	// Build request
	request := &account.GetAccountStatementRequest{
		AccountId:           accountID,
		TransactionPassword: transactionPassword,
	}

	if startDate != "" {
		formartedStartDate, err := time.Parse(dateTemplate, startDate)
		if err != nil {
			return nil, err
		}

		request.StartDate = timestamppb.New(formartedStartDate)
	}

	if endDate != "" {
		formartedEndDate, err := time.Parse(dateTemplate, endDate)
		if err != nil {
			return nil, err
		}

		request.EndDate = timestamppb.New(formartedEndDate)
	}

	// Execute RPC call

	r.logger.Info(fmt.Sprintf("Executing Get Account Statement RPC for the following range %s - %s", startDate, endDate))

	resp, err := r.AccountService.GetAccountStatement(ctx, request)

	if err != nil {
		r.logger.Error(
			"An Error Occur While Generating Statement",
			zap.Error(err),
		)
		return nil, err
	}

	// Build response
	response := &apiTypes.StatementResponse{
		AccountID:  resp.AccountId,
		PDFContent: &resp.PdfContent,
		StartDate:  resp.StartDate.AsTime().Format(dateTemplate),
		EndDate:    resp.EndDate.AsTime().Format(dateTemplate),
	}

	return response, nil
}

// Faqs is the resolver for the faqs field.
func (r *queryResolver) Faqs(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, filter apiTypes.FilterType) (*apiTypes.FAQConnection, error) {
	helper := helpersfactory{}

	request := r.paginationDetails(keywords, first, after, last, before, filter)
	if filter != "" {
		request.Filter = helper.MapProtoFAQTypes(filter)
	}

	resp, err := r.OnBoardingService.GetFAQs(ctx, &request)
	if err != nil {
		return nil, err
	}

	nodes := make([]*apiTypes.Faq, 0)
	for _, node := range resp.Nodes {
		faq_ := helper.makeFAQFromProto(node)
		nodes = append(nodes, faq_)
	}
	pageInfo := apiTypes.PageInfo{
		HasNextPage:     resp.PaginationInfo.HasNextPage,
		HasPreviousPage: resp.PaginationInfo.HasPreviousPage,
		StartCursor:     &resp.PaginationInfo.StartCursor,
		EndCursor:       &resp.PaginationInfo.EndCursor,
	}
	return &apiTypes.FAQConnection{
		Nodes:      nodes,
		PageInfo:   &pageInfo,
		TotalCount: int64(resp.TotalCount),
	}, nil
}

func (r *Resolver) paginationDetails(keywords *string, first *int64, after *string, last *int64, before *string, filter apiTypes.FilterType) onboarding.GetFAQRequest {
	var request onboarding.GetFAQRequest
	//var helper helpersfactory
	if keywords != nil {
		request.Keywords = *keywords
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
	return request
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
