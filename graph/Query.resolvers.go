package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/onfidoService"
)

func (r *queryResolver) GetApplicantSDKToken(ctx context.Context, applicantID string, applicationID string) (*onfidoService.ApplicantSDKTokenResponse, error) {
	payload := onfidoService.ApplicantSDKTokenRequest{
		ApplicantId:   applicantID,
		ApplicationId: applicationID,
	}
	return r.onfidoClient.GenerateApplicantSDKToken(context.Background(), &payload)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
