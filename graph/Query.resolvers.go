package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/onfidoService"
)

func (r *queryResolver) GetApplicantSDKToken(ctx context.Context) (*onfidoService.ApplicantSDKTokenResponse, error) {
	// TODO: Get person's profile from JWT Token.
	// TODO: Use person's ID to get their applicant_id from kyc service.

	// Sample Valid Payload Hard-Coded.

	payload := onfidoService.ApplicantSDKTokenRequest{
		ApplicantId: "f429d65a-331f-4199-bf59-a74c75266aed",
	}
	return r.onfidoClient.GenerateApplicantSDKToken(context.Background(), &payload)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
