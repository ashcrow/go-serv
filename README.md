# go-serv
go-serv attempts to take care of common requirements for web applications while not dictating any specific Go web framework. 

**Warning**: Currently in development with no official release yet.

## Features

* Built on Go's net/http library
* Framework agnostic
* Logging via logrus (https://github.com/Sirupsen/logrus/)
* Configuration file parsing via TOML (https://github.com/toml-lang/toml/)
* Command line flags which can overrule configuration file
* Simple status/health system for exposing structs
* Run HTTP and HTTPS servers with the same binary.

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
```bash
$ ./status-example -help
Usage of ./status-example:
  -BindAddress="0.0.0.0": Bind address.
  -BindPort=80: HTTP bind port.
  -LogFile="": Log file.
  -LogLevel="info": Log level.
  -MaxHeaderBytes=1048576: Max header bytes.
  -ReadTimeout=10s: Read timeout.
  -WriteTimeout=10s: Write timeout.
```

## Building Examples
There is a Makefile provided to build the code in the examples folder.

```bash
$ make build-examples-all
```
