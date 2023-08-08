package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)
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

	// Pastikan kategori dengan ID yang diberikan sudah ada
	var category models.Category
	if err := db.First(&category, newProduct.CategoryID).Error; err != nil {
		response := map[string]string{"message": "Category not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	// Buat produk baru
	newProduct.CategoryID = category.ID
	if err := db.Create(&newProduct).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Product created successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}



// Sample Product=============================
func SampleProduct(w http.ResponseWriter, r *http.Request) {
	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	products := []map[string]interface{}{
		{
			"category": "fashion",
			"name":     "hoodie",
			"price":    70000,
			"stock":    1000,
		},
		{
			"category": "fashion",
			"name":     "jogger",
			"price":    45000,
			"stock":    1000,
		},
		{
			"category": "beauty",
			"name":     "facial wash",
			"price":    35000,
			"stock":    1000,
		},
		{
			"category": "beauty",
			"name":     "sunscreen",
			"price":    45000,
			"stock":    1000,
		},
		{
			"category": "education",
			"name":     "notebook",
			"price":    18000,
			"stock":    1000,
		},
	}

	// Loop through the product data and create each product in the database
	for _, productData := range products {
		categoryName := productData["category"].(string)
		name := productData["name"].(string)
		price := int(productData["price"].(int))
		stock := int(productData["stock"].(int))

		// Fetch the category by name
		var category models.Category
		if err := db.Where("name = ?", categoryName).First(&category).Error; err != nil {
			response := map[string]string{"message": "Category not found"}
			responses.ResponseJSON(w, http.StatusNotFound, response)
			return
		}

		// Create the product and associate it with the category
		newProduct := models.Product{
			CategoryID: category.ID,
			Name:       name,
			Price:      price,
			Stock:      stock,
		}
		if err := db.Create(&newProduct).Error; err != nil {
			response := map[string]string{"message": err.Error()}
			responses.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	response := map[string]string{"message": "Products created successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}

// Get product by category
func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category_id"]

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var products []models.Product
	if err := db.Where("category_id = ?", category).Find(&products).Error; err != nil {
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

// Get all product
func GetAllProduct(w http.ResponseWriter, r *http.Request) {
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
