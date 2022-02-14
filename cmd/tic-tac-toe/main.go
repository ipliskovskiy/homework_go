package main

import (
	"gb_golang/internal/cli"
	"gb_golang/internal/game"
)

func main() {
	var g game.Game
	g = game.New(cli.CLI_PrepareToGame())
	g.Start()

	for g.IsRun() {
		cli.CLI_GetCoordinate(g)

	}

}
