package seed_data

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Company struct {
    CompanyName       string
    DateCertified     string
    Industry          string
    IndustryCategory  string
    Country           string
    State             string
    City              string
    Sector            string
    Size              string
    Website           string
}

func main() {
    // Open the CSV file
    file, err := os.Open("all-bcorps.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create a CSV reader
    reader := csv.NewReader(file)

    // Connect to the database
    dsn := "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Read each row of the CSV file and insert it into the database
    for {
        row, err := reader.Read()
        if err != nil {
            break
        }

        company := Company{
            CompanyName:      row[0],
            DateCertified:    row[1],
            Industry:         row[2],
            IndustryCategory: row[3],
            Country:          row[4],
            State:            row[5],
            City:             row[6],
            Sector:           row[7],
            Size:             row[8],
            Website:          row[9],
        }

        // Insert the company into the database
        result := db.Create(&company)
        if result.Error != nil {
            fmt.Println(result.Error)
        }
    }

    fmt.Println("CSV file seeded into database.")
}