help:
	@echo "Target			Description"
	@echo "=====			==========="
	@echo "get-deps		Downloads and installs dependencies"
	@echo "test			Run unittests."
	@echo "clean			Clean up."
	@echo "build			Fully builds the source and dependencies."
	@echo "build-fast		Builds just this packages code."
	@echo "install			Installs the library."
	@echo "build-examples		Fully builds the source and produces the examples in this package."
	@echo "build-examples-fast	Builds and produces the examples in this package."

get-deps:
	godep restore

test:
	godep go test -v -cover

build: clean
	go build -race -x -a .

build-fast: clean
	go build .

install:
	go install .

clean:
	rm -rf bin/
	go clean

build-examples-all:
	mkdir bin/
	go build -o bin/status-example examples/status.go
	go build -o bin/gorilla-example examples/gorilla.go
	go build -o bin/flag-example examples/flag.go
	go build -o bin/httpandhttps-example examples/httpandhttps.go
	@echo "See bin/ for examples."

build-examples-fast: build-fast build-examples-all

build-examples: build build-examples-all
