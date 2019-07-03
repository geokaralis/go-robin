package core

import (
	"fmt"
)

type Connection struct {
	id string
	Type string
	status int
}

func (conn *Connection) Check() {
	if conn.status != 0 {
		fmt.Printf("Connection %s is alive. Status: %d\n", conn.id, conn.status)
		return
	}

	fmt.Printf("Connection %s lost. Status: %d\n", conn.id, conn.status)
}