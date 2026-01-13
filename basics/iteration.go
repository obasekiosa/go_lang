package main

import "fmt"

func forLoop() {
	size := 5
	for i := 0; i < size; i++ {
		fmt.Print(i)
		if i < size-1 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

func whileLoop() {
	nsize := 5
	for ; ; nsize-- {
		if nsize > 0 {
			fmt.Print(nsize)
			if nsize != 1 {
				fmt.Print(", ")
			}
		} else {
			fmt.Println()
			break
		}
	}
}

func infiniteLoop() {
	size := 5
	for {
		if size > 0 {
			fmt.Print(size)
			if size != 1 {
				fmt.Print(", ")
			}
		} else {
			fmt.Println()
			break
		}
		size--
	}
}
