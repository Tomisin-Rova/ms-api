package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
	"ms.api/types"
)

func (r *mutationResolver) SubmitKYCApplication(ctx context.Context, applicationID string) (*types.Result, error) {
	if _, err := r.kycClient.StartApplicationCDD(ctx, &kycService.ApplicationIdRequest{
		ApplicationId: applicationID,
	}); err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: "Successfully started CDD check, you'll be notified once completed.",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
