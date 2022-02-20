package game

import (
	"gb_golang/internal/player"
)

func (g *Game) run() {
	var p player.Player
	for g.IsRun() {
		p = g.WhichPlayer()
		x_coordinate, y_coordinate := g.programmInterface.GetCoordinate(p)
		err := g.field.SetCoordinate(g.goSide, x_coordinate, y_coordinate)
		if err != nil {
			g.programmInterface.SendMessage("", err)
		} else {
			if g.isСheckViner() {
				str := ("\nПобеда!!!\nПобеду одержал игрок: " + p.GetName())
				p.AddVictory()
				g.programmInterface.SendMessage(str, nil)
				g.Stop()
			} else if g.field.IsDeskFull() {
				str := ("\nНичья.")
				g.programmInterface.SendMessage(str, nil)
				g.Stop()
			} else {
				if g.goSide == "X" {
					g.goSide = "0"
				} else {
					g.goSide = "X"
				}
			}
		}
	}
}
