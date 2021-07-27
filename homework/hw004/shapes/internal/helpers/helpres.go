package helpers

import "math"

// UsingInvalidArgs ...
// 	to use if args u get do not satisfy requirements
const UsingInvalidArgs = "Error: vars must be positive and != 0"

// UsedArgsIncludeInvalid ...
// 	Returns true if one of arguments is:
// 	- negative
// 	- null
func UsedArgsIncludeInvalid(args []float64) (b bool) {
	for _, arg := range args {
		if arg == 0 {
			return true
		}
		if math.Signbit(arg) {
			return true
		}
	}
	return
}
