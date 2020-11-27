package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	// read the input file, modify it to a slice of numbers
	inputFile := util.ReadFile("../input.txt")
	splitStrings := strings.Split(inputFile, ",")

	inputNumbers := make([]int, len(splitStrings))
	for i, v := range splitStrings {
		inputNumbers[i], _ = strconv.Atoi(v)
	}

	input := 1
	step(inputNumbers, 0, input)
}

// step will read the next 4 values in the input `sli` and make updates
// according to the opcodes
func step(sli []int, index, input int) int {
	// read the instruction, opcode and the indexes where the params point to
	opcode, paramIndexes := getOpCodeAndIndexes(sli, index)

	var output, jump int
	switch opcode {
	// 99: Terminates program
	case 99:
		fmt.Println("Terminating")
		return -1
	case 1: // 1: Add next two paramIndexes, store in third
		sli[paramIndexes[2]] = sli[paramIndexes[0]] + sli[paramIndexes[1]]
		jump = 4
	case 2: // 2: Multiply next two and store in third
		sli[paramIndexes[2]] = sli[paramIndexes[0]] * sli[paramIndexes[1]]
		jump = 4
	case 3: // 3: Takes one input and saves it to position of one parameter
		sli[paramIndexes[0]] = input
		jump = 2
	case 4: // 4: outputs its input value
		output = sli[paramIndexes[0]]
		fmt.Printf("Opcode 4 output: %v\n", output)
		jump = 2
	default:
		log.Fatal("Error: unknown opcode: ", opcode)
	}
	// recursively call itself & jump index value...
	return step(sli, index+jump, output)
}

/*
getOpCodeAndIndexes will parse the instruction at sli[index]
- opcode will be the left two digits, mod by 100 will get that
- rest of instructions will be grabbed via mod 10
	- these also have to be parsed for the
*/
func getOpCodeAndIndexes(sli []int, index int) (int, [3]int) {
	instruction := sli[index]

	// opcode is the lowest two digits, so mod by 100
	opcode := instruction % 100
	instruction /= 100

	// assign the indexes that need to be read by reading the parameter modes
	var paramIndexes [3]int
	for i := 1; i <= 3 && index+i < len(sli); i++ {
		// grab the mode with a mod, last digit
		mode := instruction % 10
		instruction /= 10

		switch mode {
		case 1: // immediate mode, the index itself
			paramIndexes[i-1] = index + i
		case 0: // position mode, index will be the value at the index
			paramIndexes[i-1] = sli[index+i]
		}
	}

	return opcode, paramIndexes
}
