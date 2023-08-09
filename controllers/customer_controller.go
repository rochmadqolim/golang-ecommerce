package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
	"golang.org/x/crypto/bcrypt"
)

// Register customer
func Register(w http.ResponseWriter, r *http.Request) {
	
	var newCustomer models.Customer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCustomer); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// Hash the customer's password before saving it
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newCustomer.Password), bcrypt.DefaultCost)
    if err != nil {
        response := map[string]string{"message": "Failed to hash password"}
        responses.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }
    newCustomer.Password = string(hashedPassword)

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	// Begin a transaction to ensure both customer and cart are created together
	tx := db.Begin()

	// Create the customer
	if err := tx.Create(&newCustomer).Error; err != nil {
		tx.Rollback()
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Create the cart associated with the customer
	newCart := models.Cart{CustomerID: newCustomer.ID}
	if err := tx.Create(&newCart).Error; err != nil {
		tx.Rollback()
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Commit the transaction
	tx.Commit()

	response := map[string]string{"message": "Customer registered successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}

// Login customer
func Login(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var customer models.Customer
	if err := db.Where("email = ?", request.Email).First(&customer).Error; err != nil {
		response := map[string]string{"message": "Customer not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		response := map[string]string{"message": "Invalid password"}
		responses.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Create a response containing customer information
	response := map[string]interface{}{
		"id":       customer.ID,
		"fullname": customer.Fullname,
		"email":    customer.Email,
	}

	responses.ResponseJSON(w, http.StatusOK, response)
}

// Deleta customer
func DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerIDStr := vars["id"]
	customerID, _ := strconv.Atoi(customerIDStr)

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var customer models.Customer
	if err := db.First(&customer, customerID).Error; err != nil {
		response := map[string]string{"message": "Customer not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	// Hapus Cart dan CartItem terkait dengan Customer
	var cart models.Cart
	if err := db.Where("customer_id = ?", customer.ID).First(&cart).Error; err != nil {
		response := map[string]string{"message": "Failed to retrieve cart"}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if err := db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		response := map[string]string{"message": "Failed to delete cart items"}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if err := db.Delete(&cart).Error; err != nil {
		response := map[string]string{"message": "Failed to delete cart"}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Hapus Customer
	if err := db.Delete(&customer).Error; err != nil {
		response := map[string]string{"message": "Failed to delete customer"}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Customer deleted successfully"}
	responses.ResponseJSON(w, http.StatusOK, response)
}