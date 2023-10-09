package di

import (
	"github.com/samber/do"
	"github.com/t-kuni/go-graphql-template/domain/service"
	"github.com/t-kuni/go-graphql-template/graph/directives"
	"github.com/t-kuni/go-graphql-template/infrastructure/api"
	"github.com/t-kuni/go-graphql-template/infrastructure/db"
	"github.com/t-kuni/go-graphql-template/loaders"
	"github.com/t-kuni/go-graphql-template/middleware"
	"github.com/t-kuni/go-graphql-template/validator"
)

func NewApp() *do.Injector {
	injector := do.New()

	// Directives
	do.Provide(injector, directives.NewDirectives)

	// Validator
	do.Provide(injector, validator.NewCustomValidator)

	// Middleware
	do.Provide(injector, middleware.NewAuthentication)
	do.Provide(injector, middleware.NewRecover)
	do.Provide(injector, middleware.NewAccessLog)

	// Handler
	//do.Provide(injector, handler.NewHelloHandler)
	//do.Provide(injector, handler.NewPostUserHandler)

	// Service
	do.Provide(injector, service.NewExampleService)

	// Infrastructure
	do.Provide(injector, db.NewConnector)
	do.Provide(injector, api.NewBinanceApi)

	// UseCase
	//do.Provide(injector, todos.NewFind)
	//do.Provide(injector, companies.NewGetCompanies)

	// Loaders
	do.Provide(injector, loaders.NewLoader)
	do.Provide(injector, loaders.NewLoaderImpl)

	return injector
}
