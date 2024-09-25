package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// AddProductRequest represents the request structure for adding a new product
type AddProductRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Department   string `json:"department"`
	PurchaseInfo struct {
		Price float64 `json:"price"`
		Link  string  `json:"link"`
	} `json:"purchase_info"`
	ImageLinks     []string `json:"image_links"`
	ProductCertifications []struct {
		Name string `json:"name"`
	} `json:"product_certifications"`
	CompanyCertifications []struct {
		Name string `json:"name"`
	} `json:"company_certifications"`
}

func SeedProductsFromCSV() {
	// Load environment variables from .env file
	_ = godotenv.Load("/Users/lilchichie/src/moralAuthority/backend/.env")

	// Filepath of the CSV file to read
	filePath := "affiliate_products_blueland_products1.csv"

	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to open CSV file:", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header line (skip)
	_, err = reader.Read()
	if err != nil {
		log.Fatal("Failed to read CSV header:", err)
	}

	// Iterate through each row in the CSV file
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		// Convert the CSV row to an AddProductRequest
		productRequest, err := createProductRequestFromCSV(row)
		if err != nil {
			logrus.Error("Failed to create product request from CSV:", err)
			continue
		}

		// Send the product data to the API endpoint
		err = sendProductToAPI(productRequest)
		if err != nil {
			logrus.Error("Failed to send product to API:", err)
		}
	}
}

// createProductRequestFromCSV converts a CSV row into an AddProductRequest
func createProductRequestFromCSV(row []string) (*AddProductRequest, error) {
	price, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid price format: %v", err)
	}

	product := &AddProductRequest{
		Title:       row[3],
		Description: row[11], // Assuming description is at index 11
		Department:  row[0],
		PurchaseInfo: struct {
			Price float64 `json:"price"`
			Link  string  `json:"link"`
		}{
			Price: price,
			Link:  row[5], // Assuming purchase link is at index 5
		},
		ImageLinks: []string{row[6]}, // Assuming image is at index 6
		CompanyCertifications: []struct {
			Name string `json:"name"`
		}{
			{Name: row[8]},  // Certification 1
			{Name: row[9]},  // Certification 2
			{Name: row[10]}, // Certification 3
		},
	}

	return product, nil
}

// sendProductToAPI sends the product data to the API as a POST request
func sendProductToAPI(product *AddProductRequest) error {
	apiURL := "http://localhost:8080/api/products" // Replace with your actual API URL

	// Convert the product request to JSON
	jsonData, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product request to JSON: %v", err)
	}

	// Send a POST request with the JSON data
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send product to API: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API responded with status: %s", resp.Status)
	}

	fmt.Printf("Product '%s' successfully sent to API\n", product.Title)
	return nil
}
