package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	rerrors "ms.api/libs/errors"
	"ms.api/protos/pb/cddService"

	"ms.api/graph/generated"
	"ms.api/protos/pb/onfidoService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *queryResolver) GetApplicantSDKToken(ctx context.Context) (*onfidoService.ApplicantSDKTokenResponse, error) {
	_, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	//applicant, err := r.kycClient.GetKycApplicantByPersonId(ctx, &kycService.PersonIdRequest{PersonId: personId})
	//if err != nil {
	//	return nil, err
	//}
	//return r.onfidoClient.GenerateApplicantSDKToken(ctx, &onfidoService.ApplicantSDKTokenRequest{ApplicantId: applicant.ApplicantId})

	panic("Awaiting new implementation with new kyc/onfido codebase")
}

func (r *queryResolver) GetCDDReportSummary(ctx context.Context) (*types.CDDSummary, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	resp, err := r.cddService.GetCDDSummaryReport(ctx, &cddService.PersonIdRequest{PersonId: personId})
	if err != nil {
		return nil, rerrors.NewFromGrpc(err)
	}

	output := &types.CDDSummary{}

	documents := make([]*types.CDDSummaryDocument, 0)

	for _, document := range resp.Documents {
		documents = append(documents, &types.CDDSummaryDocument{
			Name:   document.Name,
			Status: document.Status,
			Reason: document.Reason,
		})
	}

	output.Status = resp.Status
	output.Documents = documents

	return output, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
