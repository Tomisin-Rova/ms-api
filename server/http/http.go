package httpServer

import (
	"context"
	"fmt"
	"ms.api/libs/db/mongo"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"ms.api/config"
	"ms.api/graph"
	"ms.api/graph/generated"
	rerrors "ms.api/libs/errors"
	"ms.api/server/http/handlers"
	"ms.api/server/http/middlewares"
)

func MountServer(secrets *config.Secrets, logger *zap.Logger) *chi.Mux {
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
	opts, err := graph.ConnectServiceDependencies(secrets)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to setup service dependencies: %v", err))
	}

	store, mongoClient, err := mongo.New(secrets.Database.URL, secrets.Database.Name, logger)
	if err != nil {
		logger.Fatal("failed to open database", zap.Error(err))
	}

	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			logger.Error("failed to close database", zap.Error(err))
		}
	}()

	mw := middlewares.NewAuthMiddleware(opts.AuthService, logger)
	router.Use(mw.Middeware)
	opts.AuthMw = mw
	opts.DataStore = store

	if secrets.Service.Environment != string(config.Production) {
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	} else {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("content-type", "text/html")
			_, _ = writer.Write([]byte("Welcome to Roava API. Please use our APP for a better experience.</a>"))
		})
	}

	resolvers := graph.NewResolver(opts, logger)
	httpHandlers := handlers.New(opts.OnBoardingService, logger)
	// API Server
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	server.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return rerrors.FormatGqlTError(e, err)
	})

	router.Handle("/graphql", server)
	router.Get("/verify_email", httpHandlers.VerifyMagicLinkHandler)
	return router
}
