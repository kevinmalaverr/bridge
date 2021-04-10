package bridge

import (
	"fmt"
	"net/http"
)

//
type Server struct {
	port               string
	router             *router
	generalMiddlewares stack
}

// creates a server instance
func NewServer(port string) *Server {
	return &Server{
		port:               port,
		router:             newRouter(),
		generalMiddlewares: stack{},
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

// applies a middleware
func (s *Server) Use(middleware MiddlewareFunc) {
	s.generalMiddlewares.push(middleware)
}

func (s *Server) Handle(path string, method string, handler http.HandlerFunc, middlewares ...MiddlewareFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	handler = s.addMiddleware(handler, middlewares...)
	h := s.applyMiddlewares(handler)
	s.router.rules[path][method] = h
}

func (s *Server) Get(path string, handler http.HandlerFunc, middlewares ...MiddlewareFunc) {
	s.Handle(path, "GET", handler, middlewares...)
}

func (s *Server) Post(path string, handler http.HandlerFunc, middlewares ...MiddlewareFunc) {
	s.Handle(path, "POST", handler, middlewares...)
}

func (s *Server) Put(path string, handler http.HandlerFunc, middlewares ...MiddlewareFunc) {
	s.Handle(path, "PUT", handler, middlewares...)
}

func (s *Server) Delete(path string, handler http.HandlerFunc, middlewares ...MiddlewareFunc) {
	s.Handle(path, "DELETE", handler, middlewares...)
}

func (s *Server) applyMiddlewares(handler http.HandlerFunc) http.HandlerFunc {
	s.generalMiddlewares.forEach(func(m interface{}) {
		handler = setMiddleware(m.(MiddlewareFunc))(handler)
	})
	return handler
}

func (s *Server) addMiddleware(f http.HandlerFunc, middlewares ...MiddlewareFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		m := middlewares[i]
		f = setMiddleware(m)(f)
	}
	return f
}
