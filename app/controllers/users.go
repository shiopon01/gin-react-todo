package controllers

import (
	"reflect"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"

	"log"
)

// User ...
type User struct {
}

// NewUser ...
func NewUser() User {
	return User{}
}

// Get ...
func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(n)
	log.Println("LOG:", user)
	return user
}

// UserRegister router
func UserRegister(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ROUTER PAGE"})
	})

	router.GET("/:id", func(c *gin.Context) {
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

		ctrl := NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{"status": "Not found"})
			return
		}

		c.JSON(200, result)
	})
}
