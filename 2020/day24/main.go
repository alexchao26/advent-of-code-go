package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := part1(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

// 1240/668
func part1(input string, part int) int {
	instructions := strings.Split(input, "\n")
	tileIsBlack := map[[6]int]bool{}
	for _, inst := range instructions {
		// tally the steps taken in each of the hex directions
		var tallyDirections [6]int

		// iterate thorugh the NONDELIMITED instructions
		for i := 0; i < len(inst); {
			// set the direction to the next two characters if possible
			dir := string(inst[i])
			if i+2 <= len(inst) {
				dir = inst[i : i+2]
			}

			// if the two character direction is se, sw, nw or ne, both characters
			// must be used, because using one is invalid ("n" or "s")
			switch dir {
			case "se", "sw", "nw", "ne":
				tallyDirections[dirIndices[dir]]++
				i += 2
			default:
				tallyDirections[dirIndices[dir[0:1]]]++
				i++
			}
			// collapse and directions that cancel out, such as sw/nw -> w
			tallyDirections = zeroOutHexDirections(tallyDirections)
		}
		// flip that tile
		tileIsBlack[tallyDirections] = !tileIsBlack[tallyDirections]
	}

	// for part 2, play game of life 100 times
	if part == 2 {
		for i := 0; i < 100; i++ {
			// flip based on neighbors
			nextState := map[[6]int]bool{}

			// collect all coordinates to check
			toCheck := map[[6]int]bool{}
			for i := 0; i < 6; i++ {
				for k := range tileIsBlack {
					k[i]++
					toCheck[zeroOutHexDirections(k)] = true
				}
			}

			for coord := range toCheck {
				// count neighbors
				var neighbors int
				for i := 0; i < 6; i++ {
					clone := coord // don't want to modify the original coord
					clone[i]++     // generates the six directions around coord
					clone = zeroOutHexDirections(clone)
					if tileIsBlack[clone] {
						neighbors++
					}
				}
				// flipping logic:
				// back with zero or more than 2 neighbors becomes white
				if tileIsBlack[coord] && (neighbors == 0 || neighbors > 2) {
					nextState[coord] = false
				} else if !tileIsBlack[coord] && neighbors == 2 {
					// white with exactly 2 neighbors becomes black
					nextState[coord] = true
				} else {
					// stays the same
					nextState[coord] = tileIsBlack[coord]
				}
			}
			tileIsBlack = nextState
		}
	}

	var count int
	for _, b := range tileIsBlack {
		if b {
			count++
		}
	}

	return count

}

var dirIndices = map[string]int{
	"e":  0,
	"se": 1,
	"sw": 2,
	"w":  3,
	"nw": 4,
	"ne": 5,
}

// borrowed from my 2017 day 11 code which calculated hex coordinate manhattan distances
func zeroOutHexDirections(tally [6]int) [6]int {
	// zero out opposite indices
	for i := range tally {
		if tally[i] != 0 {
			oppositeIndex := (i + 3) % 6
			smaller := mathutil.MinInt(tally[oppositeIndex], tally[i])
			tally[oppositeIndex] -= smaller
			tally[i] -= smaller
		}
	}

	// handle neighbors which collapse into the current direction
	// e.g. sw,se == s
	for i := range tally {
		toLeft := (i + 5) % 6
		toRight := (i + 1) % 6
		if tally[toLeft] > 0 && tally[toRight] > 0 {
			smaller := mathutil.MinInt(tally[toLeft], tally[toRight])
			tally[toLeft] -= smaller
			tally[toRight] -= smaller
			tally[i] += smaller
		}
	}

	return tally
}
