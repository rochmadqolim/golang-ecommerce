package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/controllers"
	"github.com/rochmadqolim/golang-ecommerce/database"
	"github.com/rochmadqolim/golang-ecommerce/middleware"
)

func Run() {

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	r := mux.NewRouter()

	// Middleware
	protectedRouter := r.PathPrefix("/protected").Subrouter()
	protectedRouter.Use(middleware.JWTMiddleware)
	
	// Customer router
	r.HandleFunc("/register", controllers.RegisterCustomer).Methods("POST")
	r.HandleFunc("/login", controllers.LoginCustomer).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutCustomer).Methods("GET")
	protectedRouter.HandleFunc("/customers/{id}", controllers.DeleteCustomerByID).Methods("DELETE")
	
	// Product router
	protectedRouter.HandleFunc("/products", controllers.CreateProduct).Methods("POST")   
	protectedRouter.HandleFunc("/products", controllers.GetAllProducts).Methods("GET") 
	
	// Category router
	protectedRouter.HandleFunc("/categories", controllers.GetAllCategories).Methods("GET")
	protectedRouter.HandleFunc("/categories/{category_name}", controllers.GetProductsByCategory).Methods("GET")
	
	// Cart router
	protectedRouter.HandleFunc("/cart/{id}", controllers.GetCartCustomerByID).Methods("GET")
	
	// Item router
	protectedRouter.HandleFunc("/item", controllers.AddCartItem).Methods("POST")
	protectedRouter.HandleFunc("/item/{id}", controllers.DeleteCartItemByID).Methods("DELETE")
	
	// Order Router
	protectedRouter.HandleFunc("/checkout", controllers.CreateOrder).Methods("POST")
	protectedRouter.HandleFunc("/checkout/{id}", controllers.GetOrderStatusByID).Methods("GET")

	fmt.Println("Listen to port: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))

}