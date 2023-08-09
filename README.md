# Golang E-Commerce

Basic E-commerce Web App built with GO

## Features

- Login and register customers
- Logout and delete customer
- Customer can view product list by product category
- Customer can add product to shopping cart
- Customers can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customers can checkout and make payment transactions
- Database configuration can be customized
- There is already a product sample

## Technology Stack

- Golang :
  - Gorilla Mux (https://github.com/gorilla/mux)
  - Gorm (http://gorm.io)
  - GoDotEnv (github.com/joho/godotenv)
  - Crypto (golang.org/x/crypto)
  - JWT (github.com/golang-jwt/jwt/v4)
- PostgreSQL
- MySql

## Installation

To use using Golang E-Commerce you need to clone the repository first, you can use the command below

```sh
git clone https://github.com/rochmadqolim/golang-ecommerce.git
```

### Create & Cofig DB Postgres or MySQL

- Create your database
- Setup .ENV Database configuration if using Postgres

  ```sh
  # Config to database
  DB_DRIVER=postgres
  DB_HOST=<your_host_db>
  DB_USER=<your_user_db>
  DB_PASSWORD=<ypur_password_db>
  DB_NAME=<your_name_db>
  DB_PORT=5432 # Default port db
  ```

- Or setup .ENV Database configuration if using MySql
  ```sh
  # Config to database
  DB_DRIVER=mysql
  DB_HOST=<your_host_db>
  DB_USER=<your_user_db>
  DB_PASSWORD=<ypur_password_db>
  DB_NAME=<your_name_db>
  DB_PORT=3306 # Default port db
  ```

## Running App

```sh
go run main.go
```

Run `go run main.go` for a dev server. Navigate to `http://localhost:8000/` can use `Postman`
