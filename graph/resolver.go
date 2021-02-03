package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/roava/zebra/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/cddService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/onfidoService"
	"ms.api/protos/pb/payeeService"
	"ms.api/protos/pb/paymentService"
	"ms.api/protos/pb/personService"
	"ms.api/protos/pb/productService"
	"ms.api/protos/pb/verifyService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var (
	ErrUnAuthenticated = errors.NewTerror(
		7012, "InvalidOrExpiredTokenError", "user not authenticated", "user not authenticated")
)

//nolint
func (r *mutationResolver) validateAddress(addr types.AddressInput) error {
	if *addr.Country == "" {
		return errors.NewTerror(7013, "InvalidCountryData", "country data is missing from address", "")
	}
	if *addr.City == "" {
		return errors.NewTerror(7014, "InvalidCityData", "city data is missing from address", "")
	}
	if *addr.Street == "" {
		return errors.NewTerror(7015, "InvalidStreetData", "street data is missing from address", "")
	}
	return nil
}

type ResolverOpts struct {
	PayeeService      payeeService.PayeeServiceClient
	OnfidoClient      onfidoService.OnfidoServiceClient
	cddClient         cddService.CddServiceClient
	productService    productService.ProductServiceClient
	OnBoardingService onboardingService.OnBoardingServiceClient
	verifyService     verifyService.VerifyServiceClient
	AuthService       authService.AuthServiceClient
	paymentService    paymentService.PaymentServiceClient
	AuthMw            *middlewares.AuthMiddleware
	personService     personService.PersonServiceClient
}

type Resolver struct {
	PayeeService      payeeService.PayeeServiceClient
	cddService        cddService.CddServiceClient
	onBoardingService onboardingService.OnBoardingServiceClient
	productService    productService.ProductServiceClient
	personService     personService.PersonServiceClient
	verifyService     verifyService.VerifyServiceClient
	onfidoClient      onfidoService.OnfidoServiceClient
	authService       authService.AuthServiceClient
	paymentService    paymentService.PaymentServiceClient
	authMw            *middlewares.AuthMiddleware
	logger            *zap.Logger
}

func NewResolver(opt *ResolverOpts, logger *zap.Logger) *Resolver {
	return &Resolver{
		PayeeService:      opt.PayeeService,
		cddService:        opt.cddClient,
		onBoardingService: opt.OnBoardingService,
		verifyService:     opt.verifyService,
		onfidoClient:      opt.OnfidoClient,
		authService:       opt.AuthService,
		authMw:            opt.AuthMw,
		paymentService:    opt.paymentService,
		productService:    opt.productService,
		personService:     opt.personService,
		logger:            logger,
	}
}

func ConnectServiceDependencies(secrets *config.Secrets) (*ResolverOpts, error) {
	opts := &ResolverOpts{}

	// OnBoarding
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err := dialRPC(ctx, secrets.OnboardingServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.OnboardingServiceURL)
	}
	opts.OnBoardingService = onboardingService.NewOnBoardingServiceClient(connection)

	// OnFido
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.OnfidoServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.OnfidoServiceURL)
	}
	opts.OnfidoClient = onfidoService.NewOnfidoServiceClient(connection)

	// CDD
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.CddServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.CddServiceURL)
	}
	opts.cddClient = cddService.NewCddServiceClient(connection)

	// Verify
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.VerifyServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.VerifyServiceURL)
	}
	opts.verifyService = verifyService.NewVerifyServiceClient(connection)

	// Auth
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.AuthServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.AuthServiceURL)
	}
	opts.AuthService = authService.NewAuthServiceClient(connection)

	// Product
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.ProductServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.ProductServiceURL)
	}
	opts.productService = productService.NewProductServiceClient(connection)

	// Payment
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.PaymentServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.PaymentServiceURL)
	}
	opts.paymentService = paymentService.NewPaymentServiceClient(connection)

	// Person
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.PersonServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.PersonServiceURL)
	}
	opts.personService = personService.NewPersonServiceClient(connection)

	return opts, nil
}

func dialRPC(ctx context.Context, address string) (*grpc.ClientConn, error) {
	connection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return connection, nil
}
