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

func parseInput(input string) (grid [][]bool) {
	lines := strings.Split(input, "\n")

	grid = make([][]bool, len(lines))
	for i, l := range lines {
		grid[i] = make([]bool, len(l))
		for j, v := range l {
			if v == '#' {
				grid[i][j] = true
			}
		}
	}

	return grid
}

func rideSlopes(grid [][]bool, right, down int) int {
	var row, col int
	var ans int

	for {
		row += down
		col += right
		col %= len(grid[0])

		if row < len(grid) {
			if grid[row][col] {
				ans++
			}
		} else {
			break
		}
	}

	return ans
}
