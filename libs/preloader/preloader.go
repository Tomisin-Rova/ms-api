package preloader

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

//go:generate mockgen -source=preloader.go -destination=../../mocks/preloader_mock.go -package=mocks
type Preloader interface {
	GetPreloads(ctx context.Context) []string
}

type GQLPreloader struct{}

//GetPreloads returns the fields that were queried for in a grapqhql request
// makes use of gqlgen collect fields functionality
func (g GQLPreloader) GetPreloads(ctx context.Context) []string {
	return getNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}
func getNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := getPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, getNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}
func getPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
