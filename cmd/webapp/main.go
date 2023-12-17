package main

import (
	"github.com/charukak/todo-app-htmx/pkg/server"
	"github.com/charukak/todo-app-htmx/internal/handlers"
)

func main() {
	s := server.NewServer()
	h := handlers.NewHandler()
	s.HandleStatic("/static/", "./static")
	s.Handle("/", h.Hello)
	s.Start()
}

