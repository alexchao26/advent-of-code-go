package main

import (
	"flag"
	"fmt"
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

// 300x300 grid
func part1(input string) string {
	gridSN := parseInputs(input)

	grid := generateGrid(gridSN)

	// find best 3x3 grid
	deltas := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	var best int
	var topLeftCorner [2]int
	for x := 2; x < 300; x++ {
		for y := 2; y < 300; y++ {
			sum := 0
			for _, d := range deltas {
				sum += grid[x+d[0]][y+d[1]]
			}
			if sum > best {
				best = sum
				topLeftCorner = [2]int{x - 1, y - 1}
			}
		}
	}

	return fmt.Sprintf("%d,%d", topLeftCorner[0], topLeftCorner[1])
}

func part2(input string) string {
	gridSN := parseInputs(input)

	grid := generateGrid(gridSN)

	var bestPower int
	var xYSize [3]int

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			var sum int
			for edge := 0; edge+x <= 300 && edge+y <= 300; edge++ {
				sum += grid[x+edge][y+edge]
				for add := 0; add < edge; add++ {
					sum += grid[x+add][y+edge]
					sum += grid[x+edge][y+add]
				}
				if sum > bestPower {
					bestPower = sum
					xYSize = [3]int{x, y, edge + 1}
				}
			}
		}
	}

	return fmt.Sprintf("%d,%d,%d", xYSize[0], xYSize[1], xYSize[2])
}

func parseInputs(input string) int {
	return util.StrToInt(strings.TrimSpace(input))
}

func generateGrid(gridSN int) [][]int {
	oneIndexedGrid := make([][]int, 301) // X, Y

	for i := 1; i <= 300; i++ {
		oneIndexedGrid[i] = make([]int, 301)
	}

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			rackID := x + 10
			power := rackID * y
			power += gridSN
			power *= rackID
			// keep hundreds
			power /= 100
			power %= 10

			power -= 5
			oneIndexedGrid[x][y] = power
		}
	}

	return oneIndexedGrid
}
