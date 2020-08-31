package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"ms.api/handlers"
	"ms.api/handlers/generated"
)

func MountGraphql(resolvers *handlers.Resolver) *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
}
