package models

// Todo ...
type Todo struct {
	ID        int    `json:"id" xorm:"'id'"`
	Title     string `json:"title" xorm:"'title'"`
	Finished  bool   `json:"finished" xorm:"'finished'"`
	CreatedAt int32  `json:"created_at" xorm:"'created_at'"`
}

// NewTodo ...
func NewTodo(title string, finished bool, createdAt int32) Todo {
	return Todo{
		Title:     title,
		Finished:  finished,
		CreatedAt: createdAt,
	}
}

// TodoRepository ...
type TodoRepository struct {
}

// NewTodoRepository ...
// Run this method to get TodoRepository
func NewTodoRepository() TodoRepository {
	return TodoRepository{}
}

// GetByID ...
// TodoRepository's method
func (m TodoRepository) GetByID(id int) *Todo {
	todo := Todo{ID: id}
	has, _ := engine.Get(&todo)

	if has {
		return &todo
	}

	return nil
}
