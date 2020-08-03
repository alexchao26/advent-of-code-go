package util

// MakePermutations will make all permutations of the numbers input
// returns a pointer to avoid copying a large number of permutations
func MakePermutations(numbers []int) *[][]int {
	result := make([][]int, 0)

	swapRecurseBacktrack(numbers, 0, &result)

	return &result
}

// helper function to generate permutations
func swapRecurseBacktrack(numbers []int, startIndex int, results *[][]int) {
	if startIndex == len(numbers) {
		// make a copy of the perm
		perm := make([]int, len(numbers))
		copy(perm, numbers)

		// assign the value at the pointer results to the appended slice (dereferenced) results w/ perm
		*results = append(*results, perm)
	}

	for i := startIndex; i < len(numbers); i++ {
		// swap numbers
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]

		// recurse with startIndex incremented
		swapRecurseBacktrack(numbers, startIndex+1, results)

		// backtrack
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]
	}
}
