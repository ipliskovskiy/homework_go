package cli

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func Hello() {
	printLOGO()
	fmt.Println("\n\n\nДоброго времени суток.", emoji.WavingHand, "\nНачинаем консольную версию игры крестики нолики. ", emoji.ManTechnologist, "\n\n")
}

func printLOGO() {
	fmt.Println("\n\n\n[̲̅Х̲̅][̲̅О̲̅][̲̅Р̲̅][̲̅О̲̅][̲̅Ш̲̅][̲̅О̲̅][̲̅Ч̲̅][̲̅Т̲̅][̲̅О̲̅][̲̅Н̲̅][̲̅Е̲̅][̲̅П̲̅][̲̅Р̲̅][̲̅О̲̅][̲̅Д̲̅][̲̅А̲̅][̲̅К̲̅][̲̅Ш̲̅][̲̅Н̲̅]")
}
