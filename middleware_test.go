package goserv

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setUp does some setting up for the middleware tests.
func setUp() (*httptest.ResponseRecorder, *http.Request, http.HandlerFunc) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://127.0.0.1/", nil)
	handler := func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK!") }
	return recorder, request, handler
}

func TestRestrictByIP(t *testing.T) {
	// Case 1. Should be blocked
	w, request, handler := setUp()
	wrapped := RestrictByIP(handler, []string{"127.0.0.1"})
	wrapped(w, request)
	if w.Code != 404 {
		t.Fatalf("RestrictByIP did not return a 404: %d, %s", w.Code, w.Body)
	}

	// Case 2: Should be allowed
	w, request, handler = setUp()
	request.RemoteAddr = "127.0.0.1"
	wrapped(w, request)
	if w.Code != 200 {
		t.Fatalf("RestrictByIP did not return a 200: %d, %s", w.Code, w.Body)
	}
}
