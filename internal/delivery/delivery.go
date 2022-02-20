package delivery

import (
	"gb_golang/internal/player"
)

type ProgrammInterface interface {
	GetPlayers() [2]player.Player
	GetCoordinate(p player.Player) (uint8, uint8)
	SendMessage(string, error)
}
