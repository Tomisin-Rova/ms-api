package graph

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/protos/pb/kycService"
	"ms.api/protos/pb/onfidoService"
	"ms.api/protos/pb/onboardingService"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	onfidoClient onfidoService.OnfidoServiceClient
	kycClient    kycService.KycServiceClient
	onboardingClient onboardingService.OnBoardingServiceClient
}

func (r *Resolver) ConnectServiceDependencies() {
	// TODO: Ensure it is secure when connecting.
	// TODO: Find a way to watch the service outage and handle response to client.
	// TODO: Read heartbeat from these services, if a heartbeat is out, buzz the admin.
	if connection := dialRPC(config.GetSecrets().OnfidoServiceURL); connection != nil {
		fmt.Print("Connected to ms.onfido \n")
		r.onfidoClient = onfidoService.NewOnfidoServiceClient(connection)
	}
	if connection := dialRPC(config.GetSecrets().KYCServiceURL); connection != nil {
		fmt.Print("Connected to ms.kyc \n")
		r.kycClient = kycService.NewKycServiceClient(connection)
	}
	if connection := dialRPC(config.GetSecrets().OnboardingServiceURL); connection != nil {
		fmt.Print("Connected to ms.onboarding \n")
		r.onboardingClient = onboardingService.NewOnBoardingServiceClient(connection)
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
