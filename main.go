package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"ms.api/api"
	"ms.api/config"
	"ms.api/middlewares"
	"ms.api/server"
	"net/http"
)

func main() {
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
	router.Use(middlewares.AuthMiddleWare)
	router.Use(middlewares.ProtectedMiddleware)

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
	resolvers := &api.Resolver{}
	resolvers.ConnectServiceDependencies()
	router.Handle("/graphql", server.MountGraphql(resolvers))

	host := "0.0.0.0"
	if config.GetSecrets().Environment == config.Local {
		host = "127.0.0.1"
	}
	address := fmt.Sprintf("%s:%s", host, config.GetSecrets().Port)
	logrus.Printf("Connect to http://%s/ for GraphQL playground", address)
	if err := http.ListenAndServe(address, router); err != nil {
		logrus.Fatalf("Could not start server on %s. Got error: %s", address, err.Error())
	}
}
