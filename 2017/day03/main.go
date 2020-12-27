package main

import (
	"flag"
	"fmt"

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

func part1(input string) int {
	inputNum := cast.ToInt(input)

	directions := [][2]int{
		{0, 1},  // right
		{-1, 0}, // north
		{0, -1}, // west
		{1, 0},  // south
	}

	coordsToSquareNum := map[[2]int]int{
		[2]int{0, 0}: 1,
		[2]int{0, 1}: 2,
	}

	// @ every step, check left, i.e. see if we can turn right
	// if it's not in map, turn & move there
	// otherwise continue adding in same direction
	directionIndex := 0 // start facing right
	row, col := 0, 1
	number := 2
	for len(coordsToSquareNum) < inputNum {
		leftDiff := directions[(directionIndex+1)%4]
		leftCoord := [2]int{row + leftDiff[0], col + leftDiff[1]}
		if _, ok := coordsToSquareNum[leftCoord]; !ok {
			// turn left first
			directionIndex = (directionIndex + 1) % 4
		}
		// move to coordinate in front & add it to the map
		diff := directions[directionIndex]
		row += diff[0]
		col += diff[1]
		next := [2]int{row, col}
		coordsToSquareNum[next] = number
		number++ // increment number
	}

	return mathy.ManhattanDistance(0, 0, row, col)
}

func part2(input string) int {
	inputNum := cast.ToInt(input)

	directions := [][2]int{
		{0, 1},  // right
		{-1, 0}, // north
		{0, -1}, // west
		{1, 0},  // south
	}

	allNeighborDiffs := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	coordsToNeighborSum := map[[2]int]int{
		[2]int{0, 0}: 1,
		[2]int{0, 1}: 1,
	}

	// @ every step, check left, i.e. see if we can turn right
	// if it's not in map, turn & move there
	// otherwise continue adding in same direction
	directionIndex := 0 // start facing right
	row, col := 0, 1
	for len(coordsToNeighborSum) < inputNum {
		leftDiff := directions[(directionIndex+1)%4]
		leftCoord := [2]int{row + leftDiff[0], col + leftDiff[1]}
		if _, ok := coordsToNeighborSum[leftCoord]; !ok {
			// turn left first
			directionIndex = (directionIndex + 1) % 4
		}
		// move to coordinate in front & add it to the map
		diff := directions[directionIndex]
		row += diff[0]
		col += diff[1]
		next := [2]int{row, col}

		var sum int
		for _, d := range allNeighborDiffs {
			sum += coordsToNeighborSum[[2]int{row + d[0], col + d[1]}]
		}

		if sum > inputNum {
			return sum
		}

		coordsToNeighborSum[next] = sum
	}

	panic("uh something broke, return should occur in for loop")
}
