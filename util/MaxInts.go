package util

// MaxInts takes a variable number of integers and returns the largest one
func MaxInts(numbers ...int) int {
	floats := make([]float64, len(numbers))
	for i, num := range numbers {
		floats[i] = float64(num)
	}

	maxNum := floats[0]
	for i := range floats {
		if floats[i] > maxNum {
			maxNum = floats[i]
		}
	}

	return int(maxNum)
}
