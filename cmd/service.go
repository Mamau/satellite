package cmd

import (
	"fmt"

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

		strategy := determineStrategy(s)
		fmt.Println(strategy.ToCommand())

		//if s.SkipArgs == false {
		//	if len(args) < 2 {
		//		color.Red.Printf("You should pass args")
		//		return
		//	}
		//}
		//
		//ser := entity.NewService(s, args[1:])
		//coll := collector.NewCollector(ser)
		libs.RunCommandAtPTY(Docker2(strategy))
	},
}

func determineStrategy(config *docker.Docker) strategy.Strategy {
	//return strategy.NewPullStrategy(config)
	return strategy.NewDaemonStrategy(config)
}
