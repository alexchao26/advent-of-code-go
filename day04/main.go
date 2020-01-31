package main

import (
	"fmt"
	"strconv"
)

func numMatching(sliceToEnd []rune) int {
	// fmt.Println(sliceToEnd)
	// gets in a slice from a character to the end, returns an int of the matching characters
	result := 1
	for i := 1; i < len(sliceToEnd); i++ {
		if sliceToEnd[i] == sliceToEnd[0] {
			result += 1
		} else {
			break
		}
	}
	// fmt.Println(result)
	return result
}

func testNumber(num int) bool {
	// cast to string to iterate through digits?
	strNum := strconv.Itoa(num)
	runesSlice := []rune(strNum)

	duplicate := false

	for i := 0; i < len(runesSlice)-1; {
		if runesSlice[i] > runesSlice[i+1] {
			return false
		}

		matches := numMatching(runesSlice[i:len(runesSlice)])
		if matches == 2 {
			duplicate = true
			i++
		} else if matches == 1 {
			i++
		} else if matches > 2 {
			i += (matches - 1)
		}
	}

	// if the entire for loop passes, return if there was a duplicate or not
	return duplicate
}

func main() {
	start, end := 138307, 654504
	possibleCombinations := 0

	for i := start; i <= end; i++ {
		if testNumber(i) {
			possibleCombinations++
		}
	}

	fmt.Println(possibleCombinations)
	// fmt.Println(testNumber(111122))
	// fmt.Println(testNumber(123444))
	// fmt.Println(testNumber(112233))
}
