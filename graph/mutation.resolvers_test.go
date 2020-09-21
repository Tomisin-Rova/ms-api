package graph

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"ms.api/fakes"
	onboarding "ms.api/protos/pb/onboardingService"
	verify "ms.api/protos/pb/verifyService"
	"ms.api/types"
	"testing"
)

func TestMutationResolver_CreatePhone(t *testing.T) {
	svc := fakes.NewFakeOnBoardingClient(&onboarding.SuccessResponse{Message: "phone added"},
		&onboarding.CreatePhoneResponse{Message: "phone added", EmailToken: "token"}, nil)
	r := NewResolver(&ResolverOpts{onBoardingService: svc}, logrus.StandardLogger())
	mu := &mutationResolver{r}
	resp, err := mu.CreatePhone(context.Background(), types.CreatePhoneInput{
		Phone: "09088776655", Device: &types.Device{},
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

func TestMutationResolver_CreatePhone_Error(t *testing.T) {
	svc := fakes.NewFakeOnBoardingClient(nil, nil, errors.New("error occurred"))
	r := NewResolver(&ResolverOpts{onBoardingService: svc}, logrus.StandardLogger())
	mu := &mutationResolver{r}
	resp, err := mu.CreatePhone(context.Background(), types.CreatePhoneInput{
		Phone: "09088776655", Device: &types.Device{},
	})
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestMutationResolver_VerifySmsOtp(t *testing.T) {
	svc := fakes.NewFakeVerifyClient(&verify.OtpVerificationResponse{Message: "phone added", Match: true}, nil, nil)
	r := NewResolver(&ResolverOpts{verifyService: svc}, logrus.StandardLogger())
	mu := &mutationResolver{r}
	resp, err := mu.VerifyOtp(context.Background(), "09088776655", "009988")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

func TestMutationResolver_VerifySmsOtp_Error(t *testing.T) {
	svc := fakes.NewFakeVerifyClient(nil, nil, errors.New("failed to validate OTP"))
	r := NewResolver(&ResolverOpts{verifyService: svc}, logrus.StandardLogger())
	mu := &mutationResolver{r}
	resp, err := mu.VerifyOtp(context.Background(), "09088776655", "009988")
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestMutationResolver_CreateEmail_BadEmail(t *testing.T) {

}
