package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

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

	ans := pipeMaze(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

var pipes = map[string][2][2]int{
	"|": {
		{-1, 0}, // up
		{1, 0},  // down
	},
	"-": {
		{0, -1}, // left
		{0, 1},  //right
	},
	"L": {
		{-1, 0}, // up
		{0, 1},  // right
	},
	"J": {
		{-1, 0}, // up
		{0, -1}, // left
	},
	"7": {
		{0, -1}, // left
		{1, 0},  // down
	},
	"F": {
		{0, 1}, // right
		{1, 0}, // down
	},
}

func pipeMaze(input string, part int) int {
	grid := parseInput(input)

	var r, c int
	for i, row := range grid {
		for j, val := range row {
			if val == "S" {
				r = i
				c = j
			}
		}
	}

	fillGridLocation(grid, r, c)

	// traverse entire loop to determine length
	loopCoords := map[[2]int]bool{}
	toVisit := [][2]int{
		{r, c},
	}
	for len(toVisit) > 0 {
		coords := toVisit[0]
		if loopCoords[coords] {
			break
		}
		loopCoords[coords] = true
		toVisit = toVisit[1:]

		// assumes loop is well formed, will cause a panic if not
		diffs := pipes[grid[coords[0]][coords[1]]]
		for _, diff := range diffs {
			nextRow, nextCol := coords[0]+diff[0], coords[1]+diff[1]
			if isInRange(grid, nextRow, nextCol) && !loopCoords[[2]int{nextRow, nextCol}] {
				toVisit = append(toVisit, [2]int{nextRow, nextCol})
			}
		}

	}

	if part == 1 {
		return len(loopCoords) / 2
	}

	// part 2

	// create copy of grid with all non-loop spots replaced with a period
	reducedGrid := make([][]string, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if loopCoords[[2]int{i, j}] {
				reducedGrid[i] = append(reducedGrid[i], grid[i][j])
			} else {
				reducedGrid[i] = append(reducedGrid[i], ".")
			}
		}
	}

	// expand grid to double plus 2 in both dimensions to account for squeezing between pipes
	// the plus two is to add an empty row/column on each side for easier traversing from the outside
	expandedGrid := [][]string{}
	expandedGrid = append(expandedGrid, make([]string, len(reducedGrid[0])*2+2))

	for r, rows := range reducedGrid {
		expandedGrid = append(expandedGrid, make([]string, len(reducedGrid[0])*2+2))
		for c, val := range rows {
			expandedGrid[r*2+1][c*2+1] = val
		}
		// empty row
		expandedGrid = append(expandedGrid, make([]string, len(reducedGrid[0])*2+2))
	}

	// fill gaps between loop coords so we have an encased area again
	// we can naively try to fill in every empty spot because only ones with two valid connecting
	// pipes will be filled
	for r, rows := range expandedGrid {
		for c, val := range rows {
			if val == "" {
				fillGridLocation(expandedGrid, r, c)
			}
		}
	}

	// replacing empty strings with spaces makes the printout human readable
	for r, rows := range expandedGrid {
		for c, val := range rows {
			if val == "" {
				expandedGrid[r][c] = " "
			}
		}
	}

	toVisit = [][2]int{
		{0, 0},
	}
	seen := map[[2]int]bool{}
	for len(toVisit) > 0 {
		coords := toVisit[0]
		toVisit = toVisit[1:]
		if seen[coords] {
			continue
		}
		seen[coords] = true

		// delete reachable dots
		if expandedGrid[coords[0]][coords[1]] == "." {
			expandedGrid[coords[0]][coords[1]] = " "
		}

		for _, diff := range [][2]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			nextRow := coords[0] + diff[0]
			nextCol := coords[1] + diff[1]
			if isInRange(expandedGrid, nextRow, nextCol) {
				if expandedGrid[nextRow][nextCol] == "." || expandedGrid[nextRow][nextCol] == " " {
					toVisit = append(toVisit, [2]int{nextRow, nextCol})
				}
			}
		}
	}

	// count remaining dots
	var ans int
	for _, rows := range expandedGrid {
		for _, val := range rows {
			if val == "." {
				ans++
			}
		}
	}

	// for _, rows := range expandedGrid {
	// 	fmt.Println(rows)
	// }

	return ans
}

func fillGridLocation(grid [][]string, r, c int) {
	// inputs are nice and exactly two adjacent cells that connect to S
	// check four directions from start
	leftCol := c - 1
	rightCol := c + 1
	upRow := r - 1
	downRow := r + 1

	var combinedString string
	// check left for inRange and possible valid pipe types
	if isInRange(grid, r, leftCol) &&
		(grid[r][leftCol] == "-" || grid[r][leftCol] == "L" || grid[r][leftCol] == "F") {
		combinedString += "left"
	}
	// right
	if isInRange(grid, r, rightCol) &&
		(grid[r][rightCol] == "-" || grid[r][rightCol] == "J" || grid[r][rightCol] == "7") {
		combinedString += "right"
	}
	// up
	if isInRange(grid, upRow, c) &&
		(grid[upRow][c] == "|" || grid[upRow][c] == "7" || grid[upRow][c] == "F") {
		combinedString += "up"
	}
	if isInRange(grid, downRow, c) &&
		(grid[downRow][c] == "|" || grid[downRow][c] == "J" || grid[downRow][c] == "L") {
		combinedString += "down"
	}

	switch combinedString {
	case "leftup":
		grid[r][c] = "J"
	case "leftdown":
		grid[r][c] = "7"
	case "rightup":
		grid[r][c] = "L"
	case "rightdown":
		grid[r][c] = "F"
	case "leftright":
		grid[r][c] = "-"
	case "updown":
		grid[r][c] = "|"
		// default:
		// do not panic so we can use this function more naively for the expanded grid
		// could return an error instead and choose to check it for part1 where we NEED it to find a result
		// panic("ineligible configuration: " + combinedString)
	}
}

func isInRange(grid [][]string, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
