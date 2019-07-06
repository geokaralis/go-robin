package cli

import (
	"fmt"
	"os"
	"github.com/geokaralis/go-robin/common/session"
)

type Cli struct {
	Input bool
}

func (c *Cli) Init(s *session.Session)  {
	branding(s)
	for {
		if c.Input {
			input()
		}
	}
}

func input() {
	var input string
	fmt.Printf("To make a Request type [req], to exit [exit]: ")
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
			 fmt.Println("invalid input")
			 return
	}
	switch input {
	case "req":
			// request
	case "exit":
			os.Exit(2)
	default:
			fmt.Println("def")
	}
}

func branding(s *session.Session) {
	fmt.Printf("(c) 2019 George Karalis and Robin's respected authors. Robin is a work in progress.\n")
	fmt.Printf("            __   _    \n")
	fmt.Printf("  _______  / /  (_)__ \n")
	fmt.Printf(" / __/ _ \\/ _ \\/ / _ \\ \n")
	fmt.Printf("/_/  \\___/_.__/_/_//_/ \n")
	fmt.Printf("%s, %s\n", s.Getname(), s.Getdesc())
	fmt.Printf("Starting session, id: %s, pid: %d\n", s.Getid(), s.Getpid())
}