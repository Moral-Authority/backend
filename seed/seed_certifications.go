package seed

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

func SeedCertifications(db *gorm.DB) {
	// Open the CSV file
	file, err := os.Open("./seed/Certifications_Amazon.csv")
	if err != nil {
		logrus.Fatal(err)
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
        certifiesCompany, err := strconv.ParseBool(record[4])
        if err != nil {
            logrus.Fatal(err)
        }

        certifiesProduct, err := strconv.ParseBool(record[5])
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
			logrus.Error("cert %+v", cert)
			logrus.Error(result.Error)
		}
	}

	logrus.Println("Certifications seeded into database.")
}
