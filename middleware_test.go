package goserv

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// setUp does some setting up for the middleware tests.
func setUp() (*httptest.ResponseRecorder, *http.Request) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://127.0.0.1/", nil)
	return recorder, request
}

// restrictedExpectation is reusable code for TestRestrictByIP
func restrictedExpectation(t *testing.T, expectedCode int, allowedIps []string) {
	w, request := setUp()
	wrapped := RestrictByIP(handler, allowedIps)
	request.RemoteAddr = "127.0.0.1"
	wrapped(w, request)
	if w.Code != expectedCode {
		t.Fatalf("RestrictByIP did not return a %d: %d, %s", expectedCode, w.Code, w.Body)
	}
}

// TestRestrictByIP tests the RestrictByIP middleware
func TestRestrictByIP(t *testing.T) {
	// Case 1. Should be blocked due to no IP's in the list
	restrictedExpectation(t, 404, []string{})

	// Case 2: Should be blocked due to ip not in list
	restrictedExpectation(t, 404, []string{"127.0.0.2"})

	// Case 3: Should be allowed
	restrictedExpectation(t, 200, []string{"127.0.0.1"})
}
