package main

import (
	"./app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load Frontend HTML
	r.LoadHTMLGlob("app/frontend/*")

	// Routing (/app/router.go)
	app.RouteCollections(r)

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
