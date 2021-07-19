package main

import (
	"fmt"

	"github.com/rodkevich/go-course/homework/hw003/arrayhw"
	"github.com/rodkevich/go-course/homework/hw003/maphw"
	"github.com/rodkevich/go-course/homework/hw003/slicehw"
)

func main() {
	/* ---------------------------------------------------------------------------------
	Requirements - PART 1
	Arrays:
	Implement function that returns an average value of array (sum / N)
	*/

	// input -> [1,2,3,4,5,6]
	// output -> 3.5
	var (
		sum        float64
		out        float64
		inputArray = arrayhw.Input{1, 2, 3, 4, 5, 6}
	)

	var awf = arrayhw.ArrayWithFunction{

		Input: &inputArray,
		// implement function to count average value of inputArray
		Avg: func() float64 {
			for _, each := range inputArray {
				sum += each
			}
			out = sum / float64(len(inputArray))
			return out
		},
	}

	// call function added in place
	average := awf.Avg()

	// call function added to ArrayWithFunction type in its pkg
	awf.PrintDetailedAverageForArrImp(average)

	// Results:
	//
	//	arrayshw | result | Avg() =  3.5

	/* ---------------------------------------------------------------------------------
	Requirements  - PART 2
	Slices-1:
	•Implement max([]string) string function,
	that returns the longest word from the slice of strings
	(the first if there are more than one).
	*/

	// Input -> ("one", "two", "three")
	// Output -> ("three")

	len1, max1, err1 := slicehw.LongestStrInSlice([]string{
		"one",
		"two",
		"three",
	})
	if err1 != nil {
		fmt.Printf("\nsliceshw | error | LongestStrInSlice() Output: %v", err1)
	}

	fmt.Printf("\nsliceshw | result | LongestStrInSlice() len=%v, val=%v", len1, max1)

	// Input -> ("one", "two")
	// Output -> ("one")

	len2, max2, err2 := slicehw.LongestStrInSlice([]string{
		"one",
		"two",
	})
	if err2 != nil {
		fmt.Printf("\nsliceshw | error | LongestStrInSlice() Output: %v", err2)
	}

	fmt.Printf("\nsliceshw | result | LongestStrInSlice() len=%v, val=%v", len2, max2)

	/*  Requirements
	Slices-2:
	•Implement reverse([]int64) []int64 function,
	that returns the copy of the original slice inputSlice02 reverse order.
	The type of elements is int64.
	*/

	// Input -> (1, 2, 5, 15)
	// Output -> (15, 5, 2, 1)

	inputSlice02 := []int64{1, 2, 5, 15}
	res := slicehw.ReverseSliceOfInts(inputSlice02)

	fmt.Printf("\nsliceshw | result | ReverseSliceOfInts() val=%v", res)

	// Results:
	//
	// sliceshw | result | LongestStrInSlice() len=5, val=three
	//
	// sliceshw | result | LongestStrInSlice() len=3, val=one
	//
	// sliceshw | result | ReverseSliceOfInts() val=[15 5 2 1]

	/* ---------------------------------------------------------------------------------
	Requirements - PART 3
	Maps:
	•Implement printSorted(map[int]string) function,
	that prints map values sorted inputSlice02 order of increasing keys.
	*/

	// Input -> {2: "a", 0: "b", 1: "c" }
	// Output -> ["b", "c", "a"]

	inputMap01 := map[int]string{
		2: "a",
		0: "b",
		1: "c",
	}
	_, er1 := maphw.PrintValuesSortedByIncrKeys(inputMap01)
	if er1 != nil {
		_ = fmt.Errorf("err1")
	}

	// Input -> {10: "aa", 0: "bb", 500: "cc" }
	// Output -> ["bb", "aa", "cc"]

	inputMap02 := map[int]string{
		10:  "aa",
		0:   "bb",
		500: "cc",
	}
	_, er2 := maphw.PrintValuesSortedByIncrKeys(inputMap02)
	if er2 != nil {
		_ = fmt.Errorf("err2")
	}

	// Results:
	//
	// mapshw | result | PrintValuesSortedByIncrKeys() val=[b c a]
	//
	// mapshw | result | PrintValuesSortedByIncrKeys() val=[bb aa cc]

}
