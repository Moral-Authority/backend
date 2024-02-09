package database

import (
	"log"
	"sync"
    // "fmt"
    // "net/url"
	"github.com/Moral-Authority/backend/cmd"
	"github.com/Moral-Authority/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// lock mutex
var lock = &sync.Mutex{}

type DbConn struct {
	conn *gorm.DB
}



var instance *DbConn

func Connect(dbConfig cmd.DatabaseConfig) *DbConn {

	lock.Lock()
	defer lock.Unlock()
	if instance == nil {

		dsn := dbConfig.DatabaseConnectionUrl

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		instance = &DbConn{conn: db}
		performMigrations(instance.conn)
	}
	return instance
}

func GetDbConn() *gorm.DB {
	if instance.conn == nil {
		panic("db connection not established")
	}
	return instance.conn
}

func performMigrations(db *gorm.DB) {
	// Migrate the schema
	err := db.AutoMigrate(
		&models.Category{},
		&models.Certification{},
		&models.Company{},
		&models.Favourite{},
		&models.Image{},
		&models.LoginCredentials{},
		&models.Product{},
		&models.User{},
	)
	if err != nil {
		panic("unable to perform migrations...")
	}
}
