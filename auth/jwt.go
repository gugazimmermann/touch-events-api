package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gugazimmermann/touch-events-api/models"
	"github.com/gugazimmermann/touch-events-api/utils"
	"go.uber.org/zap"
)

type JWTClaim struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		utils.Logger.Panic("JWT_SECRET is not set in the environment")
	}
	return secret
}

func GenerateJWT(login models.Login) (tokenString string, err error) {
	jwtKey := getSecret()

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		Id:    int(login.ID),
		Email: login.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(jwtKey))

	if err != nil {
		utils.Logger.Error("Error generating JWT", zap.Error(err))
		return
	}

	return
}

func ValidateJWT(signedToken string) (err error) {
	jwtKey := getSecret()

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		utils.Logger.Error("error parsing JWT", zap.Error(err))
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		utils.Logger.Error("JWTClaim", zap.Error(err))
		return err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.New("token expired")
		utils.Logger.Error("ExpiresAt", zap.Error(err))
		return err
	}

	return nil
}
