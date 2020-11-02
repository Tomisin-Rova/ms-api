package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	rerrors "ms.api/libs/errors"
	"ms.api/protos/pb/cddService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

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
		reasons := make([]*string, 0, len(document.Reasons))
		for _, r := range document.Reasons {
			reasons = append(reasons, &r)
		}
		documents = append(documents, &types.CDDSummaryDocument{
			Name:    document.Name,
			Status:  document.Status,
			Reasons: reasons,
		})
	}

	output.Status = resp.Status
	output.Documents = documents
	return output, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
