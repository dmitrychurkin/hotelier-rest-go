package server

import (
	"github.com/dmitrychurkin/hotelier-golang/connections"
	"github.com/dmitrychurkin/hotelier-golang/middleware"
	"github.com/gin-gonic/gin"
)

// Create ...
func Create(connect *connections.Connect) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Ctx(connect.DB))
	// engine.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "t"})
	// })
	return engine
}
