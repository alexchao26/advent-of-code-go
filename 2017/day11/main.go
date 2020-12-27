package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := hexEd(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

var dirIndices = map[string]int{
	"n":  0,
	"ne": 1,
	"se": 2,
	"s":  3,
	"sw": 4,
	"nw": 5,
}

func hexEd(input string, part int) int {
	steps := strings.Split(input, ",")
	tallyDirections := make([]int, 6)
	var furthest int // for part 2
	for _, step := range steps {
		tallyDirections[dirIndices[step]]++
		distanceFromStart := getDistanceFromOrigin(tallyDirections)
		furthest = mathy.MaxInt(furthest, distanceFromStart)
	}

	if part == 1 {
		return getDistanceFromOrigin(tallyDirections)
	}
	return furthest
}

func getDistanceFromOrigin(tally []int) int {
	// zero out opposite indices, after this, there will be at most 3 positive
	// values in the slice
	for i := range tally {
		if tally[i] != 0 {
			oppositeIndex := (i + 3) % 6
			smaller := mathy.MinInt(tally[oppositeIndex], tally[i])
			tally[oppositeIndex] -= smaller
			tally[i] -= smaller
		}
	}

	// handle neighbors, which collapse into the current direction
	// e.g. sw,se == s
	for i := range tally {
		toLeft := (i + 5) % 6
		toRight := (i + 1) % 6
		if tally[toLeft] > 0 && tally[toRight] > 0 {
			smaller := mathy.MinInt(tally[toLeft], tally[toRight])
			tally[toLeft] -= smaller
			tally[toRight] -= smaller
			tally[i] += smaller
		}
	}

	distanceFromOrigin := mathy.SumIntSlice(tally)

	return distanceFromOrigin
}
