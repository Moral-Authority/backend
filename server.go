package main

import (
	"encoding/json"
	"fmt"
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
	"github.com/Moral-Authority/backend/models"
	"github.com/alexflint/go-arg"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

// func loadSeedDataFromFile(db *gorm.DB) error {
// 	file, err := os.Open("seed_data.json")
// 	if err != nil {
// 		return fmt.Errorf("failed to open seed file: %v", err)
// 	}
// 	defer file.Close()

// 	var certifications []models.Certification
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&certifications); err != nil {
// 		return fmt.Errorf("failed to decode seed data: %v", err)
// 	}

// 	for _, cert := range certifications {
// 		if err := db.Create(&cert).Error; err != nil {
// 			log.Printf("Failed to insert certification: %v", err)
// 		}
// 	}

// 	log.Println("Database populated from seed file.")
// 	return nil
// }

func main() {

	// Load environment variables only in dev mode
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "dev" // Default to dev if ENVIRONMENT is not set
	}

	var seed bool
	if environment == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		seed = true
	}

	cfg := cmd.DefaultConfig()
	log.Print("CFG", cfg)
	arg.MustParse(cfg)

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize CORS with custom settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://moralauthority.co", "https://moral-authority-fe-1bd5c971d1d7.herokuapp.com", "http://localhost:3000", "https://*.moralauthority.co"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Connect to the database
	database.Connect(cfg.URL)

	// Perform migrations after dropping the database
	database.PerformMigrations()

	// Initialize Algolia client once
	algoliaClient := InitAlgoliaClient()

	// // Attempt to load from seed file if it exists; otherwise, run the seed script
	// if _, err := os.Stat("seed_data.json"); err == nil {
	// 	log.Println("Loading database from seed file...")
	// 	if err := loadSeedDataFromFile(database); err != nil {
	// 		log.Fatalf("Error loading seed data from file: %v", err)
	// 	}
	// } else {
	// 	log.Println("Seed file not found; seeding database and creating seed file.")
	// 	cmd.SeedDatabase(cfg.URL, algoliaClient)
	// }

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

	if seed {
		cmd.SeedDatabase(cfg.URL, algoliaClient)
	}
	// HTTP handlers
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", c.Handler(srv))

	logrus.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Printf("Using port: %s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
