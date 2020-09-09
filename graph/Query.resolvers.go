package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onfidoService"
)

func (r *queryResolver) GetApplicantSDKToken(ctx context.Context) (*onfidoService.ApplicantSDKTokenResponse, error) {
	// TODO: Get person's profile from JWT Token.
	//person, _ := middlewares.GetAuthenticatedUser(ctx)
	// Sample Hard Coded PersonID
	applicant, err := r.kycClient.GetApplicantByPersonId(ctx, &kycService.ApplicantByPersonIdRequest{PersonId: "X1X2X3X4X5X6X7X8X9X10"})
	if err != nil {
		return nil, err
	}

	// Sample Valid Payload Hard-Coded.
	return r.onfidoClient.GenerateApplicantSDKToken(ctx, &onfidoService.ApplicantSDKTokenRequest{ApplicantId: applicant.Id})
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
