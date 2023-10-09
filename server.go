package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"github.com/t-kuni/go-graphql-template/di"
	"github.com/t-kuni/go-graphql-template/logger"
	"github.com/t-kuni/go-graphql-template/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/t-kuni/go-graphql-template/graph"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()

	if err := logger.SetupLogger(); err != nil {
		log.Fatalf("Logger initialization failed: %+v", err)
		os.Exit(1)
	}

	logger.SimpleInfoF("APP_ENV: %s, DB_HOST: %s", os.Getenv("APP_ENV"), os.Getenv("DB_HOST"))

	app := di.NewApp()
	defer app.Shutdown()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()

	c := graph.Config{Resolvers: &graph.Resolver{App: app}}
	c.Directives = do.MustInvoke[graph.DirectiveRoot](app)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Group(func(r chi.Router) {
		r.Use(do.MustInvoke[*middleware.Authentication](app).Middleware)
		r.Handle("/query", srv)
	})

	logger.SimpleInfoF("connect to http://localhost:%s/ for GraphQL playground", port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.SimpleFatal(err, nil)
	}
}
