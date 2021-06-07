package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io"
	"time"

	"go.uber.org/zap"
	"ms.api/graph/connections"
	"ms.api/graph/generated"
	"ms.api/graph/models"
	"ms.api/protos/pb/accountService"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/cddService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *subscriptionResolver) Cdds(ctx context.Context, keywords *string, status []types.State, first *int64, after *string, last *int64, before *string) (<-chan *types.CDDConnection, error) {
	msgs := make(chan *types.CDDConnection, 1)
	// TODO: Refactor the cdds query in order to be reused on subscription as well
	go func() {
		dataConverter := NewDataConverter(r.logger)
		perPage := r.perPageCddsQuery(first, after, last, before)

		req := &cddService.CDDSRequest{
			Page:    1,
			PerPage: perPage,
			Status:  dataConverter.StateToStringSlice(status),
		}
		if keywords != nil {
			req.Keywords = *keywords
		}

		stream, err := r.cddService.CDDSStreamed(context.Background(), req)
		if err != nil {
			r.logger.With(zap.Error(err)).Error("failed to get cdds stream")
			return
		}
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				r.logger.With(zap.Error(err)).Error("failed to get stream object")
			}

			//resp, err := r.cddService.CDDS(context.Background(), req)
			if err != nil {
				r.logger.With(zap.Error(err)).Error("failed to fetch cdds")
				break
			}

			cdds := resp.Results
			cddsResult := make([]*types.Cdd, len(cdds))
			for i, cdd := range cdds {
				cddsResult[i] = dataConverter.makeCdd(cdd)
			}

			input := models.ConnectionInput{
				Before: before,
				After:  after,
				First:  first,
				Last:   last,
			}

			edger := func(cdd *types.Cdd, offset int) connections.Edge {
				return types.CDDEdge{
					Node:   cdd,
					Cursor: connections.OffsetToCursor(offset),
				}
			}

			conn := func(edges []*types.CDDEdge, nodes []*types.Cdd, info *types.PageInfo, totalCount int) (*types.CDDConnection, error) {
				var cddNodes []*types.Cdd
				cddNodes = append(cddNodes, nodes...)
				count := int64(totalCount)
				return &types.CDDConnection{
					Edges:      edges,
					Nodes:      cddNodes,
					PageInfo:   info,
					TotalCount: &count,
				}, nil
			}
			cddConn, err := connections.CddLookupCon(cddsResult, edger, conn, input)
			if err != nil {
				r.logger.With(zap.Error(err)).Error("failed to fetch cdds")
				break
			}
			msgs <- cddConn
			time.Sleep(1 * time.Second)
		}
	}()
	return msgs, nil
}

func (r *subscriptionResolver) Cdd(ctx context.Context, id string) (<-chan *types.Cdd, error) {
	_, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	cddChannel := make(chan *types.Cdd, 1)

	go func() {
		stream, err := r.cddService.CDDByIdStreamed(context.Background(), &cddService.CddIdRequest{
			Id: id,
		})
		if err != nil {
			r.logger.With(zap.Error(err)).Error("failed to get cdd stream")
			return
		}
		for {
			cddDto, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				r.logger.With(zap.Error(err)).Error(err.Error())
				break
			}
			dataConverter := NewDataConverter(r.logger)
			cdd := dataConverter.makeCdd(cddDto)
			cddChannel <- cdd
		}
	}()

	return cddChannel, nil
}

func (r *subscriptionResolver) Accounts(ctx context.Context, first *int64, after *string, last *int64, before *string, token string) (<-chan *types.AccountConnection, error) {
	claims, err := r.authService.ValidateToken(context.Background(),
		&authService.ValidateTokenRequest{Token: token})
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	stream, err := r.accountService.GetAccountsStream(ctx, &accountService.GetAccountsRequest{
		IdentityId: claims.IdentityId,
	})

	if err != nil {
		return nil, err
	}

	accountsChan := make(chan *types.AccountConnection, 1)

	go func() {
		for {
			accounts, err := stream.Recv()
			if err != nil {
				break
			}
			accountsRes, err := r.getAccountsResponse(ctx, first, after, last, before, accounts)
			if err != nil {
				break
			}
			accountsChan <- accountsRes
		}
	}()
	return accountsChan, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
