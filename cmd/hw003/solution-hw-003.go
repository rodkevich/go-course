package main

import (
	"fmt"
	"github.com/rodkevich/go-course/homework/hw003"
	"strconv"
)

func main() {

	// PART 1: Arrays (average):
	var input = []float64{1, 2, 3, 4, 5, 6}

	avg := hw003.CountAverageOfArray(input)
	avgFormatted := strconv.FormatFloat(avg, 'f', -1, 64)

	fmt.Printf("arrayshw | result | Avg()=%v", avgFormatted)

	// PART 2: Slices - 1 (longest string in slice):
	var sliceCases = [][]string{{"one", "two", "three"}, {"one", "two"}}

	for _, each := range sliceCases {
		max, err := hw003.LongestStrInSlice(each)
		if err != nil {
			fmt.Printf("\nsliceshw | error | LongestStrInSlice() Output: %v", err)
		}
		fmt.Printf("\nsliceshw | result | LongestStrInSlice() val=%v", max)
	}

	// PART 2: Slices - 2 (reversed ints):
	inputSlice := []int64{1, 2, 5, 15}
	res := hw003.ReverseSliceOfInts(inputSlice)

	fmt.Printf("\nsliceshw | result | ReverseSliceOfInts() val=%v", res)

	// PART 3: Maps (sorting):
	var mapCases = []struct{ input map[int]string }{
		{map[int]string{2: "a", 0: "b", 1: "c"}},
		{map[int]string{10: "aa", 0: "bb", 500: "cc"}},
	}
	for _, each := range mapCases {
		hw003.PrintValuesSortedByIncrKeys(each.input)
	}
}
