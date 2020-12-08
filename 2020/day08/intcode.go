package main

import (
	"fmt"
	"strings"
)

func newComputerFromInput(input string) computer {
	var instructions []instruction

	for _, l := range strings.Split(input, "\n") {
		inst := instruction{}
		fmt.Sscanf(l, "%s %d", &inst.operation, &inst.argument)
		instructions = append(instructions, inst)
	}

	return computer{instructions: instructions}
}

type computer struct {
	instructions []instruction
	index        int
	accumulator  int
}

type instruction struct {
	operation string
	argument  int
}

func (c *computer) step() {
	switch inst := c.instructions[c.index]; inst.operation {
	case "acc":
		c.accumulator += inst.argument
		c.index++
	case "jmp":
		c.index += inst.argument
	case "nop":
		c.index++
	default:
		panic("unhandled operation type" + inst.operation)
	}
}

// func isInfiniteLoop(comp computer) (finalAccumulatorVal int, isLoop bool) {
// 	ranInstructionsIndices := map[int]bool{}
// 	for comp.index < len(comp.instructions) {
// 		nextInst := comp.index
// 		// is an infinite loop, return out
// 		if ranInstructionsIndices[nextInst] {
// 			return 0, true
// 		}
// 		ranInstructionsIndices[nextInst] = true

// 		comp.step()
// 	}

// 	// instructions finished, return final accumulator & indicate it was not an
// 	// infinite loop
// 	return comp.accumulator, false
// }
