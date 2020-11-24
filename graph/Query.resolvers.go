package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	rerrors "ms.api/libs/errors"
	"ms.api/protos/pb/cddService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *queryResolver) GetCDDReportSummary(ctx context.Context) (*types.CDDSummary, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.cddService.GetCDDSummaryReport(ctx, &cddService.PersonIdRequest{PersonId: personId})
	if err != nil {
		return nil, rerrors.NewFromGrpc(err)
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

func (r *queryResolver) GetCountries(ctx context.Context) (*types.FetchCountriesResponse, error) {
	resp, err := r.onBoardingService.FetchCountries(ctx, &onboardingService.FetchCountriesRequest{})
	if err != nil {
		return nil, rerrors.NewFromGrpc(err)
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
			ISO4217CurrencyNumericCode:    c.ISO4217CurrencyNumericCode,
			IsIndependent:                 c.IsIndependent,
			Languages:                     c.Languages,
			OfficialNameEnglish:           c.OfficialNameEnglish,
		})
	}

	countriesRes.Countries = countries
	return countriesRes, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
