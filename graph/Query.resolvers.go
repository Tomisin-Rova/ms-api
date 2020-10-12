package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onfidoService"
	"ms.api/server/http/middlewares"
)

func (r *queryResolver) GetApplicantSDKToken(ctx context.Context) (*onfidoService.ApplicantSDKTokenResponse, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	applicant, err := r.kycClient.GetKycApplicantByPersonId(ctx, &kycService.PersonIdRequest{PersonId: personId})
	if err != nil {
		return nil, err
	}
	return r.onfidoClient.GenerateApplicantSDKToken(ctx, &onfidoService.ApplicantSDKTokenRequest{ApplicantId: applicant.ApplicantId})
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
