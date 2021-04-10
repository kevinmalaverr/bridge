package bridge

import (
	"net/http"
)

type router struct {
	rules map[string]map[string]http.HandlerFunc
}

func newRouter() *router {
	return &router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *router) findHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, exist, methodExist
}

func (r *router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist, methodExist := r.findHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
