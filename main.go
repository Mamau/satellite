package main

import (
	"github.com/mamau/starter/cmd"
	"github.com/mamau/starter/libs"
)

func main() {
	libs.LoadEnv()
	cmd.Execute()
}
