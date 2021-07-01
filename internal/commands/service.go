package commands

import (
	"context"
	"os/exec"
	"strings"

	"github.com/mamau/satellite/pkg"

	"github.com/mamau/satellite/internal/config"
	"github.com/mamau/satellite/internal/config/docker"
	"github.com/mamau/satellite/internal/strategy"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			color.Red.Printf("You should pass service name\n")
			return
		}
		serviceName := args[0]
		color.Cyan.Printf("Start %s\n", serviceName)
		s := config.GetConfig().GetService(serviceName)

		createNetwork(s)
		strategy := determineStrategy(s, args[1:])

		pkg.RunCommandAtPTY(Docker(strategy))
	},
}

func determineStrategy(config *docker.Docker, args []string) strategy.Strategy {
	parent := context.Background()
	if config.GetDockerCommand() == strategy.PullType {
		ctx := context.WithValue(parent, "type", strategy.PullType)
		return strategy.NewPullStrategy(ctx, config)
	}

	if config.Detach {
		ctx := context.WithValue(parent, "type", strategy.DaemonType)
		return strategy.NewDaemonStrategy(ctx, config)
	}

	ctx := context.WithValue(parent, "type", strategy.RunType)
	return strategy.NewRunStrategy(ctx, config, args)
}

func createNetwork(config *docker.Docker) {
	if config.Network == "" {
		return
	}

	data := []string{
		"network",
		"inspect",
		config.Network,
	}

	cmd := exec.Command(commandName, data...)
	network, _ := cmd.Output()

	if strings.TrimSuffix(string(network), "\n") == "[]" {
		color.Yellow.Printf("Creating network %s\n", config.Network)
		data := []string{
			"network",
			"create",
			config.Network,
		}
		cmd := exec.Command(commandName, data...)
		out, _ := cmd.Output()

		color.Cyan.Printf("%s\n", out)
	}
}
