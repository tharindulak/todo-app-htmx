package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/charukak/todo-app-htmx/internal/app"
	"github.com/charukak/todo-app-htmx/web/templates"
)

type Handler struct {
	todoAppClient *app.TodoAppClient
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	component := templates.Hello("yay")
	component.Render(context.Background(), w)
}

func (h *Handler) TodoList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TodoList")
	todos, err := h.todoAppClient.GetTodos()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := templates.TodoList(todos)
	component.Render(context.Background(), w)
}

func NewHandler() *Handler {

	return &Handler{
		todoAppClient: app.NewTodoAppClient(os.Getenv("TODO_APP_URL")),
	}
}
