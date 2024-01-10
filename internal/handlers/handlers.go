package handlers

import (
	"context"
	"net/http"

	"github.com/charukak/todo-app-htmx/web/templates"
)

type Handler struct{}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	component := templates.Hello("yay")
    component.Render(context.Background(),w)
}

func NewHandler() *Handler {
	return &Handler{}
}
