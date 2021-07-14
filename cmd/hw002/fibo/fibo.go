package fibo

import (
	"errors"
	"fmt"
	"log"
)

var (
	errWrongArg      = "error: wrong arg passed"
	err64bitOverflow = "error: 64 bit overflowed : rtn size maximum equals 93"
)

// SizedSequence64bit returns Fibonacci array of required length
func SizedSequence64bit(size int64) ([]int64, error) {
	if size <= 0 {
		_ = log.Output(1, "fibo: wrong input")
		return nil, errors.New(errWrongArg)
	}
	var (
		counted int64
		a       int64
		b       int64 = 1
		rtn           = []int64{a, b} // init with first two
	)
	for i := int64(2); i <= size; i++ {
		counted = a + b
		if counted < 0 {
			_ = log.Output(1, "fibo: error: 64 bit overflow")
			return nil, errors.New(err64bitOverflow)
		}
		// prepare vars for new iteration
		a, b = b, counted
		rtn = append(rtn, counted)

	}
	fmt.Println("> Required sequence part counted")
	return rtn[0:size], nil
}
