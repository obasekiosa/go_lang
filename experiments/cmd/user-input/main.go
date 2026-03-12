package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	buffScan()
	directScan()
	indirectScan()
}

func directScan() {
	var input float32
	var extra string

	count, err := fmt.Scanf("%f%s", &input, &extra)

	fmt.Println("Read", count, "Values")
	if err != nil {
		fmt.Println("There was an error", err)
	}

	fmt.Println("Input was:", input)
}

func indirectScan() {
	var input float32
	var extra string

	rawInput, err := io.ReadAll(os.Stdin)
	if err != nil {
		err = fmt.Errorf("failed to read terminal input %w", err)
	}
	count, err2 := fmt.Sscanf(string(rawInput), "%f%s", &input, &extra)

	fmt.Println("Read", count, "Values")
	if err2 != nil {
		err = errors.Join(err, err2)
		fmt.Println("There was an error", err)
	}

	fmt.Println("Input was:", input)
	fmt.Println("From raw input \"", string(rawInput), "\"")
}

func buffScan() {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')

	if err != nil && err != io.EOF {
		fmt.Println("Error occured", err)
	}

	input, err := strconv.ParseFloat(strings.TrimSpace(line), 32)
	if err != nil {
		fmt.Println("input parse error", err)
	}
	input32 := float32(input)
	fmt.Println("Input was ", input32)
}
