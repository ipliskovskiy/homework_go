package cli

import (
	"fmt"
	"gb_golang/internal/game"
	"strconv"
)

type Cli struct{}

func StartGameCLI(g game.Game) {
	var runCLI bool = true
	for runCLI {
		g.Start()
		statisticGame(g)
		runCLI = RepeatGame()
	}
}

func statisticGame(g game.Game) {
	str := ("\n=========== Статистика по игре ===========\n")
	str = str + ("Количество партий: " + strconv.Itoa(int(g.GetCountGame())) + "\n")
	str = str + ("------------------\n")
	for i, v := range g.GetPlayers() {
		str = str + ("")
		str = str + ("Пользователь " + v.GetName() + "\nСторона: " + v.GetSide() + "\n")
		str = str + ("Всего игр: " + strconv.Itoa(v.GetCountGame()) + "\n")
		str = str + ("Количество побед: " + strconv.Itoa(v.GetCountVictory()) + "\n")
		if i != len(g.GetPlayers())-1 {
			str = str + ("------------------\n")
		}
	}
	fmt.Println(str)
}
