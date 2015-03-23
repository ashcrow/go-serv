package main

import (
	"flag"
	"github.com/ashcrow/go-serv"
	"net/http"
	"time"
)

// Type to hold configuration. We embed BaseConfiguration
// to take advantage of the HTTP related flags. We can
// also extend it by adding other fields (like Test)
type MyConfiguration struct {
	goserv.BaseConfiguration
	Test bool
}

// Our configuration
var Conf *MyConfiguration

// Main entry point
func main() {
	// Set defaults for the server configuration with BaseConfiguration
	// inside of MyConfiguration.
	Conf = &MyConfiguration{
		BaseConfiguration: goserv.BaseConfiguration{
			BindAddress:    "0.0.0.0",
			BindPort:       80,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			LogLevel:       "debug",
		},
		Test: false,
	}

	// Add a flag for our custom Test addition to the Configuration
	flag.BoolVar(&Conf.Test, "Test", Conf.Test, "Example of extending the Configuration")

	// Simple handler to say hello!
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})

	// Make a new server.
	server := goserv.NewServer(&Conf.BaseConfiguration)
	// Log that we are listing
	goserv.Logger.Infof("Listening for connections on %s", server.Addr)
	// And listen for connections
	server.ListenAndServe()
}
