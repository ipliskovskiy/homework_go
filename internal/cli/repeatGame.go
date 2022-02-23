package cli

import (
	"fmt"
)

func RepeatGame() bool {
	var answer string
	var result bool = false

	fmt.Println("Нажмите Enter для повтора игры или exit для выхода: ")
	fmt.Scanf("%s\n", &answer)
	if answer == "exit" {
		fmt.Println("Game over. By.")
		result = false
	} else {
		result = true
	}

	return result
}
