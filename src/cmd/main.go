package main

import (
	"fmt"
	"gym-app/internal/db"
	"gym-app/internal/routes"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	c := routes.NewContainer(dbConnection)

	routes.RegisterRoutes(server, c)

	fmt.Println("Server running on port", PORT)
	server.Run(":" + PORT)
}
