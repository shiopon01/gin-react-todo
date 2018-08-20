package main

import (
	"./app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("app/frontend/*")

	// Routing (/app/router.go)
	router.RouteRegister(r.Group("/"))

	r.Run() // listen and serve on 0.0.0.0:8080
}
