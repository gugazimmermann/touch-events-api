package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gugazimmermann/touch-events-api/auth"
	"github.com/gugazimmermann/touch-events-api/database"
	"github.com/gugazimmermann/touch-events-api/models"
	"github.com/gugazimmermann/touch-events-api/utils"
	"go.uber.org/zap"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Generate(context *gin.Context) {
	var request TokenRequest
	var login models.Login

	if err := context.ShouldBindJSON(&request); err != nil {
		utils.Logger.Error("Error binding JSON", zap.Error(err))
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Where("email = ?", request.Email).First(&login)

	if record.Error != nil {
		utils.Logger.Error("Error querying the database", zap.Error(record.Error))
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := login.CheckPwd(request.Password)

	if credentialError != nil {
		utils.Logger.Error("Invalid credentials")
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(login)

	if err != nil {
		utils.Logger.Error("Error generating JWT", zap.Error(err))
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
