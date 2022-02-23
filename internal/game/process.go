package game

func (g *Game) run() {
	var idPlayer int
	for g.IsRun() {
		idPlayer = g.whichPlayer()
		x_coordinate, y_coordinate := g.programmInterface.GetCoordinate(g.players[idPlayer])
		err := g.field.SetCoordinate(g.goSide, x_coordinate, y_coordinate)
		if err != nil {
			g.programmInterface.SendMessage("", err)
		} else {
			if g.isСheckViner() {
				g.players[idPlayer].AddVictory()
				str := ("\nПобеда!!!\nПобеду одержал игрок: " + g.players[idPlayer].GetName())
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
