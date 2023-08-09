package controllers

import (
	"net/http"

	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)

// Get all categories
func GetAllCategories(w http.ResponseWriter, r *http.Request) {

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	categoryNames := make([]string, len(categories))
	for i, category := range categories {
		categoryNames[i] = category.Name
	}

	response := map[string]interface{}{
		"categories": categoryNames,
	}
	responses.ResponseJSON(w, http.StatusOK, response)
}