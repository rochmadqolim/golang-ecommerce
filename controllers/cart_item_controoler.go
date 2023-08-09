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

    // Cart id check
    var cart models.Cart
    if err := db.First(&cart, newItem.CartID).Error; err != nil {
        response := map[string]string{"message": "Cart not found"}
        responses.ResponseJSON(w, http.StatusNotFound, response)
        return
    }

    // Get product details by product name
    var product models.Product
    if err := db.Where("name = ?", newItem.ProductName).First(&product).Error; err != nil {
        response := map[string]string{"message": "Product not found"}
        responses.ResponseJSON(w, http.StatusNotFound, response)
        return
    }

    // Fill in Price based on product prices
    newItem.Price = product.Price

    // Calculate subtotals and product item updates
    newItem.SubTotal = newItem.Quantity * newItem.Price

    // Add product item in cart
    if err := db.Create(&newItem).Error; err != nil {
        response := map[string]string{"message": err.Error()}
        responses.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    response := map[string]string{"message": "Item added to cart successfully"}
    responses.ResponseJSON(w, http.StatusOK, response)
}


// Delate cart items
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

	// Hapus cart item dari database
	if err := db.Delete(&cartItem).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Cart item deleted successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}