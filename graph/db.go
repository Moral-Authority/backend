// db.go

package graph

import (
	"fmt"
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
		host: "postgres",
		port: "5432",
		user: "local",
		name: "local",
		pass: "postgres",
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbEnvConfig.host, dbEnvConfig.user, dbEnvConfig.pass, dbEnvConfig.name, dbEnvConfig.port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	return db
}
