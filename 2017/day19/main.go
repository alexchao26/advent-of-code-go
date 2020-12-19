package main

import (
	"flag"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	visitedChars, steps := movePacket(util.ReadFile("./input.txt"))
	if part == 1 {
		fmt.Println("Output:", visitedChars)
	} else {
		fmt.Println("Output:", steps)
	}
}

var dirs = [4][2]int{
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
	{0, 1},  // right
}

var pathRegexp = regexp.MustCompile("^[-|+A-Z]$")
var capsRegexp = regexp.MustCompile("^[A-Z]$")

func movePacket(input string) (visitedChars string, steps int) {
	grid := parseInput(input)
	// finding starting point in first row
	var row, col int
	for c := 0; c < len(grid[0]); c++ {
		if grid[0][c] == "|" {
			col = c
			break
		}
	}

	// track which index in dirs is the current directions, start facing down
	var dirIndex int

	// include starting tile...
	steps = 1

	// basically an infinite loop...
	for i := 0; i < math.MaxInt64; i++ {
		inFrontVal := getNextValue(grid, row, col, dirs[dirIndex])

		if pathRegexp.MatchString(inFrontVal) {
			row += dirs[dirIndex][0]
			col += dirs[dirIndex][1]
			steps++

			// also check if it's a letter
			if capsRegexp.MatchString(inFrontVal) {
				visitedChars += inFrontVal
			}
		} else if inFrontVal == " " {
			// just try to turn right then left, assuming no 3 way intersections...
			dirIndex = (dirIndex + 1) % 4
			if pathRegexp.MatchString(getNextValue(grid, row, col, dirs[dirIndex])) {
				continue
			}

			// if right isn't a path, then try to turn left from original direction
			// i.e. add 2 more & mod
			dirIndex = (dirIndex + 2) % 4
			if pathRegexp.MatchString(getNextValue(grid, row, col, dirs[dirIndex])) {
				continue
			}

			// can't turn, break out of loop
			break
		} else {
			panic("unhandled char " + inFrontVal)
		}
	}

	return visitedChars, steps
}

func getNextValue(grid [][]string, row, col int, diff [2]int) string {
	inFrontRow := row + diff[0]
	inFrontCol := col + diff[1]
	// if not in bounds, just return a space
	if inFrontRow < 0 || inFrontRow >= len(grid) || inFrontCol < 0 || inFrontCol >= len(grid[0]) {
		return " "
	}
	return grid[inFrontRow][inFrontCol]
}

func parseInput(input string) [][]string {
	var grid [][]string

	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}
