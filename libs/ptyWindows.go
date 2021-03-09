// +build windows

package libs

import (
	"log"
	"os"
	"os/exec"
)

func RunCommandAtPTY(c *exec.Cmd) {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		log.Fatalf("cant start command %s, cause: %s\n", c.String(), err)
	}
}
