/*
Only change from part1 is adding a decorator to the digit slices to shrink large
groups (more than 2 digits) into a single digit. Then pass that into the same
checks for a duplicate & all decreasing
*/

package main

import "fmt"

func main() {
	start, end := 138307, 654504
	possibleCombinations := 0

	for i := start; i <= end; i++ {
		digits := shrinkLargeGroups(makeDigitsSlice(i))
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

/*
shrinkLargeGroups will return a new slice with any groups larger than 2 shrinked
down to 1. e.g. 111223 -> 1223
*/
func shrinkLargeGroups(digits []int) []int {
	// from start of number & ensure i+2 is within bounds of the digits slice
	for i := 0; i+2 < len(digits); i++ {
		if digits[i] == digits[i+1] && digits[i] == digits[i+2] {
			// figure out how many items to remove
			removeUpTo := i + 1
			for removeUpTo < len(digits) && digits[i] == digits[removeUpTo] {
				removeUpTo++
			}

			// copy the values into a new slice
			newSli := make([]int, 0)
			for j := 0; j <= i; j++ {
				newSli = append(newSli, digits[j])
			}
			for removeUpTo < len(digits) {
				newSli = append(newSli, digits[removeUpTo])
				removeUpTo++
			}
			digits = newSli
		}
	}
	return digits
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
