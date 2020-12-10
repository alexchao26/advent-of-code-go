package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
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
		ans := part2(util.ReadFile("./input.txt"), 10000)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	coords := parseInputCoords(input)

	boundLeft, boundRight, boundTop, boundBottom := getBounds(coords)

	// maps coordinates to the number (count) of cells that they are closest too
	coordsToMinDistanceCounts := map[[2]int]int{}

	for r := boundTop; r <= boundBottom; r++ {
		for c := boundLeft; c <= boundRight; c++ {
			bestManhattan := math.MaxInt16
			var coordsToBest [2]int
			distCounts := map[int]int{} // dedeupe equidistant cells
			for _, coord := range coords {
				man := mathutil.ManhattanDistance(r, c, coord[0], coord[1])
				if man <= bestManhattan {
					bestManhattan = man
					coordsToBest = coord
					distCounts[bestManhattan]++
				}
			}

			// do not increment anything if there were two equidistant coords
			if distCounts[bestManhattan] == 1 {
				coordsToMinDistanceCounts[coordsToBest]++
			}
		}
	}

	var largest int
	for coord, val := range coordsToMinDistanceCounts {
		// exclude edges
		if coord[0] == boundTop || coord[0] == boundBottom {
			continue
		}
		if coord[1] == boundLeft || coord[1] == boundRight {
			continue
		}

		if val > largest {
			largest = val
		}
	}
	return largest
}

func part2(input string, dist int) int {
	coords := parseInputCoords(input)

	boundLeft, boundRight, boundTop, boundBottom := getBounds(coords)

	coordsToTotalDist := map[[2]int]int{}
	var area int
	for r := boundTop; r <= boundBottom; r++ {
		for c := boundLeft; c <= boundRight; c++ {
			point := [2]int{r, c}
			for _, coord := range coords {
				coordsToTotalDist[point] += mathutil.ManhattanDistance(point[0], point[1], coord[0], coord[1])
			}
			if coordsToTotalDist[point] < dist {
				area++
			}
		}
	}

	return area
}

func parseInputCoords(input string) [][2]int {
	lines := strings.Split(input, "\n")
	coords := [][2]int{} // [row, col]
	for _, l := range lines {
		c := strings.Split(l, ", ")
		if len(c) == 2 {
			coords = append(coords, [2]int{
				mathutil.StrToInt(c[0]),
				mathutil.StrToInt(c[1]),
			})
		}
	}
	return coords
}

func getBounds(coords [][2]int) (left int, right int, top int, bottom int) {
	var (
		boundLeft   = math.MaxInt16
		boundRight  = -math.MaxInt16
		boundTop    = math.MaxInt16
		boundBottom = -math.MaxInt16
	)
	for _, c := range coords {
		if c[0] < boundTop {
			boundTop = c[0]
		}
		if c[0] > boundBottom {
			boundBottom = c[0]
		}
		if c[1] < boundLeft {
			boundLeft = c[1]
		}
		if c[1] > boundRight {
			boundRight = c[1]
		}
	}
	return boundLeft, boundRight, boundTop, boundBottom
}
