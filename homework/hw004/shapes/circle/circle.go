package circle

import (
	"fmt"
	"math"

	h "github.com/rodkevich/go-course/homework/hw004/shapes/internal/helpers"
)

// Circle struct describing circle figure
type Circle struct {
	Radius float64
}

// New ...
func New(radius float64) Circle {
	return Circle{Radius: radius}
}

// Area S=πR²
func (c Circle) Area() (float64, error) {
	if h.UsedArgsIncludeInvalid([]float64{c.Radius}...) {
		return 0, h.ErrInvalidArgs
	}
	return math.Pi * math.Pow(c.Radius, 2), nil
}

// Perimeter P=2πR
func (c Circle) Perimeter() (float64, error) {
	if h.UsedArgsIncludeInvalid([]float64{c.Radius}...) {
		return 0, h.ErrInvalidArgs
	}
	return 2 * math.Pi * c.Radius, nil
}

// String prints to output
func (c Circle) String() string {
	return fmt.Sprintf("Circle: Radius %.2f", c.Radius)
}
