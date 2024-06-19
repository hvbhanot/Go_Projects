package main

import (
	"RestAPI/db"
	"RestAPI/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

// main is the entry point of the application. It initializes the database connection,
// creates an instance of the Gin web framework, registers all routes, and starts the server.
// If any error occurs during the server startup, it prints an error message and exits the function.
// The server runs on http://localhost:8080.
func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	} // localhost:8080
}
