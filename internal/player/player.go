package player

import (
	"errors"
	"gb_golang/internal/tools/stringWorker"
)

type Player struct {
	name               string
	side               string
	victory, countGame int
}

func New(name string, side string) (Player, error) {
	if stringWorker.Checkname(name) && stringWorker.CheckSide(side) {
		p := Player{side: side, name: name, victory: 0, countGame: 0}
		return p, nil
	} else {
		return Player{}, errors.New("Incorrect parameters for creating a player")
	}
}

// func New(name string, side string) Player {
// 	p := Player{side: side, name: name, victory: 0, countGame: 0}
// 	return p
// }

func (p *Player) GetSideBool() string {
	return p.side
}

func (p *Player) GetName() string {
	return p.name
}
