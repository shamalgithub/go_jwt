package initializer

import (
	
	"github.com/shamalgithub/go_jwt.git/models"
	"gorm.io/gorm"
)

func SyncDatabase(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}