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
	// input := "12345678"
	// input := "03036732577212944063491565474664"
	characters := strings.Split(input, "")

	// digits := make([]int, len(characters))
	digits := make([]int, len(characters)*10000)
	for i := 0; i < 10000; i++ {
		for j, v := range characters {
			digits[i*len(characters)+j], _ = strconv.Atoi(v)
			// digits[j], _ = strconv.Atoi(v)
		}
	}

	var offsetIndex int
	for i := 0; i < 7; i++ {
		offsetIndex *= 10
		offsetIndex += digits[i]
	}

	// fmt.Println("offsetIndex", offsetIndex)

	// run through 100 phases, overwriting digits
	for i := 0; i < 100; i++ {
		digits = getNextOutputNumber(digits)
		fmt.Printf("output received at %v, %v to go\n", time.Now(), 100-i-1)
	}

	// Transform into AoC output
	var firstEightDigits int
	for i := 0; i < 8; i++ {
		firstEightDigits *= 10
		// firstEightDigits += digits[i]
		firstEightDigits += digits[i+offsetIndex]
	}
	fmt.Printf("First 8 digits after 100 phases: %v\n", firstEightDigits)
}

// takes in digits and all patterns, generates next set of digits
func getNextOutputNumber(digits []int) []int {
	output := make([]int, len(digits))
	// bottom up approach?

	var positiveBit int

	// the easy (back) part of the next digit
	for index := len(digits) - 1; index*2+1 >= len(digits); index-- {
		positiveBit += digits[index]
		positiveBit %= 10
		output[index] = positiveBit
	}

	for index := 0; index <= (len(digits)-1)/2; index++ {
		var sum int
		for jump := index; jump < len(digits); {
			for i := 0; i < index+1 && jump+i < len(digits); i++ {
				sum += digits[jump+i]
			}
			jump += (index + 1) * 2
			for i := 0; i < index+1 && jump+i < len(digits); i++ {
				sum -= digits[jump+i]
			}
			jump += (index + 1) * 2
		}
		if sum < 0 {
			sum *= -1
		}
		sum %= 10
		output[index] = sum
	}

	return output
}

// lower for loop
// positiveBit += digits[index]
// if index*2+1 < len(digits) {
// 	positiveBit -= digits[index*2+1]
// }
// if index*2+2 < len(digits) {
// 	positiveBit -= digits[index*2+2]
// }

// fmt.Println(positiveBit, "below, starting")
// sum := positiveBit

// jumpIndex := index*3 + 2
// for jumpIndex < len(digits) {
// 	for i := 0; i < index+1 && jumpIndex+i < len(digits); i++ {
// 		sum -= digits[jumpIndex+i]
// 	}
// 	jumpIndex += (index + 1) * 2
// 	for i := 0; i < index+1 && jumpIndex+i < len(digits); i++ {
// 		sum += digits[jumpIndex+i]
// 	}
// 	jumpIndex += (index + 1) * 2
// 	sum %= 10
// }
// if sum < 0 {
// 	sum *= -1
// }
// sum %= 10
// output[index] = sum
