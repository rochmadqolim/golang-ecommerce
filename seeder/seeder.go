// seeder/seeder.go

package seeder

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rochmadqolim/golang-ecommerce/models" // Ubah dengan import yang sesuai
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	seedData, err := os.ReadFile("seeder/products.json")
	if err != nil {
		return fmt.Errorf("failed to read seed data: %w", err)
	}

	var products []models.Product
	if err := json.Unmarshal(seedData, &products); err != nil {
		return fmt.Errorf("failed to unmarshal seed data: %w", err)
	}

	for _, product := range products {
		// Check if the corresponding category already exists
		var category models.Category
		if err := db.Where("name = ?", product.CategoryName).First(&category).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return fmt.Errorf("failed to query category: %w", err)
			}
			// Category doesn't exist, create a new one
			newCategory := models.Category{Name: product.CategoryName}
			if err := db.Create(&newCategory).Error; err != nil {
				return fmt.Errorf("failed to create category: %w", err)
			}
			category = newCategory
		}

		product.CategoryName = category.Name

		if err := db.Create(&product).Error; err != nil {
			return fmt.Errorf("failed to seed product: %w", err)
		}
	}

	return nil
}

