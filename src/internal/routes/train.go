package routes

import "github.com/gin-gonic/gin"

func RegisterTrainRoutes(router *gin.RouterGroup, c *Container) {
	// router.GET("/trains", c.TrainController.GetTrains)

	router.POST("/train", c.TrainController.CreateTrain)
	// router.PUT("/train", c.TrainController.UpdateTrain)

	// router.GET("/train/:trainId", c.TrainController.GetTrainByID)
	// router.DELETE("/train/:trainId", c.TrainController.DeleteTrain)
}
