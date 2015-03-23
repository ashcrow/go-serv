package goserv

import (
	"net/http"
	"sort"
)

// Restricts access to a handler by IP
func RestrictByIP(handler http.HandlerFunc, allowedIps []string) http.HandlerFunc {
	sort.Strings(allowedIps)

	wrapper := func(w http.ResponseWriter, r *http.Request) {
		if sort.SearchStrings(allowedIps, r.RemoteAddr) > 0 {
			handler(w, r)
		} else {
			http.NotFound(w, r)
		}
	}
	return wrapper
}
