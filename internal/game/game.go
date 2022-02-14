package game

import (
	"gb_golang/internal/field"
	"gb_golang/internal/player"
)

type Game struct {
	isGameRun bool
	goSide    string
	countGame uint8
	field     field.Field
	players   [2]player.Player
}

func New(field field.Field, p [2]player.Player) Game {
	var g Game
	g.countGame = 0
	g.isGameRun = false
	g.field = field
	g.players = p
	return g
}

func (g *Game) IsRun() bool {
	return g.isGameRun
}

func (g *Game) Start() bool {
	g.isGameRun = true
	g.countGame++
	g.goSide = "X"
	return g.isGameRun
}

func (g *Game) WhichPlayer() player.Player {
	if g.players[0].GetSide() == g.goSide {
		return g.players[0]
	} else {
		return g.players[1]
	}
}

func (g *Game) MakeMove(coordinate string) {
	g.field.SetCoordinate(g.goSide, coordinate)
}
