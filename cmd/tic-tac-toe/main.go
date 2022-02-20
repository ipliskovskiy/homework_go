package main

import (
	"gb_golang/internal/cli"
	"gb_golang/internal/delivery"
	"gb_golang/internal/game"
)

func main() {
	var g game.Game
	var i delivery.ProgrammInterface = cli.Cli{}

	g = game.New(i)
	g.IsRun()

	// g = game.New(cli.CLI_PrepareToGame())
	// g.Start()

	// for g.IsRun() {
	// 	x_coordinate, y_coordinate := cli.CLI_GetCoordinate(g.WhichPlayer())
	// 	err := g.MakeMove(uint8(x_coordinate), uint8(y_coordinate))
	// 	if err != nil {
	// 		cli.PrintError(err)
	// 	}
	// }

}
