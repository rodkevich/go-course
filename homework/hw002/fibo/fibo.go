package fibo

import (
	"errors"
	"fmt"
	"math"
)

//errors
var (
	errorLimitOf64bitSeq error = errors.New(
		"\nmax 64bit `size` is 93. \nXX Exiting app: wrong number was provided XX",
	)
)

//limits
const (
	limitOf64bitSeq int64 = 93
	maxInt64        int64 = math.MaxInt64
)

//Solution prints a sequence of fibonacci numbers up to required length
func Solution(n int) error {
	fmt.Println("Solution part started >>")
	defer fmt.Println("<< Solution part ended")
	var sequence, err = SizedSequence64bit(int64(n))
	if err != nil {
		return err
	}
	fmt.Println("Result:")
	fmt.Println(sequence)
	return nil
}

// SizedSequence64bit returns a new slice of fibonacci numbers
//with required length if possible. Breaks on 64 bit overflow
func SizedSequence64bit(size int64) ([]int64, error) {
	defer fmt.Println("\nCounted fresh sequence part")
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
		return nil, errorLimitOf64bitSeq
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

//////////////////  Experimental   /////////////////////
//////////////////      part       /////////////////////
//☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣

// Cache custom type to play with
type Cache []int64

// NewCache function returns prepared cache with first two numbers
func NewCache() Cache {
	return Cache{0: 0, 1: 1}
}

func SequenceWithCaching(n int64, cache *Cache) error {
	var sequence, err = SizedSequence64bitWithCache(n, cache)
	if err != nil {
		return err
	}
	fmt.Println("Result:")
	fmt.Println(sequence)
	return nil
}

// SizedSequence64bitWithCache uses cache or returns a new slice of fibonacci numbers
//with required length if possible. Breaks on 64 bit overflow
func SizedSequence64bitWithCache(size int64, cache *Cache) ([]int64, error) {
	if size > int64(len(*cache)) { // convert type to be comparable
		out, err := SizedSequence64bit(size)
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
