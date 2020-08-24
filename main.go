package main

import (
	"fmt"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/recover"
	"github.com/joho/godotenv"
	"log"
	"ms.api/config"
	"ms.api/middlewares"
	"ms.api/playground"
	"ms.api/server"
)

var _ = godotenv.Load()

func main() {
	app := fiber.New(&fiber.Settings{
		StrictRouting: false,
		CaseSensitive: false,
		ServerHeader:  "API",
		//Concurrency:   1000,
	})

	allowedOrigins := []string{"*"}

	// *************** MiddleWares ********** //
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
	}))

	app.Use(helmet.New())
	app.Use(middlewares.AuthMiddleWare)
	app.Use(middlewares.ProtectedMiddleware)
	// ************** MiddleWares ********** //

	/**
	Mount Playgrounds only on any env, except production.
	*/
	if config.GetSecrets().Environment != config.Production {
		// ********************* Playgrounds ****************** //
		app.Get("/", playground.MountPlayground)
		//app.Get("/visualise", GraphQLViews.MountVisualDependencyGraph)
		// ********************* Playgrounds ****************** //
	} else {
		app.Settings.Prefork = true
		// Route => handler
		app.Get("/", func(c *fiber.Ctx) {
			c.Set("content-type", "text/html")
			c.SendString("Welcome to Roava API. Please use our APP for a better experience.</a>")
		})
	}

	app.All("/graphql", Server.GraphQL().ServeGraphQL)

	// ***************** API Server **********************//
	host := "0.0.0.0"
	if config.GetSecrets().Environment == config.Local {
		host = "127.0.0.1"
	}
	address := fmt.Sprintf("%s:%s", host, config.GetSecrets().Port)
	log.Printf("Connect to http://%s/ for GraphQL playground", address)
	if err := app.Listen(address); err != nil {
		log.Fatalf("Could not start server on %s. Got error: %s", address, err.Error())
	}
	// ****************** API Server **************** //
}
