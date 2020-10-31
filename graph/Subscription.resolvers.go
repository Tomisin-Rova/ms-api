package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/99designs/gqlgen/handler"
	"ms.api/graph/generated"
	"ms.api/protos/pb/onboardingService"
	"ms.api/types"
)

func (r *subscriptionResolver) CreateApplication(ctx context.Context) (<-chan *types.CreateApplicationResponse, error) {
	personId, err := r.validateToken(ctx)
	if err != nil {
		return nil, err
	}
	stream, err := r.onBoardingService.CreateApplication(context.Background(), &onboardingService.CreateApplicationRequest{PersonId: personId})
	if err != nil {
		r.logger.WithError(err).Error("onBoarding.createApplication() failed")
		return nil, err
	}
	respChan := make(chan *types.CreateApplicationResponse, 1)
	go func(ss onboardingService.OnBoardingService_CreateApplicationClient) {
		for {
			rr, err := ss.Recv()
			if err != nil {
				return
			}

			if err == io.EOF {
				break
			}
			r.logger.WithError(err).Info("error reading stream from sever")
			if rr.Token != "" {
				respChan <- &types.CreateApplicationResponse{Token: rr.Token}
				break
			}
		}
	}(stream)
	return respChan, nil
}


func (r *subscriptionResolver) validateToken(ctx context.Context) (string, error) {
	bearerToken := handler.GetInitPayload(ctx).Authorization()
	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 {
		r.logger.WithField("token_parts", parts).Info("invalid token supplied")
		return "", errors.New("invalid authorization token")
	}

	personId, err := r.authMw.ValidateToken(parts[1])
	if err != nil {
		r.logger.WithError(err).Error("failed to authorize account")
		return "", err
	}
	return personId, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
