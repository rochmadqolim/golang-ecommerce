package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rochmadqolim/golang-ecommerce/database"
)

func main() {

	db := database.DatabaseConnection()
	database.CloseConnection(db)

	fmt.Println("database close connection")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":"+ os.Getenv("PORT"))
}