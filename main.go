package main

import (
	"log"

	api "github.com/bing-byte-9527/go-todo-list-api/api/todo"
	"github.com/bing-byte-9527/go-todo-list-api/server"
	"github.com/bing-byte-9527/go-todo-list-api/store/memory"
)

func main() {

	// create a new todo store
	todoStore := memory.NewTodoStore()

	// create a new http server with todo handler
	server := server.NewServer(api.NewTodoHandler(todoStore))

	// start the server
	log.Fatal(server.Start(":8000"))
}
