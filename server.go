package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alexflint/go-arg"
	"github.com/gorilla/websocket"
	"github.com/howstrongiam/backend/cmd"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/generated"
	r "github.com/howstrongiam/backend/graph/resolvers"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultPort = "8080"

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	// Initialize config struct and populate it from env vars and flags.
	cfg := cmd.DefaultConfiguration()
	log.Print("CFG", cfg)
	arg.MustParse(cfg)

	//port := ":" + cfg.Port

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Use a default port if $PORT is not set
	}

	c := cors.Default()

	// initialize database
	database.Connect(cfg.DatabaseConfig) // Pass the entire DatabaseConfig struct

	// Initialize config struct and populate it from env vars and flags.
	//cfg := cmd.DefaultConfiguration()
	//arg.MustParse(cfg)
	//
	//port := ":" + cfg.Port
	//
	//c := cors.Default()
	//
	//// initialize database
	//database.Connect(cfg.DatabaseConfig)

	// setup graphql server
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &r.Resolver{}}))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	if port == "" {
		port = defaultPort
	}

	srv = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	logrus.Printf("connect to http://localhost%s/ for GraphQL playground", port)
	logrus.Fatal(http.ListenAndServe(port, nil))
}
