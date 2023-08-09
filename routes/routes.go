package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rochmadqolim/golang-ecommerce/controllers"
	"github.com/rochmadqolim/golang-ecommerce/database"
)

func Run() {

	db := database.DatabaseConnection()
	defer database.CloseConnection(db)

	r := mux.NewRouter()
	// Router customer
	r.HandleFunc("/register", controllers.RegisterCustomer).Methods("POST")
	r.HandleFunc("/login", controllers.LoginCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", controllers.DeleteCustomerByID).Methods("DELETE")

	// Router product
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")    // Endpoint to create products
	r.HandleFunc("/products", controllers.GetAllProducts).Methods("GET") // Endpoint to gel all product
	
	// Router category
	r.HandleFunc("/categories", controllers.GetAllCategories).Methods("GET")
	r.HandleFunc("/categories/{category_name}", controllers.GetProductsByCategory).Methods("GET")

	// Router cart item
	r.HandleFunc("/item", controllers.AddCartItem).Methods("POST")
	r.HandleFunc("/item/{id}", controllers.DeleteCartItemByID).Methods("DELETE")

	//Router cart
	r.HandleFunc("/carts/{id}", controllers.GetCartCustomerByID).Methods("GET")

	// Router order
	r.HandleFunc("/checkout", controllers.CreateOrder).Methods("POST")
	r.HandleFunc("/checkout/{id}", controllers.GetOrderStatusByID).Methods("GET")

	fmt.Println("Listen to port: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))

}