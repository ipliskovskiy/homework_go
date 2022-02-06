package main

import (
	"fmt"
	"gb_golang/internal/tools"
	"os"
)

func main() {
	var coordinates string
	var HaveCoordinates bool = false

	for !HaveCoordinates {
		fmt.Print("Введите координаты X и Y или exit для выхода: ")
		fmt.Scanf("%s\n", &coordinates)

		if coordinates == "exit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		if tools.CheckCoordinates(coordinates) {
			HaveCoordinates = true
		} else {
			fmt.Println("Вы ввели не корректное значение координат, коордитинаты должны быть [0-2].[0-2]")
		}
	}

	Xcoordinate, Ycoordinate, err := tools.GetCoordinate(coordinates)

	if err == nil {
		fmt.Println("Ваш X: ", Xcoordinate)
		fmt.Println("Ваш Y: ", Ycoordinate)
	} else {
		fmt.Println("Error: ", err)
	}

}
