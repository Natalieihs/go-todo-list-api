package server

import "net/http"

type Server struct {
	Handler http.Handler
}

func NewServer(handler http.Handler) *Server {
	return &Server{Handler: handler}
}

func (s *Server) Start(addr string) error {

	return http.ListenAndServe(addr, s.Handler)
}
