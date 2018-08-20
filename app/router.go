package router

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"

	"./controllers"
)

// RouteRegister is application router (good naming)
func RouteRegister(router *gin.RouterGroup) {

	// Root
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ROUTER PAGE"})
	})

	router.GET("/user/:id", func(c *gin.Context) {
		n := c.Param("id")

		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}

		if id <= 0 {
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})
			return
		}

		ctrl := controllers.NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{"status": "Not found"})
			return
		}

		c.JSON(200, result)
	})
}
