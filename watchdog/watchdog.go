package watchdog

import (
	"fmt"
	"os"
	"syscall"
)

type State int

const (
	Running State = iota
	Killed
	Unavailable
	Null
)

func WatchProcess() {
	fmt.Printf("pid: %d\n", os.Getpid())

	pid := os.Getpid()

	process, err := os.FindProcess(int(pid))
	if err != nil {
		fmt.Printf("Failed to find process: %s\n", err)
		os.Exit(0)
	} else {
		sig := process.Signal(syscall.Signal(0))
		fmt.Printf("process.Signal on pid %d returned: %v\n", pid, sig)
	}
}

func Sig(signal int) {
	state := Null
	switch state {
	case Running:
		fmt.Printf("Process is running\n")
	}
}