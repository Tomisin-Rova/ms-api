package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/transport"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"ms.api/config"
	"ms.api/graph"
	"ms.api/graph/generated"
	rerrors "ms.api/libs/errors"
	"ms.api/server/http/middlewares"
)

// A Websocket transport is already added when using the NewDefaultServer function.
// So it's required to initialize the server by using almost the same implementation
// but with a custom WebSocket transport.
func NewCustomServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)

	// Configure WebSocket
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			return ctx, nil
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv
}

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

	mw := middlewares.NewAuthMiddleware(opts.AuthService, logger)
	router.Use(mw.Middeware)

	opts.AuthMw = mw

	if secrets.Service.Environment != string(config.Production) {
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	} else {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("content-type", "text/html")
			_, _ = writer.Write([]byte("Welcome to Roava API. Please use our APP for a better experience.</a>"))
		})
	}

	resolvers := graph.NewResolver(opts, logger)
	// API Server
	server := NewCustomServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	server.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return rerrors.FormatGqlTError(e, err)
	})

	// Cors setup
	corsSetup := cors.New(cors.Options{
		AllowCredentials: true,
		Debug:            false,
	})

	router.Handle("/graphql", corsSetup.Handler(server))

	return router
}
