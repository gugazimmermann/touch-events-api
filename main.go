package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gugazimmermann/touch-events-api/controllers"
	"github.com/gugazimmermann/touch-events-api/database"
	"github.com/gugazimmermann/touch-events-api/middlewares"
	"github.com/gugazimmermann/touch-events-api/utils"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func initBD() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Fatal("Failed to load env", zap.Error(err))
	}

	db_u := os.Getenv("MYSQL_USER")
	db_pwd := os.Getenv("MYSQL_PASSWORD")
	db_h := os.Getenv("MYSQL_HOST")
	db_p := os.Getenv("MYSQL_PORT")
	db_db := os.Getenv("MYSQL_DATABASE")

	db_conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_u, db_pwd, db_h, db_p, db_db)

	database.Connect(db_conn)
	database.Migrate()
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	api := router.Group("/api")

	api.GET("/public", controllers.Public)
	api.POST("/login/register", controllers.Register)
	api.POST("/login", controllers.Generate)

	protected := router.Group("/api/secured")
	protected.Use(middlewares.Auth())
	protected.GET("/ping", controllers.Protected)

	return router
}

func main() {
	utils.InitializeLogger()
	utils.Logger.Info("Init Touch Events")
	initBD()
	router := initRouter()
	router.Run(":5000")
}
