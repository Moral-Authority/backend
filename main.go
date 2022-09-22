package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph"
	"github.com/howstrongiam/backend/graph/generated"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {

	port := defaultPort

	database.ConnectDatabase()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: database.GetDatabase()}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
