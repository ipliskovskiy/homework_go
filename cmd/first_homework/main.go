package main

import (
	"fmt"
	"gb_golang/internal/tools"
	"os"

	"github.com/enescakir/emoji"
)

func main() {
	var name string
	var HaveAnswer bool = false

	fmt.Println("Доброго времени суток.", emoji.WavingHand)

	for !HaveAnswer {
		fmt.Print(emoji.ManTechnologist, " Введите имя или exit для выхода: ")
		fmt.Scanf("%s\n", &name)
		if name == "exit" {
			fmt.Println("Goodbye ", emoji.DizzyFace)
			os.Exit(0)
		}
		if tools.Checkname(name) {
			HaveAnswer = true
			// break
		} else {
			fmt.Println("Извините, но Вы ввыли не корректное имя. Имя может содержать цифры, строчные и заглавные буквы русского и английского языка, но не может начинаться с цифры или содержать знаки припенания.")
		}
	}

	fmt.Println("Приветствую тебя, ", name, emoji.PartyingFace)

}
