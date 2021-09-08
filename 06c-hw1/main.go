package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	square := square{sideLength: 6}
	triangle := triangle{32.1, 5.24}

	printArea(square)
	printArea(triangle)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println("Area is", area)
}
