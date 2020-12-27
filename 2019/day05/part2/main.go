package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	// read the input file, modify it to a slice of numbers
	inputFile := util.ReadFile("../input.txt")
	// inputFile := "3,3,1107,-1,8,3,4,3,99"

	splitStrings := strings.Split(inputFile, ",")

	inputNumbers := make([]int, len(splitStrings))
	for i, v := range splitStrings {
		inputNumbers[i], _ = strconv.Atoi(v)
	}

	// system ID is 5, that is the input to the intcode computer
	input := 5
	step(inputNumbers, 0, input)
}

// step will read the next 4 values in the input `sli` and make updates
// according to the opcodes
func step(sli []int, index, input int) int {
	// read the instruction, opcode and the indexes where the params point to
	opcode, paramIndexes := getOpCodeAndIndexes(sli, index)

	var output int
	switch opcode {
	case 99: // 99: Terminates program
		fmt.Println("Terminating...")
		return input
	case 1: // 1: Add next two paramIndexes, store in third
		sli[paramIndexes[2]] = sli[paramIndexes[0]] + sli[paramIndexes[1]]
		return step(sli, index+4, output)
	case 2: // 2: Multiply next two and store in third
		sli[paramIndexes[2]] = sli[paramIndexes[0]] * sli[paramIndexes[1]]
		return step(sli, index+4, output)
	case 3: // 3: Takes one input and saves it to position of one parameter
		sli[paramIndexes[0]] = input
		return step(sli, index+2, output)
	case 4: // 4: outputs its input value
		output = sli[paramIndexes[0]]
		fmt.Printf("Opcode 4 output: %v\n", output)
		return step(sli, index+2, output)
	// 5: jump-if-true: if first param != 0, move pointer to second param, else nothing
	case 5:
		if sli[paramIndexes[0]] != 0 {
			return step(sli, sli[paramIndexes[1]], output)
		}
		return step(sli, index+3, output)
	// 6: jump-if-false, if first param == 0 then set instruction pointer to 2nd param, else nothing
	case 6:
		if sli[paramIndexes[0]] == 0 {
			return step(sli, sli[paramIndexes[1]], output)
		}
		return step(sli, index+3, output)
	// 7: less-than, if param1 < param2 then store 1 in postion of 3rd param, else store 0
	case 7:
		if sli[paramIndexes[0]] < sli[paramIndexes[1]] {
			sli[paramIndexes[2]] = 1
		} else {
			sli[paramIndexes[2]] = 0
		}
		return step(sli, index+4, output)
	// 8: equals, if param1 == param2 then set position of 3rd param to 1, else store 0
	case 8:
		if sli[paramIndexes[0]] == sli[paramIndexes[1]] {
			sli[paramIndexes[2]] = 1
		} else {
			sli[paramIndexes[2]] = 0
		}
		return step(sli, index+4, output)
	default:
		log.Fatal("Error: unknown opcode: ", opcode)
	}
	// this should never be called b/c switch statement will always return
	return -1
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
