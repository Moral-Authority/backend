package seed_data

import (
    "encoding/csv"
    "os"

    "github.com/Moral-Authority/backend/database"
    "github.com/sirupsen/logrus"
)

func SeedCertifications(db *database.DbConn) {
    // Open the CSV file
    file, err := os.Open("/Users/angelapurcell/src/startups/moralAuthority/backend/database/seed_data/Certifications_Amazon.csv")
    if err != nil {
        logrus.Fatal(err)
    }
    defer file.Close()

    // Create a new CSV reader
    reader := csv.NewReader(file)

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

	gormDB := database.GetDbConn()

    // Iterate over the records
    for _, record := range records {
        // Prepare a SQL statement
        stmt := `INSERT INTO certifications(Name, Logo, Website, Description, CertifiesCompany, CertifiesProduct) VALUES (?, ?, ?, ?, ?, ?)`

        // Execute the SQL statement
		gormDB.Exec(stmt, record[0], record[1], record[2], record[3], record[4], record[5])
    }

    logrus.Println("Data imported successfully.")
}