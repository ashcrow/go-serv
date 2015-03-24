package main

import (
    "github.com/ashcrow/go-serv"
    "net/http"
    "time"
)


// Our configuration
var Conf *goserv.BaseConfiguration

// Main entry point
func main() {
    Conf = &goserv.BaseConfiguration{
        BindAddress:    "0.0.0.0",
        BindPort:       80,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
        LogLevel:       "info",
    }


    // Simple handler to say hello!
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello!"))
    })

    // Make a new server.
    server := goserv.NewServer(Conf)
    // Log that we are listing
    goserv.Logger.Infof(
        "Listening for connections on %s ports %d and %d", Conf.BindAddress, Conf.BindPort, Conf.BindHttpsPort)
    // And listen for connections
    goserv.RunHttpAndHttps(&server, Conf)
}
