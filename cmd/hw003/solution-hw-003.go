package main

import (
	"fmt"
	"strconv"

	"github.com/rodkevich/go-course/homework/hw003/arrayhw"
	"github.com/rodkevich/go-course/homework/hw003/maphw"
	"github.com/rodkevich/go-course/homework/hw003/slicehw"
)

func main() {

	// PART 1: Arrays (average):
	var input = []float64{1, 2, 3, 4, 5, 6}

	// переделал на преобразование в array внутри функции потому что
	// не знаю как отвечать на заданный вопрос почему input ограничен до [6]float64  =)
	// по заданию нужно работать с array a не slice

	arr, l, avg := arrayhw.CountAverageOfArray(input)
	value := strconv.FormatFloat(avg, 'f', -1, 64)

	fmt.Printf("arrayshw | result | array=%v len=%v, Avg()=%v", arr, l, value)

	// PART 2: Slices - 1 (longest string in slice):
	var sliceCases = []struct{ input []string }{
		{[]string{"one", "two", "three"}},
		{[]string{"one", "two"}},
	}
	for _, each := range sliceCases {
		l, max, err := slicehw.LongestStrInSlice(each.input)
		if err != nil {
			fmt.Printf("\nsliceshw | error | LongestStrInSlice() Output: %v", err)
		}
		fmt.Printf("\nsliceshw | result | LongestStrInSlice() length=%v, val=%v", l, max)
	}

	// PART 2: Slices - 2 (reversed ints):
	inputSlice := []int64{1, 2, 5, 15}
	res := slicehw.ReverseSliceOfInts(inputSlice)

	fmt.Printf("\nsliceshw | result | ReverseSliceOfInts() val=%v", res)

	// PART 3: Maps (sorting):
	var mapCases = []struct{ input map[int]string }{
		{map[int]string{2: "a", 0: "b", 1: "c"}},
		{map[int]string{10: "aa", 0: "bb", 500: "cc"}},
	}
	for _, each := range mapCases {
		maphw.PrintValuesSortedByIncrKeys(each.input)
	}
}
