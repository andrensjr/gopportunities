package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSports(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "GET sportsssss"})
}
