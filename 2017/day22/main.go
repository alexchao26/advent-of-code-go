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

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	infectedMap := newStateFromInput(input)
	mid := len(strings.Split(input, "\n")) / 2
	current := [2]int{mid, mid} // assume square inputs

	// direction starts at 0, i.e. facing up
	var dirIndex, countBursts int

	for i := 0; i < 10000; i++ {
		switch infectedMap[current] {
		case infected:
			// is infected, turn right
			dirIndex = (dirIndex + 1) % 4
			infectedMap[current] = clean
		case clean:
			// not infected, turn left
			dirIndex = (dirIndex + 3) % 4
			infectedMap[current] = infected

			countBursts++
		}
		// move forward
		current = [2]int{current[0] + dirs[dirIndex][0], current[1] + dirs[dirIndex][1]}
	}

	return countBursts
}

func part2(input string) int {
	state := newStateFromInput(input)
	mid := len(strings.Split(input, "\n")) / 2
	current := [2]int{mid, mid}

	var dirIndex, countBursts int

	for i := 0; i < 10000000; i++ {
		switch state[current] {
		case clean:
			dirIndex = (dirIndex + 3) % 4
			state[current] = weakened
		case weakened:
			// keep going in same direction
			state[current] = infected
			countBursts++
		case infected:
			dirIndex = (dirIndex + 1) % 4
			state[current] = flagged
		case flagged:
			dirIndex = (dirIndex + 2) % 4
			state[current] = clean
		default:
			panic(fmt.Sprintf("unhandled infection type %v", state[current]))
		}
		// move forward
		current = [2]int{current[0] + dirs[dirIndex][0], current[1] + dirs[dirIndex][1]}
	}

	return countBursts
}

var dirs = [][2]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

type infectedState int

const (
	clean infectedState = iota
	weakened
	infected
	flagged
)

// this probably makes some cool visuals if a 2D slice is used... but this is easier
func newStateFromInput(input string) map[[2]int]infectedState {
	ans := map[[2]int]infectedState{}
	for r, line := range strings.Split(input, "\n") {
		for c, v := range strings.Split(line, "") {
			if v == "#" {
				ans[[2]int{r, c}] = infected
			} else if v == "." {
				ans[[2]int{r, c}] = clean
			}
		}
	}
	return ans
}
