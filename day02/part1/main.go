package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// read the input file and place into slice of strings
	inputFile := util.ReadFile("../input.txt")
	splitStrings := strings.Split(inputFile, ", ")

	// convert to slice of numbers
	inputNumbers := make([]int, len(splitStrings))
	for i, v := range splitStrings {
		inputNumbers[i], _ = strconv.Atoi(v)
	}

	// start running the step function
	step(inputNumbers, 0)
	fmt.Println("Final value at address 0 is:", inputNumbers[0])
}

// step will read the next 4 values in the input `sli` and make updates according to the opcodes
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

// this read function may be necessary later as the intcode thingy becomes more complex
func read(sli []int, index int) (int, int, int, int) {
	return sli[index], sli[index+1], sli[index+2], sli[index+3]
}
