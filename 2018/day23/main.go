package main

import (
	"flag"
	"fmt"
	"math"
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
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	bots := parseInput(input)

	strongestBot := bots[0]
	for _, b := range bots {
		if b.strength > strongestBot.strength {
			strongestBot = b
		}
	}

	var withinRange int
	for _, b := range bots {
		if manhattanDist(b.coords, strongestBot.coords) <= strongestBot.strength {
			withinRange++
		}
	}

	return withinRange
}

func part2(input string) int {
	bots := parseInput(input)

	// get the bounds of the cube that all bots are inside of
	// the answer coordinate will be within this space
	var minCoord, maxCoord [3]int
	for i := range minCoord {
		minCoord[i] = math.MaxInt16
	}
	for _, b := range bots {
		for i := 0; i < 3; i++ {
			if minCoord[i] > b.coords[i] {
				minCoord[i] = b.coords[i]
			}
			if maxCoord[i] < b.coords[i] {
				maxCoord[i] = b.coords[i]
			}
		}
	}

	var origin [3]int

	// width of the box
	boxWidth := maxCoord[0] - minCoord[0]

	// 1. width is used to calculate the eight corners to check.
	// 2. the reachable bots are counted from each corner
	// 3. on each iteration, the box is centered around the best corner and the
	//    width is cut in half until it reaches zero
	var bestGrid [3]int
	for boxWidth > 0 {
		var maxCount int

		for x := minCoord[0]; x < maxCoord[0]+1; x += boxWidth {
			for y := minCoord[1]; y < maxCoord[1]+1; y += boxWidth {
				for z := minCoord[2]; z < maxCoord[2]+1; z += boxWidth {
					current := [3]int{x, y, z}
					var countInRange int
					for _, b := range bots {
						if b.canReach(current) {
							countInRange++
						}
					}
					if maxCount < countInRange ||
						(maxCount == countInRange && manhattanDist(bestGrid, origin) > manhattanDist(current, origin)) {
						maxCount = countInRange
						bestGrid = current
					}
				}
			}
		}

		// adjust box size, i.e. min and max coords
		for i := 0; i < 3; i++ {
			minCoord[i] = bestGrid[i] - boxWidth
			maxCoord[i] = bestGrid[i] + boxWidth
		}

		// shrink searchable box size
		boxWidth /= 2
	}

	return manhattanDist(bestGrid, origin)
}

type nanobot struct {
	coords   [3]int
	strength int
}

func (b nanobot) canReach(coord [3]int) bool {
	return manhattanDist(coord, b.coords) <= b.strength
}

func parseInput(input string) []nanobot {
	var bots []nanobot

	lines := strings.Split(input, "\n")
	for _, l := range lines {
		bot := nanobot{}
		_, err := fmt.Sscanf(l, "pos=<%d,%d,%d>, r=%d", &bot.coords[0], &bot.coords[1], &bot.coords[2], &bot.strength)
		if err != nil {
			panic("parsing input line " + err.Error())
		}
		bots = append(bots, bot)
	}

	return bots
}

func manhattanDist(one, two [3]int) int {
	var dist int
	for i := range one {
		diff := one[i] - two[i]
		if diff < 0 {
			diff *= -1
		}
		dist += diff
	}
	return dist
}
