package player

import (
	"errors"
	"gb_golang/internal/tools/stringWorker"
)

var BadParametrsForCreatingPlayer error = errors.New("incorrect parameters for creating a player")

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
		return Player{}, BadParametrsForCreatingPlayer
	}
}

func (p *Player) GetSide() string {
	return p.side
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetCountVictory() int {
	return p.victory
}

func (p *Player) GetCountGame() int {
	return p.countGame
}

func (p *Player) AddVictory() {
	p.victory++
}

func (p *Player) AddCountGame() {
	p.countGame++
}
