package commands

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"satellite/internal/config"
	"syscall"
)

var serviceCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, arguments := prepareArgs(args)

		color.Cyan.Printf("Start %q\n", serviceName)

		s := config.GetConfig().FindService(serviceName)

		eCmd := Docker(s, arguments)
		eCmd.Stderr = os.Stderr
		eCmd.Stdout = os.Stdout
		eCmd.Stdin = os.Stdin

		if err := eCmd.Start(); err != nil {
			log.Fatalf("cmd.Start: %v", err)
		}
		if err := eCmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {

				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					if status.ExitStatus() == FAILURE {
						os.Exit(FAILURE)
					}
				}

			} else {
				log.Fatalf("cmd.Wait: %v", err)
			}
		}
	},
}

func prepareArgs(args []string) (string, []string) {
	var serviceName string
	var arguments []string

	if len(args) < 1 {
		color.Red.Printf("You must pass a service name\n")
		os.Exit(1)
	}

	serviceName = args[0]

	if len(args) >= 2 {
		arguments = args[1:]
	}

	return serviceName, arguments
}
