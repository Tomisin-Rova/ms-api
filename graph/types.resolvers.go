package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/types"
)

func (r *cDDResolver) Validations(ctx context.Context, obj *types.Cdd, validationType *types.ValidationType, status []types.State) ([]*types.Validation, error) {
	filteredValidations := obj.Validations
	if validationType != nil {
		filteredValidations = make([]*types.Validation, 0)
		for _, validation := range obj.Validations {
			if validation.ValidationType == *validationType {
				filteredValidations = append(filteredValidations, validation)
			}
		}
	}
	if len(status) > 0 {
		n := 0
	loopStatusFilter:
		for _, v := range filteredValidations {
			for _, s := range status {
				if s == v.Status {
					filteredValidations[n] = v
					n += 1
					continue loopStatusFilter
				}
			}
		}
		filteredValidations = filteredValidations[:n]
	}

	return filteredValidations, nil
}

// CDD returns generated.CDDResolver implementation.
func (r *Resolver) CDD() generated.CDDResolver { return &cDDResolver{r} }

type cDDResolver struct{ *Resolver }
