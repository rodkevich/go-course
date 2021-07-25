package hw003

// CountAverageOfArray function to count average from input using array
func CountAverageOfArray(input [6]float64) float64 {
	var sum, avg float64
	for _, each := range input {
		sum += each
	}
	avg = sum / float64(len(input))

	return avg
}
