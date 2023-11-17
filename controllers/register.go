package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gugazimmermann/touch-events-api/database"
	"github.com/gugazimmermann/touch-events-api/models"
)

func Register(context *gin.Context) {

	var login models.Login

	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	login.TrimEmail(login.Email)

	if err := login.HashPwd(login.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&login)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"id": login.ID, "email": login.Email})
}
