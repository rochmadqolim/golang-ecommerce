package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)

// Get cart customer by id
func GetCartCustomerByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cartIDStr := vars["id"]

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var cart models.Cart
	if err := db.Preload("CartItems").First(&cart, cartIDStr).Error; err != nil {
		response := map[string]string{"message": "Cart not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)

		return
	}

	// Calculate the total amount in cart
	totalAmount := 0
	for _, cartItem := range cart.CartItems {
		totalAmount += cartItem.SubTotal
	}

	// Update total amount in cart
	cart.TotalAmount = totalAmount
	if err := db.Save(&cart).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"cart":    cart,
	}

	responses.ResponseJSON(w, http.StatusOK, response)
}