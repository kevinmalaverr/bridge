package main

func main() {
	server := NewServer(":3000")
	server.Use(Logging())
	server.Get("/", HandleRoot)
	server.Handle("/api", "POST", server.AddMiddleware(HandleHome, CheckAuth()))
	server.Listen()
}
