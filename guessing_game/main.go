package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const MIN_GUESS_RANGE, MAX_GUESS_RANGE int = 0, 100

func main() {

	// generate guess
	expected_guess := MIN_GUESS_RANGE + int(float32(MAX_GUESS_RANGE)*rand.Float32())
	fmt.Println(expected_guess)

	input := bufio.NewScanner(os.Stdin)
	// loop
	fmt.Print("$ ")
	for ; input.Scan(); fmt.Print("$ ") {

		line := input.Text()
		guess, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%s is not a valid integer number\nEnter another value.\n", line)
			continue
		}

		if expected_guess > guess {
			fmt.Println("Too Low.")
		} else if expected_guess < guess {
			fmt.Println("Too Hig")
		} else {
			fmt.Printf("You got it!!\nThe Guess was %d.\n", expected_guess)
			break
		}
	}
	// buff := bytes.Buffer{}

	// fmt.Fprint(&buff, "hello\n")
	// str := buff.String()
	// fmt.Println(str)
	// os.Stdout.Write(buff.Bytes())
}
