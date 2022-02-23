package game

import (
	"errors"
	"gb_golang/internal/delivery"
	"gb_golang/internal/field"
	"gb_golang/internal/player"
)

var ErrProgrammInterfaceIsNil error = errors.New("programm interface is nil")

type Game struct {
	isGameRun         bool
	goSide            string
	countGame         uint8
	field             field.Field
	players           [2]player.Player
	programmInterface delivery.ProgrammInterface
}

func New(i delivery.ProgrammInterface) Game {
	if i == nil {
		panic(ErrProgrammInterfaceIsNil)
	}
	var g Game
	g.countGame = 0
	g.isGameRun = false
	g.players = i.GetPlayers()
	g.field = field.New()
	g.programmInterface = i
	return g
}

func (g *Game) IsRun() bool {
	return g.isGameRun
}

func (g *Game) Start() {
	g.isGameRun = true
	g.countGame++
	g.goSide = "X"
	for i := range g.players {
		g.players[i].AddCountGame()
	}
	g.run()
}

func (g *Game) Stop() {
	g.isGameRun = false
}

func (g *Game) GetPlayers() [2]player.Player {
	return g.players
}

func (g *Game) GetCountGame() uint8 {
	return g.countGame
}
