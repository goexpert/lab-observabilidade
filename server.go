package labobservalidadade

import (
	"log/slog"
	"net/http"
)

type Server struct {
	Port string
	Mux  *http.ServeMux
}

func NewServer(port string) *Server {
	slog.Info("servidor criado")
	return &Server{
		Port: port,
		Mux:  http.NewServeMux(),
	}
}

func (s *Server) AddHandler(path string, handler http.HandlerFunc) {
	s.Mux.Handle(path, handler)
	slog.Info("nova rota", "path", path)
}

func (s *Server) Run() error {
	slog.Info("servidor em execução", "port", s.Port)
	err := http.ListenAndServe(":"+s.Port, s.Mux)
	if err != nil {
		return err
	}
	return nil
}
