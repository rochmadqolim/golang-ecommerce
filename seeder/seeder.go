// seeder/seeder.go

package seeder

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rochmadqolim/golang-ecommerce/models"
	"gorm.io/gorm"
)

// Seeding product to database
func SeedProducts(db *gorm.DB) error {

	seedData, err := os.ReadFile("seeder/products.json")
	if err != nil {
		return fmt.Errorf("failed to read seed data: %w", err)
	}

	var products []models.Product
	if err := json.Unmarshal(seedData, &products); err != nil {
		return fmt.Errorf("failed to unmarshal seed data: %w", err)
	}

	// Iterate product data and insert into database.
	for _, product := range products {
		var category models.Category
		if err := db.Where("name = ?", product.CategoryName).First(&category).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return fmt.Errorf("failed to query category: %w", err)
			}

			newCategory := models.Category{Name: product.CategoryName}
			if err := db.Create(&newCategory).Error; err != nil {
				return fmt.Errorf("failed to create category: %w", err)
			}
			
			category = newCategory
		}

		// Save category name to category
		product.CategoryName = category.Name

		// Insert product into database
		if err := db.Create(&product).Error; err != nil {
			return fmt.Errorf("failed to seed product: %w", err)
		}
	}

	return nil
}

