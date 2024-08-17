package seed

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"

    "github.com/Moral-Authority/backend/models"
    "github.com/volatiletech/null/v8"
    "gorm.io/gorm"
)

func SeedCompanies(db *gorm.DB) {
    // Open the CSV file
    file, err := os.Open("./seed/all-bcorps.csv")
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
