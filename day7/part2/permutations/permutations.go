package permutations

// CreatePermutations docz
func CreatePermutations(start, end int) [][]int {
	orig := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		orig[i-start] = i
	}
	results := make([][]int, 0)
	resPointer := &results
	perm(resPointer, orig, 0)

	// fmt.Println(len(results)) // 120 = 5!
	return results
}

func perm(resPointer *[][]int, slice []int, index int) {
	if index == len(slice) {
		*resPointer = append(*resPointer, slice)
		// fmt.Println("appending", resPointer)
	} else {
		temp := slice[index]
		for i := index; i < len(slice); i++ {
			newSlice := make([]int, len(slice))
			copy(newSlice, slice)
			// fmt.Println(newSlice, slice)
			newSlice[index] = newSlice[i]
			newSlice[i] = temp
			perm(resPointer, newSlice, index+1)
		}
	}
}

// CreatePermutations will return a 2D slice containing permutations from zero to the inputted int
// func CreatePermutations(size int) [][]int {
// 	result := make([][]int, 0)

// helper := func(digits []int, builder []int) {

// 	newDigits := make([]int, len(digits))
// 	copy(newDigits, digits)

// 	if len(digits) == 0 {
// 		// use the reference point in memory
// 		result = append(result, builder)
// 	} else {
// 		for index, digit := range newDigits {
// 			fmt.Println("append", append(newDigits[:index], newDigits[index+1:]...))
// 			fmt.Println("builder", append(builder, digit))
// 			helper(append(newDigits[:index], newDigits[index+1:]...), append(builder, digit))
// 			// fmt.Println("looping")
// 		}
// 	}

// }

// helper([]int{0, 1, 2, 3, 4}, make([]int, 0))

// 	return result
// }

// func nextPerm(p []int) {
// 	for i := len(p) - 1; i >= 0; i-- {
// 		if i == 0 || p[i] < len(p)-i-1 {
// 			p[i]++
// 			return
// 		}
// 		p[i] = 0
// 	}
// }

// func getPerm(orig, p []int) []int {
// 	result := append([]int{}, orig...)
// 	for i, v := range p {
// 		result[i], result[i+v] = result[i+v], result[i]
// 	}
// 	return result
// }

// CreatePermutations docz
// func CreatePermutations() [][]int {
// 	orig := []int{5, 6, 7, 8, 9}
// 	results := make([][]int, 0)

// 	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
// 		// fmt.Println(getPerm(orig, p))
// 		results = append(results, getPerm(orig, p))
// 	}
// 	return results
// }
