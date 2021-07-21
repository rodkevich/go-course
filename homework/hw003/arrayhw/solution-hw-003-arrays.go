package arrayhw

// CountAverageOfArray function to count average from input using array
func CountAverageOfArray(input []float64) ([]float64, int, float64) {
	var (
		sum float64
		avg float64
	)
	arrayLength := len(input)
	array := make([]float64, arrayLength)

	for i := 0; i < arrayLength; i++ {
		array[i] = input[i]
	}

	for _, each := range array {
		sum += each
	}

	avg = sum / float64(arrayLength)

	return array, arrayLength, avg
}
