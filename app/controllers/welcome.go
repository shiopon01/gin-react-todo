package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeRegister top
func WelcomeRegister(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

}
