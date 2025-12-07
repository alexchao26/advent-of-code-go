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

	ans := laboratories(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func laboratories(input string, part int) int {
	grid := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	cols := map[int]int{}
	for c := range len(grid[0]) {
		if grid[0][c] == "S" {
			cols[c] = 1
		}
	}

	splits := 0 // for part 1

	for _, line := range grid {
		nextCols := map[int]int{}
		for i, val := range line {
			if val == "^" && cols[i] > 0 {
				nextCols[i-1] += cols[i]
				nextCols[i+1] += cols[i]
				delete(cols, i) // blocked = just delete it
				splits++
			}
		}

		for c, val := range nextCols {
			cols[c] += val
		}
	}

	if part == 1 {
		return splits
	}

	ans := 0
	for _, val := range cols {
		ans += val
	}
	return ans
}
