package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine, c *Container) {
	api := server.Group("/api")

	RegisterExerciseRoutes(api, c)
	RegisterTagRoutes(api, c)
	RegisterTrainRoutes(api, c)
}
