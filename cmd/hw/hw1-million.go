package main

import (
	"strconv"
	"os"
	"fmt"
)

func createFile(filename string) () {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		f.Close()
	}()
}


func main() {
	path := "/Users/ipliskovskiy/go/gb_golang/test/"
	for i := 1; i <= 1000000; i++ {
		createFile(path + "file-" + strconv.Itoa(i))
	}
	fmt.Println("OKAY")
}