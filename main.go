package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ms.api/config"
	httpServer "ms.api/server/http"
	"net/http"
	"os"
)

func main() {
	logger := setupLogger()
	secrets, err := config.LoadSecrets()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}
	address := fmt.Sprintf("0.0.0.0:%s", secrets.Port)
	logger.Infof("Connect to http://%s/ for GraphQL playground", address)
	if err := http.ListenAndServe(address, httpServer.MountServer(secrets, logger)); err != nil {
		logrus.Fatalf("Could not start server on %s. Got error: %s", address, err.Error())
	}
}

func setupLogger() *logrus.Logger {
	var logFormatter logrus.Formatter
	if os.Getenv("env") == "dev" {
		logFormatter = &logrus.TextFormatter{}
	} else {
		logFormatter = &logrus.JSONFormatter{}
	}
	logger := logrus.New()
	logger.Formatter = logFormatter
	return logger
}
