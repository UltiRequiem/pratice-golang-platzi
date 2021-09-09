package main

import "net/http"

type Server struct {
	port string
}

func (s *Server) Listen() error {
	error := http.ListenAndServe(s.port, nil)

	if error != nil {
		return error
	}

        return nil
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}
