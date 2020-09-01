package handlers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ms.api/handlers/generated"
	"ms.api/services/kycService"
	"ms.api/types"
)

func (r *applicantResolver) Address(ctx context.Context, obj *kycService.Applicant) (*types.Address, error) {
	panic(fmt.Errorf("not implemented"))
}

// Applicant returns generated.ApplicantResolver implementation.
func (r *Resolver) Applicant() generated.ApplicantResolver { return &applicantResolver{r} }

type applicantResolver struct{ *Resolver }
