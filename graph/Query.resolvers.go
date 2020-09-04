package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
)

func (r *queryResolver) HelloWorld(ctx context.Context) (*kycService.Applicant, error) {
	void := new(kycService.Void)
	return r.kycClient.HelloWorld(context.Background(), void)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
