package stringWorker

import "errors"

func GetCoordinate(coordinates string) (int, int, error) {
	if len(coordinates) != 3 {
		return -1, -1, errors.New("Bad coordinates! Error: no correct string")
	}
	return int(coordinates[0]), int(coordinates[2]), nil
}
