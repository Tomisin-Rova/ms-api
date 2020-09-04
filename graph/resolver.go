package graph

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/protos/pb/kycService"
	onboarding "ms.api/protos/pb/onboardingService"
	verify "ms.api/protos/pb/verifyService"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type ResolverOpts struct {
	kycClient kycService.KycServiceClient
	OnBoardingService onboarding.OnBoardingServiceClient
	VerifyService verify.VerifyServiceClient
	Logger *logrus.Logger
}

type Resolver struct {
	kycClient kycService.KycServiceClient
	onBoardingService onboarding.OnBoardingServiceClient
	verifyService verify.VerifyServiceClient
	logger *logrus.Logger
}

func NewResolver(opt ResolverOpts) *Resolver {
	return &Resolver{
		kycClient: opt.kycClient,
		onBoardingService: opt.OnBoardingService,
		verifyService: opt.VerifyService,
		logger: opt.Logger,
	}
}

func (r *Resolver) ConnectServiceDependencies() {
	// TODO: Ensure it is secure when connecting.
	// TODO: Find a way to watch the service outage and handle response to client.
	// TODO: Read heartbeat from these services, if a heartbeat is out, buzz the admin.
	if connection := dialRPC(config.GetSecrets().KYCServiceURL); connection != nil {
		fmt.Print("Connected to ms.kyc \n")
		r.kycClient = kycService.NewKycServiceClient(connection)
	}
}

func dialRPC(address string) *grpc.ClientConn {
	//cred := new(tls.Config) // TODO: Find a way to read this from the right source.
	//connection, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(cred)))
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		return nil
	}

	return connection
}
