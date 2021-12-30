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
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

// naive brute force for part1
func part1(input string) int {
	cubes := parseInput(input)

	onCoords := map[[3]int]bool{}

	for _, c := range cubes {
		if part1OutOfBounds(c.x1, c.x2, c.y1, c.y2, c.z1, c.z2) {
			continue
		}

		for x := c.x1; x <= c.x2; x++ {
			for y := c.y1; y <= c.y2; y++ {
				for z := c.z1; z <= c.z2; z++ {
					coord := [3]int{x, y, z}
					onCoords[coord] = c.isOn
				}
			}
		}
	}

	var count int
	for _, b := range onCoords {
		if b {
			count++
		}
	}

	return count
}

func part1OutOfBounds(nums ...int) bool {
	for _, n := range nums {
		if n < -50 || n > 50 {
			return true
		}
	}
	return false
}

func part2(input string) int {
	cubes := parseInput(input)

	var finalList []cube
	// iterate through all cubes, keep a final list of cubes
	// as new cubes are added, check against the entire final list and add any
	// intersections that are found
	for _, c := range cubes {
		// add these at the end of the step to prevent duplicate checks
		var toAdd []cube

		for _, finalCube := range finalList {
			intersection, didIntersect := finalCube.getIntersection(c)
			if didIntersect {
				toAdd = append(toAdd, intersection)
			}
		}

		// if cube is an "on" cube, it needs to be added to final list
		if c.isOn {
			toAdd = append(toAdd, c)
		}

		finalList = append(finalList, toAdd...)
	}

	var total int
	for _, c := range finalList {
		total += c.volume()
	}

	return total
}

type cube struct {
	isOn   bool
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

// NOTE: must be called in correct order (cube_from_final_list).getIntersection(cube_being_added)
// because of how the isOn bool is determined
func (c cube) getIntersection(c2 cube) (intersection cube, hasIntersection bool) {
	// larger of x1s has to be smaller than smaller of x2s for there to be an overlap
	x1 := mathy.MaxInt(c.x1, c2.x1)
	x2 := mathy.MinInt(c.x2, c2.x2)
	y1 := mathy.MaxInt(c.y1, c2.y1)
	y2 := mathy.MinInt(c.y2, c2.y2)
	z1 := mathy.MaxInt(c.z1, c2.z1)
	z2 := mathy.MinInt(c.z2, c2.z2)

	if x1 > x2 || y1 > y2 || z1 > z2 {
		return cube{}, false
	}

	var intersectionState bool
	if c.isOn && c2.isOn {
		intersectionState = false
	} else if !c.isOn && !c2.isOn {
		intersectionState = true
	} else {
		// ! default to second cube's on/off state. This makes the order of which cube's method is
		// called very important. but that's what unit tests are for :)
		// alternatively the caller could deal with it.. that might be more clear...
		intersectionState = c2.isOn
	}

	return cube{
		isOn: intersectionState,
		x1:   x1, x2: x2,
		y1: y1, y2: y2,
		z1: z1, z2: z2,
	}, true
}

func (c cube) volume() int {
	vol := (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
	if c.isOn {
		return vol
	}
	return -vol
}

func parseInput(input string) (ans []cube) {
	for _, line := range strings.Split(input, "\n") {
		// off x=-29..-12,y=-13..5,z=-17..-3
		parts := strings.Split(line, " ")

		var x1, x2, y1, y2, z1, z2 int
		n, err := fmt.Sscanf(parts[1], "x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
		if err != nil || n != 6 {
			panic(fmt.Sprintf("parsing error %v, vals parsed %d", err, n))
		}

		if x1 > x2 || y1 > y2 || z1 > z2 {
			// note: they can be equal
			panic("didn't expect input to have backwards coords, sort them...")
		}

		ans = append(ans, cube{
			isOn: parts[0] == "on",
			x1:   x1,
			x2:   x2,
			y1:   y1,
			y2:   y2,
			z1:   z1,
			z2:   z2,
		})
	}
	return ans
}
