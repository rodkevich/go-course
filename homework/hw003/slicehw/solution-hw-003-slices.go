package slicehw

import (
	"errors"
	"unicode/utf8"
)

type results struct {
	list   []string
	length int // set length to be 0
}

// LongestStrInSlice returns the longest word from the slice of strings
func LongestStrInSlice(sliceOfStrings []string) (length int, s string, err error) {
	if len(sliceOfStrings) < 1 {
		return 0, "", errors.New("error: Check your input")
	}
	result := results{}

	// process depending of sliceOfStrings length
	switch len(sliceOfStrings) {
	case 0:
		return 0, "", err
	case 1:
		return len(sliceOfStrings[0]), sliceOfStrings[0], err
	default:
		// count length in symbols not bytes
		length := utf8.RuneCountInString

		for _, element := range sliceOfStrings {
			// skip smaller values
			if length(element) < result.length {
				continue
			}
			// clear smaller values in resulting obj
			if length(element) > result.length {
				result.length = length(element)
				result.list = result.list[:0]
			}
			result.list = append(result.list, element)
		}

		return result.length, result.list[0], err
	}
}

// ReverseSliceOfInts returns a reversed slice of ints
func ReverseSliceOfInts(intSlice []int64) []int64 {
	var rtn []int64
	ind := len(intSlice) // find length of slice which is max index

	for range intSlice {
		ind-- // decrease index starting from it's max value
		rtn = append(rtn, intSlice[ind])
	}

	return rtn
}
