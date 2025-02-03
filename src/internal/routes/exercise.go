package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterExerciseRoutes(router *gin.RouterGroup, c *Container) {
	router.GET("/exercises", c.ExerciseController.GetExercises)

	router.POST("/exercise", c.ExerciseController.CreateExercise)
	router.PUT("/exercise", c.ExerciseController.UpdateExercise)

	router.GET("/exercise/:exerciseId", c.ExerciseController.GetExerciseByID)
	router.DELETE("/exercise/:exerciseId", c.ExerciseController.DeleteExercise)
}
