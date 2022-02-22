package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLength: 2}
	printArea(s)

	t := triangle{height: 1, base: 1}
	printArea(t)

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
