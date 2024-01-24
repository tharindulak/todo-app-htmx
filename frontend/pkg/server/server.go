package server

import (
	"net/http"
	"os"

	"github.com/charukak/todo-app-htmx/frontend/pkg/log"
	"github.com/go-chi/chi/v5"
)

var port = ":8080"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(r chi.Router) {
	// start server
	log.Info("Server started on port: ", port)
	http.ListenAndServe(port, r)
}

// register Handle function {
func (s *Server) Handle(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, handler)
}

func (s *Server) HandleStatic(path string, dir string) {
	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

func init() {
	// register Handle function
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = ":" + envPort
	}
}
