package routes

import "github.com/gin-gonic/gin"

func RegisterTrainRoutes(router *gin.RouterGroup, c *Container) {
	router.POST("/train", c.TrainController.CreateTrain)
}
