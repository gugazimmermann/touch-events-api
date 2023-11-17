package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Public(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "public pong"})
}
