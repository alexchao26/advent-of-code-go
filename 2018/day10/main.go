package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		part1(util.ReadFile("./input.txt"))
	} else {
		part2(util.ReadFile("./input.txt"))
	}
}

func part1(input string) {
	postions, velocities := parseInputs(input)

	steps := 0

	prints := 0
	hasPrinted := false
	for hasLoneIsland(postions) {
		move(postions, velocities)
		steps++
		fmt.Println("\nsteps run", steps)

		printable := printGrid(postions)
		if printable != "" {
			fmt.Println("Steps: ", steps)
			fmt.Println(printable)
			time.Sleep(time.Millisecond * 500)
			hasPrinted = true
		}
		// stop printing after things have collided
		if printable == "" && hasPrinted {
			return
		}
		fmt.Println(prints)
	}
}

func part2(input string) {
	// Note, reused part 1 with a print for the number of steps
	part1(input)
}

func parseInputs(input string) (positions [][2]int, velocities [][2]int) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		posX := strings.TrimSpace(l[10:16])
		posY := strings.TrimSpace(l[18:24])
		velX := strings.TrimSpace(l[36:38])
		velY := strings.TrimSpace(l[40:42])
		positions = append(positions, [2]int{mathutil.StrToInt(posX), mathutil.StrToInt(posY)})
		velocities = append(velocities, [2]int{mathutil.StrToInt(velX), mathutil.StrToInt(velY)})
	}

	return positions, velocities
}

// this didn't work unfortunately, letters like K do have "lone islands" along
// the diagonals
func hasLoneIsland(grid [][2]int) bool {
	coords := map[[2]int]bool{}
	// generate map
	for _, pos := range grid {
		coords[pos] = true
	}

	// iterate through coords again and check if each has a neighbor in the map
	delta := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, pos := range grid {
		hasNeighbor := false
		for _, d := range delta {
			pos[0] += d[0]
			pos[1] += d[1]

			if coords[pos] {
				hasNeighbor = true
			}

			pos[0] -= d[0]
			pos[1] -= d[1]
		}
		if !hasNeighbor {
			return true // there is a lone island
		}
	}
	return false
}

func move(positions [][2]int, velocities [][2]int) {
	for i := range positions {
		positions[i][0] += velocities[i][0]
		positions[i][1] += velocities[i][1]
	}
}

func printGrid(positions [][2]int) string {
	// get bounds
	left := math.MaxInt16
	right := -math.MaxInt16
	top := math.MaxInt16
	bottom := -math.MaxInt16

	coords := map[[2]int]bool{}

	for _, p := range positions {
		coords[p] = true

		if p[0] < top {
			top = p[0]
		}
		if p[0] > bottom {
			bottom = p[0]
		}
		if p[1] < left {
			left = p[1]
		}
		if p[1] > right {
			right = p[1]
		}
	}

	if right-left > 20 && bottom-top > 20 {
		return ""
	}

	ans := ""
	for row := top; row <= bottom; row++ {
		for col := left; col <= right; col++ {
			if coords[[2]int{row, col}] {
				ans += "0"
			} else {
				ans += " "
			}
		}
		ans += "\n"
	}
	return ans
}
