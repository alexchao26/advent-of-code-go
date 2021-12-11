package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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

	ans := flashingOctopiLol(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func flashingOctopiLol(input string, part int) int {
	grid := parseInput(input)

	var flashed int

	adjacentDiffs := [][2]int{
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	steps := 100
	if part == 2 {
		// assume it'll run in less steps than 2^31-1...
		steps = math.MaxInt32
	}

	for s := 0; s < steps; s++ {
		// initial increment and store who will flash
		var queueToFlash [][2]int
		for i, row := range grid {
			for j := range row {
				grid[i][j]++
				if grid[i][j] > 9 {
					queueToFlash = append(queueToFlash, [2]int{i, j})
				}
			}
		}

		// map tracks who has flashed, bc it's a map it'll also help dedupe and its length will be
		// how many flashed on this step
		seen := map[[2]int]bool{}
		for len(queueToFlash) > 0 {
			front := queueToFlash[0]
			queueToFlash = queueToFlash[1:]
			if seen[front] {
				continue
			}
			seen[front] = true

			// increment neighbors
			for _, d := range adjacentDiffs {
				r, c := front[0]+d[0], front[1]+d[1]
				// check in bounds
				if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
					continue
				}
				grid[r][c]++
				// check if neighbor should flash, seen map will dedupe so don't safeguard here
				if grid[r][c] > 9 {
					queueToFlash = append(queueToFlash, [2]int{r, c})
				}
			}
		}

		// part1, track how many have flashed in total
		flashed += len(seen)
		// set them all to zero
		for c := range seen {
			grid[c[0]][c[1]] = 0
		}

		// if all octopi flashed, return the step number (1-indexed)
		if part == 2 && len(seen) == len(grid)*len(grid[0]) {
			return s + 1
		}
	}

	// for part 1
	return flashed
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, char := range strings.Split(line, "") {
			row = append(row, cast.ToInt(char))
		}
		ans = append(ans, row)
	}
	return ans
}
