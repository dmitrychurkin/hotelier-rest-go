package service

import (
	"github.com/dmitrychurkin/hotelier-golang/migrations"
	"github.com/dmitrychurkin/hotelier-golang/util"
	"github.com/jinzhu/gorm"
)

// Database ...
func Database(db *gorm.DB) {
	db.LogMode(util.GetEnv("ENV") != "production")
	migrations.RunMigrations(db)
}
