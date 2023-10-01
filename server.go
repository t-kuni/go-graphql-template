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

	println("DB_USER: " + os.Getenv("DB_USER"))

	app := di.NewApp()
	// TODO: app.Shutdown()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{App: app}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
