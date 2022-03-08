package stringWorker

import (
	"errors"
)

var ErrBadCoordinates error = errors.New("bad coordinates! error: no correct string")

func GetCoordinate(coordinates string) (uint8, uint8, error) {
	if len(coordinates) != 3 {
		return 0, 0, ErrBadCoordinates
	}
	return uint8(coordinates[0] - 48), uint8(coordinates[2] - 48), nil //не нашел как конвертировать byte в int без 48
}
