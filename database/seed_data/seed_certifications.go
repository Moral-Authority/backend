package seed_data

import (
	"encoding/csv"
	"os"

	"github.com/Moral-Authority/backend/database"
	"github.com/sirupsen/logrus"
)

type Certification struct {
	Name             string
	Logo             string
	Website          string
	Description      string
	CertifiesCompany string
	CertifiesProduct string
}

func SeedCertifications(db *database.DbConn) {
	// Open the CSV file
	file, err := os.Open("/Users/angelapurcell/src/startups/moralAuthority/backend/database/seed_data/Certifications_Amazon.csv")
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
		logrus.Error("here")
		logrus.Fatal(err)
	}

	gormDB := database.GetDbConn()

	// Iterate over the records
	for _, record := range records {
		cert := Certification{
			Name:             record[0],
			Logo:             record[1],
			Website:          record[2],
			Description:      record[3],
			CertifiesCompany: record[4],
			CertifiesProduct: record[5],
		}

		// Insert the certification into the database
		result := gormDB.Create(&cert)
		if result.Error != nil {
			logrus.Error("cert %+v", cert)
			logrus.Error(result.Error)
		}
	}

	logrus.Println("Data imported successfully.")
}
