package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	handler := middleware(hello)
	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	expectedBody := `{"message": "Hello World"}`
	body := rec.Body.String()

	if body != expectedBody {
		t.Fatalf("expected body %q, got %q", expectedBody, body)
	}

}
