package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
)

// Add item product to cartitem
func AddCartItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.CartItem

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newItem); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	// Pastikan cart dengan ID yang diberikan sudah ada
	var cart models.Cart
	if err := db.First(&cart, newItem.CartID).Error; err != nil {
		response := map[string]string{"message": "Cart not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	// Hitung subtotal dan update item produk
	newItem.SubTotal = newItem.Quantity * newItem.Price

	

	// Buat item produk baru dalam cart
	if err := db.Create(&newItem).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Item added to cart successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}



// Delate cart item
func DeleteCartItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemIDStr := vars["id"]

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var cartItem models.CartItem
	if err := db.First(&cartItem, cartItemIDStr).Error; err != nil {
		response := map[string]string{"message": "Cart item not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	// Mendapatkan cart yang sesuai
	var cart models.Cart
	if err := db.Preload("CartItems").First(&cart, cartItem.CartID).Error; err != nil {
		response := map[string]string{"message": "Cart not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	// Menghapus cart item
	if err := db.Delete(&cartItem).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Hitung total amount baru berdasarkan cart items yang tersisa
	totalAmount := 0
	for _, ci := range cart.CartItems {
		totalAmount += ci.SubTotal
	}

	// Update total amount pada cart dan simpan ke dalam database
	cart.TotalAmount = totalAmount
	if err := db.Save(&cart).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Cart item deleted successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}



