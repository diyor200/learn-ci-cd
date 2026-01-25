package main

import "net/http"

func hello(w http.ResponseWriter, _ *http.Request) {
	resp := []byte(`{"message": "Hello World"}`)

	w.Write(resp)
}

func ping(w http.ResponseWriter, _ *http.Request) {
	resp := []byte(`{"message": "Pong"}`)
	w.Write(resp)
}

func health(w http.ResponseWriter, _ *http.Request) {
	resp := []byte(`{"status": "OK"}`)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
