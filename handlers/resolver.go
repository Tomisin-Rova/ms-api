package handlers

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/services/kycService"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	kycClient kycService.KycServiceClient
}

func (r *Resolver) ConnectServiceDependencies() {
	// TODO: Ensure it is secure when connecting.
	// TODO: Find a way to watch the service outage and handle response to client.
	// TODO: Read heartbeat from these services, if a heartbeat is out, buzz the admin.
	if connection := dialRPC(config.GetSecrets().KYCServiceURL); connection != nil {
		r.kycClient = kycService.NewKycServiceClient(connection)
	}
}

func dialRPC(address string) *grpc.ClientConn {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		return nil
	}

	return connection
}
