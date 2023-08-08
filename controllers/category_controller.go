package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)

// Create category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCategory); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	// Buat kategori baru
	if err := db.Create(&newCategory).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Category created successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}

// Get categories
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"message":   "Categories retrieved successfully",
		"categories": categories,
	}
	responses.ResponseJSON(w, http.StatusOK, response)
}
