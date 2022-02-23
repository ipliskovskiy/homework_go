package cli

import (
	"fmt"
	"gb_golang/internal/game"
)

type Cli struct{}

func StartGame(g game.Game) {
	var runCLI bool = true
	for runCLI {
		// g.Start()
		statisticGame(g)
		runCLI = false
	}
}

func statisticGame(g game.Game) {
	str := ("\n======================\n")
	str = str + ("Количество партий: " + string(g.GetCountGame()) + "\n")
	for i, v := range g.GetPlayers() {
		str = str + ("")
		str = str + ("Пользователь " + v.GetName() + "\nСторона: " + v.GetSide() + ":\n")
		str = str + ("Всего игр: " + string(v.GetCountGame()) + "\n")
		str = str + ("Количество побед:" + string(v.GetVictory()) + "\n")
		if i != len(g.GetPlayers())-1 {
			str = str + ("\n------------------\n")
		}
	}
	fmt.Println(str)
}
