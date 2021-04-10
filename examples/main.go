package main

import (
	"fmt"
	"net/http"

	"github.com/kevinmalaverr/bridge"
)

func firstMid(request *http.Request, response http.ResponseWriter, next http.HandlerFunc) {
	fmt.Println("first")
	next(response, request)
}
func secondMid(request *http.Request, response http.ResponseWriter, next http.HandlerFunc) {
	fmt.Println("second")
	next(response, request)
}

func main() {
	server := bridge.NewServer(":3000")
	server.Use(firstMid)
	server.Use(secondMid)
	server.Get("/", server.AddMiddleware(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.Path)
	}, firstMid, secondMid))
	server.Listen()
}
