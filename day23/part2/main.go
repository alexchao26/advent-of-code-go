/*
Intcode struct is defined within this file
Network struct to store 50 instances of Intcode computers and 50 queues for their inputs
  NAT variables are just stored in main goroutine instead of in another struct
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

	// make the network
	network := Network{}
	network.Init(inputNumbers)

	// NAT packet (gets overwritten anytime a write to 255 happens)
	// lastNatY value stored here and compared everytime the NAT will write to 0
	natPacket := [2]int{}
	var lastNatY int

	// run indefinitely...
	// assuming that I can step through each computer one at a time,
	// that they don't truly need to be running concurrently
	for {
		// declare a boolean flag to signal when all computers are waiting for input
		allNatsWaiting := true

		// iterate over all 50 computers
		for i := 0; i < 50; i++ {
			// if this computer's queue is empty, use -1 as an input
			if len(network.queues[i]) == 0 {
				network.computers[i].Step(-1)
			} else {
				// Flip boolean for allNatsWaiting if any
				allNatsWaiting = false

				// Process off of the front of this computer's queue
				front := network.queues[i][0]
				network.computers[i].Step(front[0])
				network.computers[i].Step(front[1])

				// dequeue
				network.queues[i] = network.queues[i][1:]
			}

			// while there are unhandled outputs of this computer, add them to the
			// receiving computers's queues OR the NAT
			for len(network.computers[i].Outputs) > 2 {
				destination := network.computers[i].Outputs[0]
				packet := [2]int{network.computers[i].Outputs[1],
					network.computers[i].Outputs[2]}

				// remove three from Outputs slice
				network.computers[i].Outputs = network.computers[i].Outputs[3:]

				// if destination is 255, overwrite NAT packet
				if destination == 255 {
					natPacket = packet
				} else {
					// otherwise, add to queue of correct NIC computer
					network.queues[destination] = append(network.queues[destination], packet)
				}
			}
		}

		// if all nat computers are waiting for inputs, write the natpacket to
		// the zero-th computer's queue
		if allNatsWaiting {
			// check if this packet has a duplicate Y value, if so print AoC output
			if lastNatY == natPacket[1] {
				fmt.Println(lastNatY, "written to NAT twice")
				// stop the infinite loop
				return
			}
			network.queues[0] = append(network.queues[0], natPacket)
			lastNatY = natPacket[1]
		}
	}
}

// Network will hold all 50 NIC computers
type Network struct {
	computers []*Intcode
	queues    [][][2]int // each element will be a packet for the same-index computer to handle
}

// Init sets up the 50 computers and queues
func (network *Network) Init(puzzleInput []int) {
	network.computers, network.queues = make([]*Intcode, 50), make([][][2]int, 50)
	for i := 0; i < 50; i++ {
		// Make and prime computer with its NIC number
		network.computers[i] = MakeComputer(puzzleInput)
		network.computers[i].Step(i)

		// setup queue for this NIC computer, slice of [2]int
		network.queues[i] = [][2]int{}
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
			// NOTE: making a big assumption that -2 will never be an input...
			// Note: changed the exit number to -2 because -1 is used in these computers for no-input/empty queues
			if input == -2 {
				return
			}

			// else recurse with a -1 to signal the initial input has been processed
			comp.PuzzleInput[param1] = input
			comp.InstructionIndex += 2

			// change the input value so the next time a 3 opcode is hit, will return out
			input = -2
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
