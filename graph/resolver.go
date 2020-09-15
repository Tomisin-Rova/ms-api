package graph

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/onfidoService"
	"ms.api/protos/pb/verifyService"
	"time"
)

type ResolverOpts struct {
	onfidoClient      onfidoService.OnfidoServiceClient
	kycClient         kycService.KycServiceClient
	onBoardingService onboardingService.OnBoardingServiceClient
	verifyService     verifyService.VerifyServiceClient
}

type Resolver struct {
	kycClient         kycService.KycServiceClient
	onBoardingService onboardingService.OnBoardingServiceClient
	verifyService     verifyService.VerifyServiceClient
	onfidoClient      onfidoService.OnfidoServiceClient
	logger            *logrus.Logger
}

func NewResolver(opt *ResolverOpts, logger *logrus.Logger) *Resolver {
	return &Resolver{
		kycClient:         opt.kycClient,
		onBoardingService: opt.onBoardingService,
		verifyService:     opt.verifyService,
		onfidoClient:      opt.onfidoClient,
		logger:            logger,
	}
}

func ConnectServiceDependencies(secrets *config.Secrets) (*ResolverOpts, error) {
	// TODO: Ensure it is secure when connecting.
	// TODO: Find a way to watch the service outage and handle response to client.
	// TODO: Read heartbeat from these services, if a heartbeat is out, buzz the admin.
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	opt := &ResolverOpts{}

	// OnBoarding
	connection, err := dialRPC(ctx, secrets.OnboardingServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.OnboardingServiceURL)
	}
	opt.onBoardingService = onboardingService.NewOnBoardingServiceClient(connection)

	// OnFido
	/*connection, err = dialRPC(ctx, secrets.OnfidoServiceURL)
	if err != nil {
		return nil, errors.Wrap(err, secrets.OnfidoServiceURL)
	}
	opt.onfidoClient = onfidoService.NewOnfidoServiceClient(connection)*/

	// KYC
	/*connection, err = dialRPC(ctx, secrets.KYCServiceURL)
	if err != nil {
		return nil, errors.Wrap(err, secrets.KYCServiceURL)
	}
	opt.kycClient = kycService.NewKycServiceClient(connection)*/

	// Verify
	ctx, cancel = context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.VerifyServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.VerifyServiceURL)
	}
	opt.verifyService = verifyService.NewVerifyServiceClient(connection)
	return opt, nil
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
