package main

import (
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"strconv"
    "strings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
)

func main() {
	// Read the database URL from the environment variable
	dsn := os.Getenv("HEROKU_DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the certifications table exists by migrating the schema
	err = db.AutoMigrate(&models.Certification{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	log.Println("Seeding the database...")
	seedCertifications(db)
    // log.Println("Seeding the companies...")
    seedCompanies(db)
	log.Println("Database seeding complete.")
}

func seedCertifications(db *gorm.DB) {

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)
		
	// Use the directory to open files
	filePath := fmt.Sprintf("%s/seed/Certifications_Amazon.csv", dir)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	reader.LazyQuotes = true

	// Read the first line (header)
	_, err = reader.Read()
	if err != nil {
		logrus.Fatal(err)
	}

	// Read the rest of the file
	records, err := reader.ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// Iterate over the records
	for _, record := range records {
        // log.Println("records", record)
		certifiesCompany, err := strconv.ParseBool(strings.TrimSpace(record[4]))
		if err != nil {
			logrus.Fatal(err)
		}

		certifiesProduct, err := strconv.ParseBool(strings.TrimSpace(record[5]))
		if err != nil {
			logrus.Fatal(err)
		}

		cert := models.Certification{
			Name:             null.StringFrom(record[0]),
			Logo:             null.StringFrom(record[1]),
			Website:          null.StringFrom(record[2]),
			Description:      null.StringFrom(record[3]),
			CertifiesCompany: null.BoolFrom(certifiesCompany),
			CertifiesProduct: null.BoolFrom(certifiesProduct),
		}

		// Insert the certification into the database
		result := db.Create(&cert)
		if result.Error != nil {
            log.Println("records", record)
			logrus.Error("cert", cert)
			logrus.Error(result.Error)
		}
	}

	logrus.Println("Certifications seeded into database.")
}


func seedCompanies(db *gorm.DB) {

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)
		
	// Use the directory to open files
	filePath := fmt.Sprintf("%s/seed/all-bcorps.csv", dir)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Read each row of the CSV file and insert it into the database
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		// Parse data from CSV row into appropriate types
		userId, _ := strconv.ParseInt(row[9], 10, 64)
		isVerified, _ := strconv.ParseBool(row[10])

		company := models.Company{
			Name:        row[0],
			City:        null.StringFrom(row[1]),
			State:       null.StringFrom(row[2]),
			Country:     null.StringFrom(row[3]),
			Url:         null.StringFrom(row[4]),
			Description: null.StringFrom(row[5]),
			UserId:      null.Int64From(userId),
			IsVerified:  null.BoolFrom(isVerified),
			ImageId:     null.Int64From(0), // Or some default value if not available in CSV
		}

		// Insert the company into the database
		result := db.Create(&company)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
	}

	fmt.Println("Companies seeded into database.")
}
