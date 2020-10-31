package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jinzhu/copier"
	"ms.api/graph/generated"
	rerrors "ms.api/libs/errors"
	"ms.api/protos/pb/cddService"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onfidoService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *queryResolver) GetApplicantSDKToken(ctx context.Context) (*onfidoService.ApplicantSDKTokenResponse, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	applicant, err := r.kycClient.GetKycApplicantByPersonId(ctx, &kycService.PersonIdRequest{PersonId: personId})
	if err != nil {
		return nil, err
	}
	return r.onfidoClient.GenerateApplicantSDKToken(ctx, &onfidoService.ApplicantSDKTokenRequest{ApplicantId: applicant.ApplicantId})
}

func (r *queryResolver) GetCddSummary(ctx context.Context) (*types.CddSummary, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}
	summary, err := r.cddService.GetCDDSummaryReport(ctx, &cddService.PersonIdRequest{PersonId: personId})
	if err != nil {
		r.logger.WithError(err).Error("cdd.getSummary() failed")
		return nil, rerrors.NewFromGrpc(err)
	}
	docs := make([]*types.CddDocument, 0, len(summary.Documents))
	for _, next := range summary.Documents {
		docs = append(docs, &types.CddDocument{Name: next.Name, Status: next.Status, Reason: next.Reason})
	}
	return &types.CddSummary{
		Status:    summary.Status,
		Documents: docs,
	}, nil
}

func (r *queryResolver) GetCdd(ctx context.Context, id string) (*types.Cdd, error) {
	cdd, err := r.cddService.GetCDDById(ctx, &cddService.CddIdRequest{Id: id})
	if err != nil {
		r.logger.WithError(err).Error("cdd.getCddById() failed")
		return nil, rerrors.NewFromGrpc(err)
	}
	value := &types.Cdd{}
	if err := copier.Copy(value, cdd); err != nil {
		r.logger.WithError(err).Error("failed to copy values")
		return nil, err
	}
	return value, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
