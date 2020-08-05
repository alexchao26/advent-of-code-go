package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// parse input file into a slice of numbers
	input := util.ReadFile("../input.txt")
	// input := "03036732577212944063491565474664" // test should output "84462026"
	characters := strings.Split(input, "")

	digits := make([]int, len(characters)*10000)
	for i := 0; i < 10000; i++ {
		for j, v := range characters {
			digits[i*len(characters)+j], _ = strconv.Atoi(v)
		}
	}

	var offsetIndex int
	for i := 0; i < 7; i++ {
		offsetIndex *= 10
		offsetIndex += digits[i]
	}
	fmt.Println("offsetIndex", offsetIndex)

	// run through 100 phases, overwriting digits
	for i := 0; i < 100; i++ {
		digits = getNextOutputNumber(digits)
		fmt.Printf("output received at %v, %v to go\n", time.Now(), 100-i-1)
	}

	// Transform into AoC output
	var firstEightDigits int
	for i := 0; i < 8; i++ {
		firstEightDigits *= 10
		firstEightDigits += digits[i+offsetIndex]
	}
	fmt.Printf("Offset 8 digits after 100 phases: %v\n", firstEightDigits)
	// fmt.Println("Expect 84462026 for test")
}

// takes in digits and all patterns, generates next set of digits
func getNextOutputNumber(digits []int) []int {
	// calculate the sum of partial subsets digits[0:i]
	partials := make([]int, len(digits))
	for i := range digits {
		// add previous partial subset if i is not zero
		if i > 0 {
			partials[i] += partials[i-1]
		}
		// add digit on as well
		partials[i] += digits[i]
	}

	output := make([]int, len(digits))
	for i := range digits {
		chunkLength := i + 1

		adding := true
		for start := i; start < len(digits); start += chunkLength * 2 {
			// calculate chunk sum
			startOfChunkIndex := start - 1
			if startOfChunkIndex < 0 {
				startOfChunkIndex = 0
			}
			endOfChunkIndex := start + chunkLength - 1
			if endOfChunkIndex >= len(digits) {
				endOfChunkIndex = len(digits) - 1
			}
			chunkSum := partials[endOfChunkIndex] - partials[startOfChunkIndex]
			if chunkLength == 1 {
				chunkSum = digits[start]
			}

			if adding {
				// fmt.Printf("adding: %v\n", chunkSum)
				output[i] += chunkSum
			} else {
				// fmt.Printf("subtracting: %v\n", chunkSum)
				output[i] -= chunkSum
			}
			adding = !adding
		}
	}

	// make all output digits positive
	for i, v := range output {
		if v < 0 {
			output[i] *= -1
		}
		output[i] %= 10
	}

	return output
}
