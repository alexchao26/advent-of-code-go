package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	digInstructions := parseInput(input)

	trenchCoords := getTrenchCoords(digInstructions)

	containedCoords := getContainedCoords(trenchCoords)

	return len(containedCoords) + len(trenchCoords)
}

func getTrenchCoords(digInstructions []digInstruction) map[[2]int]bool {
	trenchCoords := map[[2]int]bool{
		{0, 0}: true,
	}

	var row, col int
	diffs := map[string][2]int{
		"L": {0, -1},
		"R": {0, 1},
		"U": {-1, 0},
		"D": {1, 0},
	}

	for _, inst := range digInstructions {
		for i := 1; i <= inst.length; i++ {
			row += diffs[inst.dir][0]
			col += diffs[inst.dir][1]
			trenchCoords[[2]int{row, col}] = true
		}
	}
	return trenchCoords
}

func getContainedCoords(trenchCoords map[[2]int]bool) map[[2]int]bool {
	// check around a coordinate that's part of a straight line for a cell that _could_ be contained
	// straight lines will have one side that's in and one that is out
	// we'll only check vertical lines to make it easier to code...
	var testCoords [][2]int

	for coord := range trenchCoords {
		upCoords := [2]int{coord[0] - 1, coord[1]}
		downCoords := [2]int{coord[0] + 1, coord[1]}
		leftCoords := [2]int{coord[0], coord[1] - 1}
		rightCoords := [2]int{coord[0], coord[1] + 1}

		if trenchCoords[upCoords] && trenchCoords[downCoords] &&
			!trenchCoords[leftCoords] && !trenchCoords[rightCoords] {
			// part of vertical line
			testCoords = append(testCoords, leftCoords, rightCoords)
			break
		}
	}

	// calculate the max size that can be contained (equal to the box containing all the coordinates)
	var (
		left   = testCoords[0][1]
		right  = testCoords[0][1]
		top    = testCoords[0][0]
		bottom = testCoords[0][0]
	)
	for coords := range trenchCoords {
		left = min(left, coords[1])
		right = max(right, coords[1])
		top = min(top, coords[0])
		bottom = max(bottom, coords[0])
	}

	maxContainedSize := (right - left + 1) * (bottom - top + 1)

	for _, coord := range testCoords {
		queue := [][2]int{coord}
		seen := map[[2]int]bool{}

		for len(queue) > 0 && len(seen) < maxContainedSize {
			current := queue[0]
			queue = queue[1:]

			if seen[current] {
				continue
			}
			seen[current] = true

			for _, diff := range [][2]int{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			} {
				nextRow := current[0] + diff[0]
				nextCol := current[1] + diff[1]
				nextCoord := [2]int{nextRow, nextCol}
				// if already seen or it's part of the trench, skip
				if trenchCoords[nextCoord] || seen[nextCoord] {
					continue
				}
				// otherwise add it to be searched
				queue = append(queue, nextCoord)
			}
		}

		if len(queue) == 0 {
			return seen
		}
	}
	panic("should return from loop")
}

func part2(input string) int {
	digInstructions := parseInput(input)

	vertices := [][2]int{}
	currentPoint := [2]int{0, 0}

	for _, inst := range digInstructions {
		hex := inst.color[1 : len(inst.color)-1]
		dirCode := inst.color[len(inst.color)-1:]

		convInt, err := strconv.ParseInt(hex, 16, 0)
		if err != nil {
			panic(err.Error())
		}

		switch dirCode {
		case "0": // R
			currentPoint[1] += int(convInt)
		case "1": // D
			currentPoint[0] += int(convInt)
		case "2": // L
			currentPoint[1] -= int(convInt)
		case "3": // U
			currentPoint[0] -= int(convInt)
		}
		vertices = append(vertices, currentPoint)
	}

	return shoelace(vertices) + 1
}

func shoelace(coordinates [][2]int) int {
	area := 0

	for i := 0; i < len(coordinates); i++ {
		coordA := coordinates[i]
		coordB := coordinates[(i+1)%(len(coordinates))]

		area += (coordA[1] * coordB[0]) - (coordB[1] * coordA[0]) +
			max(abs(coordA[0]-coordB[0]), abs(coordA[1]-coordB[1]))
	}

	return area / 2
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type digInstruction struct {
	dir    string
	length int
	color  string
}

func parseInput(input string) (ans []digInstruction) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		ans = append(ans, digInstruction{
			dir:    parts[0],
			length: cast.ToInt(parts[1]),
			color:  parts[2][1 : len(parts[2])-1],
		})
	}
	return ans
}
