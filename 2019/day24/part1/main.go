package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")

	lines := strings.Split(input, "\n")

	grid := make([][]string, len(lines))
	for i, v := range lines {
		grid[i] = strings.Split(v, "")
	}

	// can map biodiversity scores because they're essentially bitmaps
	previousBiodiversities := map[int]bool{getBiodiversity(grid): true}

	// run indefinitely
	for {
		// step through a minute
		grid = stepMinute(grid)
		newBiodiversity := getBiodiversity(grid)

		// if new biodiversity score is already in the map, print it and exit
		if previousBiodiversities[newBiodiversity] {
			fmt.Println("Repeated biodiversity score", newBiodiversity)
			return
		}

		// set biodiversity score into in map
		previousBiodiversities[newBiodiversity] = true
	}
}

// steps through one minute and returns the next grid
func stepMinute(grid [][]string) [][]string {
	result := make([][]string, len(grid))

	dRow := [4]int{0, 0, -1, 1}
	dCol := [4]int{-1, 1, 0, 0}

	for row, rowSli := range grid {
		// initialize the rows for the result slice
		result[row] = make([]string, len(grid[0]))

		for col, val := range rowSli {
			// count up the neighbors that are bugs
			var countNeighborBugs int
			for i := 0; i < 4; i++ {
				nextRow, nextCol := row+dRow[i], col+dCol[i]
				isInbounds := nextRow >= 0 && nextCol >= 0 && nextRow < len(grid) && nextCol < len(grid[0])
				if isInbounds && grid[nextRow][nextCol] == "#" {
					countNeighborBugs++
				}
			}

			// determine future state of cell
			switch val {
			case "#":
				// if bug has ONE neighbor only, it lives, otherwise becomes empty
				if countNeighborBugs == 1 {
					result[row][col] = "#"
				} else {
					result[row][col] = "."
				}
			case ".":
				// if empty, becomes a bug if has ONE or TWO neighbors only
				if countNeighborBugs == 1 || countNeighborBugs == 2 {
					result[row][col] = "#"
				} else {
					result[row][col] = "."
				}
			}
		}
	}

	return result
}

func getBiodiversity(grid [][]string) int {
	var biodiversity int
	for i, row := range grid {
		for j, val := range row {
			// dumb cheeky way to get power of two... 1 is 2^0, then shift the power
			// 1 << 0 == 2^0 == 1; 1 << 1 == 2^1
			powerOfTwo := 1 << (5*i + j)
			if val == "#" {
				biodiversity += powerOfTwo
			}
		}
	}

	return biodiversity
}
