# go-serv
[![GoDoc](http://godoc.org/gopkg.in/ashcrow/go-serv.v0?status.png)](http://godoc.org/gopkg.in/ashcrow/go-serv.v0)
[![Travis](https://travis-ci.org/ashcrow/go-serv.svg?branch=master)](https://travis-ci.org/ashcrow/go-serv)

go-serv attempts to take care of common requirements for web applications while not dictating any specific Go web framework. 

Repo: https://github.com/ashcrow/go-serv/

**Warning**: Currently in development with no official release yet.

## Features

* Built on Go's net/http library
* Framework agnostic
* Logging via logrus (https://github.com/Sirupsen/logrus/)
* Configuration file parsing via TOML (https://github.com/toml-lang/toml/)
* Command line flags which can overrule configuration file
* Simple status/health system for exposing structs
* Run HTTP and HTTPS servers with the same binary.

### Installation
```bash
$ go get gopkg.in/ashcrow/go-serv.v0
```

### Unittesting
```bash
$ go test -v -cover
```

or

```bash
$ make test
```

### Configuration File Example
```plain
# Note that the names are the same across the BaseConfiguration
# struct, this config file, and command line flags.
BindAddress = "127.0.0.1"
BindPort    = 8000
LogLevel    = "info"
LogFile     = "/tmp/out.log"
```

### Default Command Line Flags

#### Application Defaults
```bash
$ ./status-example -help
Usage of ./status-example:
  -BindAddress="0.0.0.0": Bind address.
  -BindHttpsPort=443: HTTPS bind port.
  -BindPort=80: HTTP bind port.
  -CertFile="": Cert file.
  -KeyFile="": Key file.
  -LogFile="": Log file.
  -LogLevel="info": Log level.
  -MaxHeaderBytes=1048576: Max header bytes.
  -ReadTimeout=10s: Read timeout.
  -WriteTimeout=10s: Write timeout.
```

#### Configuration File Defaults
```bash
$ ./status-example -help /path/to/conf.toml
Usage of ./status-example:
  -BindAddress="127.0.0.1": Bind address.
  -BindHttpsPort=8181: HTTPS bind port.
  -BindPort=8000: HTTP bind port.
  -CertFile="./cert.pem": Cert file.
  -KeyFile="./key.pem": Key file.
  -LogFile="/tmp/out.log": Log file.
  -LogLevel="info": Log level.
  -MaxHeaderBytes=1048576: Max header bytes.
  -ReadTimeout=10s: Read timeout.
  -WriteTimeout=10s: Write timeout.
```

## Examples
Examples can be found in the [examples folder](https://github.com/ashcrow/go-serv/tree/master/examples)

### Building Examples

There is a Makefile provided to build the code in the examples folder.

```bash
$ make build-examples-all
```
