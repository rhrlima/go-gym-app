package main

import (
	"gym-app/controller"
	"gym-app/db"
	"gym-app/repository"
	"gym-app/usecase"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Repository
	ExerciseRepository := repository.NewExerciseRepository(dbConnection)
	TagRepository := repository.NewTagRepository(dbConnection)

	// Usecase
	ExerciseUsecase := usecase.NewExerciseUsecase(ExerciseRepository)
	TagUsecase := usecase.NewTagUsecase(TagRepository)

	// Controller
	ExerciseController := controller.NewExerciseController(ExerciseUsecase)
	TagController := controller.NewTagController(TagUsecase)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	server.POST("/exercise", ExerciseController.CreateExercise)
	server.GET("/exercises", ExerciseController.GetExercises)
	server.GET("/exercise/:exerciseId", ExerciseController.GetExerciseById)

	server.POST("/tag", TagController.CreateTag)
	server.GET("/tags", TagController.GetTags)
	server.GET("/tag/:tagName", TagController.GetTagByName)

	server.Run(":"+PORT)
}
