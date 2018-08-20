package app

import (
	"github.com/gin-gonic/gin"

	c "./controllers"
)

// RouteCollections is app main router
func RouteCollections(router *gin.Engine) {
	v1 := router.Group("/")

	// Routing (import controller route)

	c.WelcomeRegister(v1.Group("/"))
	c.UserRegister(v1.Group("/user"))
}
