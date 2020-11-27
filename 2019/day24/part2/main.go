package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"strings"
)

// RecursiveWorld stores a big 3D matrix & will have associated methods
type RecursiveWorld struct {
	// 401 so there are 200 layers above and below initial layer
	// using ints to expedite calculating neighbor sums and initializing as 0s
	levels [401][5][5]int
}

func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(input, "\n")

	var initialGrid [5][5]int
	for i, line := range lines {
		for j, v := range line {
			if v == '#' {
				initialGrid[i][j] = 1
			}
		}
	}

	// initialize recursive world - zero values of array will start every cell at 0
	var world RecursiveWorld
	// set the "middle" layer at 200 - this works because we're only running for 200 minutes
	world.levels[200] = initialGrid

	// run for 200 minutes
	for i := 0; i < 200; i++ {
		world.minute()
	}

	// print the final count
	fmt.Println("Final count", world.countBugs())
}

func (world *RecursiveWorld) minute() {
	nextMinuteLevels := [401][5][5]int{}

	for i := 0; i < 401; i++ {
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				sumNeighbors := world.getSumOfNeighbors(i, row, col)
				nextMinuteLevels[i][row][col] = world.nextCellValue(
					world.levels[i][row][col], sumNeighbors)
			}
		}
	}

	// reassign levels
	world.levels = nextMinuteLevels
}

// get the sum of neighbors, including recursive layers
func (world *RecursiveWorld) getSumOfNeighbors(i, j, k int) int {
	// center should always remain zero
	if j == 2 && k == 2 {
		return 0
	}

	dx, dy := [4]int{0, 0, -1, 1}, [4]int{-1, 1, 0, 0}

	var sumOfNeighbors int
	for d := 0; d < 4; d++ {
		nextRow, nextCol := j+dx[d], k+dy[d]
		isInBounds := nextRow >= 0 && nextCol >= 0 && nextRow < 5 && nextCol < 5

		// if not in bounds, this cell is trying to access the layer "outside" of it
		if !isInBounds {
			sumOfNeighbors += world.getNeighborsOut(i+1, nextRow, nextCol)
		} else if nextRow == 2 && nextCol == 2 {
			// if a neighbor cell has 2,2 coordinates, it is trying to recurse "in"
			sumOfNeighbors += world.getNeighborsIn(i-1, j, k)
		} else if isInBounds {
			// otherwise if it is inbounds, add from this layer
			sumOfNeighbors += world.levels[i][nextRow][nextCol]
		}
	}

	return sumOfNeighbors
}

// Assuming going outwards moves UP level indexes
// nextRow and nextCol are the requested coordinates from the origin/cell calling this function
func (world *RecursiveWorld) getNeighborsOut(level, nextRow, nextCol int) int {
	// edge case for "recursive" calls asking for -1 layer or 401 layer
	if level == -1 || level == 401 {
		return 0
	}
	currentLevel := world.levels[level]

	// origin cell is asking for "above" itself
	if nextRow == -1 {
		return currentLevel[1][2]
	}
	// origin cell asking for "below" itself
	if nextRow == 5 {
		return currentLevel[3][2]
	}
	// asking for "left"
	if nextCol == -1 {
		return currentLevel[2][1]
	}
	// asking for "right"
	return currentLevel[2][3]
}

// Assume going inwards moves DOWN level indices
// originRow and Col are coordinates of the cell requesting its neighboring values
func (world *RecursiveWorld) getNeighborsIn(level, originRow, originCol int) int {
	// edge case for "recursive" calls asking for -1 layer or 401 layer
	if level == -1 || level == 401 {
		return 0
	}
	// sum up 5 values of the argument level
	currentLevel := world.levels[level]
	var left, right, top, bottom int
	for i := 0; i < 5; i++ {
		left += currentLevel[i][0]
		right += currentLevel[i][4]
		top += currentLevel[0][i]
		bottom += currentLevel[4][i]
	}

	// if originRow is 1, then it was above this layer, return top values
	if originRow == 1 {
		return top
	}
	// if originRow is 3, it is below this layer
	if originRow == 3 {
		return bottom
	}
	// if originCol is 1, it is left of this layer
	if originCol == 1 {
		return left
	}
	// otherwise only remaining direction is right
	return right
}

// gets the next values of a cell given the old value of the cell & sum of its neighbors
func (world *RecursiveWorld) nextCellValue(oldVal, sumNeighbors int) int {
	if oldVal == 1 && sumNeighbors == 1 {
		return 1
	}
	if oldVal == 0 && (sumNeighbors == 1 || sumNeighbors == 2) {
		return 1
	}

	return 0
}

// count up the bugs in every level and return it
func (world *RecursiveWorld) countBugs() int {
	var bugs int
	for _, grid := range world.levels {
		for _, row := range grid {
			for _, val := range row {
				bugs += val
			}
		}
	}

	return bugs
}
