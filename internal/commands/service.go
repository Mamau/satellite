package commands

import (
	"context"

	"github.com/mamau/satellite/pkg"

	config2 "github.com/mamau/satellite/internal/config"
	docker2 "github.com/mamau/satellite/internal/config/docker"

	strategy2 "github.com/mamau/satellite/internal/strategy"

	"github.com/gookit/color"
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
		s := config2.GetConfig().GetService(serviceName)

		strategy := determineStrategy(s, args[1:])

		pkg.RunCommandAtPTY(Docker(strategy))
	},
}

func determineStrategy(config *docker2.Docker, args []string) strategy2.Strategy {
	parent := context.Background()
	if config.GetDockerCommand() == strategy2.PullType {
		ctx := context.WithValue(parent, "type", strategy2.PullType)
		return strategy2.NewPullStrategy(ctx, config)
	}

	if config.Detach {
		ctx := context.WithValue(parent, "type", strategy2.DaemonType)
		return strategy2.NewDaemonStrategy(ctx, config)
	}

	ctx := context.WithValue(parent, "type", strategy2.RunType)
	return strategy2.NewRunStrategy(ctx, config, args)
}
