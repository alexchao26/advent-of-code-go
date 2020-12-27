package main

import (
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	ans := part1(util.ReadFile("./input.txt"))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	steps, stateRules := parseInput(input)

	// lazy, use a huge array and just start in the middle
	bigArray := make([]int, steps)
	index := steps / 2
	currentStateName := "A"

	for i := 0; i < steps; i++ {
		currentVal := bigArray[index]
		rulesToFollow := stateRules[currentStateName][currentVal]
		// write
		bigArray[index] = rulesToFollow.valToWrite
		if rulesToFollow.direction == "left" {
			index--
		} else {
			index++
		}
		currentStateName = rulesToFollow.nextState
	}

	return mathy.SumIntSlice(bigArray)
}

type ruleset struct {
	name       string // for debugging only
	valToWrite int
	direction  string
	nextState  string
}

// assume all programs start in state A for now, one less thing to parse...
func parseInput(input string) (steps int, states map[string][2]ruleset) {
	// a manual parse here would be faster...
	blocks := strings.Split(input, "\n\n")

	fmt.Sscanf(strings.Split(blocks[0], "\n")[1], "Perform a diagnostic checksum after %d steps.", &steps)

	states = map[string][2]ruleset{}
	for _, block := range blocks[1:] {
		lines := strings.Split(block, "\n")
		var stateName string
		fmt.Sscanf(lines[0], "In state %1s:", &stateName)

		rulesIfZero := ruleset{name: stateName}
		fmt.Sscanf(strings.Trim(lines[2], " -."), "Write the value %d", &rulesIfZero.valToWrite)
		fmt.Sscanf(strings.Trim(lines[3], " -."), "Move one slot to the %s", &rulesIfZero.direction)
		fmt.Sscanf(strings.Trim(lines[4], " -."), "Continue with state %1s", &rulesIfZero.nextState)

		rulesIfOne := ruleset{name: stateName}
		fmt.Sscanf(strings.Trim(lines[6], " -."), "Write the value %d", &rulesIfOne.valToWrite)
		fmt.Sscanf(strings.Trim(lines[7], " -."), "Move one slot to the %s", &rulesIfOne.direction)
		fmt.Sscanf(strings.Trim(lines[8], " -."), "Continue with state %1s", &rulesIfOne.nextState)

		states[stateName] = [2]ruleset{rulesIfZero, rulesIfOne}
	}

	return steps, states
}
