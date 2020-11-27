/*
Intcode struct is defined within this file
Game struct maintains information about the ball and paddle coordinates & state of the screen
NOTE: the "computer" and "game" are decoupled in this solution which is probably not _ideal_,
NOTE  but reusing the old code, it was simpler like this

NOTE: The "hold for an input" value was changed from -1 to -2 in the Intcode computer
*/

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

	// set first puzzle input value to 2 to "play the game for free"
	inputNumbers[0] = 2

	// initialize a computer
	comp := MakeComputer(inputNumbers)

	// Get the setup Outputs for the game screen
	// NOTE inputting -2 because this is the input that "holds" the computer when it's asking for an input
	comp.Step(-2)

	// initialize game
	game := MakeGame(comp.Outputs)

	// loop while the intcode computer is asking for an input
	for comp.IsRunning {
		// clear previous outputs of computer
		comp.Outputs = []int{}

		// Find joystick input based on ball vs paddle x coord: 0 (zero value), -1 (left) or 1 (right)
		var joystickInput int
		if game.ballX < game.paddleX {
			joystickInput = -1
		} else if game.ballX > game.paddleX {
			joystickInput = 1
		}

		// move joystick by giving computer input
		comp.Step(joystickInput)

		// handle all trios of outputs from the computer
		for i := 0; i < len(comp.Outputs); i += 3 {
			game.HandleOutput(comp.Outputs[i], comp.Outputs[i+1], comp.Outputs[i+2])
		}

		// // Uncomment to print board at every "step"
		// for _, level := range game.screen {
		// 	// pretty print game screen
		// 	var strLevel string
		// 	for _, tileID := range level {
		// 		if tileID == 0 {
		// 			strLevel += " "
		// 		} else {
		// 			strLevel += strconv.Itoa(tileID)
		// 		}
		// 	}
		// 	fmt.Println(strLevel)
		// }
		// fmt.Println("") // print an extra line to separate out the new game screens
		// time.Sleep(time.Millisecond * 10)
	}

	fmt.Println("Final score", game.score)
}

// Game stores the screen appearance of the game
type Game struct {
	screen           [][]int
	score            int
	ballX, ballY     int
	paddleX, paddleY int
}

// MakeGame returns an instance of a Game struct
func MakeGame(setupInstructions []int) Game {
	// determine the size of the screen, i.e. iterate through all instructions 3 at a time and find
	// the max x (2nd of triplet) and max y (3rd of triplet)
	// note: assuming all x and y values are positive, i.e. drawing down & right from top left corner
	var maxX, maxY int
	for i := 0; i+2 < len(setupInstructions); i += 3 {
		x := setupInstructions[i]
		y := setupInstructions[i+1]
		if x < 0 || y < 0 {
			continue
		}
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
	}

	// make screen, x is distance from left, y is distance from top, so Y is the number of rows
	// X is number of columns. +1 to convert max index to length of slice
	screen := make([][]int, maxY+1)
	for i := range screen {
		screen[i] = make([]int, maxX+1)
	}

	// initialize Game
	game := Game{screen, 0, 0, 0, 0, 0}

	// fill screen with initial outputs
	for i := 0; i+2 < len(setupInstructions); i += 3 {
		game.HandleOutput(setupInstructions[i], setupInstructions[i+1], setupInstructions[i+2])
	}

	return game
}

// HandleOutput handles one trio of outputs from intcode
func (game *Game) HandleOutput(fromLeft, fromTop, tileID int) {
	if fromLeft == -1 && fromTop == 0 {
		game.score = tileID
	} else {
		game.screen[fromTop][fromLeft] = tileID
		if tileID == 3 {
			game.paddleX = fromLeft
			game.paddleY = fromTop
		}
		if tileID == 4 {
			game.ballX = fromLeft
			game.ballY = fromTop
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
	// Note: need to optimize this to not resize if the params are not being used
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
		comp.Step(input)
	case 2: // 2: Multiply next two and store in third
		comp.PuzzleInput[param3] = comp.PuzzleInput[param1] * comp.PuzzleInput[param2]
		comp.InstructionIndex += 4
		comp.Step(input)
	case 3: // 3: Takes one input and saves it to position of one parameter
		// check if input has already been used (i.e. input == -1)
		// if it's been used, return out to prevent further Steps
		// NOTE: making a big assumption that -1 will never be an input...
		if input == -2 {
			return
		}

		// else recurse with a -1 to signal the initial input has been processed
		comp.PuzzleInput[param1] = input
		comp.InstructionIndex += 2
		comp.Step(-2)
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
	if maxArg >= len(comp.PuzzleInput) {
		// make empty slice to copy into, of the new, larger size
		resizedPuzzleInput := make([]int, maxArg+1)
		// copy old puzzle input values in
		copy(resizedPuzzleInput, comp.PuzzleInput)

		// overwrite puzzle input
		comp.PuzzleInput = resizedPuzzleInput
	}
}
