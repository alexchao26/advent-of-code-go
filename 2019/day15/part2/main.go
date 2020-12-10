/*
Intcode struct is defined within this file
Robot struct houses an Intcode computer and its RecursiveMove method populates a map of
	coordinates to the floor type (-1: wall, 1: hallway, 2: O2 tank, 5: origin)
	That map is converted into a 2D grid (slice)

	The shortest length is no longer needed

	Two functions added that determine if the space is full of O2 and to spreadOxygen for one minute
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algos"
	"github.com/alexchao26/advent-of-code-go/mathutil"
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

	robot := MakeRobot(inputNumbers)

	// fire off recursive move function to populate the robot's floorDetails property
	robot.RecursiveMove()

	// make grid from the map of coordinates to floor types
	grid := Draw(robot.floorDetails)

	// overwrite the origin with a 1 (open floor space)
	for y, row := range grid {
		for x, tileType := range row {
			if tileType == 5 {
				grid[y][x] = 1
			}
		}
	}

	// while the grid is not full of oxygen, spread oxygen and increment minutes
	var minutes int
	for !isFullOfOxygen(grid) {
		spreadOxygen(grid)
		minutes++
	}

	fmt.Println("Minutes elapsed", minutes)
}

func isFullOfOxygen(grid [][]int) bool {
	for _, row := range grid {
		for _, tile := range row {
			// if there is a hallway space that is not filled with O2, return false
			if tile == 1 {
				return false
			}
		}
	}
	// if entire looping passes, return true
	return true
}

var dRow []int = []int{0, 0, -1, 1}
var dCol []int = []int{-1, 1, 0, 0}

// spreadOxygen will spread all oxygen to one neighboring cell
// returns boolean true if O2 has not spread everywhere (i.e. run again), false if O2 is everywhere
func spreadOxygen(grid [][]int) {
	// traverse through grid and mark all cells that are a 1 and have a neighboring 2
	// tag with -1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// traverse around neighbors
				for d := 0; d < 4; d++ {
					neighborRow := i + dRow[d]
					neighborCol := j + dCol[d]
					inBounds := neighborRow >= 0 && neighborRow < len(grid) && neighborCol >= 0 && neighborCol < len(grid[0])
					// if a neighboring cell is a 2 (i.e. filled with oxygen), then mark this cell
					if inBounds && grid[neighborRow][neighborCol] == 2 {
						grid[i][j] = -1
						break
					}
				}
			}
		}
	}

	// then iterate through again changing all -1's to a 2
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == -1 {
				grid[i][j] = 2
			}
		}
	}
}

// Robot struct to maintain detail's on the Robot's coordinates, path
type Robot struct {
	fromTop, fromLeft int
	floorDetails      map[string]int // maps coordinates and type of tile (0 == wall, 1 == path, 2 == oxygen)
	computer          *Intcode
}

// MakeRobot returns an instance of a Robot
func MakeRobot(intcodeInput []int) *Robot {
	return &Robot{
		0,
		0,
		map[string]int{"0,0": 5}, // mark the origin specially with a 5
		MakeComputer(intcodeInput),
	}
}

var backtrack map[int]int = map[int]int{
	1: 2, // north (1), south (2)
	2: 1,
	3: 4, // west (3), east(4)
	4: 3,
}

// dx is the difference to add when traveling in the given direction
// i.e. add 0 for north and south, for west decrement 1, for east add 1
var dx map[int]int = map[int]int{
	1: 0,
	2: 0,
	3: -1,
	4: 1,
}

// dy is the vertical distance traveled
var dy map[int]int = map[int]int{
	1: 1,
	2: -1,
	3: 0,
	4: 0,
}

// RecursiveMove will populate a robot's floor details property by traveling in all directions
// and
func (robot *Robot) RecursiveMove() {
	for i := 1; i <= 4; i++ {
		// if next coordinates have already been detailed, skip all calculations
		nextCoords := fmt.Sprintf("%v,%v", robot.fromTop+dy[i], robot.fromLeft+dx[i])

		if robot.floorDetails[nextCoords] == 0 {
			robot.computer.Step(i)
			computerOutput := robot.computer.Outputs[len(robot.computer.Outputs)-1]

			switch computerOutput {
			case 0: // hit a wall, do not recurse
				// update robot's wall coords to include the wall
				// note representing walls with a -1 to avoid the zero value detection
				robot.floorDetails[nextCoords] = -1
			case 1, 2: // walked and hit the O2 tank or not
				// update floorDetails
				robot.floorDetails[nextCoords] = computerOutput

				// continue to walk the robot. walk the robot into the nextCoords spot
				robot.fromLeft += dx[i]
				robot.fromTop += dy[i]

				// recurse
				robot.RecursiveMove()

				// backtrack so the robot walks in the remainder of directions from this output
				robot.fromLeft -= dx[i]
				robot.fromTop -= dy[i]
				// backtrack the computer
				robot.computer.Step(backtrack[i])
			}
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
	if maxArg >= len(comp.PuzzleInput) {
		// make empty slice to copy into, of the new, larger size
		resizedPuzzleInput := make([]int, maxArg+1)
		// copy old puzzle input values in
		copy(resizedPuzzleInput, comp.PuzzleInput)

		// overwrite puzzle input
		comp.PuzzleInput = resizedPuzzleInput
	}
}

// Draw was copied from day11. It converts a map of points mapped from a (0,0) origin to a 2D grid
// The origin loses its reference...
func Draw(mapCoordsToType map[string]int) [][]int {
	var lowX, highX, lowY, highY int
	for key := range mapCoordsToType {
		coords := strings.Split(key, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		switch {
		case x < lowX:
			lowX = x
		case x > highX:
			highX = x
		}
		switch {
		case y < lowY:
			lowY = y
		case y > highY:
			highY = y
		}
	}

	// Determine the bounds of the grid
	edgeLength := 2 * mathutil.MaxInt(-lowY, -lowX, highY, highX)

	grid := make([][]int, edgeLength)
	for i := 0; i < edgeLength; i++ {
		// each character will initialize as a space character
		grid[i] = make([]int, edgeLength)
	}

	// Iterate through all coordinates and transcribe x,y onto a 2D grid
	// where the math is a little different...
	for key, val := range mapCoordsToType {
		// key is string coords
		coords := strings.Split(key, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		x += edgeLength / 2
		y += edgeLength / 2
		// val is color to paint (1 or 0)
		if val != -1 {
			grid[x][y] = val
		}
	}

	// trim off due to making the initial grid too large
	grid = trim(grid)
	// rotate it because of how I coded up the robot's coordinates :/
	grid = algos.RotateIntGrid(grid)
	// retrim
	grid = trim(grid)

	return grid
}

// helper function for Draw to remove whitespace from overestimating the size
// of the drawing space
func trim(grid [][]int) [][]int {
	// remove all empty rows at top and bottom
removeRowsTop:
	for i := 0; i < len(grid); {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				break removeRowsTop
			}
		}
		grid = grid[1:]
	}

	// remove empty columns on left
removeColsLeft:
	for i := 0; i < len(grid[0]); {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] != 0 {
				break removeColsLeft
			}
		}
		// if loop hasn't broken out, iterate over first "column" and slice off "0-index"
		for j := 0; j < len(grid); j++ {
			grid[j] = grid[j][1:]
		}
	}

	return grid
}
