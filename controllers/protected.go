package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Protected(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "protected pong"})
}
