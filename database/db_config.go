package database

import (
	"sync"

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

func Connect(dsn string) *DbConn {

	lock.Lock()
	defer lock.Unlock()
	if instance == nil {

		// Open a database connection
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		instance = &DbConn{conn: db}
	}
	return instance
}

func GetDbConn() *gorm.DB {
	if instance.conn == nil {
		panic("db connection not established")
	}
	return instance.conn
}

func PerformMigrations() {
	db := GetDbConn()
	// Migrate the schema

	_  = db.Migrator().DropTable(
		&models.Category{},
		&models.Certification{},
		&models.Company{},
		&models.Favorite{},
		&models.Image{},
		&models.PurchaseInfo{},
		&models.User{},
		&models.CompanyCertification{},
		&models.ProductCertification{},
		&models.ProductCategories{},
		&models.HomeGardenProduct{},
		&models.ClothingAccessoriesProduct{},
		&models.HealthBathBeautyProduct{},
		&models.ToysKidsBabiesProduct{},
	)

	err := db.AutoMigrate(
		&models.Category{},
		&models.Certification{},
		&models.Company{},
		&models.Favorite{},
		&models.Image{},
		&models.PurchaseInfo{},
		&models.User{},
		&models.CompanyCertification{},
		&models.ProductCertification{},
		&models.ProductCategories{},
		&models.HomeGardenProduct{},
		&models.ClothingAccessoriesProduct{},
		&models.HealthBathBeautyProduct{},
		&models.ToysKidsBabiesProduct{},
	)

	if err != nil {
		panic("unable to perform migrations...")
	}
}
