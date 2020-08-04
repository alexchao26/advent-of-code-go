/*
IntcodeY struct is defined within this file
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

	robotBrain := MakeComputerY(inputNumbers)
	robotBrain.StepY(0)
	robotBrain.StepY(0)

	fmt.Println(robotBrain.Outputs)
}

// RobotY struct, x and y are coordinate system based, NOT 2D array 0-indexed
type RobotY struct {
	x                int
	y                int
	Direction        string
	MapCoordsToColor map[string]int
}

// MakeRobotY holds info on the location and direction of the robot only
func MakeRobotY(startX, startY int) *RobotY {
	return &RobotY{
		startX,
		startY,
		"up",
		make(map[string]int),
	}
}

// MoveRobotY moves the RobotY
func (robot *RobotY) MoveRobotY(direction int) {
	// direction is the same as the output from the robot brain
	// i.e. 0 to turn left, 1 to turn right, then step forward 1 space
	turnLeft := map[string]string{
		"up":    "left",
		"left":  "down",
		"down":  "right",
		"right": "up",
	}
	turnRight := map[string]string{
		"up":    "right",
		"right": "down",
		"down":  "left",
		"left":  "up",
	}

	if direction == 0 {
		robot.Direction = turnLeft[robot.Direction]
	} else {
		robot.Direction = turnRight[robot.Direction]
	}

	switch robot.Direction {
	case "up":
		robot.y++
	case "down":
		robot.y--
	case "left":
		robot.x--
	case "right":
		robot.x++
	}
}

/*
IntcodeY is an OOP approach *************************************************
MakeComputerY is equivalent to the constructor
StepY takes in an input int and updates properties in the computer:
	- InstructionIndex: where to read the next instruction from
	- LastOutput, what the last opcode 4 outputted
	- PuzzleIndex based if the last instruction modified the puzzle at all
****************************************************************************/
type IntcodeY struct {
	PuzzleInput      []int // file/puzzle input parsed into slice of ints
	InstructionIndex int   // stores the index where the next instruction is
	Outputs          []int // all outputs stored in order
	IsRunning        bool  // will be true until a 99 opcode is hit
}

// MakeComputerY initializes a new comp
func MakeComputerY(PuzzleInput []int) IntcodeY {
	puzzleInputCopy := make([]int, len(PuzzleInput))
	copy(puzzleInputCopy, PuzzleInput)

	comp := IntcodeY{
		puzzleInputCopy,
		0,
		make([]int, 0),
		true,
	}

	return comp
}

// StepY will read the next 4 values in the input `sli` and make updates
// according to the opcodes
func (comp *IntcodeY) StepY(input int) {
	// read the instruction, opcode and the indexes where the params point to
	opcode, paramIndexes := comp.GetOpCodeAndParamIndexesY()
	param1, param2, param3 := paramIndexes[0], paramIndexes[1], paramIndexes[2]

	switch opcode {
	case 99: // 99: Terminates program
		// fmt.Println("Terminating...")
		comp.IsRunning = false
	case 1: // 1: Add next two paramIndexes, store in third
		comp.PuzzleInput[param3] = comp.PuzzleInput[param1] + comp.PuzzleInput[param2]
		comp.InstructionIndex += 4
		comp.StepY(input)
	case 2: // 2: Multiply next two and store in third
		comp.PuzzleInput[param3] = comp.PuzzleInput[param1] * comp.PuzzleInput[param2]
		comp.InstructionIndex += 4
		comp.StepY(input)
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
		comp.StepY(-1)
	case 4: // 4: outputs its input value
		// set LastOutput of the computer & log it
		comp.Outputs = append(comp.Outputs, comp.PuzzleInput[param1])
		// fmt.Printf("Opcode 4 output: %v\n", comp.LastOutput)
		comp.InstructionIndex += 2

		// continue running until terminates or asks for another input
		comp.StepY(input)
	// 5: jump-if-true: if first param != 0, move pointer to second param, else nothing
	case 5:
		if comp.PuzzleInput[param1] != 0 {
			comp.InstructionIndex = comp.PuzzleInput[param2]
		} else {
			comp.InstructionIndex += 3
		}
		comp.StepY(input)
	// 6: jump-if-false, if first param == 0 then set instruction pointer to 2nd param, else nothing
	case 6:
		if comp.PuzzleInput[param1] == 0 {
			comp.InstructionIndex = comp.PuzzleInput[param2]
		} else {
			comp.InstructionIndex += 3
		}
		comp.StepY(input)
	// 7: less-than, if param1 < param2 then store 1 in postion of 3rd param, else store 0
	case 7:
		if comp.PuzzleInput[param1] < comp.PuzzleInput[param2] {
			comp.PuzzleInput[param3] = 1
		} else {
			comp.PuzzleInput[param3] = 0
		}
		comp.InstructionIndex += 4
		comp.StepY(input)
	// 8: equals, if param1 == param2 then set position of 3rd param to 1, else store 0
	case 8:
		if comp.PuzzleInput[param1] == comp.PuzzleInput[param2] {
			comp.PuzzleInput[param3] = 1
		} else {
			comp.PuzzleInput[param3] = 0
		}
		comp.InstructionIndex += 4
		comp.StepY(input)
	default:
		log.Fatalf("Error: unknown opcode %v at index %v", opcode, comp.PuzzleInput[comp.InstructionIndex])
	}
}

/*
GetOpCodeAndParamIndexesY will parse the instruction at comp.PuzzleInput[comp.InstructionIndex]
- opcode will be the left two digits, mod by 100 will get that
- rest of instructions will be grabbed via mod 10
	- these also have to be parsed for the
*/
func (comp *IntcodeY) GetOpCodeAndParamIndexesY() (int, [3]int) {
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
		}
	}

	return opcode, paramIndexes
}
