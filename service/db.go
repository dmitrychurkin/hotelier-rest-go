package service

import (
	"github.com/dmitrychurkin/hotelier-golang/migration"
	"github.com/dmitrychurkin/hotelier-golang/util"
	"github.com/jinzhu/gorm"
)

// Database ...
func Database(db *gorm.DB) {
	db.LogMode(util.GetEnv("GIN_MODE") != "release")
	migration.RunMigrations(db)
}
