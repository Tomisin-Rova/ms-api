package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"ms.api/config"
	httpServer "ms.api/server/http"

	coreLogger "github.com/roava/zebra/logger"
	"go.uber.org/zap"
)

func main() {
	logger := setupLogger()
	secrets, err := config.LoadSecrets()
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to load config: %v", err))
	}

	address := fmt.Sprintf("0.0.0.0:%s", secrets.Port)
	logger.Info(fmt.Sprintf("Connect to http://%s/ for GraphQL playground", address))
	if err := http.ListenAndServe(address, httpServer.MountServer(secrets, logger)); err != nil {
		logger.Fatal(fmt.Sprintf("Could not start server on %s. Got error: %s", address, err.Error()))
	}
}

func setupLogger() *zap.Logger {
	var loggerConfigs = coreLogger.Config{
		Name:       "ms.api",
		WithCaller: true,
	}

	switch strings.ToLower(os.Getenv("ENVIRONMENT")) {
	case "dev", "local":
		loggerConfigs.Level = zap.DebugLevel
		loggerConfigs.Debug = true
	case "stg", "prod":
		loggerConfigs.Level = zap.InfoLevel
	default:
		loggerConfigs.Level = zap.DebugLevel
		loggerConfigs.Debug = true
	}

	logger := coreLogger.New(loggerConfigs)
	return logger
}
