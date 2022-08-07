package main

import (
	"log"
	"net/http"
	"os"
)

func getListenAddr() string {
	if bindAddress := os.Getenv("BIND_ADDRESS"); bindAddress != "" {
		return bindAddress
	}
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return "127.0.0.1:3000"
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		_, _ = w.Write([]byte("Hello, world!"))
	case "/bye":
		_, _ = w.Write([]byte("👋"))
	default:
		http.NotFound(w, r)
	}
}

func main() {
	server := &http.Server{
		Addr:    getListenAddr(),
		Handler: http.HandlerFunc(httpHandler),
	}
	log.Println("starting server on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
