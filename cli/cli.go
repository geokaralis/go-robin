package cli

import (
	"fmt"
	"os"
)

type Cli struct {

}

func (c *Cli) Init()  {
	for {
		input()
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