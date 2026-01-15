package main

import "fmt"

func increment(x *int) {
	*x++
}

func showNumber(x *int, isAddress bool) {
	if isAddress {
		fmt.Println("Current address is", x)
	} else {
		fmt.Println("Current number is", *x)
	}

}

func doIncrement() {
	num := 34
	num_ptr := &num
	showNumber(num_ptr, true)
	showNumber(num_ptr, false)
	increment(num_ptr)
	showNumber(num_ptr, false)
}
