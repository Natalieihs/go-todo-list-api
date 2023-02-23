package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/bing-byte-9527/go-todo-list-api/store"
)

type TodoStore struct {
	mu    sync.RWMutex
	todos map[string]*store.Todo
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos: make(map[string]*store.Todo),
	}
}

func (s *TodoStore) GetTodos() ([]*store.Todo, error) {
	s.mu.Lock()

	defer s.mu.Unlock()
	todos := make([]*store.Todo, 0, len(s.todos))
	for _, t := range s.todos {
		todos = append(todos, t)
	}
	return todos, nil
}

func (s *TodoStore) AddTodo(todo *store.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[todo.ID]; ok {
		return errors.New("todo already exists")
	}
	todo.CreatedAt = time.Now()
	s.todos[todo.ID] = todo
	return nil
}

func (s *TodoStore) RemoveTodoByID(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.todos[id]; !ok {
		return errors.New("todo not found")
	}
	delete(s.todos, id)
	return nil
}
