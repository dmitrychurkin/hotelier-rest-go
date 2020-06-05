package migration

import (
	"github.com/dmitrychurkin/hotelier-golang/model"
	"github.com/jinzhu/gorm"
)

// RunMigrations ...
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
