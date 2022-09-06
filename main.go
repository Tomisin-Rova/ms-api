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

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

func initTracer() func(context.Context) error {
	// OpenTelemetry
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("Could not set resources: ", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}

func main() {
	//Initialize tracer
	cleanup := initTracer()
	defer cleanup(context.Background())

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	// Connect to database
	models.ConnectDatabase()

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
