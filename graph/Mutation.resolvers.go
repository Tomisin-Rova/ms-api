package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/types"
)

func (r *mutationResolver) SubmitKYCApplication(ctx context.Context, applicantID string) (*types.Result, error) {
	payload := kycService.ApplicationRequest{
		ApplicantId: applicantID,
	}

	res, err := r.kycClient.SubmitKYCApplication(ctx, &payload)
	if err != nil {
		return nil, err
	}
	return &types.Result{
		Success: res.Success,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) CreatePasscode(ctx context.Context, userID string, passcode string) (*types.Result, error) {
	payload := onboardingService.CreatePasscodeRequest{
		PersonId: userID,
		Passcode: passcode,
	}
	res, err := r.onboardingClient.CreatePasscode(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) UpdatePersonBiodata(ctx context.Context, personID string, address string, firstName string, lastName string, dob string) (*types.Result, error) {
	payload := onboardingService.UpdatePersonRequest{
		PersonId:  personID,
		Address:   address,
		FirstName: firstName,
		LastName:  lastName,
		Dob:       dob,
	}
	res, err := r.onboardingClient.UpdatePersonBiodata(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

func (r *mutationResolver) AddReasonsForUsingRoava(ctx context.Context, personID string, reasons string) (*types.Result, error) {
	payload := onboardingService.RoavaReasonsRequest{
		PersonId: personID,
		Reasons:  reasons,
	}
	res, err := r.onboardingClient.AddReasonsForUsingRoava(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	return &types.Result{
		Success: true,
		Message: res.Message,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
