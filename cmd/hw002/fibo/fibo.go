package fibo

import "fmt"

//SizedSequence64bit returns Fibonacci array of required length
func SizedSequence64bit(size int64) ([]int64, error) {
	var (
		counted int64
		a       int64
		b       int64 = 1
		rtn           = []int64{a, b} // init with first two
	)
	for i := int64(2); i <= size; i++ {
		counted = a + b
		if counted < 0 {
			fmt.Printf("! Warn: 64 bit overflow on %d iteration \n", i+1)
			size = i
			break
		} else {
			//prepare vars for new iteration
			a, b = b, counted
			rtn = append(rtn, counted)
		}
	}
	fmt.Println("> Required sequence part counted")
	return rtn[0:size], nil
}
