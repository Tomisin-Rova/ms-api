package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ms.api/api/generated"
	"ms.api/models"
)

func (r *mutationResolver) SubmitLiveVideo(ctx context.Context, id primitive.ObjectID) (*models.Result, error) {
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
