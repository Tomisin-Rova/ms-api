package Server

import (
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	fiberGQL "github.com/Just4Ease/fiber-gqlgen"
	"ms.api/api"
	"ms.api/api/generated"
	secrets "ms.api/config"
)

func GraphQL() *fiberGQL.Server {
	cfg := generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{}})
	gateway := fiberGQL.New(cfg)

	gateway.SetQueryCache(lru.New(3000))

	if secrets.GetSecrets().Environment != secrets.Production {
		gateway.Use(extension.Introspection{})
	}
	gateway.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100000),
	})
	return gateway
}
