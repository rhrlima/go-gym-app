package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterExerciseRoutes(router *gin.RouterGroup, c *Container) {
	router.POST("/exercise", c.ExerciseController.CreateExercise)
	router.GET("/exercise/:exerciseId", c.ExerciseController.GetExerciseById)
	router.GET("/exercises", c.ExerciseController.GetExercises)
}
