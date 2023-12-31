package main

import (
	"github.com/charukak/todo-app-htmx/internal/handlers"
	"github.com/charukak/todo-app-htmx/pkg/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := server.NewServer()
	h := handlers.NewHandler()
	s.HandleStatic("/static/", "./static")
	s.Handle("/", h.Hello)
	s.Start()

}

