package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	// east cucumbers try to move first
	// then south cucumbers
	// individually they only move if the space in front is empty
	// EVEN if the guy in front is facing the same direction, don't move forward
	// ALSO wrap right->left, bottom->top

	// find first step when they stop moving

	var grid [][]string
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	for steps := 1; steps < math.MaxInt64; steps++ {
		hash := fmt.Sprint(grid)
		next := map[[2]int]string{}

		// traverse east
		for r := range grid {
			for c := range grid[0] {
				if grid[r][c] == ">" {
					nextC := (c + 1) % len(grid[0])
					if grid[r][nextC] == "." {
						next[[2]int{r, nextC}] = ">"
						next[[2]int{r, c}] = "."
					} else {
						// do not move forward
						next[[2]int{r, c}] = ">"
					}
				}
			}
		}

		var nextGrid [][]string
		for range grid {
			nextGrid = append(nextGrid, make([]string, len(grid[0])))
		}
		for c, v := range next {
			nextGrid[c[0]][c[1]] = v
		}
		for r := range grid {
			for c := range grid[0] {
				if nextGrid[r][c] == "" {
					nextGrid[r][c] = grid[r][c]
				}
			}
		}
		grid = nextGrid

		next = map[[2]int]string{}
		// traverse south
		for r := range grid {
			for c := range grid[0] {
				if grid[r][c] == "v" {
					nextR := (r + 1) % len(grid)
					if grid[nextR][c] == "." {
						next[[2]int{nextR, c}] = "v"
						next[[2]int{r, c}] = "."
					} else {
						// do not move forward
						next[[2]int{r, c}] = "v"
					}
				}
			}
		}

		nextGrid = [][]string{}
		for range grid {
			nextGrid = append(nextGrid, make([]string, len(grid[0])))
		}
		for c, v := range next {
			nextGrid[c[0]][c[1]] = v
		}
		for r := range grid {
			for c := range grid[0] {
				if nextGrid[r][c] == "" {
					nextGrid[r][c] = grid[r][c]
				}
			}
		}

		if hash == fmt.Sprint(nextGrid) {
			// nothing moved, return the STEP NUMBER IT IS
			return steps
		}
		grid = nextGrid
	}

	panic("should return from loop")
}
