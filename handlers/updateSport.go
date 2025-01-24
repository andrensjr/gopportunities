package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSport(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Update sport"})
}
