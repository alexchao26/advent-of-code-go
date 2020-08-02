package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// read the input file, modify it to a slice of numbers
	inputFile := util.ReadFile("../input.txt")
	splitStrings := strings.Split(inputFile, ", ")
	inputNumbers := make([]int, len(splitStrings))

	for i, v := range splitStrings {
		inputNumbers[i], _ = strconv.Atoi(v)
	}

	// brute force to try all options for nouns and verbs
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// crete a copy of the inputNumbers slice
			clone := make([]int, 120)
			copy(clone, inputNumbers)
			clone[1] = i
			clone[2] = j

			// run step on the cloned slice
			step(clone, 0)
			// check if the zero address is equal to the AoC value
			if clone[0] == 19690720 {
				// print answers to console (manually add to advent of code)
				fmt.Println("noun is", i, "verb is", j)
				fmt.Println("actual result value noun * 10 + verb = ", i*100+j)

				// return to end main function
				return
			}
		}
	}
}

// step will read the next 4 values in the input `sli` and make updates
// according to the opcodes
func step(sli []int, index int) bool {
	if sli[index] == 99 {
		return false
	}
	opcode, two, three, four := read(sli, index)
	switch opcode {
	case 1:
		sli[four] = sli[two] + sli[three]
	case 2:
		sli[four] = sli[two] * sli[three]
	}
	// recursively call itself & increment index value...
	return step(sli, index+4)
}

func read(sli []int, index int) (int, int, int, int) {
	return sli[index], sli[index+1], sli[index+2], sli[index+3]
}
