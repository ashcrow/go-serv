package main

import (
	"github.com/ashcrow/go-serv"
	"net/http"
	"runtime"
	"time"
)

// Our configuration
var Conf *goserv.BaseConfiguration

// Main entry point
func main() {
	// Set defaults for the server configuration
	Conf = &goserv.BaseConfiguration{
		BindAddress:    "0.0.0.0",
		BindPort:       80,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		LogLevel:       "info",
	}

	// Create a function that returns a struct for use with FuncHandler
	GetMemStats := func() interface{} {
		ms := runtime.MemStats{}
		runtime.ReadMemStats(&ms)
		return ms
	}

	// Bind it
	http.HandleFunc("/status/memory", goserv.FuncHandler(GetMemStats, true))

	// Create a function that returns an interface for use with FuncHandler
	SysStats := func() interface{} {
		return map[string]interface{}{
			"CPUCount":       runtime.NumCPU(),
			"GoroutineCount": runtime.NumGoroutine(),
			"CGoCallCount":   runtime.NumCgoCall(),
			"GoVersion":      runtime.Version(),
		}
	}

	// Bind it and log
	http.HandleFunc("/status/sys", goserv.FuncHandler(SysStats, true))

	// Bind a struct with StructHandler and use RestrictByIP
	// middleware to only allow specified IP addresses
	http.HandleFunc(
		"/status/config",
		goserv.RestrictByIP(
			goserv.StructHandler(Conf, true), []string{"127.0.0.1"}))

	// Make a new server.
	server := goserv.NewServer(Conf)
	// Log that we are listing
	goserv.Logger.Infof("Listening for connections on %s", server.Addr)
	// And listen for connections
	server.ListenAndServe()
}
