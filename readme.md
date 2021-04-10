# Bridge.go

Bridge.go is a middleware based library to create web applications

## Installation

```shell
go get github.com/kevinmalaverr/bridge
```

## Example

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/kevinmalaverr/bridge"
)

func firstMid(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	fmt.Println("middleware")
	next(response, request)
}

func main() {
	server := bridge.NewServer(":3000")

	server.Use(firstMid)

	server.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "hello world!")
	})

	server.Listen()
}

```
