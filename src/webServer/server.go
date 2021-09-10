package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	error := http.ListenAndServe(s.port, nil)

	if error != nil {
		return error
	}

	return nil
}

func NewServer(port int) *Server {
	return &Server{
		port:   fmt.Sprintf(":%d", port),
		router: NewRouter(),
	}
}
