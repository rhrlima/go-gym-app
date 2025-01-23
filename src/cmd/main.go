package main

import (
	"gym-app/db"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	server.Run(":"+PORT)
}
