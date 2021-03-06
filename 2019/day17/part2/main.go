package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	// read the input file, modify it to a slice of numbers
	inputFile := util.ReadFile("../input.txt")

	splitStrings := strings.Split(inputFile, ",")

	inputNumbers := make([]int, len(splitStrings))
	for i, v := range splitStrings {
		inputNumbers[i], _ = strconv.Atoi(v)
	}

	// Modify the first address from a 1 to a 2 to start the vacuum robot
	inputNumbers[0] = 2

	robot := MakeRobot(inputNumbers)
	// fire off function to populate the robot's floorGrid property
	robot.GetFloorGrid()

	// NOTE the computer accepts numbers as inputs, but these numbers correlate to ASCII
	// computer will ask for main movement routine. allows: [A-C,] i.e. A B C separated by ,
	// end with a newline, e.g. integer 10 in ASCII
	A, B, C := int('A'), int('B'), int('C')
	comma, newline := int(','), int('\n')

	mainRoutineOrder := []int{A, B, A, C, B, C, B, C, A, C}
	for i, routine := range mainRoutineOrder {
		robot.computer.Step(routine)
		if i != len(mainRoutineOrder)-1 {
			robot.computer.Step(comma)
		} else {
			robot.computer.Step(newline)
		}
	}

	// computer then asks for details of each movement function
	// accepts [LR0-9,] ended with a newline
	// NOTE I determined these paths manually after printing the floor and finding a pattern by eye...
	L, R := int('L'), int('R')
	numbers := make([]int, 10)
	for i := range numbers {
		numbers[i] = int('0') + i
	}
	// A is L10 R12 R12
	patternA := []int{
		L, comma, numbers[1], numbers[0], comma,
		R, comma, numbers[1], numbers[2], comma,
		R, comma, numbers[1], numbers[2],
		newline,
	}
	// B is R6 R10 L10
	patternB := []int{
		R, comma, numbers[6], comma,
		R, comma, numbers[1], numbers[0], comma,
		L, comma, numbers[1], numbers[0],
		newline,
	}
	// C is R10 L10 L12 R6
	patternC := []int{
		R, comma, numbers[1], numbers[0], comma,
		L, comma, numbers[1], numbers[0], comma,
		L, comma, numbers[1], numbers[2], comma,
		R, comma, numbers[6],
		newline,
	}
	// fmt.Println("pattern A: ")
	for _, v := range patternA {
		// fmt.Printf("%v,", v)
		robot.computer.Step(v)
	}
	for _, v := range patternB {
		robot.computer.Step(v)
	}
	for _, v := range patternC {
		robot.computer.Step(v)
	}

	// finally, asks for continuous video feed or not 'y' or 'n' and a new line
	robot.computer.Step(int('y'))
	robot.computer.Step(newline)

	fmt.Println("Dust collected", robot.computer.Outputs[len(robot.computer.Outputs)-1])
}

// Robot struct to maintain detail's on the Robot's coordinates, path
type Robot struct {
	row, col  int
	floorGrid [][]string
	computer  *Intcode
}

// MakeRobot returns an instance of a Robot
func MakeRobot(intcodeInput []int) *Robot {
	return &Robot{
		0,
		0,
		make([][]string, 0),
		MakeComputer(intcodeInput),
	}
}

// GetFloorGrid will fire off the computer and populate the robot's floor details
func (robot *Robot) GetFloorGrid() {
	robot.computer.Step(-1)
	robot.floorGrid = append(robot.floorGrid, []string{})
	row := 0

	for _, v := range robot.computer.Outputs {
		switch v {
		case 10:
			row++
			robot.floorGrid = append(robot.floorGrid, []string{})
		default:
			tileType := cast.ASCIIIntToChar(v)
			robot.floorGrid[row] = append(robot.floorGrid[row], tileType)
		}
	}

	// parse off empty slices @ end
	for i := len(robot.floorGrid) - 1; i >= 0; i-- {
		if len(robot.floorGrid[i]) == 0 {
			robot.floorGrid = robot.floorGrid[:len(robot.floorGrid)-1]
		}
	}
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
			fmt.Println("Terminating...")
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

			// NOTE this is specific to day17 to print the robot walking around the scaffold
			// if the last two outputs are newlines (ASCII 10's), print out the output
			// then just clear the output to make life easy
			if output == 10 && comp.Outputs[len(comp.Outputs)-2] == 10 {
				Print2DGrid(comp.Outputs)
				// clear outputs slice, sleep for 100ms
				comp.Outputs = []int{}
				time.Sleep(time.Millisecond * 100)
			}

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

// Print2DGrid is day17 specific. allValues are ASCII ints including 10 for newline
func Print2DGrid(allValues []int) {
	var row int
	floorGrid := [][]string{[]string{}}
	for _, v := range allValues {
		switch v {
		case 10:
			row++
			floorGrid = append(floorGrid, []string{})
		default:
			tileType := cast.ASCIIIntToChar(v)
			floorGrid[row] = append(floorGrid[row], tileType)
		}
	}

	for _, v := range floorGrid {
		fmt.Println(v)
	}
}
