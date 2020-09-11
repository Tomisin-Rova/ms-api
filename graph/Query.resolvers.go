package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onfidoService"
)

func (r *queryResolver) GetApplicantSDKToken(ctx context.Context, personID string) (*onfidoService.ApplicantSDKTokenResponse, error) {
	// TODO: Get person's profile from JWT Token.
	//person, _ := middlewares.GetAuthenticatedUser(ctx)
	// Sample Hard Coded PersonID
	application, err := r.kycClient.GetApplicationByPersonId(ctx, &kycService.PersonIdRequest{PersonId: personID})
	if err != nil {
		return nil, err
	}

	applicant := application.Applicant
	// Sample Valid Payload Hard-Coded.
	return r.onfidoClient.GenerateApplicantSDKToken(ctx, &onfidoService.ApplicantSDKTokenRequest{ApplicantId: applicant.ApplicantId})
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
