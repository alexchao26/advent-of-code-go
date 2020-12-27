package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
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

// globals for part 1
var directions = []string{"N", "E", "S", "W"}
var directionsToDiff = map[string][2]int{
	// !! X and Y like a coordinate system! Not typical 2D matrices in algos
	"N": [2]int{0, 1},
	"E": [2]int{1, 0},
	"S": [2]int{0, -1},
	"W": [2]int{-1, 0},
}

func part1(input string) int {
	instructions := parseInput(input)

	// X and Y like a coordinate system
	var shipX, shipY int
	dirIndex := 1 // index in directions slice
	for _, inst := range instructions {
		switch inst.action {
		case "N":
			shipY += inst.value
		case "S":
			shipY -= inst.value
		case "E":
			shipX += inst.value
		case "W":
			shipX -= inst.value
		case "L":
			// rotate ship left, this is equivalent
			// -1 + 4 to keeping dirIndex positive for the modding
			dirIndex += (-1 + 4) * inst.value / 90
			dirIndex %= 4
		case "R":
			dirIndex += inst.value / 90
			dirIndex %= 4
		case "F":
			d := directionsToDiff[directions[dirIndex]]
			shipX += d[0] * inst.value
			shipY += d[1] * inst.value
		default:
			panic("unexpected action")
		}
	}

	return mathy.ManhattanDistance(0, 0, shipX, shipY)
}

func part2(input string) int {
	instructions := parseInput(input)

	// X and Y like a coordinate system
	waypointX := 10
	waypointY := 1
	var shipX, shipY int

	for _, inst := range instructions {
		switch inst.action {
		case "N":
			waypointY += inst.value
		case "S":
			waypointY -= inst.value
		case "E":
			waypointX += inst.value
		case "W":
			waypointX -= inst.value
		case "L":
			// rotate waypoint left around ship (origin)
			turns := inst.value / 90
			for turns > 0 {
				// this simple bit is all it needs to rotate around the origin
				// I had a 20 line if/else block...
				waypointX, waypointY = -waypointY, waypointX
				turns--
			}
		case "R":
			turns := inst.value / 90
			for turns > 0 {
				waypointX, waypointY = waypointY, -waypointX
				turns--
			}
		case "F":
			shipX += inst.value * waypointX
			shipY += inst.value * waypointY
		default:
			panic("unexpected action")
		}
	}

	return mathy.ManhattanDistance(0, 0, shipX, shipY)
}

type instruction struct {
	action string
	value  int
}

func parseInput(input string) []instruction {
	var ans []instruction

	lines := strings.Split(input, "\n")
	for _, l := range lines {
		inst := instruction{
			action: l[:1],
			value:  cast.ToInt(l[1:]),
		}
		ans = append(ans, inst)
	}

	return ans
}
