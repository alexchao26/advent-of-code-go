package main

import (
	_ "embed"
	"flag"
	"fmt"
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

	if part == 1 {
		ans := part1(input, [2]float64{200000000000000, 400000000000000})
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string, testRange [2]float64) int {
	hailstones := parseInput(input)

	ans := 0

	// only move forward (per the velocity)
	for i, hs1 := range hailstones {
		for _, hs2 := range hailstones[i+1:] {
			intersection := getIntersectingCoordinates(hs1, hs2)
			if intersection == nil {
				continue
			}

			// ensure intersection is in the right direction
			// solve for time to reach intersection?
			if solveForTimeToReachPoint(hs1, intersection) < 0 || solveForTimeToReachPoint(hs2, intersection) < 0 {
				continue
			}

			if testRange[0] <= intersection[0] && intersection[0] <= testRange[1] &&
				testRange[0] <= intersection[1] && intersection[1] <= testRange[1] {

				ans++
			}
		}
	}

	return ans
}

// does not check for if the lines are the same line
func areHailstonesParallel(hs1, hs2 hailstone) bool {
	if hs1.hasVerticalPath && hs2.hasVerticalPath {
		return true
	}
	if hs1.hasVerticalPath || hs2.hasVerticalPath {
		return false
	}

	return hs1.slope == hs2.slope
}

// returns nil slice if the lines do not intersect
func getIntersectingCoordinates(hs1, hs2 hailstone) []float64 {
	if areHailstonesParallel(hs1, hs2) {
		return nil
	}
	// assume not the exact same line and that there is only one intersection point

	// point-slope line formula
	// y - y1 = m(x - x1)
	x := (hs1.slope*hs1.x - hs2.slope*hs2.x + hs2.y - hs1.y) / (hs1.slope - hs2.slope)
	y := hs1.slope*(x-hs1.x) + hs1.y

	return []float64{x, y}
}

func solveForTimeToReachPoint(hs hailstone, point []float64) float64 {
	if len(point) != 2 {
		panic("expected len == 2 for point slice")
	}
	// x = vx * t + x0
	// t = (intersection_x - x_0) / vx
	t := (point[0] - hs.x) / hs.vx

	return t
}

func part2(input string) int {
	hailstones := parseInput(input)

	var possibleRockVelX, possibleRockVelY, possibleRockVelZ []int
	for i, hs1 := range hailstones {
		for _, hs2 := range hailstones[i+1:] {

			if hs1.vx == hs2.vx {
				possibilities := getPossibleVelocities(int(hs2.x), int(hs1.x), int(hs1.vx))
				if len(possibleRockVelX) == 0 {
					possibleRockVelX = possibilities
				} else {
					possibleRockVelX = getIntersection(possibleRockVelX, possibilities)
				}
			}
			if hs1.vy == hs2.vy {
				possibilities := getPossibleVelocities(int(hs2.y), int(hs1.y), int(hs1.vy))
				if len(possibleRockVelY) == 0 {
					possibleRockVelY = possibilities
				} else {
					possibleRockVelY = getIntersection(possibleRockVelY, possibilities)
				}
			}
			if hs1.vz == hs2.vz {
				possibilities := getPossibleVelocities(int(hs2.z), int(hs1.z), int(hs1.vz))
				if len(possibleRockVelZ) == 0 {
					possibleRockVelZ = possibilities
				} else {
					possibleRockVelZ = getIntersection(possibleRockVelZ, possibilities)
				}
			}
		}
	}

	if len(possibleRockVelX) == 1 && len(possibleRockVelY) == 1 && len(possibleRockVelZ) == 1 {
		rockVelX := float64(possibleRockVelX[0])
		rockVelY := float64(possibleRockVelY[0])
		rockVelZ := float64(possibleRockVelZ[0])

		hailstoneA, hailstoneB := hailstones[0], hailstones[1]
		mA := (hailstoneA.vy - rockVelY) / (hailstoneA.vx - rockVelX)
		mB := (hailstoneB.vy - rockVelY) / (hailstoneB.vx - rockVelX)
		cA := hailstoneA.y - (mA * hailstoneA.x)
		cB := hailstoneB.y - (mB * hailstoneB.x)
		rockX := (cB - cA) / (mA - mB)
		rockY := mA*rockX + cA
		time := (rockX - hailstoneA.x) / (hailstoneA.vx - rockVelX)
		rockZ := hailstoneA.z + (hailstoneA.vz-rockVelZ)*time
		return int(rockX + rockY + rockZ)
	}

	panic("more than one possible velocity in a direction")
}

func getPossibleVelocities(pos1, pos2 int, vel int) []int {
	match := []int{}
	for possibleVel := -1000; possibleVel < 1000; possibleVel++ {
		if possibleVel != vel && (pos1-pos2)%(possibleVel-vel) == 0 {
			match = append(match, possibleVel)
		}
	}
	return match
}

func getIntersection(sli1, sli2 []int) []int {
	result := []int{}

	map2 := map[int]bool{}
	for _, val := range sli2 {
		map2[val] = true
	}

	for _, val := range sli1 {
		if map2[val] {
			result = append(result, val)
		}
	}
	return result
}

func parseInput(input string) (ans []hailstone) {
	for _, line := range strings.Split(input, "\n") {
		positions := []float64{}
		vels := []float64{}

		line = strings.ReplaceAll(line, ",", "")
		parts := strings.Split(line, " @ ")

		for _, posStr := range strings.Fields(parts[0]) {
			positions = append(positions, float64(cast.ToInt(posStr)))
		}
		for _, velStr := range strings.Fields(parts[1]) {
			vels = append(vels, float64(cast.ToInt(velStr)))
		}

		ans = append(ans, makeHailstone(positions[0], positions[1], positions[2], vels[0], vels[1], vels[2]))
	}

	return ans
}

type hailstone struct {
	x, y, z         float64
	vx, vy, vz      float64
	hasVerticalPath bool
	slope           float64
}

func makeHailstone(x, y, z, vx, vy, vz float64) hailstone {
	hs := hailstone{
		x:               x,
		y:               y,
		z:               z,
		vx:              vx,
		vy:              vy,
		vz:              vz,
		hasVerticalPath: vx == 0,
		slope:           0,
	}
	if !hs.hasVerticalPath {
		hs.slope = vy / vx
	}
	return hs
}
