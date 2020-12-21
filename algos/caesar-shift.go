package algos

// CaesarShift performs a forwards caesar shift of a given shiftAmount
func CaesarShift(in string, shiftAmount int) string {
	var result string
	for _, char := range in {
		// char to ascii number
		ascii := int(char) - int('a')
		ascii += shiftAmount
		ascii %= 26
		ascii += int('a')
		result += string(rune(ascii))
	}
	return result
}
