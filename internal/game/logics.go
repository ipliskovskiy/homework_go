package game

import (
	"fmt"
	"gb_golang/internal/player"
)

func (g *Game) WhichPlayer() player.Player {
	if g.players[0].GetSide() == g.goSide {
		return g.players[0]
	} else {
		return g.players[1]
	}
}

func (g *Game) isСheckViner() bool {
	f := g.field.GetField()
	//== for test
	fmt.Println("|", f[0][0], "|", f[0][1], "|", f[0][2])
	fmt.Println("|", f[1][0], "|", f[1][1], "|", f[1][2])
	fmt.Println("|", f[2][0], "|", f[2][1], "|", f[2][2])
	//== for test

	if f[1][1] == g.goSide {
		switch {
		case f[1][0] == g.goSide && f[1][0] == f[1][2]:
			return true
		case f[0][0] == g.goSide && f[0][0] == f[2][2]:
			return true
		case f[0][1] == g.goSide && f[0][1] == f[2][1]:
			return true
		case f[0][2] == g.goSide && f[0][2] == f[2][0]:
			return true
		default:
			return false
		}
	} else if f[0][0] == g.goSide {
		switch {
		case f[0][2] == g.goSide && f[0][1] == f[0][2]:
			return true
		case f[2][0] == g.goSide && f[1][0] == f[2][0]:
			return true
		default:
			return false
		}
	} else if f[2][0] == g.goSide && f[2][0] == f[2][1] && f[2][1] == f[2][2] {
		return true
	} else if f[2][2] == g.goSide && f[2][2] == f[1][2] && f[1][2] == f[0][2] {
		return true
	}

	return false
}
