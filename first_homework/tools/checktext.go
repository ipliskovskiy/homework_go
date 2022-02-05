package tools

import (
	"regexp"
)

const userNameRegexp = `^[A-zА-я]+[A-zА-я0-9]*$`

func Checkname(name string) bool {
	matched, _ := regexp.Match(userNameRegexp, []byte(name))

	if matched == true { // думаю указать matched == true более наглядно чем просто matched
		return true
	}
	return false
}
