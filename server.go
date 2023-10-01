package main

import (
	"github.com/joho/godotenv"
	"github.com/t-kuni/go-graphql-template/di"
	"github.com/t-kuni/go-graphql-template/logger"
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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{App: app}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.SimpleInfoF("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.SimpleFatal(http.ListenAndServe(":"+port, nil), nil)
}
