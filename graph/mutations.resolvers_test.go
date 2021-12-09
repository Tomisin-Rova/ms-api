package graph

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"ms.api/mocks"
	"ms.api/types"
)

func TestMutationResolver_RequestOtp(t *testing.T) {
	verificationServiceClient := new(mocks.VerificationServiceClient)
	resolverOpts := &ResolverOpts{
		VerificationService: verificationServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	expire := int64(3600)

	resp, err := resolver.RequestOtp(context.Background(), types.DeliveryMode(""), "", &expire)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_VerifyOtp(t *testing.T) {
	verificationServiceClient := new(mocks.VerificationServiceClient)
	resolverOpts := &ResolverOpts{
		VerificationService: verificationServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.VerifyOtp(context.Background(), "", "")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_Signup(t *testing.T) {
	authServiceClient := new(mocks.AuthServiceClient)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()

	resp, err := resolver.Signup(context.Background(), types.CustomerInput{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_ResetLoginPassword(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.ResetLoginPassword(context.Background(), "", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CheckCustomerEmail(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}

	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CheckCustomerEmail(context.Background(), "", types.DeviceInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CheckCustomerData(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CheckCustomerData(context.Background(), types.CheckCustomerDataInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_Register(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.Register(context.Background(), types.CustomerDetailsInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_SubmitCdd(t *testing.T) {
	onboardingServiceClient := new(mocks.OnboardingServiceClient)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.SubmitCdd(context.Background(), types.CDDInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_AnswerQuestionary(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.AnswerQuestionary(context.Background(), types.QuestionaryAnswerInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_SetTransactionPassword(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.SetTransactionPassword(context.Background(), "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_ResetTransactionPassword(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.ResetTransactionPassword(context.Background(), "", "", "", "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_Login(t *testing.T) {
	authServiceClient := new(mocks.AuthServiceClient)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.Login(context.Background(), types.AuthInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_RefreshToken(t *testing.T) {
	authServiceClient := new(mocks.AuthServiceClient)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.RefreshToken(context.Background(), "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_SetDeviceToken(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.SetDeviceToken(context.Background(), []*types.DeviceTokenInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_SetDevicePreferences(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}

	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.SetDevicePreferences(context.Background(), []*types.DevicePreferencesInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CheckBvn(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CheckBvn(context.Background(), "", "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateAccount(t *testing.T) {
	accountServiceClient := new(mocks.AccountServiceClient)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateAccount(context.Background(), types.AccountInput{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateVaultAccount(t *testing.T) {
	accountServiceClient := new(mocks.AccountServiceClient)
	resolverOpts := &ResolverOpts{
		AccountService: accountServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateVaultAccount(context.Background(), types.VaultAccountInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateBeneficiary(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateBeneficiary(context.Background(), types.BeneficiaryInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_AddBeneficiaryAccount(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.AddBeneficiaryAccount(context.Background(), "", types.BeneficiaryAccountInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_DeleteBeneficaryAccount(t *testing.T) {
	customerServiceClient := new(mocks.CustomerServiceClient)
	resolverOpts := &ResolverOpts{
		CustomerService: customerServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.DeleteBeneficaryAccount(context.Background(), "", "", "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_CreateTransfer(t *testing.T) {
	paymentServiceClient := new(mocks.PaymentServiceClient)
	resolverOpts := &ResolverOpts{
		PaymentService: paymentServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.CreateTransfer(context.Background(), types.TransactionInput{}, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_RequestResubmit(t *testing.T) {
	onboardingServiceClient := new(mocks.OnboardingServiceClient)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}

	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	message := ""

	resp, err := resolver.RequestResubmit(context.Background(), "", []string{}, &message)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_StaffLogin(t *testing.T) {
	authServiceClient := new(mocks.AuthServiceClient)
	resolverOpts := &ResolverOpts{
		AuthService: authServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.StaffLogin(context.Background(), "", types.AuthTypeGoogle)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_UpdateKYCStatus(t *testing.T) {
	onboardingServiceClient := new(mocks.OnboardingServiceClient)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.UpdateKYCStatus(context.Background(), "", types.KYCStatusesApproved, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestMutationResolver_UpdateAMLStatus(t *testing.T) {
	onboardingServiceClient := new(mocks.OnboardingServiceClient)
	resolverOpts := &ResolverOpts{
		OnboardingService: onboardingServiceClient,
	}
	resolver := NewResolver(resolverOpts, zaptest.NewLogger(t)).Mutation()
	resp, err := resolver.UpdateAMLStatus(context.Background(), "", types.AMLStatusesPending, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
