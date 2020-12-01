package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

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
	initialZeroIndex := 25 // lazy placement of empty pots to left and right of inputs
	state, changesMap := parseInputs(input, initialZeroIndex)

	// fmt.Printf("%d\t%v\n", 0, state)

	for i := 0; i < 20; i++ {
		state = step(state, changesMap)
		// fmt.Printf("%d\t%v\n", i+1, state)
	}

	ans := sumOfPotNumbers(state, initialZeroIndex)

	return ans
}

func part2(input string) int {
	initialZeroIndex := 300 // lazy placement of empty pots to left and right of inputs
	state, changesMap := parseInputs(input, initialZeroIndex)

	patterns := map[string]int{}
	sums := []int{sumOfPotNumbers(state, initialZeroIndex)}

	var patternIndices [2]int

	for gens := 1; gens < initialZeroIndex-100; gens++ {
		state = step(state, changesMap)
		// fmt.Printf("%d\t%v\n", gens, state)

		currentSum := sumOfPotNumbers(state, initialZeroIndex)
		sums = append(sums, currentSum)

		trimmedPattern := strings.Trim(stringify(state), ".")
		// fmt.Printf("%d\t%s\n", gens, trimmedPattern)

		if lastIndex, ok := patterns[trimmedPattern]; ok {
			patternIndices = [2]int{lastIndex, gens}
			break // break once a pattern is found
		} else {
			// store pattern to index so the pattern frequency can be found
			patterns[trimmedPattern] = gens
		}
	}

	// calc to 50000000000
	// find the frequency and the sum's diff, then add that diff * number of generations left
	freq := patternIndices[1] - patternIndices[0]
	patternDiff := sums[patternIndices[1]] - sums[patternIndices[0]]
	if freq != 1 {
		log.Fatal("Pattern frequency is assumed to be 1, part2() needs to be updated to handle != 1 cases")
	}

	fiveBillion := 50000000000
	ans := sums[patternIndices[1]] + (fiveBillion-patternIndices[1])*patternDiff

	return ans
}

func parseInputs(input string, initialZeroIndex int) (state []string, changesMap map[string]string) {
	lines := strings.Split(input, "\n")
	first := strings.Split(lines[0], "state: ")

	// increment it by 3 to account for empty nodes at start
	for i := 0; i < initialZeroIndex; i++ {
		state = append(state, ".")
	}

	for _, val := range first[1] {
		state = append(state, string(val))
	}

	// add 3 onto end also
	for i := 0; i < initialZeroIndex; i++ {
		state = append(state, ".")
	}

	changesMap = make(map[string]string)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line != "" {
			splitStep := strings.Split(line, " => ")
			changesMap[splitStep[0]] = splitStep[1]
		}
	}

	return state, changesMap
}

func step(state []string, changesMap map[string]string) (nextState []string) {
	for i := range state {
		fiveStr := ""
		for index := i - 2; index >= 0 && index < len(state) && len(fiveStr) < 5; index++ {
			fiveStr += state[index]
		}

		if v, ok := changesMap[fiveStr]; ok {
			nextState = append(nextState, v)
		} else {
			// missing case should only apply to examples or len < 5
			nextState = append(nextState, ".")
		}
	}

	return nextState
}

func sumOfPotNumbers(state []string, zeroIndex int) int {
	ans := 0
	for i, v := range state {
		if v == "#" {
			ans += i - zeroIndex
		}
	}
	return ans
}

func stringify(state []string) string {
	ans := ""
	for _, v := range state {
		ans += v
	}
	return ans
}
