package cli

import (
	"fmt"
	"gb_golang/internal/tools/stringWorker"
	"os"
)

func GetSide() string {
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
			fmt.Println("Вы ввели не корректное значение координат, коордитинаты должны быть [0-2].[0-2]")
		}
	}

	return side
}
