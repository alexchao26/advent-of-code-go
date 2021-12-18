package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/data-structures/heap"
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
	grid := parseInput(input)
	_ = grid

	for i, rows := range grid {
		for j := range rows {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i-1][j] < grid[i][j-1] {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1] - grid[0][0]

}

func part2(input string) int {
	// make the grid 5 times larger
	// add 1 to every cell when moving right OR down (so diagonal is +2)
	// 9 wraps back to 1
	grid := parseInput(input)
	bigGrid := make([][]int, len(grid)*5)
	for i := range bigGrid {
		bigGrid[i] = make([]int, len(grid[0])*5)
	}
	for r, row := range grid {
		for c, v := range row {
			bigGrid[r][c] = v
		}
	}

	assignGrid := func(baseGrid [][]int, newGrid [][]int, r, c int) {
		for i := 0; i < len(newGrid); i++ {
			for j := 0; j < len(newGrid[0]); j++ {
				baseGrid[r+i][c+j] = newGrid[i][j]
			}
		}
	}

	incrementGrid := func(baseGrid [][]int, by int) [][]int {
		newGrid := make([][]int, len(baseGrid))
		for i := range newGrid {
			newGrid[i] = make([]int, len(baseGrid[0]))
		}
		for i := range baseGrid {
			for j := range baseGrid[0] {
				newGrid[i][j] = baseGrid[i][j] + by
				for newGrid[i][j] > 9 {
					newGrid[i][j] -= 9
				}
			}
		}
		return newGrid
	}

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if r == 0 && c == 0 {
				continue
			}
			nextGrid := incrementGrid(grid, r+c)
			assignGrid(bigGrid, nextGrid, r*len(grid), c*len(grid[0]))
		}
	}

	minHeap := heap.NewMinHeap()
	minHeap.Add(node{0, 0, 0})
	visited := map[[2]int]bool{}
	for minHeap.Length() > 0 {
		front := minHeap.Remove().(node)
		coord := [2]int{front.row, front.col}
		// moving this check here instead of in the heap.Add() makes a HUGE difference
		// by reducing the swell of the queue and following computations
		if visited[coord] {
			continue
		}
		visited[[2]int{front.row, front.col}] = true

		if front.row == len(bigGrid)-1 && front.col == len(bigGrid[0])-1 {
			return front.risk
		}

		// travel to all of front's neighbors
		for _, d := range [][2]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		} {
			nr, nc := front.row+d[0], front.col+d[1]
			if nr >= 0 && nr < len(bigGrid) && nc >= 0 && nc < len(bigGrid[0]) {
				minHeap.Add(node{
					row:  nr,
					col:  nc,
					risk: front.risk + bigGrid[nr][nc],
				})
			}
		}
	}

	panic("should return from loop")
}

// A* node
type node struct {
	row, col int
	risk     int
}

func (n node) Value() int {
	return n.risk
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, char := range strings.Split(line, "") {
			row = append(row, cast.ToInt(char))
		}
		ans = append(ans, row)
	}
	return ans
}
