package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gugazimmermann/touch-events-api/auth"
)

func Auth() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(401, gin.H{"error": "missing access token"})
			context.Abort()
			return
		}

		err := auth.ValidateJWT(tokenString)

		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}
