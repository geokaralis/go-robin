build:
	go build ./...

run:
	go run ./...

compile:
	GOOS=linux GOARCH=386 go build -o bin/robin-linux-x86 *.go
	GOOS=windows GOARCH=386 go build -o bin/robin-windows-x86 *.go

all: build
	go install github.com/geokaralis/go-robin/binaries/robin_core

install: build
	go install github.com/geokaralis/go-robin/binaries/robin_core
	go install github.com/geokaralis/go-robin/binaries/robin_watchdog