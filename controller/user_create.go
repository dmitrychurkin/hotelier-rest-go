package controller

import (
	"github.com/dmitrychurkin/hotelier-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserCreate ...
func UserCreate(c *gin.Context) {
	if db, ok := c.Keys["DB"].(*gorm.DB); ok {
		db.Create(&models.User{Name: "Dima", Age: 32})
	}
}
