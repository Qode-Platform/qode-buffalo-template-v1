package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func code(path string) int {
	w := httptest.NewRecorder()
	newApp().ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
	return w.Code
}

func TestServesAtRoot(t *testing.T) {
	if got := code("/health"); got != http.StatusOK {
		t.Fatalf("GET /health = %d", got)
	}
}
