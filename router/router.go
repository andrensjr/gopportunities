package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	InitializeRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
