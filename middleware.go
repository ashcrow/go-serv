package goserv

import (
	"net/http"
	"sort"
)

// RestrictByIP restricts access to a handler by IP
func RestrictByIP(handler http.HandlerFunc, allowedIps []string) http.HandlerFunc {
	sort.Strings(allowedIps)

	wrapper := func(w http.ResponseWriter, r *http.Request) {
		index := sort.SearchStrings(allowedIps, r.RemoteAddr)
		if len(allowedIps) > 0 && allowedIps[index] == r.RemoteAddr {
			handler(w, r)
		} else {
			Logger.Warnf("%s was not in the list of allowed IP's for %v", r.RemoteAddr, handler)
			http.NotFound(w, r)
		}
	}
	return wrapper
}
