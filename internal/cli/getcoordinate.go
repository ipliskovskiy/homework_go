package cli

import (
	"fmt"
	"gb_golang/internal/game"
	"gb_golang/internal/player"
	"gb_golang/internal/tools/stringWorker"
	"os"
)

func CLI_GetCoordinate(g game.Game) (int, int) {
	HaveCoordinates := false
	var coordinates string
	var p player.Player
	p = g.WhichPlayer()

	fmt.Println("Ходит ", p.GetName(), "\nВаша сторона: ", p.GetSide())
	for !HaveCoordinates {
		fmt.Print("Введите координаты X и Y или exit для выхода: ")
		fmt.Scanf("%s\n", &coordinates)

		if coordinates == "exit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		if stringWorker.CheckCoordinates(coordinates) {
			HaveCoordinates = true
		} else {
			fmt.Println("Вы ввели не корректное значение координат, коордитинаты должны быть [0-2].[0-2]")
		}
	}

	Xcoordinate, Ycoordinate, err := stringWorker.GetCoordinate(coordinates)

	if err == nil {
		return int(Xcoordinate), int(Ycoordinate)
	} else {
		fmt.Println("Error: ", err)
		return int(Xcoordinate), int(Ycoordinate)
	}
}
