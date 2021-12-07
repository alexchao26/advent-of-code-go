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

	ans := calcMinFuel(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func calcMinFuel(input string, part int) int {
	var startingPositions []int
	for _, val := range strings.Split(input, ",") {
		startingPositions = append(startingPositions, cast.ToInt(val))
	}

	// horizontal positions of each crab
	// limited fuel, so all horizontal positions need to match
	// part1: 1 fuel to move 1 space
	// part2: 1 fuel to move 1st space, 2 for 2nd, 3 for 3rd, etc...

	// find bounds to run loop through
	lowest, highest := startingPositions[0], startingPositions[0]
	for _, v := range startingPositions {
		if v < lowest {
			lowest = v
		}
		if v > highest {
			highest = v
		}
	}

	bestFuelCost := math.MaxInt64
	for finalIndex := lowest; finalIndex <= highest; finalIndex++ {
		// calculate diffs to all
		cost := 0
		for _, startIndex := range startingPositions {
			horizDiff := int(math.Abs(float64(startIndex - finalIndex)))
			if part == 1 {
				cost += horizDiff
			} else {
				cost += calcSummationFromOneToEnd(horizDiff)
			}
		}

		if cost < bestFuelCost {
			bestFuelCost = cost
		}
	}

	return bestFuelCost
}

func calcSummationFromOneToEnd(end int) int {
	// 1 2 3 4 5
	// (1 + 5) * 2.5
	ans := float64(end+1) * float64(end) / 2
	return int(ans)
}
