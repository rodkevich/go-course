package rectangle

import "fmt"

// Rectangle struct describing rectangle figure
type Rectangle struct {
	Height float64
	Width  float64
}

// NewRectangle make a new shape of rectangle type
func NewRectangle(height float64, width float64) Rectangle {
	return Rectangle{Height: height, Width: width}
}

// Area S=a*b
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter P=2(a+b)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with Height %.2f and Width %.2f", r.Height, r.Width)
}
