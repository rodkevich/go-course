package main

import (
	"fmt"
	"github.com/rodkevich/go-course/homework/hw004/shapes"
	"github.com/rodkevich/go-course/homework/hw004/shapes/circle"
	"github.com/rodkevich/go-course/homework/hw004/shapes/rectangle"
)

// DescribeShape pretty print for type
func DescribeShape(s shapes.Shaper) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// choose your own dimensions
	// c := circle.Circle{Radius: 8}
	// r := rectangle.Rectangle{
	// 	Height: 9,
	// 	Width:  3,
	// }

	c := circle.New(12)
	r := rectangle.New(7, 6)

	DescribeShape(c)
	DescribeShape(r)
}
