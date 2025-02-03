package routes

import "github.com/gin-gonic/gin"

func RegisterTrainSessionRoutes(router *gin.RouterGroup, c *Container) {
	router.GET("/train-sessions", c.TrainSessionController.GetTrainSessions)

	router.POST("/train-session", c.TrainSessionController.CreateTrainSession)
	router.PUT("/train-session", c.TrainSessionController.UpdateTrainSession)

	router.GET("/train-session/:trainSessionId", c.TrainSessionController.GetTrainSessionByID)
	router.DELETE("/train-session/:trainSessionId", c.TrainSessionController.DeleteTrainSession)
}
