package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ms.api/graph/generated"
	"ms.api/types"
)

func (r *mutationResolver) SubmitKYCApplication(ctx context.Context, applicationID string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePasscode(ctx context.Context, userID string, passcode string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePersonBiodata(ctx context.Context, input *types.UpdateBioDataInput) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddReasonsForUsingRoava(ctx context.Context, personID string, reasons string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePhone(ctx context.Context, input types.CreatePhoneInput) (*types.CreatePhoneResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, phone string, code string) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateEmail(ctx context.Context, input *types.CreateEmailInput) (*types.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
