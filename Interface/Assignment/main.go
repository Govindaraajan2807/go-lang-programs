package main

import "fmt"

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

	tr := triangle{base: 10, height: 10}
	sq := square{sideLength: 10}

	printArea(tr)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	// height := 3.0
	// base := 3.0
	fmt.Print("Area of triangle is : ")
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	// sideLength := 2.0
	fmt.Print("Area of Square is : ")
	return s.sideLength * s.sideLength
}
