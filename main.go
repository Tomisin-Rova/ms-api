package main

import (
	"fmt"
	"github.com/roava/zebra/logger"
	"go.uber.org/zap"
	"ms.api/config"
	httpServer "ms.api/server/http"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Load secrets
	logService := setupLogger()
	secrets, err := config.LoadSecrets()
	if err != nil {
		logService.Fatal("load secrets", zap.Error(err))
	}

	address := fmt.Sprintf("localhost:%s", secrets.Service.Port)
	logService.Info(fmt.Sprintf("Connect to http://%s/ for GraphQL playground", address))
	if err := http.ListenAndServe(address, httpServer.MountServer(secrets, logService)); err != nil {
		logService.Fatal(fmt.Sprintf("Could not start server on %s. Got error: %s", address, err.Error()))
	}
}

func setupLogger() *zap.Logger {
	var loggerConfigs = logger.Config{
		Name:       config.ServiceName,
		WithCaller: true,
	}

	switch strings.ToLower(os.Getenv("ENVIRONMENT")) {
	case config.DevEnvironment, config.LocalEnvironment:
		loggerConfigs.Level = zap.DebugLevel
		loggerConfigs.Debug = true
	default:
		loggerConfigs.Level = zap.InfoLevel
	}

	newLogger := logger.New(loggerConfigs)
	return newLogger
}
