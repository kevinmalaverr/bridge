package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port               string
	router             *Router
	generalMiddlewares []Middleware
}

func NewServer(port string) *Server {
	return &Server{
		port:               port,
		router:             NewRouter(),
		generalMiddlewares: make([]Middleware, 0),
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
	s.applyMiddlewares(handler)
	s.router.rules[path][method] = handler
}

func (s *Server) Use(middleware MiddlewareFunc) {
	s.generalMiddlewares = append(s.generalMiddlewares, SetMiddleware(middleware))
}

func (s *Server) applyMiddlewares(handler http.HandlerFunc) {
	for _, m := range s.generalMiddlewares {
		handler = m(handler)
	}
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...MiddlewareFunc) http.HandlerFunc {

	for _, m := range middlewares {
		// pass handler to each middleware
		f = SetMiddleware(m)(f)
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
