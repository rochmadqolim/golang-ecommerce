package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)

// Create product + category
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newProduct); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	// Check if the specified category exists
	var category models.Category
	if err := db.Where("name = ?", newProduct.CategoryName).First(&category).Error; err != nil {
		// Category doesn't exist, create a new one
		newCategory := models.Category{
			Name: newProduct.CategoryName,
		}
		if err := db.Create(&newCategory).Error; err != nil {
			response := map[string]string{"message": err.Error()}
			responses.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
		category = newCategory
	}

	// Create a new product
	product := models.Product{
		CategoryName: category.Name,
		Name:         newProduct.Name,
		Price:        newProduct.Price,
		Stock:        newProduct.Stock,
	}

	if err := db.Create(&product).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Product created successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}

// Get product by category
func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryName := vars["category_name"]

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var products []models.Product
	if err := db.Where("category_name = ?", categoryName).Find(&products).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"message":  "Products retrieved successfully",
		"products": products,
	}
	responses.ResponseJSON(w, http.StatusOK, response)
}

// Get all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"message":  "Products retrieved successfully",
		"products": products,
	}
	responses.ResponseJSON(w, http.StatusOK, response)
}