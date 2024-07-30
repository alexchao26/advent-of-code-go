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

	ans := pipeMaze(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

var pipes = map[string][2][2]int{
	"|": {
		{-1, 0}, // up
		{1, 0},  // down
	},
	"-": {
		{0, -1}, // left
		{0, 1},  //right
	},
	"L": {
		{-1, 0}, // up
		{0, 1},  // right
	},
	"J": {
		{-1, 0}, // up
		{0, -1}, // left
	},
	"7": {
		{0, -1}, // left
		{1, 0},  // down
	},
	"F": {
		{0, 1}, // right
		{1, 0}, // down
	},
}

func pipeMaze(input string, part int) int {
	grid := parseInput(input)

	var r, c int
	for i, row := range grid {
		for j, val := range row {
			if val == "S" {
				r = i
				c = j
			}
		}
	}

	// need to find which directions are connected to the initial cell
	// could traverse along the four directions and see which two lead back to
	// the start... all pipes only connect two coordinates so it's actually
	// fairly easy to traverse without creating a huge search space

	// note this soultion was overkill for part 1
	// inputs are nice and exactly two adjacent cells that connect to S
	var loopCoords map[[2]int]bool
	for pipeType := range pipes {
		// assign the start square to a random pipeType, then see if it will loop
		grid[r][c] = pipeType

		seen := map[[2]int]bool{}
		toAnalyze := [][2]int{
			{r, c},
		}
		didLoop := false
		for len(toAnalyze) > 0 {
			coords := toAnalyze[0]
			if seen[coords] {
				didLoop = true
				break
			}
			seen[coords] = true
			toAnalyze = toAnalyze[1:]

			if diffs, ok := pipes[grid[coords[0]][coords[1]]]; ok {
				for _, diff := range diffs {
					nextRow, nextCol := coords[0]+diff[0], coords[1]+diff[1]
					if isInRange(grid, nextRow, nextCol) && !seen[[2]int{nextRow, nextCol}] {
						toAnalyze = append(toAnalyze, [2]int{nextRow, nextCol})
					}
				}
			}

		}
		if didLoop {
			loopCoords = seen
			break
		}
	}

	if part == 1 {
		return len(loopCoords) / 2
	}

	// part 2

	newGrid := make([][]string, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if loopCoords[[2]int{i, j}] {
				newGrid[i] = append(newGrid[i], grid[i][j])
			} else {
				newGrid[i] = append(newGrid[i], ".")
			}
		}
	}

	for _, r := range newGrid {
		fmt.Println(r)
	}

	return -1
}

func isInRange(grid [][]string, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col <= len(grid)
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
