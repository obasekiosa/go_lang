package main

import "fmt"

var r int = 2

const pi float32 = 3.142

func circleCircumference() {
	circum := 2 * pi * float32(r)
	fmt.Println(circum)
}

const name string = "Seki🗡️"

func sayName() {
	fmt.Println(name)
}
