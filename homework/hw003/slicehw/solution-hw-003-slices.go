package slicehw

import (
	"errors"
)

type results struct {
	list []string
	len  int // set len to be 0
}

// LongestStrInSlice returns the longest word from the slice of strings
func LongestStrInSlice(sliceOfStrings []string) (l int, s string, err error) {
	if len(sliceOfStrings) < 1 {
		return 0, "", errors.New("error: Check your input")
	}
	res := results{}

	// process depending of sliceOfStrings length
	switch len(sliceOfStrings) {
	case 0:
		return 0, "", err
	case 1:
		return len(sliceOfStrings[0]), sliceOfStrings[0], err
	default:
		// skip smaller values
		for _, value := range sliceOfStrings {
			if len(value) < res.len {
				continue
			}
			// clear smaller values in resulting obj
			if len(value) > res.len {
				res.len = len(value)
				res.list = res.list[:0]
			}
			res.list = append(res.list, value)
		}
		return res.len, res.list[0], err
	}
}

// ReverseSliceOfInts returns a reversed slice of ints
func ReverseSliceOfInts(intSlice []int64) []int64 {
	var rtn []int64
	ind := len(intSlice) // find len of slice which is max index

	for range intSlice {
		ind-- // decrease index starting from it's max value
		rtn = append(rtn, intSlice[ind])
	}

	return rtn
}
