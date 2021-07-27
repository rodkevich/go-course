package helpers

import "math"

// UsingInvalidArgs ...
// 	to use if args u get do not satisfy requirements
const UsingInvalidArgs = "error: vars must be positive and != 0"

// UsedArgsIncludeInvalid ...
// 	Returns true if one of arguments is:
// 	- negative
// 	- null
func UsedArgsIncludeInvalid(args []float64) (b bool) {
	// check for empty slice
	if len(args) == 0 {
		return true
	}
	for _, arg := range args {
		// check for 0
		if arg == 0 {
			return true
		}
		// check for any negative
		if math.Signbit(arg) {
			return true
		}
	}
	return
}
