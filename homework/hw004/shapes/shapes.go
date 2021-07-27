package shapes

// Shaper abstract interface to implement in a pkg
type Shaper interface {
	Area() (float64, error)
	Perimeter() (float64, error)
	String() string
}
