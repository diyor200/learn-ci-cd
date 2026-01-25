package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "server:", log.Lshortfile)

	http.HandleFunc("/hello", middleware(hello))
	http.HandleFunc("/ping", middleware(ping))
	http.HandleFunc("/health", middleware(health))

	server := &http.Server{Addr: ":8080"}

	logger.Println("Server is running on :8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
