// A simple example using the Gorilla toolkit http://www.gorillatoolkit.org/
package main

import (
	"net/http"
	"time"

	"github.com/ashcrow/go-serv"
	"github.com/gorilla/mux"
)

// Conf is our configuration
var Conf goserv.BaseConfiguration

// Main entry point
func main() {
	// Set defaults for the server configuration with BaseConfiguration
	// inside of MyConfiguration.
	Conf = goserv.BaseConfiguration{
		BindAddress:    "0.0.0.0",
		BindPort:       80,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		LogLevel:       "debug",
	}

	// New gorilla mux router
	router := mux.NewRouter()
	// Don't enforce strict slashing
	router.StrictSlash(false)

	// Map /hello/ with a variable called name.
	router.HandleFunc("/hello/{name}/", func(w http.ResponseWriter, r *http.Request) {
		// Pull the variable and use it in write (note: this is not safe in real life ;-)
		name := mux.Vars(r)["name"]
		w.Write([]byte("hello " + name + "!"))
	})

	// Map a struct with StructHandler as a status
	router.HandleFunc("/status/config/", goserv.StructHandler(Conf, true))

	// Make a new server.
	server := goserv.NewServer(&Conf)
	// Set the handler to our gorilla mux router
	server.Handler = router
	// Only serve on 127.0.0.1
	router.Host(server.Addr)
	// Log that we are listing
	goserv.Logger.Infof("Listening for connections on %s", server.Addr)
	// And listen for connections
	server.ListenAndServe()
}
