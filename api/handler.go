package api

import (
	"net/http"

	api "github.com/bing-byte-9527/go-todo-list-api/api/todo"
)

// Handler represents an HTTP handler
type Handler struct {
	todo *api.TodoHandler
}

// NewHandler creates a new HTTP handler
func NewHandler(todo *api.TodoHandler) *Handler {
	return &Handler{todo: todo}
}

// ServeHTTP serves HTTP requests
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/todos":
		h.todo.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
