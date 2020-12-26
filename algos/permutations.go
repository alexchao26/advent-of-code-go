package algos

import "strings"

// PermuteIntSlice will make all permutations of the numbers input
func PermuteIntSlice(numbers []int) [][]int {
	return recurseIntSlice(numbers, 0)
}

// helper function to generate permutations
func recurseIntSlice(numbers []int, startIndex int) [][]int {
	if startIndex == len(numbers) {
		// makes a copy using append
		return [][]int{append([]int{}, numbers...)}
	}

	var perms [][]int
	for i := startIndex; i < len(numbers); i++ {
		// swap, append perms, backtrack
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]
		perms = append(perms, recurseIntSlice(numbers, startIndex+1)...)
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]
	}
	return perms
}

// PermuteString generates all permutations for a given string
func PermuteString(str string) []string {
	return recurseString(strings.Split(str, ""), 0)
}

func recurseString(sli []string, index int) []string {
	if index == len(sli) {
		return []string{strings.Join(sli, "")}
	}

	var perms []string
	for i := index; i < len(sli); i++ {
		sli[i], sli[index] = sli[index], sli[i]
		perms = append(perms, recurseString(sli, index+1)...)
		sli[i], sli[index] = sli[index], sli[i]
	}
	return perms
}

// PermuteStringSlice will make all permutations of a string slice
func PermuteStringSlice(in []string) [][]string {
	return recurseStringsSlice(in, 0)
}

// helper function to generate permutations
func recurseStringsSlice(in []string, startIndex int) [][]string {
	if startIndex == len(in) {
		// makes a copy using append
		return [][]string{append([]string{}, in...)}
	}

	var perms [][]string
	for i := startIndex; i < len(in); i++ {
		// swap, append perms, backtrack
		in[startIndex], in[i] = in[i], in[startIndex]
		perms = append(perms, recurseStringsSlice(in, startIndex+1)...)
		in[startIndex], in[i] = in[i], in[startIndex]
	}
	return perms
}
