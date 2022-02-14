package cli

import (
	"fmt"
	"gb_golang/internal/tools/stringWorker"
	"os"
)

func CLI_GetSide() string {
	var HaveSide = false
	var side string
	for !HaveSide {
		fmt.Print("Введите наименование стороны за которую вы играете X/0 или exit для выхода: ")
		fmt.Scanf("%s\n", &side)

		if side == "exit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		if stringWorker.CheckSide(side) {
			HaveSide = true
		} else {
			fmt.Println("Вы ввели не корректное значение стороны, сторона может быть X или 0")
		}
	}

	return side
}
