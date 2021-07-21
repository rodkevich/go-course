package test

// Нема больше тестов бо нет времени переделывать

// import (
// 	"github.com/rodkevich/go-course/homework/hw003/arrayhw"
// 	"github.com/rodkevich/go-course/homework/hw003/slicehw"
// 	"testing"
// )
//
// func TestPrintDetailedAverageForArrImp(t *testing.T) {
// 	var (
// 		sumInputNumbers float64
// 		expectedAverage = 3.5
// 		in              = arrayhw.Input{1, 2, 3, 4, 5, 6}
// 	)
//
// 	array := arrayhw.ArrayWithFunction{
// 		Input: &in,
// 		Avg: func() float64 {
// 			for _, ent := range &in {
// 				sumInputNumbers += ent
// 			}
// 			rtn := sumInputNumbers / float64(len(&in))
// 			return rtn
// 		},
// 	}
// 	average := array.Avg()
// 	array.printDetailedAverageForArrImp(average)
//
// 	if average != expectedAverage {
// 		t.Errorf("Error in comapiring. array = %v; want= %v", average, expectedAverage)
// 	}
// }
//
// func TestSolutionForStringSlices(t *testing.T) {
// 	cases := []struct {
// 		input     []string
// 		got, want string
// 	}{
// 		{input: []string{"one", "two", "three"}, want: "three"},
// 		{input: []string{"one", "two"}, want: "one"},
// 		{input: []string{"one"}, want: "one"},
// 		{input: []string{"a"}, want: "a"},
// 		{input: []string{}, want: ""},
// 	}
//
// 	for _, c := range cases {
// 		_, got, _ := slicehw.LongestStrInSlice(c.input)
// 		if got != c.want {
// 			t.Errorf("LongestStrInSlice(%q) == %q, want %q", c.input, got, c.want)
// 		}
// 	}
// }
//
// func TestReverseIntSlice(t *testing.T) {
// 	in := []int64{1, 2, 5, 15}
// 	want := []int64{15, 5, 2, 1}
// 	res := slicehw.ReverseSliceOfInts(in)
// 	for ind, i := range res {
// 		if i != want[ind] {
// 			t.Errorf("Error in comparing. slice = %v; want= %v", res, want)
// 		}
// 	}
// }
//
// // Does not work yet ;)
//
// // func TestPrintValuesSortedByIncrKeys(t *testing.T) {
// //
// // 	cases := []struct {
// // 		input     map[int]string
// // 		got, want []string
// // 	}{
// // 		// case 1
// // 		{input:
// // 		map[int]string{10: "aa", 0: "bb", 500: "cc",
// // 		}, want:
// // 		[]string{"bb", "aa", "cc"}},
// // 		// case 2
// // 		{input:
// // 		map[int]string{2: "a", 0: "b", 1: "c",
// // 		}, want:
// // 		[]string{"b", "c", "a"}},
// // 	}
// // 	for _, c := range cases {
// // 		assert := assert.New(t)
// // 		got, _ := maphw.PrintValuesSortedByIncrKeys(c.input)
// // 		assert.ElementsMatch(t, c.want, got)
// // 	}
// // }
