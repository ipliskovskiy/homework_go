package stringWorker

import "regexp"

var reUserNameRegexp = regexp.MustCompile(`^[A-zА-я]+[A-zА-я0-9]*$`)
var reCoordinates = regexp.MustCompile(`^[0-2]+(.|,|-|:)+[0-2]$`)
var reSide = regexp.MustCompile(`X|0`)

func Checkname(name string) bool {
	return reUserNameRegexp.MatchString(name)
}

func CheckCoordinates(coordinates string) bool {
	return reCoordinates.MatchString(coordinates)
}

func CheckSide(side string) bool {
	return reSide.MatchString(side)
}
