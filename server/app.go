package server

import (
	"github.com/gin-gonic/gin"
)

// Create ...
func Create() *gin.Engine {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "t"})
	})
	return server
}
