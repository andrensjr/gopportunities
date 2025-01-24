package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSport(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Create sport"})
}
