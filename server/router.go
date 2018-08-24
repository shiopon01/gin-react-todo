package main

import (
	"github.com/gin-gonic/gin"

	"github.com/shiopon01/gin-react-todo/server/controllers"
)

// RouteCollections is app main router
func RouteCollections(router *gin.Engine) {
	// Favicon
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// 404 Page
	// router.NoRoute(func(c *gin.Context) {
	// })

	// Add API Routes
	v1 := router.Group("/api")
	controllers.TodoRegister(v1.Group("/todo"))
}
