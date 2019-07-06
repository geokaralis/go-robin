package main

import (
	"github.com/geokaralis/go-robin/common/session"
	"github.com/geokaralis/go-robin/cli"
)

func main() {
	/*
	 * Create a new session for robin core and pass it to the cli 
	 * to keep a reference to the instance. Upon failure robin logs
	 * the instance and falls back to another, while trying to restart
	 * the failed session.
	 */
	session := session.NewSession()
	cli := &cli.Cli {
		Input: false, // at the current stage accept no user input
	}
	cli.Init(session)
}