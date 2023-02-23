package api

import (
	"encoding/json"
	"net/http"

	"github.com/bing-byte-9527/go-todo-list-api/store"
)

type TodoHandler struct {
	store store.TodoStore
}

func NewTodoHandler(store store.TodoStore) *TodoHandler {
	return &TodoHandler{store: store}
}
func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.HandlerGetTodos(w, r)
	case http.MethodPost:
		h.HandlerCreateTodo(w, r)
	case http.MethodDelete:
		h.HandlerRemoveTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) HandlerGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.store.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
func (h *TodoHandler) HandlerCreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo store.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.store.AddTodo(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) HandlerRemoveTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	if err := h.store.RemoveTodoByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
