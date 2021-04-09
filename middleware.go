package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("checking auth")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}
