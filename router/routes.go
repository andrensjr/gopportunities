package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	routerApi := router.Group("/api")

	{
		routerApi.GET("/esporte", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "GET"})
		})
		routerApi.POST("/esporte", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "POST"})
		})
		routerApi.DELETE("/esporte", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "DELETE"})
		})
		routerApi.PUT("/esporte", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "UPDATE"})
		})
		routerApi.GET("/esportes", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "GET ESPORTESSSS"})
		})
	}

	// router.GET("/healthcheck", func(context *gin.Context) {
	// 	context.String(http.StatusOK, "HEALTHCHECK OK !!!")
	// })

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

}
