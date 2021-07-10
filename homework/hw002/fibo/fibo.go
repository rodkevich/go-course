package fibo

import (
	"errors"
	"fmt"
	"math"
)

//errors
var (
	errorWrongUserInput error = errors.New(
		"\nmax 64bit `size` is 93. \nExiting app because wrong input was provided",
	)
)

//limits
const (
	limitOf64bitSeq int64 = 93
	maxInt64        int64 = math.MaxInt64
)

// SizedSequence64bit returns a new slice of fibonacci numbers
//with required length if possible. Breaks on 64 bit overflow
func SizedSequence64bit(size int64) ([]int64, error) {
	defer fmt.Println("\nCounting fresh sequence part finished")

	var (
		counted int64
		a       int64   = 0
		b       int64   = 1
		rtn     []int64 = []int64{a, b} // init with first two
	)

	// check input arg to be expected.if not -> return
	if size > limitOf64bitSeq {
		// TODO: почитать про nil
		// TODO: почитать type assertion
		return nil, errorWrongUserInput
	}
	for i := int64(2); i <= size; i++ {
		counted = a + b
		// check if max value reached

		if counted > maxInt64 {
			size = i
			break
		} else {
			//prepare vars for new iteration
			a, b = b, counted
			rtn = append(rtn, counted)
		}
	}

	return rtn[0:size], nil
}

//☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣
////////////////////  Experimental   ///////////////////////
////////////////////      part       ///////////////////////
//☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣

// Cache to play with custom types
type Cache []int64

// NewCache ...
func NewCache() Cache {
	return Cache{0: 0, 1: 1}
}

// SizedSequence64bitUsingCache uses cache or returns a new slice of fibonacci numbers
//with required length if possible. Breaks on 64 bit overflow
func SizedSequence64bitUsingCache(size int64, cache *Cache) ([]int64, error) {
	s64 := SizedSequence64bit
	if size > int64(len(*cache)) { // convert type to be comparable
		out, err := s64(size)
		if err != nil {
			return nil, err
		}
		*cache = out
		return out, err
	} else {
		fmt.Println("\nUsing cache")
		var cachedArray []int64
		for _, v := range *cache {
			cachedArray = append(cachedArray, v)
		}
		newObject := cachedArray[0:size]
		return newObject, nil
	}
}
