package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/samber/do"
	"github.com/t-kuni/go-graphql-template/graph"
	"github.com/t-kuni/go-graphql-template/logger"
)

type Directives struct {
	graph.DirectiveRoot
}

func NewDirectives(i *do.Injector) (graph.DirectiveRoot, error) {
	dir := &Directives{}
	return graph.DirectiveRoot{
		Test: dir.Test,
	}, nil
}

func (*Directives) Test(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	logger.SimpleInfoF("Test directive called")
	return next(ctx)
}
