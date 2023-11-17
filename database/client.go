package database

import (
	"github.com/gugazimmermann/touch-events-api/models"
	"github.com/gugazimmermann/touch-events-api/utils"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {

	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if dbError != nil {
		utils.Logger.Panic("Cannot connect to DB", zap.Error(dbError))
	}
}

func Migrate() {
	Instance.AutoMigrate(&models.Login{})
}
