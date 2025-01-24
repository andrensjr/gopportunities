package router

import (
	"gopportunities/router/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	routerApi := router.Group("/api")

	{
		routerApi.GET("/esporte", handlers.GetSport)
		routerApi.POST("/esporte", handlers.CreateSport)
		routerApi.DELETE("/esporte", handlers.DeleteSport)
		routerApi.PUT("/esporte", handlers.UpdateSport)
		routerApi.GET("/esportes", handlers.GetSports)
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
