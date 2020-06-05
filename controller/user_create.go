package controller

import (
	"log"
	"net/http"

	"github.com/dmitrychurkin/hotelier-golang/glob"
	"github.com/dmitrychurkin/hotelier-golang/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type userCreateDTO struct {
	Name string `form:"name" json:"name" binding:"required,max=255,min=3"`
	Age  uint8  `form:"age" json:"age" binding:"required,lt=200,gt=0"`
}

// UserCreate ...
func UserCreate(c *gin.Context) {
	if db, ok := c.Keys[glob.DB].(*gorm.DB); ok {

		var userCreate userCreateDTO
		if err := c.ShouldBindJSON(&userCreate); err != nil {
			// TODO: replace with c.Writer.WriteHeader(http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&model.User{Name: userCreate.Name, Age: userCreate.Age})

		c.Writer.WriteHeader(http.StatusCreated)
	} else {
		log.Println("Something went wrong")
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}
