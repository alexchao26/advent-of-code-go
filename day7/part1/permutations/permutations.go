package permutations

// CreatePermutations will return a 2D slice containing permutations from zero to the inputted int
func CreatePermutations(size int) []string {
	result := make([]string, 0)

	pointerToResult := &result
	onePerm(pointerToResult, "01234")

	return result
}

func onePerm(resultSlice *[]string, substring string) {
	// if len(*resultSlice) == 0 {
	// 	for i := 0; i < len(substring); i++ {
	// 		*resultSlice = append(*resultSlice, string(substring[i]))
	// 	}
	// 	return
	// }
	// for i := 0; i < len(*resultSlice); i++ {
	// 	newSubstring := *resultSlice[i]
	// }
	// modify the actual data @ the location in memory
	// fmt.Println(resultSlice, *resultSlice)

}
