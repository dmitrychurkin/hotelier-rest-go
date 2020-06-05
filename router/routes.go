package router

import (
	"github.com/dmitrychurkin/hotelier-golang/controller"
	"github.com/gin-gonic/gin"
)

// Config ...
func Config(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/user", controller.UserCreate)
	}
}
