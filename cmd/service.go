package cmd

import (
	"context"

	"github.com/gookit/color"
	"github.com/mamau/satellite/config"
	"github.com/mamau/satellite/config/docker"
	"github.com/mamau/satellite/libs"
	"github.com/mamau/satellite/strategy"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			color.Red.Printf("You should pass service name")
			return
		}
		serviceName := args[0]
		color.Cyan.Printf("Start %s\n", serviceName)
		s := config.GetConfig().GetService(serviceName)

		strategy := determineStrategy(s, args[1:])

		libs.RunCommandAtPTY(Docker(strategy))
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
