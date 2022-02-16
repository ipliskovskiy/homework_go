package game

import (
	"fmt"
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
		fmt.Println(g.checkViner())
		if g.goSide == "X" {
			g.goSide = "0"
		} else {
			g.goSide = "X"
		}
		return nil
	}
	return err
}

func (g *Game) checkViner() bool {
	f := g.field.GetField()
	//== for test
	fmt.Println("|", f[0][0], "|", f[0][1], "|", f[0][2])
	fmt.Println("|", f[1][0], "|", f[1][1], "|", f[1][2])
	fmt.Println("|", f[2][0], "|", f[2][1], "|", f[2][2])
	//== for test

	if f[1][1] == g.goSide {
		switch {
		case f[1][0] == f[1][2] && f[1][2] == g.goSide:
			return true
		case f[0][0] == f[2][2] && f[2][2] == g.goSide:
			return true
		case f[2][0] == f[0][2] && f[0][2] == g.goSide:
			return true
		case f[0][1] == f[2][2] && f[2][2] == g.goSide:
			return true
		default:
			return false
		}
	} else if f[0][0] == g.goSide {
		switch {
		case f[0][1] == f[0][2] && f[0][2] == g.goSide:
			return true
		case f[1][0] == f[2][0] && f[2][0] == g.goSide:
			return true
		default:
			return false
		}
	} else if f[2][0] == f[2][1] && f[2][1] == f[2][2] && f[2][2] == g.goSide {
		return true
	} else if f[0][2] == f[1][2] && f[1][2] == f[2][2] && f[2][2] == g.goSide {
		return true
	}

	return false
}
