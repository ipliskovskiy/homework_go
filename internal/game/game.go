package game

import (
	"gb_golang/internal/cli"
	"gb_golang/internal/field"
	"gb_golang/internal/player"
)

type Game struct {
	countGame int
	field     field.Field
	players   [2]player.Player
}

func New(f field.Field) Game {
	var g Game
	g.countGame = 0
	g.field = field.New()
	for x := 0; x < 2; x++ {
		g.players[x], _ = player.New(cli.GetPlayerName(), cli.GetSide())
	}
	return g
}