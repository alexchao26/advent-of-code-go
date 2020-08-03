/*
Intcode struct is defined within this file
MakePermutations is in the util package as that will likely be reused
*/

package main

import (
	"adventofcode/util"
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

	// Make perms via a util function
	perms := util.MakePermutations([]int{0, 1, 2, 3, 4})

	// iterate over all perms and run through a single pass of the Amps
	// if the final output (from Amp E) is higher, update the highestOutput variable
	highestOutput := 0
	for _, perm := range *perms {
		// initialize 5 computers
		ampA := MakeComputer(inputNumbers, perm[0])
		ampB := MakeComputer(inputNumbers, perm[1])
		ampC := MakeComputer(inputNumbers, perm[2])
		ampD := MakeComputer(inputNumbers, perm[3])
		ampE := MakeComputer(inputNumbers, perm[4])

		// first input (besides phase setting) to Amp A is zero
		ampA.Step(0)
		ampB.Step(ampA.LastOutput)
		ampC.Step(ampB.LastOutput)
		ampD.Step(ampC.LastOutput)
		ampE.Step(ampD.LastOutput)

		if ampE.LastOutput > highestOutput {
			highestOutput = ampE.LastOutput
		}
	}

	// print highest output found
	fmt.Printf("Highest output is %v\n", highestOutput)
}

/*
Intcode is an OOP approach *************************************************
MakeComputer is equivalent to the constructor
Step takes in an input int and updates properties in the computer:
	- InstructionIndex: where to read the next instruction from
	- LastOutput, what the last opcode 4 outputted
	- PuzzleIndex based if the last instruction modified the puzzle at all
****************************************************************************/
type Intcode struct {
	PhaseSetting     int   // initial input: ID or number used to "prime"/setup the comp
	PuzzleInput      []int // file/puzzle input parsed into slice of ints
	InstructionIndex int   // stores the index where the next instruction is
	LastOutput       int   // last output from an opcode 4
}

// MakeComputer initializes a new comp
func MakeComputer(PuzzleInput []int, PhaseSetting int) Intcode {
	puzzleInputCopy := make([]int, len(PuzzleInput))
	copy(puzzleInputCopy, PuzzleInput)

	comp := Intcode{
		PhaseSetting,
		puzzleInputCopy,
		0,
		0,
	}

	// Prime the computer by running its initial phase setting through it
	// This will update the comp's InstructionIndex so it's pointing to the next command
	// will also update the PuzzleInput itself via opcode 3's insert
	// AND will run the computer until it asks for the next input, _comp is now primed_
	comp.Step(PhaseSetting)
	return comp
}

// Step will read the next 4 values in the input `sli` and make updates
// according to the opcodes
func (comp *Intcode) Step(input int) int {
	// read the instruction, opcode and the indexes where the params point to
	opcode, paramIndexes := comp.GetOpCodeAndIndexes()

	switch opcode {
	case 99: // 99: Terminates program
		// fmt.Println("Terminating...")
		return input
	case 1: // 1: Add next two paramIndexes, store in third
		comp.PuzzleInput[paramIndexes[2]] = comp.PuzzleInput[paramIndexes[0]] + comp.PuzzleInput[paramIndexes[1]]
		comp.InstructionIndex += 4
		return comp.Step(input)
	case 2: // 2: Multiply next two and store in third
		comp.PuzzleInput[paramIndexes[2]] = comp.PuzzleInput[paramIndexes[0]] * comp.PuzzleInput[paramIndexes[1]]
		comp.InstructionIndex += 4
		return comp.Step(input)
	case 3: // 3: Takes one input and saves it to position of one parameter
		// check if input has already been used (i.e. input == -1)
		// if it's been used, return the LastOutput
		// NOTE: making a big assumption that -1 will never be an input...
		if input == -1 {
			return comp.LastOutput
		}

		// otherwise use the input, then recurse with a -1 to signal the initial input has been used
		comp.PuzzleInput[paramIndexes[0]] = input
		comp.InstructionIndex += 2
		return comp.Step(-1)
	case 4: // 4: outputs its input value
		// set LastOutput of the computer & log it
		comp.LastOutput = comp.PuzzleInput[paramIndexes[0]]
		// fmt.Printf("Opcode 4 output: %v\n", comp.LastOutput)
		comp.InstructionIndex += 2

		// continue running until terminates or asks for another input
		return comp.Step(input)
	// 5: jump-if-true: if first param != 0, move pointer to second param, else nothing
	case 5:
		if comp.PuzzleInput[paramIndexes[0]] != 0 {
			comp.InstructionIndex = comp.PuzzleInput[paramIndexes[1]]
		} else {
			comp.InstructionIndex += 3
		}
		return comp.Step(input)
	// 6: jump-if-false, if first param == 0 then set instruction pointer to 2nd param, else nothing
	case 6:
		if comp.PuzzleInput[paramIndexes[0]] == 0 {
			comp.InstructionIndex = comp.PuzzleInput[paramIndexes[1]]
		} else {
			comp.InstructionIndex += 3
		}
		return comp.Step(input)
	// 7: less-than, if param1 < param2 then store 1 in postion of 3rd param, else store 0
	case 7:
		if comp.PuzzleInput[paramIndexes[0]] < comp.PuzzleInput[paramIndexes[1]] {
			comp.PuzzleInput[paramIndexes[2]] = 1
		} else {
			comp.PuzzleInput[paramIndexes[2]] = 0
		}
		comp.InstructionIndex += 4
		return comp.Step(input)
	// 8: equals, if param1 == param2 then set position of 3rd param to 1, else store 0
	case 8:
		if comp.PuzzleInput[paramIndexes[0]] == comp.PuzzleInput[paramIndexes[1]] {
			comp.PuzzleInput[paramIndexes[2]] = 1
		} else {
			comp.PuzzleInput[paramIndexes[2]] = 0
		}
		comp.InstructionIndex += 4
		return comp.Step(input)
	default:
		log.Fatal("Error: unknown opcode: ", opcode)
	}
	// this should never be called b/c switch statement will always return
	return -1
}

/*
GetOpCodeAndIndexes will parse the instruction at comp.PuzzleInput[comp.InstructionIndex]
- opcode will be the left two digits, mod by 100 will get that
- rest of instructions will be grabbed via mod 10
	- these also have to be parsed for the
*/
func (comp *Intcode) GetOpCodeAndIndexes() (int, [3]int) {
	instruction := comp.PuzzleInput[comp.InstructionIndex]

	// opcode is the lowest two digits, so mod by 100
	opcode := instruction % 100
	instruction /= 100

	// assign the indexes that need to be read by reading the parameter modes
	var paramIndexes [3]int
	for i := 1; i <= 3 && comp.InstructionIndex+i < len(comp.PuzzleInput); i++ {
		// grab the mode with a mod, last digit
		mode := instruction % 10
		instruction /= 10

		switch mode {
		case 1: // immediate mode, the index itself
			paramIndexes[i-1] = comp.InstructionIndex + i
		case 0: // position mode, index will be the value at the index
			paramIndexes[i-1] = comp.PuzzleInput[comp.InstructionIndex+i]
		}
	}

	return opcode, paramIndexes
}
