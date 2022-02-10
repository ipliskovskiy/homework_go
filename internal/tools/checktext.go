package tools

import "regexp"

var reUserNameRegexp = regexp.MustCompile(`^[A-zА-я]+[A-zА-я0-9]*$`)
var reCoordinates = regexp.MustCompile(`^[0-2]+(.|,|-|:)+[0-2]$`)

func Checkname(name string) bool {
	return reUserNameRegexp.MatchString(name)
}

func CheckCoordinates(coordinates string) bool {
	return reCoordinates.MatchString(coordinates)
}
