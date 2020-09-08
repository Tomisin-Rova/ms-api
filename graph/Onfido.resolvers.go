package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ms.api/graph/generated"
	"ms.api/protos/pb/onfidoService"
)

func (r *applicantSDKTokenRequestResolver) ApplicaitonID(ctx context.Context, obj *onfidoService.ApplicantSDKTokenRequest) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// ApplicantSDKTokenRequest returns generated.ApplicantSDKTokenRequestResolver implementation.
func (r *Resolver) ApplicantSDKTokenRequest() generated.ApplicantSDKTokenRequestResolver {
	return &applicantSDKTokenRequestResolver{r}
}

type applicantSDKTokenRequestResolver struct{ *Resolver }
