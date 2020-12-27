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

	ans := gameOfLight(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func gameOfLight(input string, part int) int {
	var grid [][]string
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}

	for i := 0; i < 100; i++ {
		grid = tick(grid)
		if part == 2 {
			grid[0][0] = "#"
			grid[0][len(grid[0])-1] = "#"
			grid[len(grid)-1][0] = "#"
			grid[len(grid)-1][len(grid[0])-1] = "#"
		}
	}

	var count int
	for _, row := range grid {
		for _, c := range row {
			if c == "#" {
				count++
			}
		}
	}

	return count
}

func tick(grid [][]string) [][]string {
	var nextGrid [][]string
	for r, row := range grid {
		nextGrid = append(nextGrid, make([]string, len(grid[0])))
		for c, cell := range row {
			var neighbors int
			for rDiff := -1; rDiff <= 1; rDiff++ {
				for cDiff := -1; cDiff <= 1; cDiff++ {
					if !(rDiff == 0 && cDiff == 0) {
						nextRow := r + rDiff
						nextCol := c + cDiff
						if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) &&
							grid[nextRow][nextCol] == "#" {
							neighbors++
						}
					}
				}
			}
			if cell == "#" && (neighbors == 2 || neighbors == 3) {
				nextGrid[r][c] = "#"
			} else if cell == "." && neighbors == 3 {
				nextGrid[r][c] = "#"
			} else {
				nextGrid[r][c] = "."
			}
		}
	}

	return nextGrid
}
