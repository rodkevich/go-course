/*
Lesson_04:
Implement Shape interface for Circle and Perimeter structs and use it to print information
about them with the DescribeShape function.
The result of calling DescribeShape function must output into terminal the following
for a Circle and Rectangle(the numbers may differ depending on the dimensions you picked):

Circle: radius 8.00
Area: 201.06
Perimeter: 50.27

Rectangle with height 9.00 and width 3.00
Area: 27.00
Perimeter: 24.00

You must not change the DescribeShape function!
Hint: To be able to do that you will need to implement fmt.Stringer for your shapes
Hint #2: you might need to use math.Pi and math.Pow from the math package in the standard go library:
https://pkg.go.dev/math (this is preferred way, but you definitely can do it without it)
*/

package main

import (
	"fmt"
	"math"
)

// Shape abstract interface to implement in a pkg
type Shape interface {
	Area() interface{}
	Perimeter() interface{}
	String() string
}

// Circle struct describing circle figure
type Circle struct {
	radius float64
}

// Area S=πR²
func (c Circle) Area() interface{} {
	var rtn float64 = math.Pi * math.Pow(c.radius, 2)
	return rtn
}

// Perimeter P=2πR
func (c Circle) Perimeter() interface{} {
	var rtn float64 = 2 * math.Pi * c.radius
	return rtn
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", c.radius)
}

// Rectangle struct describing rectangle figure
type Rectangle struct {
	height float64
	width  float64
}

// Area S=a*b
func (r Rectangle) Area() interface{} {
	var rtn float64 = r.width * r.height
	return rtn
}

// Perimeter P=2(a+b)
func (r Rectangle) Perimeter() interface{} {
	var rtn float64 = 2 * (r.height + r.width)
	return rtn
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", r.height, r.width)
}

// DescribeShape pretty print for type
func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// choose your own dimensions
	c := Circle{radius: 8}
	r := Rectangle{
		height: 9,
		width:  3,
	}

	DescribeShape(c)
	DescribeShape(r)
}

/*
Results:

Circle: radius 8.00
Area: 201.06
Perimeter: 50.27
Rectangle with height 9.00 and width 3.00
Area: 27.00
Perimeter: 24.00

Process finished with exit code 0
*/
