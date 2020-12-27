package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := taxicab(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

var dirs = [][2]int{
	{-1, 0}, // north
	{0, 1},  // east
	{1, 0},  // south
	{0, -1}, // west
}

func taxicab(input string, part int) int {
	var dirIndex int // start facing north
	var row, col int

	visited := map[[2]int]bool{
		{0, 0}: true,
	}

	for _, inst := range strings.Split(input, ", ") {
		var turn string
		var steps int
		fmt.Sscanf(inst, "%1s%d", &turn, &steps)
		if turn == "R" {
			dirIndex = (dirIndex + 1) % 4
		} else if turn == "L" {
			dirIndex = (dirIndex + 3) % 4
		} else {
			panic("unhandled turning direction " + turn)
		}

		for i := 0; i < steps; i++ {
			// move forward one step at a time
			row += dirs[dirIndex][0]
			col += dirs[dirIndex][1]
			if visited[[2]int{row, col}] && part == 2 {
				return mathy.ManhattanDistance(0, 0, row, col)
			}
			visited[[2]int{row, col}] = true
		}
	}

	return mathy.ManhattanDistance(0, 0, row, col)
}
