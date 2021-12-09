package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ms.api/graph/generated"
	"ms.api/protos/pb/customer"
	"ms.api/types"
)

func (r *mutationResolver) RequestOtp(ctx context.Context, typeArg types.DeliveryMode, target string, expireTimeInSeconds *int64) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, target string, otpToken string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) Signup(ctx context.Context, customer types.CustomerInput) (*types.AuthResponse, error) {
	msg := "Not implemented"
	return &types.AuthResponse{
		Message: &msg,
		Code:    int64(500),
	}, nil
}

func (r *mutationResolver) ResetLoginPassword(ctx context.Context, otpToken string, email string, loginPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CheckCustomerEmail(ctx context.Context, email string, device types.DeviceInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CheckCustomerData(ctx context.Context, customerData types.CheckCustomerDataInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) Register(ctx context.Context, customerDetails types.CustomerDetailsInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) SubmitCdd(ctx context.Context, cdd types.CDDInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) AnswerQuestionary(ctx context.Context, questionary types.QuestionaryAnswerInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) AcceptContent(ctx context.Context, contentID string) (*types.Response, error) {
	res, err := r.CustomerService.SetAcceptance(ctx, &customer.SetAcceptanceRequest{ContentId: contentID})
	if err != nil {
		return nil, err
	}

	message := "Successful"
	return &types.Response{
		Message: &message,
		Success: res.Success,
		Code:    int64(res.Code),
	}, nil
}

func (r *mutationResolver) SetTransactionPassword(ctx context.Context, password string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) ResetTransactionPassword(ctx context.Context, otpToken string, email string, newTransactionPassword string, currentTransactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, credentials types.AuthInput) (*types.AuthResponse, error) {
	msg := "Not implemented"
	return &types.AuthResponse{
		Message: &msg,
		Code:    int64(500),
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, token string) (*types.AuthResponse, error) {
	msg := "Not implemented"
	return &types.AuthResponse{
		Message: &msg,
		Code:    int64(500),
	}, nil
}

func (r *mutationResolver) SetDeviceToken(ctx context.Context, tokens []*types.DeviceTokenInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) SetDevicePreferences(ctx context.Context, preferences []*types.DevicePreferencesInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CheckBvn(ctx context.Context, bvn string, phone string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, account types.AccountInput) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateVaultAccount(ctx context.Context, account types.VaultAccountInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateBeneficiary(ctx context.Context, beneficiary types.BeneficiaryInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) AddBeneficiaryAccount(ctx context.Context, beneficiaryID string, account types.BeneficiaryAccountInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) DeleteBeneficaryAccount(ctx context.Context, beneficiaryID string, accountID string, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) CreateTransfer(ctx context.Context, transfer types.TransactionInput, transactionPassword string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) RequestResubmit(ctx context.Context, customerID string, reportIds []string, message *string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) StaffLogin(ctx context.Context, token string, authType types.AuthType) (*types.AuthResponse, error) {
	msg := "Not implemented"
	return &types.AuthResponse{
		Message: &msg,
		Code:    int64(500),
	}, nil
}

func (r *mutationResolver) UpdateKYCStatus(ctx context.Context, id string, status types.KYCStatuses, message string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

func (r *mutationResolver) UpdateAMLStatus(ctx context.Context, id string, status types.AMLStatuses, message string) (*types.Response, error) {
	msg := "Not implemented"
	return &types.Response{
		Message: &msg,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
