package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rochmadqolim/golang-ecommerce/config"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/models"
	"github.com/rochmadqolim/golang-ecommerce/responses"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var regisCustomer models.Customer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&regisCustomer); err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	newCustomer := models.Customer {
		Fullname: regisCustomer.Fullname,
		Email: regisCustomer.Email,
		Password: regisCustomer.Password,
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newCustomer.Password), bcrypt.DefaultCost)
	newCustomer.Password = string(hashPassword)

	if err := db.Create(&newCustomer).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		responses.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "signup successfully"}
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