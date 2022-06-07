package main

import (
	"log"
	"runtime/debug"
	"fmt"
)

type MyError struct {
	text string
	trace string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s", e.text, e.trace)
}

func New(text string) error {
	return &MyError{
		text: text,
		trace: string(debug.Stack()),
	}
}

var STR = "HELLO WORLD"

func whitPanic() {
	log.Println(STR[11])
}

func main() {
	var err error
	defer func() {
		if val := recover(); val != nil {
			log.Println("Panic", val)
			err = New("my error")
			log.Println(err)
		}
	}()

	whitPanic()
}

