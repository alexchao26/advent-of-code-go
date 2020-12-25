package algos

import "strings"

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

// MakeStringPermutations generates all permutations for a given string
func MakeStringPermutations(str string) []string {
	return recurse(strings.Split(str, ""), 0)
}

func recurse(sli []string, index int) []string {
	if index == len(sli) {
		return []string{strings.Join(sli, "")}
	}

	var perms []string
	for i := index; i < len(sli); i++ {
		sli[i], sli[index] = sli[index], sli[i]
		perms = append(perms, recurse(sli, index+1)...)
		sli[i], sli[index] = sli[index], sli[i]
	}
	return perms
}
