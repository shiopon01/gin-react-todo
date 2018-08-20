package controllers

import (
	"../models"

	"log"
)

// User is
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
