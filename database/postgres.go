package database

import (
	"fmt"
	"github.com/howstrongiam/backend/config"
	"github.com/howstrongiam/backend/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func ConnectDatabase() {
	if dbConnection == nil {
		dbEnvConfig := config.PostgresConfig{
			Host: "postgres",
			Port: "5432",
			User: "local",
			Name: "local",
			Pass: "postgres",
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbEnvConfig.Host, dbEnvConfig.User, dbEnvConfig.Pass, dbEnvConfig.Name, dbEnvConfig.Port)
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		if err != nil {

			fmt.Println("Error", err)
			panic(err)
		}

		fmt.Println("Connected to Postgres")

		db.AutoMigrate(&model.User{})
		dbConnection = db
	}
}

func GetDatabase() *gorm.DB {
	return dbConnection
}
