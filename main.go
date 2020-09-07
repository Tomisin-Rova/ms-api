package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ms.api/config"
	"ms.api/log"
	httpServer "ms.api/server/http"
	"net/http"
)

func main() {
	logger := log.New(log.LevelInfo, 0)
	address := fmt.Sprintf("127.0.0.1:%s", config.GetSecrets().Port)
	logger.Debug("Connect to http://%s/ for GraphQL playground", address)
	if err := http.ListenAndServe(address, httpServer.MountGraphql()); err != nil {
		logrus.Fatalf("Could not start server on %s. Got error: %s", address, err.Error())
	}
}
