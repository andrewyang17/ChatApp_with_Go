package main

import (
	"goChat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()
	log.Println("Starting channel listener")

	go handlers.ListenToWsChannel()

	log.Println("Starting webserver on port 8080")
	_ = http.ListenAndServe(":8080", mux)
}
