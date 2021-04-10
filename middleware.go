package bridge

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type MiddlewareFunc func(http.ResponseWriter, *http.Request, http.HandlerFunc)

func setMiddleware(middlewareFunc MiddlewareFunc) Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(response http.ResponseWriter, request *http.Request) {
			middlewareFunc(response, request, handlerFunc)
		}
	}
}
