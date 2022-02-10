package tools

import "errors"

func GetCoordinate(coordinates string) (string, string, error) {
	if len(coordinates) != 3 {
		return "", "", errors.New("no correct string")
	}
	return string(coordinates[0]), string(coordinates[2]), nil
}
