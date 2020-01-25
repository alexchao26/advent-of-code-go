package main

import "fmt"

func main() {
	perms := createPermutations(1, 3)
	fmt.Println(perms)
}

func createPermutations(start, end int) [][]int {
	orig := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		orig[i-start] = i
	}
	results := make([][]int, 0)

	// need to declare the function and all types in order to recurse w/ a closure
	// due to how go evaluates
	var recurse func(sliceCopy []int, start int)
	recurse = func(sliceCopy []int, start int) {
		if start == len(sliceCopy) {
			results = append(results, sliceCopy)
		} else {
			thisCopy := make([]int, len(sliceCopy))
			copy(thisCopy, sliceCopy)
			for i := start; i < len(sliceCopy); i++ {
				thisCopy[i], thisCopy[start] = thisCopy[start], thisCopy[i]
				recurse(thisCopy, start+1)
			}
		}
	}
	recurse(orig, 0)

	return results
}

// func perm(resPointer *[][]int, slice []int, index int) {
// 	if index == len(slice) {
// 		*resPointer = append(*resPointer, slice)
// 		// fmt.Println("appending", resPointer)
// 	} else {
// 		temp := slice[index]
// 		for i := index; i < len(slice); i++ {
// 			newSlice := make([]int, len(slice))
// 			copy(newSlice, slice)
// 			// fmt.Println(newSlice, slice)
// 			newSlice[index] = newSlice[i]
// 			newSlice[i] = temp
// 			perm(resPointer, newSlice, index+1)
// 		}
// 	}
// }
