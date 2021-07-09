package fibo

import (
	"errors"
	"fmt"
	"math"
)

const (
	limit64bitSeqSize int64 = 93
)

// SizedSequence64bit returns a new slice of fibonacci numbers with required length if possible. Breaks on 64 bit overflow
func SizedSequence64bit(size int64) ([]int64, error) {
	defer fmt.Println("Program stopped")
	// set vars for use in function scope
	var (
		a, b int64 = 0, 1
		rtn        = []int64{a, b}
		err        = errors.New("\nmax 64bit `size` is 93. exit app because wrong input")
	)
	// check if overflow can happen
	if size > limit64bitSeqSize {
		// TODO: почитать про nil (возвращается []int64. Хак?)
		return nil, err
	}
	for i := int64(2); i <= size; i++ {
		var c int64 = a + b
		if c > math.MaxInt64 {
			size = i
			break
		} else {
			a, b = b, c
			rtn = append(rtn, c)
		}
	}
	return rtn[0:size], nil
}

// Cache ...
type Cache map[int]int

// NewCache ...
func (c *Cache) NewCache() *Cache {
	rtn := &Cache{0: 1, 1: 1}
	return rtn
}

// Experimental ...
type Experimental interface {
	NewCache() *Cache
}
