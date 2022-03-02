package cli

import (
	"fmt"
	"gb_golang/internal/field"
)

func (с Cli) SendMessage(str string, err error) {
	if str != "" {
		fmt.Println(str)
	}
	switch err {
	case field.ValidateCoordinate:
		fmt.Println("Не корректные координаты.")
	case field.FieldIsNotEmpty:
		fmt.Println("Ячейка занята!")
	}
}
