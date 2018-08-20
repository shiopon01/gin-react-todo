package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// init ...
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "gin:password@/gin_sample")
	if err != nil {
		panic(err)
	}
}

// User is
type Users struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"username" xorm:"'username'"`
	Name     string `json:"name" xorm:"'name'"`
	Password string `json:"password" xorm:"'password'"`
}

// NewUser ...
func NewUser(id int, username string, name string, password string) Users {
	return Users{
		ID:       id,
		Username: username,
		Name:     name,
		Password: password,
	}
}

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (m UserRepository) GetByID(id int) *Users {
	var user = Users{ID: id}
	has, _ := engine.Get(&user)
	log.Println("LOG MODEL:", has)
	if has {
		return &user
	}

	return nil
}
