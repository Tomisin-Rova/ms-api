package graph

import (
	"errors"
	"testing"

	"ms.api/mocks"
	"ms.api/protos/pb/onboardingService"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/types"

	coreErrors "github.com/roava/zebra/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"golang.org/x/net/context"
)

func TestMutationResolver_CreatePhone(t *testing.T) {
	const (
		success = iota
		errorInvalidPhone
		errorOnboardingSvcCreatePhone
	)

	var tests = []struct {
		name string
		args struct {
			phone  string
			device types.DeviceInput
		}
		testType int
	}{
		{
			name: "Test create phone successfully",
			args: struct {
				phone  string
				device types.DeviceInput
			}{
				phone: "5522552255",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: success,
		},
		{
			name: "Test invalid phone",
			args: struct {
				phone  string
				device types.DeviceInput
			}{
				phone: "invalidPhone",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: errorInvalidPhone,
		},
		{
			name: "Test error calling OnboardingService.CreatePhone()",
			args: struct {
				phone  string
				device types.DeviceInput
			}{
				phone: "5522552255",
				device: types.DeviceInput{
					Identifier: "testIdentifier",
					Brand:      "testBrand",
					Os:         "testOs",
				},
			},
			testType: errorOnboardingSvcCreatePhone,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			onBoardingServiceClient := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onBoardingServiceClient,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case errorInvalidPhone:
				response, err := mutationResolver.CreatePhone(context.Background(), testCase.args.phone, testCase.args.device)

				assert.Error(t, err)
				assert.Equal(t, 7010, err.(*coreErrors.Terror).Code())
				assert.Nil(t, response)
			case errorOnboardingSvcCreatePhone:
				onBoardingServiceClient.On("CreatePhone", mock.Anything, &onboardingService.CreatePhoneRequest{
					PhoneNumber: testCase.args.phone,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(nil, errors.New(""))

				response, err := mutationResolver.CreatePhone(context.Background(), testCase.args.phone, testCase.args.device)

				assert.Error(t, err)
				assert.Nil(t, response)
			case success:
				onBoardingServiceClient.On("CreatePhone", mock.Anything, &onboardingService.CreatePhoneRequest{
					PhoneNumber: testCase.args.phone,
					Device: &protoTypes.Device{
						Identifier: testCase.args.device.Identifier,
						Brand:      testCase.args.device.Brand,
						Os:         testCase.args.device.Os,
					},
				}).Return(&onboardingService.CreatePhoneResponse{}, nil)

				response, err := mutationResolver.CreatePhone(context.Background(), testCase.args.phone, testCase.args.device)

				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, true, response.Success)
			}

			onBoardingServiceClient.AssertExpectations(t)
		})
	}
}

func TestMutationResolver_ConfirmPhone(t *testing.T) {
	const (
		success = iota
		errorOnboardingSvcVerifySmsOtp
	)

	var tests = []struct {
		name string
		args struct {
			token string
			code  string
		}
		testType int
	}{
		{
			name: "Test confirm phone correctly",
			args: struct {
				token string
				code  string
			}{
				token: "testToken",
				code:  "123456",
			},
			testType: success,
		},
		{
			name: "Test confirm phone correctly",
			args: struct {
				token string
				code  string
			}{
				token: "testToken",
				code:  "123456",
			},
			testType: errorOnboardingSvcVerifySmsOtp,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocks
			onBoardingServiceClient := new(mocks.OnBoardingServiceClient)

			resolver := NewResolver(&ResolverOpts{
				OnBoardingService: onBoardingServiceClient,
			}, zaptest.NewLogger(t))
			mutationResolver := resolver.Mutation()

			switch testCase.testType {
			case errorOnboardingSvcVerifySmsOtp:
				onBoardingServiceClient.On("VerifySmsOtp", mock.Anything, &onboardingService.OtpVerificationRequest{
					Token: testCase.args.token, Code: testCase.args.code,
				}).Return(nil, errors.New(""))

				response, err := mutationResolver.ConfirmPhone(context.Background(), testCase.args.token, testCase.args.code)

				assert.Error(t, err)
				assert.Nil(t, response)
			case success:
				onBoardingServiceClient.On("VerifySmsOtp", mock.Anything, &onboardingService.OtpVerificationRequest{
					Token: testCase.args.token, Code: testCase.args.code,
				}).Return(&onboardingService.OtpVerificationResponse{}, nil)

				response, err := mutationResolver.ConfirmPhone(context.Background(), testCase.args.token, testCase.args.code)

				assert.NoError(t, err)
				assert.NotNil(t, response)
			}

			onBoardingServiceClient.AssertExpectations(t)
		})
	}
}
