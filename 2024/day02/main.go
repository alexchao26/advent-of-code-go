package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	grid := convertInput(input)
	validLevels := 0
	for _, level := range grid {
		if testLevel(level) {
			validLevels += 1
		}
	}
	return validLevels
}

func part2(input string) int {
	grid := convertInput(input)
	// tolerate one bad level...
	validLevels := 0
	for _, level := range grid {
		for i := range len(level) {
			newLevel := []int{}
			for j := range len(level) {
				if i != j {
					newLevel = append(newLevel, level[j])
				}
			}
			if testLevel(newLevel) {
				validLevels += 1
				break
			}
		}
	}

	return validLevels
}

func testLevel(level []int) bool {
	isIncreasing := level[1] > level[0]

	// The levels are either all increasing or all decreasing.
	// Any two adjacent levels differ by at least one and at most three.
	for i := 1; i < len(level); i++ {
		if isIncreasing && level[i] <= level[i-1] {
			return false
		} else if !isIncreasing && level[i] >= level[i-1] {
			return false
		}

		diff := mathy.AbsInt(level[i] - level[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func convertInput(input string) [][]int {
	grid := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		level := []int{}
		for _, n := range strings.Split(line, " ") {
			level = append(level, cast.ToInt(n))
		}
		grid = append(grid, level)
	}
	return grid
}
