package graph

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/libs/mapper"
	"ms.api/libs/preloader"
	devicevalidator "ms.api/libs/validator/device"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/account"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/customer"
	"ms.api/protos/pb/messaging"
	"ms.api/protos/pb/onboarding"
	"ms.api/protos/pb/payment"
	"ms.api/protos/pb/pricing"
	"ms.api/protos/pb/verification"
	"ms.api/server/http/middlewares"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type ResolverOpts struct {
	AccountService      account.AccountServiceClient
	AuthService         auth.AuthServiceClient
	CustomerService     customer.CustomerServiceClient
	OnboardingService   onboarding.OnboardingServiceClient
	MessagingService    messaging.MessagingServiceClient
	PaymentService      payment.PaymentServiceClient
	PricingService      pricing.PricingServiceClient
	VerificationService verification.VerificationServiceClient
	preloader           preloader.Preloader
	mapper              mapper.Mapper
	AuthMw              *middlewares.AuthMiddleware
	EmailValidator      emailvalidator.EmailValidator
	DeviceValidator     devicevalidator.DeviceValidator
	PhoneValidator      phonenumbervalidator.PhoneNumberValidator
	Helper              Helper
}

type Resolver struct {
	AccountService      account.AccountServiceClient
	AuthService         auth.AuthServiceClient
	CustomerService     customer.CustomerServiceClient
	OnBoardingService   onboarding.OnboardingServiceClient
	PaymentService      payment.PaymentServiceClient
	PricingService      pricing.PricingServiceClient
	VerificationService verification.VerificationServiceClient
	MessagingService    messaging.MessagingServiceClient
	preloader           preloader.Preloader
	mapper              mapper.Mapper
	logger              *zap.Logger
	emailValidator      emailvalidator.EmailValidator
	deviceValidator     devicevalidator.DeviceValidator
	phoneValidator      phonenumbervalidator.PhoneNumberValidator
	helper              Helper
}

func NewResolver(opt *ResolverOpts, logger *zap.Logger) *Resolver {
	helper := opt.Helper
	if helper == nil {
		helper = &helpersfactory{}
	}
	return &Resolver{
		AccountService:      opt.AccountService,
		AuthService:         opt.AuthService,
		CustomerService:     opt.CustomerService,
		MessagingService:    opt.MessagingService,
		OnBoardingService:   opt.OnboardingService,
		PaymentService:      opt.PaymentService,
		PricingService:      opt.PricingService,
		VerificationService: opt.VerificationService,
		preloader:           opt.preloader,
		mapper:              opt.mapper,
		logger:              logger,
		emailValidator:      startEmailValidator(opt.EmailValidator),
		deviceValidator:     startDeviceValidator(opt.DeviceValidator),
		phoneValidator:      startPhoneValidator(opt.PhoneValidator),
		helper:              helper,
	}
}

func ConnectServiceDependencies(secrets *config.Secrets) (*ResolverOpts, error) {
	opts := &ResolverOpts{
		preloader: preloader.GQLPreloader{},
		mapper:    mapper.NewMapper(),
	}
	localDevEnvironment := secrets.Service.Environment == config.LocalEnvironment
	// OnBoarding
	if len(secrets.OnboardingServiceURL) > 0 || !localDevEnvironment {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.OnboardingServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.OnboardingServiceURL)
		}
		opts.OnboardingService = onboarding.NewOnboardingServiceClient(connection)
	}

	// Verification
	if len(secrets.VerificationServiceURL) > 0 || !localDevEnvironment {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.VerificationServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.VerificationServiceURL)
		}
		opts.VerificationService = verification.NewVerificationServiceClient(connection)
	}

	// Auth
	if len(secrets.AuthServiceURL) > 0 || !localDevEnvironment {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.AuthServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.AuthServiceURL)
		}
		opts.AuthService = auth.NewAuthServiceClient(connection)
	}

	// Account
	if len(secrets.AccountServiceURL) > 0 || !localDevEnvironment {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.AccountServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.AccountServiceURL)
		}
		opts.AccountService = account.NewAccountServiceClient(connection)
	}

	// Payment
	// TODO: Return local validation once payments service it's refactored
	if len(secrets.PaymentServiceURL) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.PaymentServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.PaymentServiceURL)
		}
		opts.PaymentService = payment.NewPaymentServiceClient(connection)
	}

	// Customer
	if len(secrets.CustomerServiceURL) > 0 || !localDevEnvironment {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.CustomerServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.CustomerServiceURL)
		}
		opts.CustomerService = customer.NewCustomerServiceClient(connection)
	}

	// messaging
	if len(secrets.MessagingServiceURL) > 0 || !localDevEnvironment {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.MessagingServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.MessagingServiceURL)
		}
		opts.MessagingService = messaging.NewMessagingServiceClient(connection)
	}

	// Pricing
	// TODO: Return local validation once pricing service it's refactored
	if len(secrets.PricingServiceURL) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		connection, err := dialRPC(ctx, secrets.PricingServiceURL)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, secrets.PricingServiceURL)
		}
		opts.PricingService = pricing.NewPricingServiceClient(connection)
	}

	return opts, nil
}

func dialRPC(ctx context.Context, address string) (*grpc.ClientConn, error) {
	connection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func startEmailValidator(validator emailvalidator.EmailValidator) emailvalidator.EmailValidator {
	if validator == nil {
		return &emailvalidator.Validator{}
	}
	return validator
}

func startDeviceValidator(validator devicevalidator.DeviceValidator) devicevalidator.DeviceValidator {
	if validator == nil {
		return &devicevalidator.Validator{}
	}
	return validator
}

func startPhoneValidator(validator phonenumbervalidator.PhoneNumberValidator) phonenumbervalidator.PhoneNumberValidator {
	if validator == nil {
		return &phonenumbervalidator.Validator{}
	}
	return validator
}
