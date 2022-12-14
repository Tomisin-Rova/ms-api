package preloader

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

//go:generate mockgen -source=preloader.go -destination=../../mocks/preloader_mock.go -package=mocks
type Preloader interface {
	GetPreloads(ctx context.Context) []string
	GetArgMap(ctx context.Context, field string) map[string]interface{}
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

func collectFields(ctx context.Context) []graphql.CollectedField {
	var fields []graphql.CollectedField
	if graphql.GetFieldContext(ctx) != nil {
		fields = graphql.CollectFieldsCtx(ctx, nil)

		octx := graphql.GetOperationContext(ctx)
		for _, col := range fields {
			if col.Name == "nodes" || col.Name == "edges" {
				// This endpoint is using the cursor pattern; the columns we
				// actually need to filter with are nested into the edges or nodes
				// field.
				fields = graphql.CollectFields(octx, col.SelectionSet, nil)
				break
			}
		}
	}
	return fields
}

// getArgMap returns the argument map for a specified field from the gqlgen field collector
func (g GQLPreloader) GetArgMap(ctx context.Context, field string) map[string]interface{} {
	var argMap map[string]interface{}
	for _, f := range collectFields(ctx) {
		if f.Name == field {
			argMap = f.ArgumentMap(nil)
			break
		}
	}
	return argMap
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
