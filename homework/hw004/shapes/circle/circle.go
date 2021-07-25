package circle

import (
	"fmt"
	"math"
)

// Circle struct describing circle figure
type Circle struct {
	Radius float64
}

func New(radius float64) Circle {
	return Circle{Radius: radius}
}

// Area S=πR²
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Perimeter P=2πR
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: Radius %.2f", c.Radius)
}
