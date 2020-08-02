package main

import (
	"fmt"
)

func main() {
	start, end := 138307, 654504
	possibleCombinations := 0

	for i := start; i <= end; i++ {
		digits := makeDigitsSlice(i)
		if isIncreasing(digits) && hasDuplicate(digits) {
			possibleCombinations++
		}
	}

	fmt.Println(possibleCombinations)
}

func makeDigitsSlice(num int) []int {
	result := make([]int, 6)
	for i := 5; num > 0; i-- {
		result[i] = num % 10
		num -= num % 10
		num /= 10
	}
	return result
}

func isIncreasing(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			return false
		}
	}
	return true
}

func hasDuplicate(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i-1] == digits[i] {
			return true
		}
	}
	return false
}
