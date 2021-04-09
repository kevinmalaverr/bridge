package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth()))
	server.Handle("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "inline function")
	})
	server.Listen()
}
