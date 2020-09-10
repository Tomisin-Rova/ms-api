package httpServer

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"ms.api/config"
	"ms.api/graph"
	"ms.api/graph/generated"
	"net/http"
	"os"
)

func MountGraphql() *chi.Mux {
	secrets, err := config.LoadSecrets()
	logger := setupLogger()
	if err != nil {
		logger.Fatal("failed to load secrets: %v", err)
	}

	router := chi.NewRouter()
	// Middlewares
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	if secrets.Environment != config.Production {
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	} else {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("content-type", "text/html")
			_, _ = writer.Write([]byte("Welcome to Roava API. Please use our APP for a better experience.</a>"))
		})
	}

	// API Server
	opt, err := graph.ConnectServiceDependencies(secrets)
	if err != nil {
		logger.Fatalf("failed to setup service dependencies: %v", err)
	}

	resolvers := graph.NewResolver(opt, logger)
	router.Handle("/graphql", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers})))
	return router
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
