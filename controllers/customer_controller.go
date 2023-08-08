package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/config"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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





func Login(w http.ResponseWriter, r *http.Request) {

	var loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginCredentials); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var customer models.Customer
	if err := db.Where("email = ?", loginCredentials.Email).First(&customer).Error; err != nil {

		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "invalid email or password"}
			responses.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			responses.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(loginCredentials.Password)); err != nil {
		response := map[string]string{"message": "invalid email or password"}
		responses.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}
	

	// Generate and return JWT token on successful login
	expTime := time.Now().Add(time.Minute * 15)
	claims := &config.JWTClaim{
		Email: customer.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "golang-ecommerce",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]interface{}{
		"id":      customer.ID,
		"fullname":      customer.Fullname,
		"email":      customer.Email,
	}
	
	responses.ResponseJSON(w, http.StatusOK, response)

	
}

// Validate customer
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerIDStr := vars["id"]
	customerID, err := strconv.ParseUint(customerIDStr, 10, 32)
	if err != nil {
		response := map[string]string{"message": "Invalid customer ID"}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	var customer models.Customer
	if err := db.Preload("Cart.CartItems").First(&customer, uint32(customerID)).Error; err != nil {
		response := map[string]string{"message": "Customer not found"}
		responses.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := map[string]interface{}{
		"message":  "Customer retrieved successfully",
		"customer": customer,
	}
	responses.ResponseJSON(w, http.StatusOK, response)
}
