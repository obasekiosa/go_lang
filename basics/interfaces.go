package main

import "fmt"

type House struct {
	Color  string
	Width  int
	Breath int
	Height int
}

type Area interface {
	Area() int
}

type Volume interface {
	Volume() int
}

func (u House) Volume() int {
	return u.Breath * u.Height * u.Width
}

func (u House) Area() int {
	return u.Breath * u.Width
}

func (h House) Show() {
	fmt.Printf("This house is color %s. Has an area of %d, and a volume of %d.\n", h.Color, h.Area(), h.Volume())
}

func newHouse(width, breath, height int, color string) House {
	return House{
		color,
		width,
		breath,
		height,
	}
}

func displayHouse() {
	house := newHouse(10, 20, 30, "green")
	house.Show()
}
