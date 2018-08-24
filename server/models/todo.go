package models

import (
	"log"
	"time"
)

// Todos struct
type Todos struct {
	ID        int       `json:"id" xorm:"'id'"`
	Title     string    `json:"title" xorm:"'title'"`
	Finished  bool      `json:"finished" xorm:"'finished'"`
	CreatedAt time.Time `json:"created_at" xorm:"'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"'updated_at'"`
}

// NewTodo is create todo struct function
func NewTodo(title string, finished bool, createdAt time.Time, updatedAt time.Time) Todos {
	return Todos{
		Title:     title,
		Finished:  finished,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// TodosRepository struct
type TodosRepository struct {
}

// NewTodosRepository ...
// Run this method to get TodoRepository
func NewTodosRepository() TodosRepository {
	return TodosRepository{}
}

// Methods

// GetAll is get all todos
func (m TodosRepository) GetAll() *[]Todos {

	var todos []Todos
	err := engine.Find(&todos)

	log.Println("RESULT", err, todos)

	// if results {
	// 	return results
	// }

	return &todos
}

// GetByID is get one todo
func (m TodosRepository) GetByID(id int) *Todos {
	todos := Todos{ID: id}
	has, _ := engine.Get(&todos)

	if has {
		return &todos
	}

	return nil
}

// AddTodo is get one todo
func (m TodosRepository) AddTodo(title string) *Todos {
	todo := new(Todos)
	todo.Title = title

	succeed, _ := engine.Insert(todo)

	if succeed == 1 {
		return todo
	}

	return nil
}
