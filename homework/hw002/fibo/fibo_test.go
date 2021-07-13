package fibo

import (
	"fmt"
	"testing"
)

func TestSizedSequence64bit(t *testing.T) {
	var expected = []int64{0, 1, 1, 2, 3, 5, 8}
	var got, err = sizedSequence64bit(7)
	//show slices info
	printS(got)
	printS(expected)
	///////////////////////////////////////
	//example for comparing arrays
	a := [2]int64{1, 2}
	b := [2]int64{1, 2}
	if a != b {
		t.Errorf("Error in example. got = %d; want __", got)
	}
	///////////////////////////////////////
	//compare length
	if len(got) != len(expected) {
		t.Errorf("Error in lenght. got = %d; expected = %d", len(got), len(expected))
	}
	//compare values
	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Error in values. got =  %d; want __", got)
		}
	}

	if err != nil {
		t.Errorf("Error here err. err = %d; want __", err)
	}
}
func printS(s []int64) {
	fmt.Printf("type=%T, len=%d cap=%d %v\n", s, len(s), cap(s), s)
}
