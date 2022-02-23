package main

import "fmt"

func main() {
	var str [2][5]string
	fmt.Println(len(str[1]), cap(str))
}
