package httpServer

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"ms.api/config"
	"ms.api/graph"
	"ms.api/graph/generated"
	"net/http"
)

func MountGraphql() *chi.Mux {
	router := chi.NewRouter()
	// *************** MiddleWares **********
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

	if config.GetSecrets().Environment != config.Production {
		// ********************* Playgrounds ****************** //
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	} else {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("content-type", "text/html")
			_, _ = writer.Write([]byte("Welcome to Roava API. Please use our APP for a better experience.</a>"))
		})
	}

	// ***************** API Server **********************
	resolvers := &graph.Resolver{}
	resolvers.ConnectServiceDependencies()
	router.Handle("/graphql", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers})))

	return router
}
