package goserv

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestData struct {
	Integer int
	String  string
}

func TestStructHandler(t *testing.T) {

	request, _ := http.NewRequest("GET", "http://127.0.0.1/status/", nil)
	expected := TestData{10, "OK"}

	for _, pretty := range []bool{false, true} {
		w := httptest.NewRecorder()
		handler := StructHandler(expected, pretty)
		handler(w, request)

		result := Status{}
		json.Unmarshal(w.Body.Bytes(), &result)
		data := result.Data.(map[string]interface{})
		now := time.Now()
		if result.Time.After(now) {
			t.Fatalf("StructHandler returned a time in the future: %+v > %+v", result.Time, now)
		}
		// Note the use of float64 caseting since the an interface to integer.
		if data["Integer"] != float64(expected.Integer) {
			t.Fatalf("StructHandler returned the wrong data for Integer: %+v != %+v", data["Integer"], expected.Integer)
		}
		if data["String"] != expected.String {
			t.Fatalf("StructHandler returned the wrong data for String: %+v != %+v", data["String"], expected.String)
		}
	}
}
