package permutations

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

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

// CreatePermutations docz
func CreatePermutations() [][]int {
	orig := []int{0, 1, 2, 3, 4}
	results := make([][]int, 0)

	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		// fmt.Println(getPerm(orig, p))
		results = append(results, getPerm(orig, p))
	}
	return results
}
