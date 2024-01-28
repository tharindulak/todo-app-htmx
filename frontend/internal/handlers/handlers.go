package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/charukak/todo-app-htmx/frontend/internal/app"
	"github.com/charukak/todo-app-htmx/frontend/pkg/log"
	"github.com/charukak/todo-app-htmx/frontend/web/templates"
)

type Handler struct {
	todoAppClient *app.TodoAppClient
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	component := templates.Hello("yay")
	component.Render(context.Background(), w)
}

func (h *Handler) TodoList(w http.ResponseWriter, r *http.Request) {
	log.Info("GET TodoList")
	todos, err := h.todoAppClient.GetTodos()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := templates.TodoList(todos)
	component.Render(context.Background(), w)
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Info("POST CreateTodo")

	title := r.FormValue("title")
	description := r.FormValue("description")

	todo, err := h.todoAppClient.CreateTodo(title, description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := templates.TodoItem(*todo)

	component.Render(context.Background(), w)
}

func NewHandler() *Handler {

	return &Handler{
		todoAppClient: app.NewTodoAppClient(os.Getenv("TODO_APP_URL")),
	}
}
