package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSport(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Delete sport"})
}
