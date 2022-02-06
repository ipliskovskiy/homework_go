package tools

import (
	"regexp"
)

var userNameRegexp = regexp.MustCompile(`^[A-zА-я]+[A-zА-я0-9]*$`)

func Checkname(name string) bool {
	matched := userNameRegexp.MatchString(name)

	return matched
}
