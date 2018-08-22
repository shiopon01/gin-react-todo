package main

import (
	"github.com/gin-gonic/gin"

	c "github.com/shiopon01/gin-sample/server/controllers"
)

// RouteCollections is app main router
func RouteCollections(router *gin.Engine) {
	// Favicon
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// 404 Page
	router.NoRoute(func(c *gin.Context) {
	})

	// Add Routes
	v1 := router.Group("/api")

	c.UserRegister(v1.Group("/user"))
	c.TodoRegister(v1.Group("/todo"))
}
