package hw003

import (
	"errors"
	"unicode/utf8"
)

// LongestStrInSlice returns the longest word from the slice of strings
func LongestStrInSlice(input []string) (result string, err error) {

	if len(input) < 1 {
		return result, errors.New("error: Check your input")
	}
	// process depending of input length
	switch len(input) {
	case 0:
		return
	case 1:
		result = input[0]
		return
	default:
		length := utf8.RuneCountInString
		var lengthMax int

		for _, element := range input {
			if length(element) > lengthMax {
				lengthMax = length(element)
				result = element
			}
		}
		return
	}
}

// ReverseSliceOfInts returns a reversed slice of ints
func ReverseSliceOfInts(intSlice []int64) []int64 {

	// var rtn []int64
	// ind := len(intSlice) // find length of slice which is max index

	// for _ = range intSlice {
	// 	ind-- // decrease index starting from it's max value
	// 	rtn = append(rtn, intSlice[ind])
	// }

	rtn := make([]int64, len(intSlice))
	copy(rtn, intSlice)

	for i := len(rtn)/2 - 1; i >= 0; i-- {
		c := len(rtn) - 1 - i
		rtn[i], rtn[c] = rtn[c], rtn[i]
	}

	return rtn
}
