package session

import (
	"fmt"
	"os"

	"github.com/geokaralis/go-robin/crypto"
)

func Test() {
	fmt.Println("Dddd")
}

type Session struct {
	id string
	status int
	pid int
	name string
	desc string
}

func NewSession() *Session {
	return &Session {
		id: crypto.Id(8),
		status: 1,
		pid: os.Getpid(),
		name: "Robin",
		desc: "Decentralized load balancer in go.",
	}
}

func (s *Session) Start() {
	fmt.Printf("            __   _    \n")
	fmt.Printf("  _______  / /  (_)__ \n")
	fmt.Printf(" / __/ _ \\/ _ \\/ / _ \\ \n")
	fmt.Printf("/_/  \\___/_.__/_/_//_/ \n")
	fmt.Printf("%s, %s\n\n", s.name, s.desc)
	fmt.Printf("Starting session, id: %s, pid: %d\n", s.id, s.pid)
}

func (s *Session) Watch() {
	switch s.status {
	case 0:
		fmt.Printf("Session terminated.\n")
	case 1:
		fmt.Printf("Session up and running.\n")
	}
}