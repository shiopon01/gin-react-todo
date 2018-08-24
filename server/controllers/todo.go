package controllers

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shiopon01/gin-sample/server/models"
)

// TodoRegister is todo router.
// Built in the main router by `/server/router.go`
func TodoRegister(router *gin.RouterGroup) {
	router.GET("/", top)
	router.GET("/list", list)
	router.GET("/detail/:id", detail)
	router.POST("/add", add)
}

// GET /

func top(c *gin.Context) {
	c.JSON(200, gin.H{"message": "TODOS"})
}

// GET /list

func list(c *gin.Context) {
	result := models.NewTodosRepository().GetAll()

	if result == nil || reflect.ValueOf(result).IsNil() {
		c.JSON(404, gin.H{"status": "Not found"})
		return
	}
	c.JSON(200, result)
}

// GET /detail

func detail(c *gin.Context) {
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

	result := models.NewTodosRepository().GetByID(id)

	if result == nil || reflect.ValueOf(result).IsNil() {
		c.JSON(404, gin.H{"status": "Not found"})
		return
	}

	c.JSON(200, result)
}

// POST /add

func add(c *gin.Context) {
	type JSONText struct {
		Text string `json:"text" binding:"required"`
	}

	json := new(JSONText)
	c.BindJSON(&json)

	if "" == json.Text {
		c.JSON(400, gin.H{"status": "required title....."})
		return
	}

	result := models.NewTodosRepository().AddTodo(json.Text)

	if result == nil || reflect.ValueOf(result).IsNil() {
		c.JSON(404, gin.H{"status": "Not found"})
		return
	}

	c.JSON(200, result)
}
