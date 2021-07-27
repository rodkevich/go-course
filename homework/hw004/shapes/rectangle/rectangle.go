package rectangle

import (
	"errors"
	"fmt"

	h "github.com/rodkevich/go-course/homework/hw004/shapes/internal/helpers"
)

// Rectangle struct describing rectangle figure
type Rectangle struct {
	Height float64
	Width  float64
}

// New make a new shape of rectangle type
func New(height float64, width float64) Rectangle {
	return Rectangle{Height: height, Width: width}
}

// Area S=a*b
func (r Rectangle) Area() (float64, error) {
	if h.UsedArgsIncludeInvalid([]float64{r.Width, r.Height}) {
		return 0, errors.New(h.UsingInvalidArgs)
	}
	return r.Width * r.Height, nil
}

// Perimeter P=2(a+b)
func (r Rectangle) Perimeter() (float64, error) {
	if h.UsedArgsIncludeInvalid([]float64{r.Width, r.Height}) {
		return 0, errors.New(h.UsingInvalidArgs)
	}
	return 2 * (r.Height + r.Width), nil
}

// String prints to output
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with Height %.2f and Width %.2f", r.Height, r.Width)
}
