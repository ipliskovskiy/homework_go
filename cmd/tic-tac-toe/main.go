package main

import (
	"gb_golang/internal/cli"
	"gb_golang/internal/delivery"
	"gb_golang/internal/game"
)

func main() {
	var g game.Game
	var i delivery.ProgrammInterface = cli.Cli{}
	cli.Hello()
	g = game.New(i)
	cli.StartGameCLI(g)
}
