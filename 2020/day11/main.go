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

func part1(input string) int {
	grid := parseInput(input)

	var lastState string
	for {
		grid = step(grid, 4, true)

		// check if last grid matches current, break out if they match
		str := stringify(grid)
		if str == lastState {
			break
		}
		lastState = str
	}

	var ans int
	for _, row := range grid {
		for _, v := range row {
			if v == "#" {
				ans++
			}
		}
	}

	return ans
}

func part2(input string) int {
	grid := parseInput(input)

	var lastState string
	for {
		grid = step(grid, 5, false)

		// check if last grid matches current, break out if they match
		str := stringify(grid)
		if str == lastState {
			break
		}
		lastState = str
	}

	var ans int
	for _, row := range grid {
		for _, v := range row {
			if v == "#" {
				ans++
			}
		}
	}

	return ans
}

func parseInput(input string) [][]string {
	var ans [][]string

	lines := strings.Split(input, "\n")
	for _, l := range lines {
		ans = append(ans, strings.Split(l, ""))
	}

	return ans
}

// justNeighbors differentiates part 1 (true) from part 2
// tolerance = 4 for part 1, 5 for part 2
func step(grid [][]string, tolerance int, justNeighbors bool) [][]string {
	var nextGrid [][]string

	for r, row := range grid {
		nextGrid = append(nextGrid, make([]string, len(grid[0])))
		for c, v := range row {
			if v == "." {
				nextGrid[r][c] = "."
			} else {
				neighbors := countNeighbors(grid, r, c, justNeighbors)
				// check if seats should be updated
				if v == "L" && neighbors == 0 {
					nextGrid[r][c] = "#"
				} else if v == "#" && neighbors >= tolerance {
					nextGrid[r][c] = "L"
				} else {
					nextGrid[r][c] = v
				}
			}
		}

	}
	return nextGrid
}

var directions = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func countNeighbors(grid [][]string, row, col int, justNeighbors bool) int {
	var countNeighbors int
	for _, d := range directions {
		nextR, nextC := row, col
		for {
			nextR += d[0]
			nextC += d[1]
			if nextR < 0 || nextR >= len(grid) || nextC < 0 || nextC >= len(grid[0]) {
				break
			}
			if grid[nextR][nextC] == "L" {
				break
			}
			if grid[nextR][nextC] == "#" {
				countNeighbors++
				break
			}

			// break out after first pass if only checking immediate neighbors (part 1)
			if justNeighbors {
				break
			}
		}
	}

	return countNeighbors
}

// stringifies grid so it can be compared to its former state
func stringify(grid [][]string) string {
	var str string
	for _, row := range grid {
		for _, v := range row {

			str += v
		}
	}
	return str
}
