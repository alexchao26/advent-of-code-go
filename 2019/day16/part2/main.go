package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alexchao26/advent-of-code-go/util"
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
		// output a time to make sure this is running fast enough
		fmt.Printf("output received at %v, %v to go\n", time.Now(), 100-i-1)
	}

	// Transform into github.com/alexchao26/advent-of-code-go output
	var firstEightDigits int
	for i := 0; i < 8; i++ {
		firstEightDigits *= 10
		firstEightDigits += digits[i+offsetIndex]
	}
	fmt.Printf("\nOffset 8 digits after 100 phases: %v\n", firstEightDigits)
	fmt.Println("Expect 84462026 for test, 36265589 for actual input")
}

// takes in digits and all patterns, generates next set of digits
func getNextOutputNumber(digits []int) []int {
	// ONE INDEX THE PARTIAL SUMS, so partials[0] = 0, partials[1] = digits[0]
	partials := make([]int, len(digits)+1)
	for i := range digits {
		// add previous partial subset if i is not zero
		partials[i+1] += partials[i]
		// add digit on as well
		partials[i+1] += digits[i]
	}

	output := make([]int, len(digits))
	for i := range digits {
		chunkLength := i + 1

		adding := true
		for start := i; start < len(digits); start += chunkLength * 2 {
			// calculate chunk sum
			startOfChunkIndex := chunkLength - 1
			endOfChunkIndex := start + chunkLength - 1
			if endOfChunkIndex >= len(digits) {
				endOfChunkIndex = len(digits) - 1
			}
			chunkSum := partials[endOfChunkIndex] - partials[startOfChunkIndex]

			// increment or decrements output index
			if adding {
				output[i] += chunkSum
			} else {
				output[i] -= chunkSum
			}
			adding = !adding
		}
	}

	// make all output digits positive & single digits
	// NOTE this is not equivalent to taking the mod of a negative number!
	// This has to be done because of the problem's spec to just take the 1's digit
	for i, v := range output {
		if v < 0 {
			output[i] *= -1
		}
		output[i] %= 10
	}

	return output
}
