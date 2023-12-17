package handlers

import "net/http"

type Handler struct{}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func NewHandler() *Handler {
    return &Handler{}
}

