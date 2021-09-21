package server

import "github.com/kutyrov/my_server_go/internal/app/storage"

type Server struct {
	port    string
	storage *storage.Storage
}

func New(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() error {
	// fmt.Println("пишу из старта", s.port)
	return nil
}
