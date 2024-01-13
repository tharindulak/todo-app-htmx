package main

import (
	"net/http"

	"github.com/charukak/todo-app-htmx/pkg/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("web/static/"))

	r.Handle("/", fs)

	// h := handlers.NewHandler()
	// r.Get("/", h.Hello)

	s := server.NewServer()
	// s.Handle("/", h.Hello)
	s.Start(r)

}
