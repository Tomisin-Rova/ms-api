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

func (r *subscriptionResolver) GetKYCApplicationResult(ctx context.Context, applicantID string) (resultChan <-chan *types.Result, err error) {
	payload := kycService.ApplicationRequest{
		ApplicantId: applicantID,
	}
	response, err := r.kycClient.GetKYCCheckStatus(ctx, &payload)
	if err != nil {
		return nil, err
	}

	for {
		check, err := response.Recv()
		fmt.Print(check)
		if err != nil {
			return nil, err
		}

		if check != nil {
			var result types.Result

			ch := make(chan *types.Result)

			result.Success = check.Success
			result.Message = check.Message

			ch <- &result
			resultChan = ch
		}
	}
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
