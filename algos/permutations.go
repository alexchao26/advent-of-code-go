package algos

import "strings"

// MakeIntPermutations will make all permutations of the numbers input
func MakeIntPermutations(numbers []int) [][]int {
	return recurseInts(numbers, 0)
}

// helper function to generate permutations
func recurseInts(numbers []int, startIndex int) [][]int {
	if startIndex == len(numbers) {
		// makes a copy using append
		return [][]int{append([]int{}, numbers...)}
	}

	var perms [][]int
	for i := startIndex; i < len(numbers); i++ {
		// swap, append perms, backtrack
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]
		perms = append(perms, recurseInts(numbers, startIndex+1)...)
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]
	}
	return perms
}

// MakeStringPermutations generates all permutations for a given string
func MakeStringPermutations(str string) []string {
	return recurseStrings(strings.Split(str, ""), 0)
}

func recurseStrings(sli []string, index int) []string {
	if index == len(sli) {
		return []string{strings.Join(sli, "")}
	}

	var perms []string
	for i := index; i < len(sli); i++ {
		sli[i], sli[index] = sli[index], sli[i]
		perms = append(perms, recurseStrings(sli, index+1)...)
		sli[i], sli[index] = sli[index], sli[i]
	}
	return perms
}
