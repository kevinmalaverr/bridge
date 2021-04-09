package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	fmt.Println("server listening at http://localhost" + s.port)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		// pass handler to each middleware
		f = m(f)
	}
	return f
}

func (s *Server) Get(path string, handler http.HandlerFunc) {
	s.Handle(path, "GET", handler)
}

func (s *Server) Post(path string, handler http.HandlerFunc) {
	s.Handle(path, "POST", handler)
}

func (s *Server) Put(path string, handler http.HandlerFunc) {
	s.Handle(path, "PUT", handler)
}

func (s *Server) Delete(path string, handler http.HandlerFunc) {
	s.Handle(path, "DELETE", handler)
}
