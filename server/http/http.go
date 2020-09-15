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
)

func MountGraphql(secrets *config.Secrets, logger *logrus.Logger) *chi.Mux {
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

	opt, err := graph.ConnectServiceDependencies(secrets)
	if err != nil {
		logger.Fatalf("failed to setup service dependencies: %v", err)
	}

	resolvers := graph.NewResolver(opt, logger)
	// API Server
	router.Handle("/graphql", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers})))
	return router
}
