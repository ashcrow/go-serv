package goserv

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	flag "github.com/ogier/pflag"
	"os"
	"testing"
	"time"
)

// Configuration for testing
var TestConfig = BaseConfiguration{
	"127.0.0.1",
	8080,
	10 * time.Second,
	10 * time.Second,
	1 << 20,
	"debug",
	"out.log",
	0,
	"",
	"",
}

// Required for testing
func init() {
	flag.StringP("test", "t", "", "")
	flag.Parse()
}

// tearDown removes the log file created by tests
func tearDown() {
	os.Remove(TestConfig.LogFile)
}

// Test_makeLogger tests makeLogger to ensure it sets the proper values
func Test_makeLogger(t *testing.T) {
	makeLogger(&TestConfig)

	if Logger.Level != logrus.DebugLevel {
		t.Fatalf("makeLogger did not set the correct level. Expected %s, got %s", logrus.DebugLevel, Logger.Level)
	}

	if _, err := os.OpenFile(TestConfig.LogFile, os.O_RDONLY, os.ModePerm); err != nil {
		t.Fatalf("makeLogger did not create the log file at %s", TestConfig.LogFile)
	}

	tearDown()
}

// TestNewServer tests NewServer and verifies it sets proper values
func TestNewServer(t *testing.T) {
	server := NewServer(&TestConfig)
	if server.MaxHeaderBytes != TestConfig.MaxHeaderBytes {
		t.Fatalf("NewServer set incorrect MaxHeaderBytes. Expected %s, got %s", TestConfig.MaxHeaderBytes, server.MaxHeaderBytes)
	}

	if server.Addr != fmt.Sprintf("%s:%d", TestConfig.BindAddress, TestConfig.BindPort) {
		t.Fatalf("NewServer set incorrect BindAddress. Expected %s, got %s", TestConfig.BindAddress, server.Addr)
	}

	if server.ReadTimeout != TestConfig.ReadTimeout {
		t.Fatalf("NewServer set incorrect ReadTimeout. Expected %s, got %s", TestConfig.ReadTimeout, server.ReadTimeout)
	}

	if server.WriteTimeout != TestConfig.WriteTimeout {
		t.Fatalf("NewServer set incorrect WriteTimeout. Expected %s, got %s", TestConfig.WriteTimeout, server.WriteTimeout)
	}

	tearDown()
}

// TODO(ashcrow): Finish testing
/*
func TestRunHttpAndHttps(t *testingT) {

}

func TestBackgroundRunHttps(t *testingT) {

}
func TestBackgroundRunHttp(t *testingT) {

}
*/
