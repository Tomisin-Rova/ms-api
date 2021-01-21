package graph

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	emailvalidator "ms.api/libs/validator/email"
	"testing"

	"ms.api/fakes"
	"ms.api/mocks"

	"github.com/roava/zebra/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	onboarding "ms.api/protos/pb/onboardingService"
	verify "ms.api/protos/pb/verifyService"
	"ms.api/types"
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

func TestMutationResolver_VerifyEmail(t *testing.T) {
	const (
		verifyEmailCorrectly = iota
		invalidEmail
		errorOnVerifyEmailMagicLInkCall
	)

	var tests = []struct {
		name    string
		request struct {
			email             string
			verificationToken string
		}
		testType int
	}{
		{
			name: "Test error on VerifyEmailMagicLInk call",
			request: struct {
				email             string
				verificationToken string
			}{email: "test@email.com", verificationToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"},
			testType: errorOnVerifyEmailMagicLInkCall,
		},
		{
			name: "Test invalid email",
			request: struct {
				email             string
				verificationToken string
			}{email: "invalidEmail", verificationToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"},
			testType: invalidEmail,
		},
		{
			name: "Test verify email correctly",
			request: struct {
				email             string
				verificationToken string
			}{email: "test@email.com", verificationToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"},
			testType: verifyEmailCorrectly,
		},
	}
	for _, testCase := range tests {
		onboardingService := new(mocks.OnBoardingServiceClient)

		t.Run(testCase.name, func(t *testing.T) {
			r := NewResolver(&ResolverOpts{
				OnBoardingService: onboardingService,
			}, zaptest.NewLogger(t))
			mu := &mutationResolver{r}

			switch testCase.testType {
			case invalidEmail:
				_, err := mu.VerifyEmail(context.Background(), testCase.request.email, testCase.request.verificationToken)
				assert.Error(t, err)
				assert.Equal(t, emailvalidator.ErrInvalidEmail, err)
			case errorOnVerifyEmailMagicLInkCall:
				onboardingService.On("VerifyEmailMagicLInk", mock.Anything, &onboarding.VerifyEmailMagicLInkRequest{
					Email:             testCase.request.email,
					VerificationToken: testCase.request.verificationToken,
				}).Return(nil, errors.New("error"))

				_, err := mu.VerifyEmail(context.Background(), testCase.request.email, testCase.request.verificationToken)
				assert.Error(t, err)
			case verifyEmailCorrectly:
				onboardingService.On("VerifyEmailMagicLInk", mock.Anything, &onboarding.VerifyEmailMagicLInkRequest{
					Email:             testCase.request.email,
					VerificationToken: testCase.request.verificationToken,
				}).Return(&onboarding.SuccessResponse{
					Message: "success",
				}, nil)

				response, err := mu.VerifyEmail(context.Background(), testCase.request.email, testCase.request.verificationToken)
				assert.NoError(t, err)
				assert.NotEmpty(t, response)
				assert.Equal(t, &types.Result{
					Success: true,
					Message: "success",
				}, response)
			}

			onboardingService.AssertExpectations(t)
		})
	}
}
