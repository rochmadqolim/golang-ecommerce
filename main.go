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
	r.HandleFunc("/customers/{id:[0-9]+}", controllers.GetCustomerByID).Methods("GET")

	// Router product
	r.HandleFunc("/products/sample", controllers.SampleProduct).Methods("POST") // Endpoint to create sample products
	r.HandleFunc("/products/product", controllers.CreateProduct).Methods("POST") // Endpoint to create products
	r.HandleFunc("/products/product", controllers.GetAllProduct).Methods("GET") // Endpoint to gel all product
	r.HandleFunc("/products/by_category/{category_id}", controllers.GetProductsByCategory).Methods("GET")
	
	// Router category
	r.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories", controllers.GetAllCategories).Methods("GET")

	// Router cart item
	r.HandleFunc("/addCartItem", controllers.AddCartItem).Methods("POST")
	r.HandleFunc("/cartitem/{id}", controllers.DeleteCartItemByID).Methods("DELETE")

	//Router cart
	r.HandleFunc("/carts/{id}", controllers.GetCartByID).Methods("GET")


	fmt.Println("listen in port: "+ os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+ os.Getenv("PORT"), r))
}