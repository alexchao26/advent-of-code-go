package algos

// CombinationsInts returns all combinations of an input slice of a given length
func CombinationsInts(nums []int, targetLength int) [][]int {
	if targetLength > len(nums) {
		panic("target length is greated than length of input slice")
	}

	var combos [][]int
	// loop over starting points in the nums slice
	for i := 0; i < len(nums); i++ {
		combos = append(combos, helperCombinationsInts(nums[i:], targetLength, []int{})...)
	}

	return combos
}

func helperCombinationsInts(nums []int, length int, current []int) [][]int {
	if len(current) == length {
		return [][]int{append([]int{}, current...)}
	}
	var combos [][]int
	for i := range nums {
		// add value onto the current combo
		current = append(current, nums[i])

		// recurse with only the remaining numbers, then append any valid combos
		// that were found
		recurseResult := helperCombinationsInts(nums[i+1:], length, current)
		combos = append(combos, recurseResult...)

		// backtrack
		current = current[:len(current)-1]
	}

	return combos
}
