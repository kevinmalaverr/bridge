package bridge

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// type Request *http.Request
// type Response http.Response

type Middleware func(http.HandlerFunc) http.HandlerFunc
type MiddlewareFunc func(*http.Request, http.ResponseWriter, http.HandlerFunc)

func CheckAuth(req *http.Request, res http.ResponseWriter, next http.HandlerFunc) {
	flag := true
	fmt.Println("checking auth")
	if flag {
		next(res, req)
	} else {
		return
	}
}

func Logging(request *http.Request, response http.ResponseWriter, next http.HandlerFunc) {
	start := time.Now()
	defer func() {
		log.Println(request.URL.Path, time.Since(start))
	}()
	next(response, request)
}

func SetMiddleware(middlewareFunc MiddlewareFunc) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			middlewareFunc(r, w, f)
		}
	}
}
