package handlers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/handlers/generated"
	"ms.api/services/kycService"
)

func (r *queryResolver) HelloWorld(ctx context.Context) (*kycService.Applicant, error) {
	void := new(kycService.Void)
	result, err := r.kycClient.HelloWorld(context.Background(), void)
	if err != nil {
		return nil, err
	}

	//if err := utils.Pack(result, &applicant); err != nil {
	//	return nil, err
	//}
	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
