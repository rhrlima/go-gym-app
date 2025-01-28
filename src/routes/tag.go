package routes

import "github.com/gin-gonic/gin"

func RegisterTagRoutes(router *gin.RouterGroup, c *Container) {
	router.POST("/tag", c.TagController.CreateTag)
	router.GET("/tags", c.TagController.GetTags)
	//TODO delete tag
}
