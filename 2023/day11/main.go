package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
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
		ans := part1(input, 2)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part1(input, 1000000)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string, expansionFactor int) int {
	grid := parseInput(input)

	// record which rows and cols are empty first
	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}

	for r := 0; r < len(grid); r++ {
		galaxyFound := false
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] != "." {
				galaxyFound = true
				break
			}
		}
		if !galaxyFound {
			emptyRows[r] = true
		}
	}
	for c := 0; c < len(grid[0]); c++ {
		galaxyFound := false
		for r := 0; r < len(grid); r++ {
			if grid[r][c] != "." {
				galaxyFound = true
				break
			}
		}
		if !galaxyFound {
			emptyCols[c] = true
		}
	}

	// traverse grid and calculate coordinates of each galaxy while accumulating the expanded rows/cols
	// that can be added into the galaxy's coordinates as they're found
	galaxyCoords := map[int][2]int{}
	expandedRowsToAdd := 0
	for r := 0; r < len(grid); r++ {
		if emptyRows[r] {
			expandedRowsToAdd += expansionFactor - 1
			continue
		}

		expendedColsToAdd := 0
		for c := 0; c < len(grid[0]); c++ {
			if emptyCols[c] {
				expendedColsToAdd += expansionFactor - 1
				continue
			}

			if grid[r][c] == "#" {
				galaxyCoords[len(galaxyCoords)] = [2]int{
					r + expandedRowsToAdd,
					c + expendedColsToAdd,
				}
			}
		}
	}

	// shortest distance is basically manhattan distance, helper function handles absolute values
	totalDistance := 0
	for i := 0; i < len(galaxyCoords); i++ {
		for j := i + 1; j < len(galaxyCoords); j++ {
			g1, g2 := galaxyCoords[i], galaxyCoords[j]
			totalDistance += mathy.ManhattanDistance(g1[0], g1[1], g2[0], g2[1])
		}
	}

	return totalDistance
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
