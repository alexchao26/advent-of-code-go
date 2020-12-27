package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	// parse input file into a slice of numbers
	input := util.ReadFile("../input.txt")
	characters := strings.Split(input, "")
	digits := make([]int, len(characters))
	for i, v := range characters {
		digits[i], _ = strconv.Atoi(v)
	}

	// generate all patterns once at start
	patterns := make([][]int, len(digits))
	for i := range digits {
		patterns[i] = generatePattern(i+1, len(digits))
	}

	// run through 100 phases, overwriting digits
	for i := 0; i < 100; i++ {
		digits = getNextOutputNumber(digits, patterns)
	}

	// Transform into github.com/alexchao26/advent-of-code-go output
	var firstEightDigits int
	for i := 0; i < 8; i++ {
		firstEightDigits *= 10
		firstEightDigits += digits[i]
	}
	fmt.Printf("First 8 digits after 100 phases: %v\n", firstEightDigits)
}

// generatePattern takes in the index (one indexed, not zero) that is being considered, and returns the pattern to multiply
// digits by
func generatePattern(oneIndex, lengthNeeded int) []int {
	if oneIndex < 1 {
		log.Fatal("Input to generatePattern must be a positive int")
	}
	basePattern := []int{0, 1, 0, -1}
	pattern := make([]int, 0, 4*(oneIndex+1))
	for len(pattern)-1 < lengthNeeded {
		for _, v := range basePattern {
			for i := 0; i < oneIndex; i++ {
				pattern = append(pattern, v)
			}
		}
	}
	return pattern[1 : lengthNeeded+1]
}

// takes in digits and all patterns, generates next set of digits
func getNextOutputNumber(digits []int, patterns [][]int) []int {
	output := make([]int, len(digits))
	for index := range digits {
		// for this index, sum up all products of digits and this index's pattern
		var sum int
		for i := 0; i < len(digits); i++ {
			sum += patterns[index][i] * digits[i]
		}

		// ensure the sum is positive & take the single's place digit
		if sum < 0 {
			sum *= -1
		}
		sum %= 10

		// assign to output slice for the same index
		output[index] = sum
	}

	return output
}
