package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ms.api/graph/generated"
	"ms.api/types"
)

func (r *mutationResolver) SubmitLiveVideo(ctx context.Context, id string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PingKYCService(ctx context.Context, message string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
