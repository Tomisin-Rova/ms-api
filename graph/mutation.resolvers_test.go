package graph

import (
	"context"
	"errors"
	"testing"

	"ms.api/fakes"
	onboarding "ms.api/protos/pb/onboardingService"
	verify "ms.api/protos/pb/verifyService"
	"ms.api/types"

	"github.com/roava/zebra/logger"
	"github.com/stretchr/testify/assert"
)

func TestMutationResolver_CreatePhone(t *testing.T) {
	svc := fakes.NewFakeOnBoardingClient(&onboarding.SuccessResponse{Message: "phone added"},
		&onboarding.CreatePhoneResponse{Message: "phone added", Token: "token"}, nil, nil)
	r := NewResolver(&ResolverOpts{OnBoardingService: svc}, logger.New())
	mu := &mutationResolver{r}
	resp, err := mu.CreatePhone(context.Background(), types.CreatePhoneInput{
		Phone: "09088776655", Device: &types.Device{},
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

func TestMutationResolver_CreatePhone_Error(t *testing.T) {
	svc := fakes.NewFakeOnBoardingClient(nil, nil, nil, errors.New("error occurred"))
	r := NewResolver(&ResolverOpts{OnBoardingService: svc}, logger.New())
	mu := &mutationResolver{r}
	resp, err := mu.CreatePhone(context.Background(), types.CreatePhoneInput{
		Phone: "09088776655", Device: &types.Device{},
	})
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestMutationResolver_VerifySmsOtp(t *testing.T) {
	svc := fakes.NewFakeVerifyClient(&verify.OtpVerificationResponse{Message: "phone added", Match: true}, nil, nil)
	onBoardingSvc := fakes.NewFakeOnBoardingClient(&onboarding.SuccessResponse{Message: ""},
		&onboarding.CreatePhoneResponse{}, &onboarding.OtpVerificationResponse{Match: true}, nil)
	r := NewResolver(&ResolverOpts{verifyService: svc, OnBoardingService: onBoardingSvc}, logger.New())
	mu := &mutationResolver{r}
	resp, err := mu.VerifyOtp(context.Background(), "09088776655", "009988")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

func TestMutationResolver_VerifySmsOtp_Error(t *testing.T) {
	svc := fakes.NewFakeVerifyClient(nil, nil, errors.New("failed to validate OTP"))
	onBoardingSvc := fakes.NewFakeOnBoardingClient(&onboarding.SuccessResponse{Message: ""},
		&onboarding.CreatePhoneResponse{}, nil, errors.New("failed to perform op"))
	r := NewResolver(&ResolverOpts{verifyService: svc, OnBoardingService: onBoardingSvc}, logger.New())
	mu := &mutationResolver{r}
	resp, err := mu.VerifyOtp(context.Background(), "09088776655", "009988")
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
