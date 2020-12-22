package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algos"
	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = fractalArt(util.ReadFile("./input.txt"), 5)
	} else {
		ans = fractalArt(util.ReadFile("./input.txt"), 18)
	}
	fmt.Println("Output:", ans)
}

var startingPattern = `.#.
..#
###`

func fractalArt(input string, rounds int) int {
	var state [][]string
	for _, line := range strings.Split(startingPattern, "\n") {
		state = append(state, strings.Split(line, ""))
	}

	rules := parseInput(input)

	for i := 0; i < rounds; i++ {
		state = tick(state, rules)
	}

	var count int
	for _, row := range state {
		for _, v := range row {
			if v == "#" {
				count++
			}
		}
	}
	return count
}

func parseInput(input string) map[string][][]string {
	// some helper functions for generating the rules map
	// need to parse the left sides of the enhancement rules
	//   then helper functions that rotate them (util.RotateStringGrid()) and
	//   one to mirror image it
	makeGridFromString := func(str string) [][]string {
		var grid [][]string
		for _, line := range strings.Split(str, "/") {
			grid = append(grid, strings.Split(line, ""))
		}
		return grid
	}
	stringifyGrid := func(grid [][]string) (str string) {
		for _, row := range grid {
			for _, v := range row {
				str += v
			}
		}
		return str
	}

	rules := map[string][][]string{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " => ")
		keyGrid := makeGridFromString(parts[0])
		resultGrid := makeGridFromString(parts[1])

		for i := 0; i < 4; i++ {
			keyGrid = algos.RotateStringGrid(keyGrid)
			rules[stringifyGrid(keyGrid)] = resultGrid
			rules[stringifyGrid((algos.MirrorStringGrid(keyGrid)))] = resultGrid
		}
	}
	return rules
}

func tick(grid [][]string, rules map[string][][]string) [][]string {
	var nextState [][]string

	// determine the size of break up the grid by. prioritize 2x2 grids
	var edgeSize int
	if len(grid)%2 == 0 {
		edgeSize = 2
	} else if len(grid)%3 == 0 {
		edgeSize = 3
	} else {
		panic("grid is not evenly divisible by 2 or 3, got " + cast.ToString(len(grid)))
	}

	// iterate over like a sudoku grid, r and c iterate over the top left corner
	// of each sub-square
	for r := 0; r < len(grid); r += edgeSize {
		// a new row of sub-squares is being iterated over, add edgeSize+1 number
		// of empty slices onto the nextState grid
		for i := 0; i < edgeSize+1; i++ {
			nextState = append(nextState, []string{})
		}
		for c := 0; c < len(grid[0]); c += edgeSize {
			// generate the string to match a key in the rules map
			var strToMatch string
			for i := 0; i < edgeSize; i++ {
				for j := 0; j < edgeSize; j++ {
					// r+i and c+j point at coords within the original grid
					strToMatch += grid[r+i][c+j]
				}
			}

			// finding the result of the enhancement rule for the string to match
			resulting, ok := rules[strToMatch]
			if !ok {
				panic("No matching pattern found for " + strToMatch)
			}

			// append the values from the result onto the appropriate nextState row
			for i, vals := range resulting {
				nextStateIndex := len(nextState) - edgeSize - 1 + i
				nextState[nextStateIndex] = append(nextState[nextStateIndex], vals...)
			}
		}
	}

	return nextState
}
