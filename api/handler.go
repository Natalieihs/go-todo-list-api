package api

import (
	"net/http"
)

// Handler represents an HTTP handler
type Handler struct {
	todo *TodoHandler
}

// NewHandler creates a new HTTP handler
func NewHandler(todo *TodoHandler) *Handler {
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
