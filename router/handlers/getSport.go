package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSport(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "GET SPORT"})
}
