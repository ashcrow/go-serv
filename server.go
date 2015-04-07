// Package goserv is a simple web application server bootstrap. It provides a common base
// for those wishing to focus on their application rather than on setting up
// flags, configuration files or logging. go-serv also comes with a simple
// status system for those who wish to expose structures via HTTP(s) for
// monitoring.
package goserv

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
)

//
const VERSION = "0.0.0"

// Package level logger
var Logger logrus.Logger

// Logger used with http.Server. This is an instance of Logger.Writer()
var ServerErrorLogger log.Logger

// BaseConfiguration structure which can be filled out via
// defaults passed via flags, a configuration file or via
// the defaults set by the programmer (in that order of
// precedence)
type BaseConfiguration struct {
	BindAddress    string
	BindPort       int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	LogLevel       string
	LogFile        string
	BindHttpsPort  int
	CertFile       string
	KeyFile        string
}

// Configures the package level Logger instance
func makeLogger(conf *BaseConfiguration) {
	levelName, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		defer Logger.Warnf("%s is not a valid level. Defaulting to info.", conf.LogLevel)
		levelName = logrus.InfoLevel
	}

	logOut, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		defer Logger.Warnf("%s. Defaulting to Stderr.", err)
		logOut = os.Stderr
	}

	formatter := new(logrus.TextFormatter)
	formatter.DisableColors = true
	Logger = logrus.Logger{
		Out:       logOut,
		Formatter: formatter,
		Level:     levelName,
	}
	Logger.Infof("Initialized at level %s", Logger.Level)
}

// Runs an http server in a goroutine.
func BackgroundRunHttp(server *http.Server, conf *BaseConfiguration) chan error {
	http := make(chan error)
	go func() {
		http <- server.ListenAndServe()
	}()
	return http
}

// Runs an https server in a goroutine.
func BackgroundRunHttps(server *http.Server, conf *BaseConfiguration) chan error {
	https := make(chan error)
	go func() {
		server.Addr = conf.BindAddress + ":" + strconv.FormatInt(int64(conf.BindHttpsPort), 10)
		https <- server.ListenAndServeTLS(conf.CertFile, conf.KeyFile)
	}()
	return https
}

// Runs both Http and Https servers in their own goroutines. If one server
// exists the channels are closed and execution returns to the main goroutine.
func RunHttpAndHttps(server *http.Server, conf *BaseConfiguration) error {
	http := BackgroundRunHttp(server, conf)
	https := BackgroundRunHttps(server, conf)
	var err error
LOOP:
	for {
		select {
		case err = <-http:
			Logger.Fatalf("HTTP server error: %s", err)
			break LOOP
		case err = <-https:
			Logger.Fatalf("HTTPS server error: %s", err)
			break LOOP
		}
	}
	close(http)
	close(https)
	return err
}

// NewServer creates a new http.Server instance based off the BaseConfiguration.
// NewServer also handles reading the TOML configuration file and
// providing/reading the command line flags. Because of this
// NewServer should always be called after all flags have been defined.
func NewServer(conf *BaseConfiguration) http.Server {
	// TOML configuration file can overwrite defaults
	tomlData, err := ioutil.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		defer Logger.Info("No conf. Skipping.")
	} else {
		if _, err := toml.Decode(string(tomlData), &conf); err != nil {
			defer Logger.Errorf("Configuration file could not be decoded. %s", err)
		}
	}
	// Flags can override config items
	// Server flags
	flag.StringVar(&conf.BindAddress, "BindAddress", conf.BindAddress, "Bind address.")
	flag.IntVar(&conf.BindPort, "BindPort", conf.BindPort, "HTTP bind port.")
	flag.DurationVar(&conf.ReadTimeout, "ReadTimeout", conf.ReadTimeout, "Read timeout.")
	flag.DurationVar(&conf.WriteTimeout, "WriteTimeout", conf.WriteTimeout, "Write timeout.")
	flag.IntVar(&conf.MaxHeaderBytes, "MaxHeaderBytes", conf.MaxHeaderBytes, "Max header bytes.")

	// Server Logger flags
	flag.StringVar(&conf.LogLevel, "LogLevel", conf.LogLevel, "Log level.")
	flag.StringVar(&conf.LogFile, "LogFile", conf.LogFile, "Log file.")

	// TLS related flags
	flag.IntVar(&conf.BindHttpsPort, "BindHttpsPort", conf.BindHttpsPort, "HTTPS bind port.")
	flag.StringVar(&conf.CertFile, "CertFile", conf.CertFile, "Cert file.")
	flag.StringVar(&conf.KeyFile, "KeyFile", conf.KeyFile, "Key file.")
	flag.Parse()

	// Logging specific work also injecting the logrus log into the Servers errorlog
	// BUG(ashcrow): This needs work!!!
	makeLogger(conf)
	Logger.Debugf("Final configuration: %+v", conf)
	w := Logger.Writer()
	defer w.Close()
	ServerErrorLogger = *log.New(w, "ServerErrorLogger", log.Lshortfile)
	// -------------

	// Return the configured http.Server
	return http.Server{
		Addr:           conf.BindAddress + ":" + strconv.FormatInt(int64(conf.BindPort), 10),
		ReadTimeout:    conf.ReadTimeout,
		WriteTimeout:   conf.WriteTimeout,
		MaxHeaderBytes: conf.MaxHeaderBytes,
		ErrorLog:       &ServerErrorLogger,
	}
}
