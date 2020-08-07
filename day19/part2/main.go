/*
Intcode struct is defined within this file
	- Every drone needs its own computer made, I let them get garbage collected as often
	  as possible, not sure how extendable this is going to be to part 2
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

	// x, y := 1135, 1625

	allHorizontalRanges := [][2]int{
		[2]int{0, 1},
	}

	for len(allHorizontalRanges) < 5000 {
		startOfLastRange := allHorizontalRanges[len(allHorizontalRanges)-1][0]
		horizontalRange := getHorizontalRange(inputNumbers, startOfLastRange, len(allHorizontalRanges))
		allHorizontalRanges = append(allHorizontalRanges, horizontalRange)

		if len(allHorizontalRanges) > 100 {
			// indexes of the current line and 100 rows up
			currentIndex := len(allHorizontalRanges) - 1
			index100RowsUp := len(allHorizontalRanges) - 100

			// check if a square can fit from the start of the current range to the end of the range 100 rows up
			if allHorizontalRanges[index100RowsUp][1]-allHorizontalRanges[currentIndex][0] >= 99 {
				// Print the AoC format
				fmt.Println("AoC answer:", allHorizontalRanges[currentIndex][0]*10000+index100RowsUp)

				break
			}
		}
	}
}

func getHorizontalRange(inputNumbers []int, startX, y int) [2]int {
	horizontalRange := [2]int{}

	var lastOutput int

	// step until the output is a one
	// OR exit condition of 10 cells have all returned zeroes (this will apply for some of the first levels)
	for x := startX; lastOutput == 0; x++ {
		lastOutput = makeDroneAndTest(inputNumbers, x, y)
		// when the first element is found that is "pulled", set the start of the range
		if lastOutput == 1 {
			horizontalRange[0] = x
		}

		// NOTE edge case for the first few rows if none of the drones are pulled
		if x > startX+20 {
			return [2]int{0, 0}
		}

	}

	for x := horizontalRange[0] + 1; lastOutput != 0; x++ {
		lastOutput = makeDroneAndTest(inputNumbers, x, y)

		if lastOutput == 0 {
			horizontalRange[1] = x - 1
		}
	}

	return horizontalRange
}

func makeDroneAndTest(inputNumbers []int, x, y int) int {
	drone := MakeComputer(inputNumbers)
	drone.Step(x)
	drone.Step(y)
	lastOutput := drone.Outputs[len(drone.Outputs)-1]

	return lastOutput
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
	PuzzleInput      []int // file/puzzle input parsed into slice of ints
	InstructionIndex int   // stores the index where the next instruction is
	RelativeBase     int   // relative base for opcode 9 and param mode 2
	Outputs          []int // stores all outputs
	IsRunning        bool  // will be true until a 99 opcode is hit
}

// MakeComputer initializes a new comp
func MakeComputer(PuzzleInput []int) *Intcode {
	puzzleInputCopy := make([]int, len(PuzzleInput))
	copy(puzzleInputCopy, PuzzleInput)

	comp := Intcode{
		puzzleInputCopy,
		0,
		0,
		make([]int, 0),
		true,
	}
	return &comp
}

// Step will read the next 4 values in the input `sli` and make updates
// according to the opcodes
// Update to run iteratively (while the computer is running)
// it will also return out if a -1 input is asked for
// then call Step again to provide the next input, or run with -1 from the start
//   to run the computer until it asks for an input OR terminates
func (comp *Intcode) Step(input int) {
	for comp.IsRunning {
		// read the instruction, opcode and the indexes where the params point to
		opcode, paramIndexes := comp.GetOpCodeAndParamIndexes()
		param1, param2, param3 := paramIndexes[0], paramIndexes[1], paramIndexes[2]

		// ensure params are within the bounds of PuzzleInput, resize if necessary
		switch opcode {
		case 1, 2, 7, 8:
			comp.ResizeMemory(param1, param2, param3)
		case 5, 6:
			comp.ResizeMemory(param1, param2)
		case 3, 4, 9:
			comp.ResizeMemory(param1)
		}

		switch opcode {
		case 99: // 99: Terminates program
			// fmt.Println("Terminating...")
			comp.IsRunning = false
		case 1: // 1: Add next two paramIndexes, store in third
			comp.PuzzleInput[param3] = comp.PuzzleInput[param1] + comp.PuzzleInput[param2]
			comp.InstructionIndex += 4
		case 2: // 2: Multiply next two and store in third
			comp.PuzzleInput[param3] = comp.PuzzleInput[param1] * comp.PuzzleInput[param2]
			comp.InstructionIndex += 4
		case 3: // 3: Takes one input and saves it to position of one parameter
			// check if input has already been used (i.e. input == -1)
			// if it's been used, return out to prevent further Steps
			// NOTE: making a big assumption that -1 will never be an input...
			if input == -1 {
				return
			}

			// else recurse with a -1 to signal the initial input has been processed
			comp.PuzzleInput[param1] = input
			comp.InstructionIndex += 2

			// change the input value so the next time a 3 opcode is hit, will return out
			input = -1
		case 4: // 4: outputs its input value
			output := comp.PuzzleInput[param1]
			// set LastOutput of the computer & log it
			comp.Outputs = append(comp.Outputs, output)

			comp.InstructionIndex += 2
		// 5: jump-if-true: if first param != 0, move pointer to second param, else nothing
		case 5:
			if comp.PuzzleInput[param1] != 0 {
				comp.InstructionIndex = comp.PuzzleInput[param2]
			} else {
				comp.InstructionIndex += 3
			}
		// 6: jump-if-false, if first param == 0 then set instruction pointer to 2nd param, else nothing
		case 6:
			if comp.PuzzleInput[param1] == 0 {
				comp.InstructionIndex = comp.PuzzleInput[param2]
			} else {
				comp.InstructionIndex += 3
			}
		// 7: less-than, if param1 < param2 then store 1 in postion of 3rd param, else store 0
		case 7:
			if comp.PuzzleInput[param1] < comp.PuzzleInput[param2] {
				comp.PuzzleInput[param3] = 1
			} else {
				comp.PuzzleInput[param3] = 0
			}
			comp.InstructionIndex += 4
		// 8: equals, if param1 == param2 then set position of 3rd param to 1, else store 0
		case 8:
			if comp.PuzzleInput[param1] == comp.PuzzleInput[param2] {
				comp.PuzzleInput[param3] = 1
			} else {
				comp.PuzzleInput[param3] = 0
			}
			comp.InstructionIndex += 4
		// 9: adjust relative base
		case 9:
			comp.RelativeBase += comp.PuzzleInput[param1]
			comp.InstructionIndex += 2
		default:
			log.Fatalf("Error: unknown opcode %v at index %v", opcode, comp.PuzzleInput[comp.InstructionIndex])
		}
	}
}

/*
GetOpCodeAndParamIndexes will parse the instruction at comp.PuzzleInput[comp.InstructionIndex]
- opcode will be the left two digits, mod by 100 will get that
- rest of instructions will be grabbed via mod 10
	- these also have to be parsed for the
*/
func (comp *Intcode) GetOpCodeAndParamIndexes() (int, [3]int) {
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
		case 0: // position mode, index will be the value at the index
			paramIndexes[i-1] = comp.PuzzleInput[comp.InstructionIndex+i]
		case 1: // immediate mode, the index itself
			paramIndexes[i-1] = comp.InstructionIndex + i
		case 2: // relative mode, like position mode but index is added to relative base
			paramIndexes[i-1] = comp.PuzzleInput[comp.InstructionIndex+i] + comp.RelativeBase
		}
	}

	return opcode, paramIndexes
}

// ResizeMemory will take any number of integers and resize the computer's memory appropriately
func (comp *Intcode) ResizeMemory(sizes ...int) {
	// get largest of input sizes
	maxArg := sizes[0]
	for _, v := range sizes {
		if v > maxArg {
			maxArg = v
		}
	}

	// resize if PuzzleInput's length is shorter
	if maxArg >= len(comp.PuzzleInput) {
		// make empty slice to copy into, of the new, larger size
		resizedPuzzleInput := make([]int, maxArg+1)
		// copy old puzzle input values in
		copy(resizedPuzzleInput, comp.PuzzleInput)

		// overwrite puzzle input
		comp.PuzzleInput = resizedPuzzleInput
	}
}
