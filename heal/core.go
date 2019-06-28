package heal

import (
	"os/exec"
	"time"
	"fmt"
)

func StartHealingProcess()  {
	go func() {
		cmd := exec.Command("robin", "60")
		cmd.Start()
		time.Sleep(60 * time.Second)
		fmt.Printf("Started healing process\n")
	}()
}