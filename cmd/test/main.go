package main

import (
	"fmt"
)

func f(a []int) {
	a[1] = 666
}

func main() { // slice

	a := []int{1, 2, 34, 5, 6, 7, 8}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(a, b)
	a[1] = 9999
	fmt.Println(a, b)
	f(a)
	fmt.Println(a, b)

	// q := a[4:6]
	// fmt.Println(q, len(q), cap(q))
	// a = append(a, 9, 9, 9, 9, 9, 9, 9)
	// q[1] = 999

	// fmt.Println(a)

	// fmt.Println("-------")
	// b := make([]int, 7, 10000)
	// fmt.Println(b, len(b), cap(b))

	// // слайс это указатель на элементы масива

	// fmt.Println(b[:len(b)+10]) // нет выхода

	// // copy slice
	// z := []int{1, 2, 34, 5, 6, 7, 8}

	// //c := z //не копирует

	// c := make([]int, len(z)) // копирует
	// copy(c, z) //

	// fmt.Println(z, c)
	// c[1] = 999
	// fmt.Println(z, c)

	// MAP

	m := map[string]int{}
	m["hello"] = 1
	val := m["hello"]

	if val, ok := m["hello"]; ok { // что за IF?
		fmt.Println("its here:", val)
	}
	fmt.Println(val, m["a"])

	// ========
	nm3 := make(map[string]int, 100)
	fmt.Println(nm3)

}
