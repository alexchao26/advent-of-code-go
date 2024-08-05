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
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	grid := parseInput(input)

	// tilt north, all O's roll to the top or to the next #
	tiltNorth(grid)

	// then calculate total load (total rows) - (row index) per rock
	ans := 0
	for r, row := range grid {
		for _, val := range row {
			if val == "O" {
				ans += len(grid) - r
			}
		}
	}

	return ans
}

func part2(input string) int {
	grid := parseInput(input)

	seenStates := map[string]int{}

	cycles := 1000000000
	for c := 0; c < cycles; c++ {
		key := stringifyStringGrid(grid)
		if lastIndex, ok := seenStates[key]; ok {
			cyclePeriod := c - lastIndex
			for c+cyclePeriod < cycles {
				c += cyclePeriod
			}
		}
		seenStates[key] = c

		// 1 cycle = tilt N, W, S, E
		tiltNorth(grid)
		tiltWest(grid)
		tiltSouth(grid)
		tiltEast(grid)
	}

	ans := 0
	for r, row := range grid {
		for _, val := range row {
			if val == "O" {
				ans += len(grid) - r
			}
		}
	}

	// 99841 too low
	return ans
}

func tiltNorth(grid [][]string) {
	for r, row := range grid {
		for c, val := range row {
			if val == "O" {
				for nextRow := r - 1; nextRow >= 0; nextRow-- {
					// can only fall north if nextRow is an empty space
					if grid[nextRow][c] == "." {
						grid[nextRow][c] = "O"
						grid[nextRow+1][c] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltSouth(grid [][]string) {
	for r := len(grid) - 1; r >= 0; r-- {
		for c := range len(grid[0]) {
			val := grid[r][c]
			if val == "O" {
				for nextRow := r + 1; nextRow < len(grid); nextRow++ {
					// can only fall north if nextRow is an empty space
					if grid[nextRow][c] == "." {
						grid[nextRow][c] = "O"
						grid[nextRow-1][c] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltEast(grid [][]string) {
	for c := len(grid[0]) - 1; c >= 0; c-- {
		for r := range grid {
			val := grid[r][c]

			if val == "O" {
				for nextCol := c + 1; nextCol < len(grid[0]); nextCol++ {
					// can only fall north if nextCol is an empty space
					if grid[r][nextCol] == "." {
						grid[r][nextCol] = "O"
						grid[r][nextCol-1] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltWest(grid [][]string) {
	for c := range len(grid[0]) {
		for r := range grid {
			val := grid[r][c]

			if val == "O" {
				for nextCol := c - 1; nextCol >= 0; nextCol-- {
					// can only fall north if nextCol is an empty space
					if grid[r][nextCol] == "." {
						grid[r][nextCol] = "O"
						grid[r][nextCol+1] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func stringifyStringGrid(grid [][]string) string {
	ans := ""
	for _, row := range grid {
		ans += strings.Join(row, "")
	}
	return ans
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
