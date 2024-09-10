package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Moral-Authority/backend/cmd"
	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/generated"
	r "github.com/Moral-Authority/backend/graph/resolvers"
	"github.com/alexflint/go-arg"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func InitAlgoliaClient() *search.Client {
	
	appID := os.Getenv("ALGOLIASEARCH_APPLICATION_ID")
	apiKey := os.Getenv("ALGOLIASEARCH_API_KEY")

	if appID == "" || apiKey == "" {
		log.Fatalf("Algolia credentials not set in environment variables")
	}

	client := search.NewClient(appID, apiKey)
	return client
}

func main() {
	// err := godotenv.Load("/Users/lilchichie/src/moralAuthority/backend/.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }


	// Set the log output to stdout
	log.SetOutput(os.Stdout)

	// Initialize config struct and populate it from env vars and flags.
	cfg := cmd.DefaultConfiguration()
	log.Print("CFG", cfg)
	arg.MustParse(cfg)

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort // Use a default port if $PORT is not set
	}

	// Initialize CORS with custom settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://moralauthority.co", "https://moral-authority-fe-1bd5c971d1d7.herokuapp.com", "http://localhost:3000", "https://*.moralauthority.co"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Connect to the database
	database.Connect(cfg.DatabaseConfig)

	// Drop and recreate the database schema
	// database.DropDatabase()

	// Perform migrations after dropping the database
	database.PerformMigrations()

	// Initialize Algolia client
	algoliaClient := InitAlgoliaClient()

	// Setup GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r.Resolver{
		AlgoliaClient: algoliaClient,
	}}))
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

	// HTTP handlers
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", c.Handler(srv))

	logrus.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Printf("Using port: %s", port)
	logrus.Fatal(http.ListenAndServe(":"+port, nil))
}
