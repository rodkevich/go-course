package shapes

// Shaper abstract interface to implement in a pkg
type Shaper interface {
	Area() float64
	Perimeter() float64
	String() string
}
