package main

import (
	"fmt"
)

type Session struct {
	id string
	status int
	pid int
	name string
	desc string
}

func (s *Session) Watch() {
	switch s.status {
	case 0:
		fmt.Printf("Session terminated.\n")
	case 1:
		fmt.Printf("Session up and running.\n")
	}
}