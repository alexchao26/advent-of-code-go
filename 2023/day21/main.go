package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input, 64)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input, 26501365)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string, steps int) int {
	grid := parseInput(input)
	var row, col int
	for r, rowSlice := range grid {
		for c, val := range rowSlice {
			if val == "S" {
				row = r
				col = c
				break
			}
		}
	}

	queue := map[[2]int]bool{
		{row, col}: true,
	}
	for i := 0; i < steps; i++ {
		newQueue := map[[2]int]bool{}
		for coord := range queue {
			for _, diff := range [][2]int{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			} {
				nextRow := coord[0] + diff[0]
				nextCol := coord[1] + diff[1]
				if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
					continue
				}

				if grid[nextRow][nextCol] == "." || grid[nextRow][nextCol] == "S" {
					newQueue[[2]int{nextRow, nextCol}] = true
				}
			}
		}

		queue = newQueue
	}

	return len(queue)
}

func part2(input string, steps int) int {
	grid := parseInput(input)
	var row, col int
	for r, rowSlice := range grid {
		for c, val := range rowSlice {
			if val == "S" {
				row = r
				col = c
				break
			}
		}
	}
	grid[row][col] = "."

	// keeps track of the flip-flopping coords separately
	evenSeenCoords := map[[2]int]bool{}
	oddSeenCoords := map[[2]int]bool{}

	// need a set of all coords added to the queue so that we're not re-adding the same coords
	uniqueCoords := map[[2]int]bool{}

	queue := [][2]int{
		{row, col},
	}

	// results to calculate quadratic constants with
	results := []int{}

	// perform two steps at once to always be on an even number of steps
	for s := 0; s < steps && len(results) < 3; s++ {
		activeSeenCoords := evenSeenCoords
		if s%2 == 1 {
			activeSeenCoords = oddSeenCoords
		}

		newQueue := [][2]int{}
		for _, coord := range queue {
			activeSeenCoords[coord] = true

			for _, diff := range [][2]int{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			} {
				nextRow := coord[0] + diff[0]
				nextCol := coord[1] + diff[1]
				nextCoord := [2]int{nextRow, nextCol}

				// handles infinite grid and garden space detection
				modNextRow := ((nextRow % len(grid)) + len(grid)) % +len(grid)
				modNextCol := ((nextCol % len(grid[0])) + len(grid[0])) % len(grid[0])
				if grid[modNextRow][modNextCol] != "." {
					continue
				}

				// if already seen, skip
				if uniqueCoords[nextCoord] {
					continue
				}
				uniqueCoords[nextCoord] = true

				newQueue = append(newQueue, nextCoord)
			}
		}

		queue = newQueue

		if s != 0 && s%131 == 65 {
			results = append(results, len(activeSeenCoords))
		}
	}

	// solve quadratic for a b and c constants
	a := (results[2] + results[0] - 2*results[1]) / 2
	b := results[1] - results[0] - a
	c := results[0]

	n := steps / len(grid)

	return a*n*n + b*n + c
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
