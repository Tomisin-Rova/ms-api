package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
)

func (r *subscriptionResolver) GetKYCApplicationResult(ctx context.Context, applicantID string) (<-chan *kycService.Cdd, error) {
	payload := kycService.ApplicationIdRequest{
		ApplicationId: applicantID,
	}
	response, err := r.kycClient.AwaitCDDReport(ctx, &payload)
	if err != nil {
		return nil, err
	}

	ch := make(chan *kycService.Cdd)
	go func(response kycService.KycService_AwaitCDDReportClient, ch chan *kycService.Cdd) {
		for {
			cdd, err := response.Recv()
			if err != nil {
				return
			}

			if cdd != nil {
				// Disconnect client here after they've received their data.
				ch <- cdd
			}
		}
	}(response, ch)
	return ch, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
