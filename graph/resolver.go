package graph

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/cddService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/onfidoService"
	"ms.api/protos/pb/verifyService"
	"time"
)

// All error types here, so they don't get over-written in the mutation, query or subscription resolvers when generating schema
var (
	ErrUnAuthenticated = errors.New("user not authenticated")
	ErrPayloadInvalid  = errors.New("payload is empty/invalid")
)

type ResolverOpts struct {
	OnfidoClient      onfidoService.OnfidoServiceClient
	cddClient         cddService.CddServiceClient
	onBoardingService onboardingService.OnBoardingServiceClient
	verifyService     verifyService.VerifyServiceClient
	AuthService       authService.AuthServiceClient
}

type Resolver struct {
	cddService        cddService.CddServiceClient
	onBoardingService onboardingService.OnBoardingServiceClient
	verifyService     verifyService.VerifyServiceClient
	onfidoClient      onfidoService.OnfidoServiceClient
	authService       authService.AuthServiceClient
	logger            *logrus.Logger
}

func NewResolver(opt *ResolverOpts, logger *logrus.Logger) *Resolver {
	return &Resolver{
		cddService:        opt.cddClient,
		onBoardingService: opt.onBoardingService,
		verifyService:     opt.verifyService,
		onfidoClient:      opt.OnfidoClient,
		authService:       opt.AuthService,
		logger:            logger,
	}
}

func ConnectServiceDependencies(secrets *config.Secrets) (*ResolverOpts, error) {
	// TODO: Ensure it is secure when connecting.
	// TODO: Find a way to watch the service outage and handle response to client.
	// TODO: Read heartbeat from these services, if a heartbeat is out, buzz the admin.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := &ResolverOpts{}

	// OnBoarding
	connection, err := dialRPC(ctx, secrets.OnboardingServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.OnboardingServiceURL)
	}
	opts.onBoardingService = onboardingService.NewOnBoardingServiceClient(connection)

	// OnFido
	connection, err = dialRPC(ctx, secrets.OnfidoServiceURL)
	if err != nil {
		return nil, errors.Wrap(err, secrets.OnfidoServiceURL)
	}
	opts.OnfidoClient = onfidoService.NewOnfidoServiceClient(connection)

	// CDD
	connection, err = dialRPC(ctx, secrets.CddServiceURL)
	if err != nil {
		return nil, errors.Wrap(err, secrets.CddServiceURL)
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

	//// Auth
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.AuthServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.VerifyServiceURL)
	}
	opts.AuthService = authService.NewAuthServiceClient(connection)
	return opts, nil
}

func dialRPC(ctx context.Context, address string) (*grpc.ClientConn, error) {
	//cred := new(tls.Config) // TODO: Find a way to read this from the right source.
	//connection, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(cred)))
	connection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return connection, nil
}
