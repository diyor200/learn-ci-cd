package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "server:", log.Lshortfile)

	http.HandleFunc("/hello", middleware(hello))
	server := &http.Server{Addr: ":8080"}

	logger.Println("Server is running on :8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	resp := []byte(`{"message": "Hello World"}`)

	w.Write(resp)
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		next.ServeHTTP(w, r)

		log.Println(r.Method, r.URL.String(), time.Since(now), r.RemoteAddr)
	})
}
