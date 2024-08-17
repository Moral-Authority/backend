package main

import (
    "log"
    "os"

    "github.com/Moral-Authority/backend/seed" // Ensure this is correctly imported
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Read the database URL from the environment variable
    dsn := os.Getenv("DATABASE_URL")

    // Connect to the database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // If the SEED_DB environment variable is set, seed the database
    if os.Getenv("SEED_DB") == "true" {
        log.Println("Seeding the database...")
        seed.SeedCertifications(db) // Ensure this is prefixed with the package name
        // Add other seed functions as needed
        log.Println("Database seeding complete.")
        return
    }

    // Your existing application logic goes here
    log.Println("Starting the application...")
    // Add code to start your server or other services
}
