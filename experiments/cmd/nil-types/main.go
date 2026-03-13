package main

import "fmt"

type St struct{}

func main() {
	var s []int = nil
	var m map[string]int = nil
	var ci chan int = nil
	var st *St = nil
	var sp *[]int = nil
	var mp *map[string]int = nil
	var cip *chan int = nil

	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(ci)
	fmt.Println(st)
	fmt.Println(sp)
	fmt.Println(mp)
	fmt.Println(cip)

	var val int = 8

	fmt.Println(float32(18) / float32(val))
}
