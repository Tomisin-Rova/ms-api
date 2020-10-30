package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/types"
)

func (r *subscriptionResolver) GetKYCApplicationResult(ctx context.Context, applicantID string) (<-chan *types.Cdd, error) {
	//payload := kycService.PersonIdRequest{
	//	PersonId: applicantID,
	//}
	//response, err := r.kycClient.AwaitCDDReport(ctx, &payload)
	//if err != nil {
	//	return nil, err
	//}

	//ch := make(chan *types.Cdd)
	//go func(response kycService.KycService_AwaitCDDReportClient, ch chan *types.Cdd) {
	//	for {
	//		cdd, err := response.Recv()
	//		if err != nil {
	//			return
	//		}
	//
	//		if cdd != nil {
	//			// Disconnect client here after they've received their data.
	//			ch <- &types.Cdd{
	//				ID:          cdd.Id,
	//				Owner:       cdd.Owner,
	//				Details:     cdd.Details,
	//				Status:      cdd.Status,
	//				Kyc:         cdd.Kyc,
	//				TimeCreated: cdd.TimeCreated,
	//				TimeUpdated: cdd.TimeUpdated,
	//			}
	//		}
	//	}
	//}(response, ch)
	//return ch, nil

	panic("Awaiting new implementation based on new kyc/onfido codebase")
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
