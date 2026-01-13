package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter Sentences. Type Enter and Cntr+D to end")
	fmt.Print("$ ")
	count := count()
	fmt.Println("Number of words entered is", count)
}

func count() int {
	// count, _ := countWordsStream()
	count, _ := countWordsStreamNative()
	return count
}

// from stdin
// we could get the input in a stream of tokens. and count as they come or
// we could get it all in a buffer.

func countWordsStream() (int, error) {
	input := bufio.NewScanner(os.Stdin)
	count := 0
	for input.Scan() {
		full_token := input.Text()
		tokens := strings.Fields(full_token)

		count += len(tokens)
	}

	if err := input.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func countWordsStreamNative() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), 1024)

	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
