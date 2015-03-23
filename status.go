package goserv

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// This struct is used for the output of all status items
type Status struct {
	Time time.Time
	Data interface{}
}

// A function that takes no input and returns on item
type Callable func() interface{}

// Takes an interface and bool to note if the results should be pretty
// and turns it into Status output in JSON
func StructHandler(item interface{}, pretty bool) http.HandlerFunc {
	status := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		status := Status{
			Time: time.Now().UTC(),
			Data: item,
		}
		output, _ := json.Marshal(status)
		if pretty {
			var buf bytes.Buffer
			json.Indent(&buf, output, "", "\t")
			output = buf.Bytes()
		}
		w.Write(output)
		return
	}
	return status
}

// Takes a Callable and bool to note if the results should be pretty
// and passes the resulting interface go StatusHandler.StructHandler
func FuncHandler(callable Callable, pretty bool) http.HandlerFunc {
	return StructHandler(callable(), pretty)
}
