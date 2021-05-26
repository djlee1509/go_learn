package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{height: 20, base: 25}
	s := square{sideLength: 20}

	printArea(t)
	printArea(s)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	squareArea := math.Pow(s.sideLength, 2)
	return squareArea
}

func (t triangle) getArea() float64 {
	triangleArea := t.base * t.height * 0.5
	return triangleArea
}
