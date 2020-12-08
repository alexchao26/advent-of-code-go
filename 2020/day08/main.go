package main

import (
	"flag"
	"fmt"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	comp := newComputerFromInput(input)

	// keep track of all the indices of instructions that have run
	// if it has already been run, break
	ranInstructionsIndices := map[int]bool{}
	for {
		nextInst := comp.index
		if ranInstructionsIndices[nextInst] {
			break
		}
		ranInstructionsIndices[nextInst] = true

		comp.step()
	}

	return comp.accumulator
}

func part2(input string) int {
	comp := newComputerFromInput(input)

	// iterate through instruction indices
	for i := range comp.instructions {
		// make new computer each time
		newComputer := newComputerFromInput(input)

		// flip this index's instruction if a jmp or nop
		switch newComputer.instructions[i].operation {
		case "jmp":
			newComputer.instructions[i].operation = "nop"
		case "nop":
			newComputer.instructions[i].operation = "jmp"
		case "acc":
			continue
		}

		// run isInfiniteLoop check which returns final global value
		if ans, isLoop := isInfiniteLoop(newComputer); !isLoop {
			return ans
		}
	}

	// this should never be hit
	fmt.Println("ERROR: No terminating set of instructions found")
	return -1
}

func isInfiniteLoop(comp computer) (finalAccumulatorVal int, isLoop bool) {
	ranInstructionsIndices := map[int]bool{}
	for comp.index < len(comp.instructions) {
		nextInst := comp.index
		// is an infinite loop, return out
		if ranInstructionsIndices[nextInst] {
			return 0, true
		}
		ranInstructionsIndices[nextInst] = true

		comp.step()
	}

	// instructions finished, return final accumulator & indicate it was not an
	// infinite loop
	return comp.accumulator, false
}
