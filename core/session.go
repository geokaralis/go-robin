package core

import (
	"fmt"
	"crypto/rand"
	"encoding/base64"
	"math"
	"os"
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

func RandomId(l int) string {
	buff := make([]byte, int(math.Round(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip the one extra byte we get from half the results.
}

func NewSession() *Session {
	return &Session {
		id: RandomId(8),
		status: 1,
		pid: os.Getpid(),
		name: "Robin",
		desc: "Decentralized load balancer in go.",
	}
}

func (s *Session) Start() {
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