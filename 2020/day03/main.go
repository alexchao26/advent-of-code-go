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
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	grid := parseInput(input)
	return rideSlopes(grid, 3, 1)
}

func part2(input string) int {
	grid := parseInput(input)

	slopes := [][2]int{
		// right, down
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	ans := 1
	for _, slope := range slopes {
		ans *= rideSlopes(grid, slope[0], slope[1])
	}

	return ans
}

func parseInput(input string) (grid [][]string) {
	lines := strings.Split(input, "\n")

	grid = make([][]string, len(lines))
	for i, l := range lines {
		grid[i] = strings.Split(l, "")
	}

	return grid
}

func rideSlopes(grid [][]string, right, down int) int {
	var ans int

	for row, col := 0, 0; row < len(grid); row, col = row+down, col+right {
		if grid[row][col%len(grid[0])] == "#" {
			ans++
		}
	}

	return ans
}
