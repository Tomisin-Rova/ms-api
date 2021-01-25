package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"ms.api/graph/generated"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/cddService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/payeeService"
	"ms.api/protos/pb/productService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *queryResolver) GetCDDReportSummary(ctx context.Context) (*types.CDDSummary, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.cddService.GetCDDSummaryReport(ctx, &cddService.PersonIdRequest{PersonId: claims.PersonId})
	if err != nil {
		return nil, err
	}

	output := &types.CDDSummary{
		Status: resp.Status,
	}
	documents := make([]*types.CDDSummaryDocument, 0)
	for _, document := range resp.Documents {
		documents = append(documents, &types.CDDSummaryDocument{
			Name:    document.Name,
			Status:  document.Status,
			Reasons: document.Reasons,
		})
	}

	output.Documents = documents
	return output, nil
}

func (r *queryResolver) Me(ctx context.Context) (*types.Person, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	person, err := r.authService.GetPerson(ctx, &authService.GetPersonRequest{
		PersonId:   claims.PersonId,
		IdentityId: claims.IdentityId,
		DeviceId:   claims.DeviceId,
	})
	if err != nil {
		r.logger.Error("failed to get person", zap.Error(err))
		return nil, err
	}
	p := &types.Person{}
	if err := copier.Copy(p, person); err != nil {
		r.logger.Error("copier failed", zap.Error(err))
		return nil, errors.New("failed to read profile information. please retry")
	}
	p.Dob = person.DOB
	return p, nil
}

func (r *queryResolver) GetCountries(ctx context.Context) (*types.FetchCountriesResponse, error) {
	resp, err := r.onBoardingService.FetchCountries(ctx, &onboardingService.FetchCountriesRequest{})
	if err != nil {
		return nil, err
	}

	countriesRes := &types.FetchCountriesResponse{}
	countries := make([]*types.Country, 0)
	for _, c := range resp.Countries {
		countries = append(countries, &types.Country{
			CountryID:                     c.CountryId,
			Capital:                       c.Capital,
			CountryName:                   c.CountryName,
			Continent:                     c.Continent,
			Dial:                          c.Dial,
			GeoNameID:                     c.GeoNameId,
			ISO4217CurrencyAlphabeticCode: c.ISO4217CurrencyAlphabeticCode,
			ISO4217CurrencyNumericCode:    int64(c.ISO4217CurrencyNumericCode),
			IsIndependent:                 c.IsIndependent,
			Languages:                     c.Languages,
			OfficialNameEnglish:           c.OfficialNameEnglish,
		})
	}

	countriesRes.Countries = countries
	return countriesRes, nil
}

func (r *queryResolver) Reasons(ctx context.Context) (*types.FetchReasonResponse, error) {
	resp, err := r.onBoardingService.FetchReasons(ctx, &onboardingService.EmptyRequest{})
	if err != nil {
		return nil, err
	}

	response := &types.FetchReasonResponse{}
	reasons := make([]*types.Reason, 0)
	for _, r := range resp.Reasons {
		reasons = append(reasons, &types.Reason{
			ID:          r.Id,
			Description: r.Description,
		})
	}

	response.Reasons = reasons
	return response, nil
}

func (r *queryResolver) Accounts(ctx context.Context) (*types.AccountsResult, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.productService.GetAccounts(ctx, &productService.GetAccountRequest{
		PersonId: claims.PersonId,
	})
	if err != nil {
		r.logger.Error("failed to get accounts from product service", zap.Error(err))
		return nil, err
	}
	primaryAccount := &types.Account{
		Currency:       resp.PrimaryAccount.CurrencyCode,
		CurrencySymbol: resp.PrimaryAccount.CurrencySymbol,
		AccountNumber:  resp.PrimaryAccount.AccountNumber,
		AccountName:    resp.PrimaryAccount.AccountName,
		Balance:        resp.PrimaryAccount.Balance,
	}
	accounts := make([]*types.Account, 0)
	for _, next := range resp.Accounts {
		accounts = append(accounts, &types.Account{
			Currency:       next.CurrencyCode,
			CurrencySymbol: next.CurrencySymbol,
			AccountNumber:  next.AccountNumber,
			AccountName:    next.AccountName,
			Balance:        next.Balance,
		})
	}
	return &types.AccountsResult{
		PrimaryAccount:   primaryAccount,
		CurrencyAccounts: accounts,
	}, nil
}

func (r *queryResolver) GetPayeesByPhoneNumbers(ctx context.Context, phone []string) (*types.GetPayeesByPhoneNumbers, error) {
	resp, err := r.PayeeService.GetPayeesByPhoneNumbers(ctx, &payeeService.FetchPayeeByPhoneRequest{Phone: phone})
	if err != nil {
		return nil, err
	}

	response := &types.GetPayeesByPhoneNumbers{}
	payees := make([]*types.Payee, 0)
	for _, r := range resp.Payees {
		payees = append(payees, &types.Payee{
			PersonID:    r.PersonId,
			PhoneNumber: r.Phone,
			FirstName:   r.FirstName,
			LastName:    r.LastName,
		})
	}

	response.Payees = payees
	return response, nil
}

func (r *queryResolver) SupportedCurrencies(ctx context.Context) ([]*types.Country, error) {
	resp, err := r.onBoardingService.FetchCountries(ctx, &onboardingService.FetchCountriesRequest{})
	if err != nil {
		return nil, err
	}

	countriesRes := &types.FetchCountriesResponse{}
	countries := make([]*types.Country, 0)
	for _, c := range resp.Countries {
		if c.CountryName == "US" ||
			c.CountryName == "UK" ||
			c.CountryName == "Nigeria" {
			countries = append(countries, &types.Country{
				CountryID:                     c.CountryId,
				Capital:                       c.Capital,
				CountryName:                   c.CountryName,
				Continent:                     c.Continent,
				Dial:                          c.Dial,
				GeoNameID:                     c.GeoNameId,
				ISO4217CurrencyAlphabeticCode: c.ISO4217CurrencyAlphabeticCode,
				ISO4217CurrencyNumericCode:    int64(c.ISO4217CurrencyNumericCode),
				IsIndependent:                 c.IsIndependent,
				Languages:                     c.Languages,
				OfficialNameEnglish:           c.OfficialNameEnglish,
			})
		}
	}

	countriesRes.Countries = countries
	return countries, nil
}

func (r *queryResolver) GetAddressesByText(ctx context.Context, text string) (*types.FetchAddressesResponse, error) {
	resp, err := r.onBoardingService.GetAddressesByText(ctx, &onboardingService.GetAddressesRequest{Text: text})
	if err != nil {
		return nil, err
	}

	fetchAddressRes := &types.FetchAddressesResponse{}
	addresses := make([]*types.AddressResult, 0)
	for _, c := range resp.Addresses {
		addresses = append(addresses, &types.AddressResult{
			Addressline1:      c.Addressline1,
			Addressline2:      c.Addressline2,
			Summaryline:       c.Summaryline,
			Organisation:      c.Organisation,
			Buildingname:      c.Buildingname,
			Premise:           c.Premise,
			Street:            c.Street,
			Dependentlocality: c.Dependentlocality,
			Posttown:          c.Posttown,
			County:            c.County,
			Postcode:          c.Postcode,
			Latitude:          c.Latitude,
			Longitude:         c.Longitude,
			Grideasting:       c.Grideasting,
			Gridnorthing:      c.Gridnorthing,
		})

	}

	fetchAddressRes.Addresses = addresses
	return fetchAddressRes, nil
}

func (r *queryResolver) CheckEmail(ctx context.Context, email string) (*bool, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}
	resp, err := r.onBoardingService.CheckEmailExistence(ctx, &onboardingService.CheckEmailExistenceRequest{Email: email})
	if err != nil {
		r.logger.Info("onboardingService.checkEmailExistence() failed", zap.Error(err))
		return nil, err
	}
	return &resp.Exists, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
