package server

import (
	"github.com/dmitrychurkin/hotelier-golang/connection"
	"github.com/dmitrychurkin/hotelier-golang/middleware"
	"github.com/dmitrychurkin/hotelier-golang/router"
	"github.com/gin-gonic/gin"
)

// Create ...
func Create(connect *connection.Connect) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Ctx(connect.DB))
	router.Config(engine)
	return engine
}
