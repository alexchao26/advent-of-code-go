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

	// initialize a computer with a senor boost input of `2`
	robotBrain := MakeComputer(inputNumbers)
	robot := MakeRobot(0, 0)

	for robotBrain.IsRunning {
		// get the current color from the robot's map
		currentCoords := fmt.Sprintf("%v,%v", robot.x, robot.y)
		currentColor := robot.MapCoordsToColor[currentCoords]
		robotBrain.Step(currentColor)

		// get outputs from the robot's brain (Intcode)
		lenOutputs := len(robotBrain.Outputs)
		color := robotBrain.Outputs[lenOutputs-2]
		direction := robotBrain.Outputs[lenOutputs-1]

		// "paint"/update robot's Map and move the robot
		robot.MapCoordsToColor[currentCoords] = color
		robot.MoveRobot(direction)
	}

	fmt.Printf("Tiles painted %v\n", len(robot.MapCoordsToColor))
}

// Robot struct, x and y are coordinate system based, NOT 2D array 0-indexed
type Robot struct {
	x                int
	y                int
	Direction        string
	MapCoordsToColor map[string]int
}

// MakeRobot holds info on the location and direction of the robot only
func MakeRobot(startX, startY int) *Robot {
	return &Robot{
		startX,
		startY,
		"up",
		make(map[string]int),
	}
}

// MoveRobot moves the Robot
func (robot *Robot) MoveRobot(direction int) {
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
func MakeComputer(PuzzleInput []int) Intcode {
	puzzleInputCopy := make([]int, len(PuzzleInput))
	copy(puzzleInputCopy, PuzzleInput)

	comp := Intcode{
		puzzleInputCopy,
		0,
		0,
		make([]int, 0),
		true,
	}
	return comp
}

// Step will read the next 4 values in the input `sli` and make updates
// according to the opcodes
func (comp *Intcode) Step(input int) {
	// read the instruction, opcode and the indexes where the params point to
	opcode, paramIndexes := comp.GetOpCodeAndParamIndexes()
	param1, param2, param3 := paramIndexes[0], paramIndexes[1], paramIndexes[2]

	// ensure params are within the bounds of PuzzleInput, resize if necessary
	comp.ResizeMemory(param1, param2, param3)

	switch opcode {
	case 99: // 99: Terminates program
		fmt.Println("Terminating...")
		comp.IsRunning = false
	case 1: // 1: Add next two paramIndexes, store in third
		comp.PuzzleInput[param3] = comp.PuzzleInput[param1] + comp.PuzzleInput[param2]
		comp.InstructionIndex += 4
		comp.Step(input)
	case 2: // 2: Multiply next two and store in third
		comp.PuzzleInput[param3] = comp.PuzzleInput[param1] * comp.PuzzleInput[param2]
		comp.InstructionIndex += 4
		comp.Step(input)
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
		comp.Step(-1)
	case 4: // 4: outputs its input value
		// set LastOutput of the computer & log it
		comp.Outputs = append(comp.Outputs, comp.PuzzleInput[param1])
		// fmt.Printf("Opcode 4 output: %v\n", comp.LastOutput)
		comp.InstructionIndex += 2

		// continue running until terminates or asks for another input
		comp.Step(input)
	// 5: jump-if-true: if first param != 0, move pointer to second param, else nothing
	case 5:
		if comp.PuzzleInput[param1] != 0 {
			comp.InstructionIndex = comp.PuzzleInput[param2]
		} else {
			comp.InstructionIndex += 3
		}
		comp.Step(input)
	// 6: jump-if-false, if first param == 0 then set instruction pointer to 2nd param, else nothing
	case 6:
		if comp.PuzzleInput[param1] == 0 {
			comp.InstructionIndex = comp.PuzzleInput[param2]
		} else {
			comp.InstructionIndex += 3
		}
		comp.Step(input)
	// 7: less-than, if param1 < param2 then store 1 in postion of 3rd param, else store 0
	case 7:
		if comp.PuzzleInput[param1] < comp.PuzzleInput[param2] {
			comp.PuzzleInput[param3] = 1
		} else {
			comp.PuzzleInput[param3] = 0
		}
		comp.InstructionIndex += 4
		comp.Step(input)
	// 8: equals, if param1 == param2 then set position of 3rd param to 1, else store 0
	case 8:
		if comp.PuzzleInput[param1] == comp.PuzzleInput[param2] {
			comp.PuzzleInput[param3] = 1
		} else {
			comp.PuzzleInput[param3] = 0
		}
		comp.InstructionIndex += 4
		comp.Step(input)
	// 9: adjust relative base
	case 9:
		comp.RelativeBase += comp.PuzzleInput[param1]
		comp.InstructionIndex += 2
		comp.Step(input)
	default:
		log.Fatalf("Error: unknown opcode %v at index %v", opcode, comp.PuzzleInput[comp.InstructionIndex])
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
	if maxArg > len(comp.PuzzleInput) {
		// make empty slice to copy into, of the new, larger size
		resizedPuzzleInput := make([]int, maxArg+1)
		// copy old puzzle input values in
		copy(resizedPuzzleInput, comp.PuzzleInput)

		// overwrite puzzle input
		comp.PuzzleInput = resizedPuzzleInput
	}
}
