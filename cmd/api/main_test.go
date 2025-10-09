package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	handler := handleHealth()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)

	}

	got := rr.Body.String()
	want := `{"status":"ok"}`
	if got != want {
		t.Fatalf("handler returned unexpected body: got %v want %v", got, want)
	}
}
