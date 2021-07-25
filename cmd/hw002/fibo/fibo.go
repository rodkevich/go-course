package fibo

import (
	"errors"
	"fmt"
	"log"
)

const (
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
		a   int64
		b   int64 = 1
		rtn       = make([]int64, 0, size)
	)
	rtn = append(rtn, a, b) // init with first two
	for i := int64(2); i <= size; i++ {
		var counted = a + b
		if counted < 0 {
			_ = log.Output(1, "fibo: error: 64 bit overflow")
			return nil, errors.New(err64bitOverflow)
		}
		// prepare vars for iteration
		a, b = b, counted
		rtn = append(rtn, counted)
	}
	fmt.Println("> Required sequence part counted")

	return rtn[0:size], nil
}
