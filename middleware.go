package main

import (
	"log"
	"net/http"
	"time"
)

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		next.ServeHTTP(w, r)

		log.Println(r.Method, r.URL.String(), time.Since(now), r.RemoteAddr)
	}
}
