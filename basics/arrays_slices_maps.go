package main

import (
	"fmt"
)

func arrays() {
	arraysAppend()
}

func arraysAppend() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	nums = append(nums, 4)
	fmt.Println(nums)
}

func maps() {
	mapMutate()
}

func mapMutate() {
	m := map[string]int{
		"red": 1,
	}
	fmt.Println(m)

	m["green"] = 0
	fmt.Println(m)
}
