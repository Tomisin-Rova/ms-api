package httpServer

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ms.api/config"
	"ms.api/graph"
	"ms.api/graph/generated"
	onboarding "ms.api/protos/pb/onboardingService"
	verify "ms.api/protos/pb/verifyService"
	"net/http"
	"os"
	"time"
)

func MountGraphql() *chi.Mux {
	logger := setupLogger()
	router := chi.NewRouter()
	// MiddleWares
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	//router.Use(middlewares.AuthMiddleWare) // A different session service will be used, not the gateway.
	//router.Use(middlewares.ProtectedMiddleware)

	secrets := config.GetSecrets()
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	onBoardingClient, err := createOnBoardingServiceClient(ctx, secrets.OnBoardingServiceUrl)
	if err != nil {
		logger.Fatal("failed to connect to ms.onboarding: ", err)
	}
	verifyServiceClient, err := createVerifyServiceClient(ctx, secrets.VerifyServiceUrl)
	if err != nil {
		logger.Fatal("failed to connect to ms.verify: ", err)
	}

	if secrets.Environment != config.Production {
		// ********************* Playgrounds ****************** //
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	} else {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("content-type", "text/html")
			_, _ = writer.Write([]byte("Welcome to Roava API. Please use our APP for a better experience.</a>"))
		})
	}
	// API server
	opt := graph.ResolverOpts{
		OnBoardingService: onBoardingClient,
		VerifyService: verifyServiceClient,
		Logger: logger,
	}
	resolvers := graph.NewResolver(opt)
	router.Handle("/graphql", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers})))
	return router
}

func createOnBoardingServiceClient(ctx context.Context, addr string) (onboarding.OnBoardingServiceClient, error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return onboarding.NewOnBoardingServiceClient(conn), nil
}

func createVerifyServiceClient(ctx context.Context, addr string) (verify.VerifyServiceClient, error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return verify.NewVerifyServiceClient(conn), nil
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
