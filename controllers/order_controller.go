package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)

// Create order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newOrder); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	// Ensure the specified cart exists
	var cart models.Cart
	if err := db.First(&cart, newOrder.CartID).Error; err != nil {
		response := map[string]string{"message": "Cart not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	// Create a new order
	order := models.Order{
		Fullname:    newOrder.Fullname,
		Address:     newOrder.Address,
		Cart:        cart,
		CartID:      newOrder.CartID,
		TotalAmount: cart.TotalAmount,
	}

	if err := db.Create(&order).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Checkout created successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}

// Status order
func GetOrderStatusByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIDStr := vars["id"]
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		response := map[string]string{"message": "Invalid order ID"}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var order models.Order
	if err := db.Preload("Cart.Customer").First(&order, orderID).Error; err != nil {
		response := map[string]string{"message": "Order not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := map[string]interface{}{
		"message": "Order retrieved successfully",
		"order":   order,
	}
	responses.ResponseJSON(w, http.StatusOK, response)
}