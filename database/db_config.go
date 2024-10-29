package database

import (
	"log"
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

	// _ = db.Migrator().DropTable(
	// 	&models.Category{},
	// 	&models.Certification{},
	// 	&models.Company{},
	// 	&models.Favorite{},
	// 	&models.Image{},
	// 	&models.PurchaseInfo{},
	// 	&models.User{},
	// 	&models.CompanyCertification{},
	// 	&models.ProductCertification{},
	// 	&models.ProductCategories{},
	// 	&models.HomeGardenProduct{},
	// 	&models.ClothingAccessoriesProduct{},
	// 	&models.HealthBathBeautyProduct{},
	// 	&models.ToysKidsBabiesProduct{},
	// )

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

func WipeDatabase(db *gorm.DB) {
	// Get the list of all tables in the database
	var tables []string
	err := db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables).Error
	if err != nil {
		log.Fatal("Failed to get table names:", err)
	}

	// Truncate each table
	for _, table := range tables {
		err := db.Exec("TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE").Error
		if err != nil {
			log.Fatal("Failed to truncate table", table, ":", err)
		}
	}

	log.Println("All tables wiped.")
}
