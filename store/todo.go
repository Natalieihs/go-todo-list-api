package store

import "time"

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TodoStore interface {
	GetTodos() ([]*Todo, error)
	AddTodo(todo *Todo) error
	RemoveTodoByID(id string) error
}
