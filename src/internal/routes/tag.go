package routes

import "github.com/gin-gonic/gin"

func RegisterTagRoutes(router *gin.RouterGroup, c *Container) {
	router.GET("/tags", c.TagController.GetTags)

	router.POST("/tag", c.TagController.CreateTag)
	router.PUT("/tag", c.TagController.UpdateTag)

	router.GET("/tag/:tagId", c.TagController.GetTagByID)
	router.DELETE("/tag/:tagId", c.TagController.DeleteTagByID)
}
