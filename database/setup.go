package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/seeder"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	// Load .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbDriver := os.Getenv("DB_DRIVER")

	var dsn string
	var err error
	var db *gorm.DB

	if dbDriver == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if dbDriver == "postgres" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		panic("unsupported database driver")
	}

	if err != nil {
		panic("failed to create connection to database")
	}

	// Tables migration
	err = db.AutoMigrate(&models.Cart{}, &models.Category{}, &models.CartItem{}, &models.Customer{}, &models.Order{}, &models.Product{})
	if err != nil {
		panic("failed to migrate database")
	}

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