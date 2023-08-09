package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/seeder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to create connection to database")
	}

	db.AutoMigrate(&models.Cart{}, &models.Category{},&models.CartItem{}, &models.Customer{}, &models.Order{}, &models.Product{})

	// Seed products
	err = seeder.SeedProducts(db)
	if err != nil {
		fmt.Println("Seeder failed:", err)
	}
	
	return db
	
}

func CloseConnection(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		panic("failed to close database connection")
	}

	database.Close()
}