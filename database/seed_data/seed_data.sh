package seed_data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	ID      int
	Name    string
	Age     int
	Address string
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

		age, _ := strconv.Atoi(row[2])

		person := Person{
			Name:    row[0],
			Address: row[1],
			Age:     age,
		}

		// Insert the person into the database
		result := db.Create(&person)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
	}

	fmt.Println("CSV file seeded into database.")
}
