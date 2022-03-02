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

// func (g *Game) GetStatistic() {
// 	str := ("\n======================\n")
// 	str = str + ("Количество партий: " + str(g.countGame) + "\n")
// 	for i := range g.players {
// 		str = str + ("")
// 		str = str + ("Пользователь " + g.players[i].GetName() + "\nСторона: " + g.players[i].GetSide() + ":\n")
// 		str = str + ("Всего игр: " + string(g.players[i].GetCountGame()) + "\n")
// 		str = str + ("Количество побед:" + string(g.players[i].GetVictory()) + "\n")
// 		if i != len(g.players)-1 {
// 			str = str + ("\n------------------\n")
// 		}
// 	}
// 	g.programmInterface.SendMessage(str, nil)
// }
