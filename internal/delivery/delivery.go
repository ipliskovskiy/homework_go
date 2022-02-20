package delivery

import (
	"gb_golang/internal/player"
)

type ProgrammInterface interface {
	GetPlayers() [2]player.Player
}
