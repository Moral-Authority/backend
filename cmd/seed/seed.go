package main

import (
	"log"
	"os"
	"time"

	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	"github.com/Moral-Authority/backend/handlers"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

func main() {
	// Read the database URL from the environment variable
	// dsn := os.Getenv("DATABASE_URL")
	// if dsn == "" {
	// 	log.Fatal("DATABASE_URL is not set")
	// }

	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Wipe all tables in the database
	wipeDatabase(db)

	// Ensure the certifications table exists by migrating the schema
	err = db.AutoMigrate(&models.Certification{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	log.Println("Seeding the database...")
	seedCertifications(db)
	log.Println("Seeding BCorps...")
	// Seed BCorp companies
	seedCompaniesFromCSV(db, "all-bcorps.csv", "BCorp Certified")
	log.Println("Seeding Made Safe...")
	// Seed Made Safe companies
	seedCompaniesFromCSV(db, "made_safe_companies.csv", "Made Safe")
	log.Println("Seeding Products.")
	seedProductsFromCSV(db, "affiliate_products_blueland_products.csv", "Blueland")
	log.Println("Database seeding complete.")
}

func wipeDatabase(db *gorm.DB) {
	// Get the list of all tables in the database
	var tables []string
	err := db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables).Error
	if err != nil {
		log.Fatal("Failed to get table names:", err)
	}

	// Truncate each table
	for _, table := range tables {
		err := db.Exec("TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE").Error
		if err != nil {
			log.Fatal("Failed to truncate table", table, ":", err)
		}
	}

	log.Println("All tables wiped.")
}

func seedCertifications(db *gorm.DB) {

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)

	// Use the directory to open files
	filePath := fmt.Sprintf("%s/cmd/seed/certifications/final_import.csv", dir)
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

		var certifiesCompany, certifiesProduct bool
		var err error

		if record[5] != "" {
			certifiesCompany, err = strconv.ParseBool(strings.TrimSpace(record[5]))
			if err != nil {
				logrus.Fatal(err)
			}
		}

		if record[6] != "" {
			certifiesProduct, err = strconv.ParseBool(strings.TrimSpace(record[6]))
			if err != nil {
				logrus.Fatal(err)
			}
		}

		name := null.StringFrom(strings.TrimSpace(record[0]))

		cert := models.Certification{
			Name:             name,
			Logo:             null.StringFrom(record[2]),
			Website:          null.StringFrom(record[3]),
			Description:      null.StringFrom(record[4]),
			CertifiesCompany: null.BoolFrom(certifiesCompany),
			CertifiesProduct: null.BoolFrom(certifiesProduct),
			Certifier:        null.StringFrom(record[8]),
			Industry:         null.StringFrom(record[9]),
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

func seedCompaniesFromCSV(db *gorm.DB, fileName, certificationName string) {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)

	// Use the directory to open files
	filePath := fmt.Sprintf("%s/cmd/seed/companies/%s", dir, fileName)
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

	certificationID := findCertificationID(db, certificationName)
	// Read each row of the CSV file and insert it into the database
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		company := formatCompanyFromRow(filePath, row)

		// Insert the company into the database
		result := db.Create(&company)
		if result.Error != nil {
			fmt.Println(result.Error)
		}

		dateCertified := extractDateCertified(filePath, row)
		createCompanyCertification(db, company.ID, certificationID, dateCertified)
	}

	fmt.Println("Companies seeded into database.")
}

func formatCompanyFromRow(filePath string, row []string) models.Company {
	switch filePath {
	case "cmd/seed/companies/all-bcorps.csv":
		return models.Company{
			Name:        row[0],
			Country:     null.StringFrom(row[4]),
			State:       null.StringFrom(row[5]),
			City:        null.StringFrom(row[6]),
			Url:         null.StringFrom(row[9]),
			Description: null.NewString("", false),
			UserId:      null.Int64From(0),
			IsVerified:  null.Bool{Valid: false},
			ImageId:     null.Int64From(0), // Or some default value if not available in CSV
		}
	case "cmd/seed/companies/made_safe_companies.csv":
		return models.Company{
			Name:        row[0],
			Country:     null.NewString("", false),
			State:       null.NewString("", false),
			City:        null.NewString("", false),
			Url:         null.StringFrom(row[2]),
			Description: null.NewString("", false),
			UserId:      null.Int64From(0),
			IsVerified:  null.Bool{Valid: false},
			ImageId:     null.Int64From(0), // Or some default value if not available in CSV
		}

	default:
		return models.Company{
			Name:        row[0],
			Description: null.NewString("", false),
			UserId:      null.Int64From(0),
			IsVerified:  null.Bool{Valid: false},
			ImageId:     null.Int64From(0), // Default values
		}
	}
}

func extractDateCertified(filePath string, row []string) null.Time {

	switch filePath {
	case "cmd/seed/companies/all-bcorps.csv":
		timeCertified, err := time.Parse("2006-01-02 15:04:05", row[1])
		if err != nil {
			fmt.Println("Error parsing time:", err)
			timeCertified = time.Time{} // Use zero value if parsing fails
		}

		return null.TimeFrom(timeCertified)
	default:
		return null.Time{Valid: false}
	}
}

func createCompanyCertification(db *gorm.DB, companyID uint, certificationID uint, dateCertified null.Time) error {

	// Create the CompanyCertification relationship
	companyCert := models.CompanyCertification{
		CompanyID:       companyID,
		CertificationID: certificationID,
		CertifiedAt:     dateCertified,
		ExpirationDate:  null.NewTime(time.Time{}, false),
		OtherDetails:    null.NewString("Additional details about the certification", true),
	}

	if err := db.Create(&companyCert).Error; err != nil {
		return fmt.Errorf("failed to insert company certification: %v", err)
	}

	return nil
}

func findCertificationID(db *gorm.DB, certificationName string) uint {
	// Find the Certification ID based on the certification name
	var certification models.Certification
	if err := db.Where("name = ?", certificationName).First(&certification).Error; err != nil {
		log.Printf("failed to find certification: %v", err)
		return 0
	}

	return certification.ID
}

func findCompanyID(db *gorm.DB, companyName string) uint {
	// Find the Certification ID based on the certification name
	var company models.Company
	if err := db.Where("name = ?", companyName).First(&company).Error; err != nil {
		log.Printf("failed to find company: %v", err)
		return 0
	}

	return company.ID
}

func seedProductsFromCSV(db *gorm.DB, fileName string, companyName string) {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)

	// Construct the file path
	filePath := fmt.Sprintf("%s/cmd/seed/products/%s", dir, fileName)
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

	companyID := findCompanyID(db, companyName)
	prodDeptType, isDept := handlers.IsStringValidProductDepartment("Home & Garden")
	if !isDept {	
		fmt.Println("Invalid product department")
	}
	prodDept := prodDeptType.ToInt()

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		// Create the HomeGardenProduct
		product := models.HomeGardenProduct{
			ProductBase: models.ProductBase{
				SubDepartment: row[2],
				Title:         row[3],
				Url:           row[5],
				CompanyID:     companyID,
				ProductImage:  row[6],
			},
		}

		// Insert the product into the database
		result := db.Create(&product)
		if result.Error != nil {
			fmt.Println(result.Error)
		} else {
			// Create the PurchaseInfo with ProductDepartment set to HomeGardenProductDepartment
			purchaseInfo := models.PurchaseInfo{
				ProductID:         product.ID,
				ProductDepartment: prodDept, // HomeGardenProductDepartment
				Price:             row[4],
				Url:               row[5],
			}

			// Insert the purchase info into the database
			result = db.Create(&purchaseInfo)
			if result.Error != nil {
				fmt.Println(result.Error)
			}
		}
	}

	fmt.Println("Products seeded into database.")
}