package arrayhw

import (
	"fmt"
	"strconv"
)

// Input ...
type Input = [6]float64

// Average ...
type Average func() float64

// ArrayWithFunction ...
type ArrayWithFunction struct {
	Input *Input
	Avg   Average // to try adding function in place
}

// PrintDetailedAverageForArrImp try to add function to type in a pkg
func (arr *ArrayWithFunction) PrintDetailedAverageForArrImp(res float64) {
	// prettify value when printing
	value := strconv.FormatFloat(res, 'f', -1, 64)
	fmt.Printf("arrayshw | result | Avg() =  %v", value)
}
