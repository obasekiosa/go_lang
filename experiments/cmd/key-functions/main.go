package main

import (
	"fmt"
)

type fType func() int

var fu func(string) int = func(a string) int {
	return 1
}

func add(a, b int) int {
	return a + b
}

func mod(a, b int) int {
	return a % b
}

func minus(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

type FInt2 = func(int, int) int

var funcs [4]func(int, int) int = [4]func(int, int) int{add, minus, mod, mult}

func playAndSeeFunc() {
	a, b := 10, 3
	for i, f := range funcs {
		fmt.Println("Index: ", i, "Name: ", f, "Output: ", f(a, b))
	}
	fmt.Println("Output: ", funcs)
}

func playWithVariadic() {
	applyAll(add, 1, 2, 3, 4, 5)
	input := []int{3, 4, 2, 8}
	applyAll(add, input...)
}

func applyAll(f FInt2, all ...int) {
	result := 0
	for _, a := range all {
		result = f(a, result)
		fmt.Println("Partial Result:", result, "From:", a)
	}

	fmt.Println("Result:", result, "Input:", all, "Size:", len(all))
}

func repeat(val any) {
	fmt.Println(val)
}

func deferInOrder() {
	defer repeat(1)
	repeat(0)
	defer repeat(2)

}

func panicAndRecover() {
	defer func() {
		repeat(5)
		recover()
		repeat(7)
	}()

	panic("Did not work")
}

func pointerPlayer() {
	xPtr := new(int)
	repeat(xPtr)
	repeat(*xPtr)

	*xPtr += 4
	repeat(xPtr)
	repeat(*xPtr)
}

func makeOddGenerator(start int) func() int {
	if start%2 == 0 {
		start++
	}

	return func() int {
		val := start
		start += 2
		return val
	}
}

func playOdd() {
	gen := makeOddGenerator(10)
	for i := 0; i < 10; i++ {
		repeat(gen())
	}
}

func main() {
	// playAndSeeFunc()
	// playWithVariadic()
	// deferInOrder()

	panicAndRecover()
	pointerPlayer()
	playOdd()
}
