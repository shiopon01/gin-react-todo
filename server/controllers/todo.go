package controllers

import (
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/shiopon01/gin-sample/server/models"
)

// Todo ...
type Todo struct {
}

// NewTodo ...
func NewTodo() Todo {
	return Todo{}
}

// Get ...
func (c Todo) Get(n int) interface{} {
	repo := models.NewTodosRepository()
	todo := repo.GetByID(n)
	log.Println("LOG:", todo)
	return todo
}

// Get ...
func (c Todo) GetAll() {
	repo := models.NewTodosRepository()
	repo.GetAll()
	return
}

// TodoRegister router
func TodoRegister(router *gin.RouterGroup) {

	// GET /
	// return normal message
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "TODOS"})
	})

	// GET /list
	router.GET("/list", func(c *gin.Context) {
		ctrl := NewTodo()
		ctrl.GetAll()

		c.JSON(200, gin.H{"status": "TUKKOMI"})
	})

	router.GET("/detail/:id", func(c *gin.Context) {
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

		ctrl := NewTodo()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{"status": "Not found"})
			return
		}

		c.JSON(200, result)
	})
}
