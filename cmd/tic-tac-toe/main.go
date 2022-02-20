package main

import (
	"gb_golang/internal/cli"
	"gb_golang/internal/delivery"
	"gb_golang/internal/game"
)

func main() {
	var g game.Game
	cli.Hello()
	var i delivery.ProgrammInterface = cli.Cli{}
	g = game.New(i)

	g.Start()

}
