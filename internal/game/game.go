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

func (g *Game) Stop() bool {
	g.isGameRun = false
	return !g.isGameRun
}

func (g *Game) WhichPlayer() player.Player {
	if g.players[0].GetSide() == g.goSide {
		return g.players[0]
	} else {
		return g.players[1]
	}
}

func (g *Game) MakeMove(x_coordinate uint8, y_coordinate uint8) error {
	err := g.field.SetCoordinate(g.goSide, x_coordinate, y_coordinate)
	if err == nil {
		if g.goSide == "X" {
			g.goSide = "0"
		} else {
			g.goSide = "X"
		}
		return nil
	}
	return err
}
