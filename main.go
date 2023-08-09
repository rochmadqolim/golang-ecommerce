package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/controllers"
	"github.com/rochmadqolim/golang-ecommerce/database"
)

func main() {

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)


	r := mux.NewRouter()
	// Router customer
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/customers/{id:[0-9]+}", controllers.DeleteCustomerByID).Methods("DELETE")

	// Router product
	r.HandleFunc("/productcreate", controllers.CreateProduct).Methods("POST") // Endpoint to create products
	r.HandleFunc("/products/product", controllers.GetAllProducts).Methods("GET") // Endpoint to gel all product
	r.HandleFunc("/products/category/{category_name}", controllers.GetProductsByCategoryName).Methods("GET")
	
	// Router category
	r.HandleFunc("/categories", controllers.GetAllCategories).Methods("GET")

	// Router cart item
	r.HandleFunc("/addCartItem", controllers.AddCartItem).Methods("POST")
	r.HandleFunc("/cartitem/{id}", controllers.DeleteCartItem).Methods("DELETE")
	
	//Router cart
	r.HandleFunc("/carts/{id}", controllers.GetCartByID).Methods("GET")
	
	// Router order
	r.HandleFunc("/checkout", controllers.CreateCheckout).Methods("POST")
	r.HandleFunc("/checkout/{id:[0-9]+}", controllers.GetOrderStatus).Methods("GET")

	fmt.Println("listen in port: "+ os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+ os.Getenv("PORT"), r))
}