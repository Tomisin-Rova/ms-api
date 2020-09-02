package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ms.api/config"
	"ms.api/libs/host"
	httpServer "ms.api/server/http"
	"net/http"
)

func main() {
	address := fmt.Sprintf("%s:%s", host.Host(), config.GetSecrets().Port)
	logrus.Printf("Connect to http://%s/ for GraphQL playground", address)
	if err := http.ListenAndServe(address, httpServer.MountGraphql()); err != nil {
		logrus.Fatalf("Could not start server on %s. Got error: %s", address, err.Error())
	}
}
