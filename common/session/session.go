package session

import (
	"fmt"
	"os"
	"github.com/geokaralis/go-robin/crypto"
)

type Session struct {
	id string
	status int
	pid int
	name string
	desc string
}

func NewSession() *Session {
	return &Session {
		id: crypto.Id(12),
		status: 1,
		pid: os.Getpid(),
		name: "Robin",
		desc: "Decentralized load balancer in go.",
	}
}

func (s *Session) Getid() string {
	return s.id
}

func (s *Session) Getpid() int {
	return s.pid
}

func (s *Session) Getname() string {
	return s.name
}

func (s *Session) Getdesc() string {
	return s.desc
}

func (s *Session) Start() {
	
}

func (s *Session) Watch() {
	switch s.status {
	case 0:
		fmt.Printf("Session terminated.\n")
	case 1:
		fmt.Printf("Session up and running.\n")
	}
}