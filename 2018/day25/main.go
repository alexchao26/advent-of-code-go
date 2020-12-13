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
		fmt.Println("Output:", ans)
	} else {
		panic("no part 2 :P")
	}
}

func part1(input string) int {
	allCoords := parseInput(input)

	constellations := [][][4]int{}

	for _, iterCoord := range allCoords {
		// check which constellations this coordinate is in range of
		indicesWithinRange := []int{}
		for constIndex, constellation := range constellations {
			for _, constCoord := range constellation {
				if manhattanDistance(constCoord, iterCoord) <= 3 {
					indicesWithinRange = append(indicesWithinRange, constIndex)
					break
				}
			}
		}

		// if no existing constellations are within range, append a new one with
		// just this current coordinate
		if len(indicesWithinRange) == 0 {
			constellations = append(constellations, [][4]int{iterCoord})
		} else {
			// otherwise merge all constellations together (into the first one)
			// add the current node
			// then remove any other constellations that were merged into #1
			firstIndex := indicesWithinRange[0]
			for i, indexToMerge := range indicesWithinRange {
				if i != 0 {
					current := constellations[indexToMerge]
					constellations[firstIndex] = append(constellations[firstIndex], current...)
				}
			}
			constellations[firstIndex] = append(constellations[firstIndex], iterCoord)

			// remove all but the first constellation, in reverse order b/c
			// using this method for removal affects the end of the slice, and
			// in some cases that may be the element that we want to remove
			for i := len(indicesWithinRange) - 1; i > 0; i-- {
				constIndex := indicesWithinRange[i]
				lastIndex := len(constellations) - 1
				constellations[constIndex] = constellations[lastIndex]
				constellations = constellations[:lastIndex]
			}
		}
	}

	return len(constellations)
}

func parseInput(input string) [][4]int {
	lines := strings.Split(input, "\n")
	allCoords := make([][4]int, len(lines))
	for i, l := range lines {
		newCoord := [4]int{}
		fmt.Sscanf(l, "%d,%d,%d,%d", &newCoord[0], &newCoord[1], &newCoord[2], &newCoord[3])

		allCoords[i] = newCoord
	}

	return allCoords
}

func manhattanDistance(one, two [4]int) int {
	var sum int
	for i := range one {
		diff := one[i] - two[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}
	return sum
}
