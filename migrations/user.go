package migrations

import (
	"github.com/dmitrychurkin/hotelier-golang/models"
	"github.com/jinzhu/gorm"
)

// RunMigrations ...
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
