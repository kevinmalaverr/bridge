package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kevinmalaverr/bridge"
)

func firstMid(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	fmt.Println("first")
	next(response, request)
}
func secondMid(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	fmt.Println("second")
	next(response, request)
}

func logging(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	start := time.Now()
	defer func() {
		log.Println(request.URL.Path, time.Since(start))
	}()
	next(response, request)
}

func auth(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	password := "1234"
	if request.URL.Query().Get("pass") != password {
		response.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(response, "incorrect password")
	} else {
		next(response, request)
	}
}

func main() {
	server := bridge.NewServer(":3000")

	// general middlewares
	server.Use(firstMid)
	server.Use(secondMid)
	server.Use(logging)

	server.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.Path)
	})

	// local middlewares
	server.Post("/auth", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "authorized")
	}, auth)

	server.Listen()
}
