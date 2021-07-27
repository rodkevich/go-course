package main

import (
	"fmt"

	"github.com/rodkevich/go-course/homework/hw004/shapes"
	"github.com/rodkevich/go-course/homework/hw004/shapes/circle"
	"github.com/rodkevich/go-course/homework/hw004/shapes/rectangle"
)

func main() {

	c := circle.New(12)
	r := rectangle.New(7, 6)

	DescribeShape(c)
	DescribeShape(r)
}

// DescribeShape pretty print for type
func DescribeShape(s shapes.Shaper) {
	fmt.Println(s)

	area, errA := s.Area()
	if errA != nil {
		// fmt.Println(errA.Error())
		panic(fmt.Sprintf("Error: %v", errA.Error()))
	}
	fmt.Printf("Area: %.2f\n", area)

	perimeter, errP := s.Perimeter()
	if errP != nil {
		// fmt.Println(errP.Error())
		panic(fmt.Sprintf("Error: %v", errP.Error()))
	}
	fmt.Printf("Perimeter: %.2f\n", perimeter)
}
