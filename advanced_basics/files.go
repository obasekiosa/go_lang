package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile() {
	file, err := os.Open("hello.txt")

	if err != nil {
		fmt.Printf("Error occured %s", err.Error())
		return
	}
	defer file.Close()

	var content strings.Builder
	io.Copy(&content, file)

	for line := range strings.Lines(content.String()) {
		fmt.Printf(" -> %s", line)
	}
}

func createFile() {
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Printf("Could not open: %s\n", err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()
	for i := 0; i < 26; i++ {
		w.WriteRune(rune(65 + (i % 26)))
	}
	w.WriteString("\n")
}
