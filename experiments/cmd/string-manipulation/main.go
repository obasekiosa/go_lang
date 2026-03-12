package main

import (
	"fmt"
	"unicode/utf8"
)

const strValue = "abcd"
const otherStrValue = "efgh"

func main() {
	combine := strValue + " . " + otherStrValue
	fmt.Printf("string \"%s\" of len %d\n", combine, len(combine))
	indexString(combine, 3)
}

func indexString(s string, index int) {
	var val byte = s[index]
	fmt.Println(rune(val))
	fmt.Println(string(val))
	err := utf8.RuneError
	fmt.Println("Error", err)
	r, si := utf8.DecodeRune([]byte{101})
	fmt.Println("Rune: ", r, "Size:", si)

}
