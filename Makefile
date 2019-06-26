build:
	go build -o bin/robin *.go

run:
	go run *.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/robin-linux-x86 *.go
	GOOS=windows GOARCH=386 go build -o bin/robin-windows-x86 *.go

all: build