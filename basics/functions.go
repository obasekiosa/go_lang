package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func callAdd() {
	a, b := 1, 3

	fmt.Println(add(a, b))
}

func swap(a int, b int) (int, int) {
	return b, a
}

func callSwap() {
	a, b := 10, 20

	fmt.Println(swap(a, b))
}
