package main

import (
	"github.com/geokaralis/go-robin/common/session"
	"github.com/geokaralis/go-robin/cli"
)

func main() {
	session := session.NewSession()
	session.Start()

	cli := &cli.Cli {}
	cli.Init()
}