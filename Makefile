build:
	go build -o bin/robin main.go

run:
	go run *.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/robin-linux-x86 main.go
	GOOS=windows GOARCH=386 go build -o bin/robin-windows-x86 main.go

all: build