// db.go

package graph

import (
	"fmt"
	"os"

	"github.com/howstrongiam/backend/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host string
	port string
	user string
	name string
    pass string
}

func Connect() *gorm.DB {

	dbEnvConfig := dbConfig{
		host: os.Getenv("DB_HOST"),
		port: os.Getenv("DB_PORT"),
		user: os.Getenv("DB_USER"),
		name: os.Getenv("DB_NAME"),
        pass: os.Getenv("DB_PASS"),
	}
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbEnvConfig.host, dbEnvConfig.user, dbEnvConfig.pass, dbEnvConfig.name, dbEnvConfig.port)
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN: dsn,
        PreferSimpleProtocol: true, // disables implicit prepared statement usage
      }), &gorm.Config{})
    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    db.AutoMigrate(&model.User{})

	return db
}
