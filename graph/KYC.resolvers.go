package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
	"ms.api/types"
)

func (r *applicationResolver) Applicant(ctx context.Context, obj *kycService.Application) (*types.Applicant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *applicationResolver) TimeCreated(ctx context.Context, obj *kycService.Application) (int64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *applicationResolver) TimeUpdated(ctx context.Context, obj *kycService.Application) (int64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *cDDResolver) TimeCreated(ctx context.Context, obj *kycService.Cdd) (int64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *cDDResolver) TimeUpdated(ctx context.Context, obj *kycService.Cdd) (int64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Application returns generated.ApplicationResolver implementation.
func (r *Resolver) Application() generated.ApplicationResolver { return &applicationResolver{r} }

// CDD returns generated.CDDResolver implementation.
func (r *Resolver) CDD() generated.CDDResolver { return &cDDResolver{r} }

type applicationResolver struct{ *Resolver }
type cDDResolver struct{ *Resolver }
