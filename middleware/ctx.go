package middleware

import (
	"github.com/dmitrychurkin/hotelier-golang/glob"
	"github.com/dmitrychurkin/hotelier-golang/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Ctx ...
func Ctx(db *gorm.DB) func(c *gin.Context) {
	service.Database(db)
	return func(c *gin.Context) {
		c.Set(glob.DB, db)
		c.Next()
	}
}
