package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"ms.api/api"
	"ms.api/api/generated"
)

func MountGraphql(resolvers *api.Resolver) *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
}
