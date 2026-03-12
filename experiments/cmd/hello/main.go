package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	helloWithExit()
}

func helloWithExit() {
	os.Exit(1)
}
