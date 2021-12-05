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

	ans := countIntersections(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func countIntersections(input string, part int) int {
	var coords [][4]int
	for _, line := range strings.Split(input, "\n") {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			panic("parsing error: " + err.Error())
		}
		coords = append(coords, [4]int{x1, y1, x2, y2})
	}

	// find largest x and y
	// Note: different x,y as prompt, x = rows, y = cols instead
	var endRow, endCol int
	for _, c := range coords {
		if c[0] > endRow {
			endRow = c[0]
		}
		if c[2] > endRow {
			endRow = c[2]
		}
		if c[1] > endCol {
			endCol = c[1]
		}
		if c[3] > endCol {
			endCol = c[3]
		}
	}

	grid := make([][]int, endRow+1)
	for i := range grid {
		grid[i] = make([]int, endCol+1)
	}

	for _, c := range coords {
		if c[0] == c[2] {
			// horiz line, row1 = row2
			row := c[0]
			start, end := c[1], c[3]
			// swap start and end if needed
			if c[1] > c[3] {
				start, end = end, start
			}
			for col := start; col <= end; col++ {
				grid[row][col]++
			}
		} else if c[1] == c[3] {
			// vert line, col1 = col2
			col := c[1]
			start, end := c[0], c[2]
			// swap start and end if needed
			if c[0] > c[2] {
				start, end = end, start
			}
			for row := start; row <= end; row++ {
				grid[row][col]++
			}
		} else if part == 2 {
			// check diagonals for part 2 only

			// get left most pair first by comparing column coords
			if c[1] > c[3] {
				c = [4]int{
					c[2], c[3],
					c[0], c[1],
				}
			}

			// compare row coords next, will be going rightwards regardless
			// if going down and right
			if c[0] < c[2] {
				for row := c[0]; row <= c[2]; row++ {
					col := c[1] + row - c[0]
					grid[row][col]++
				}
			} else {
				// going up and right
				for row := c[0]; row >= c[2]; row-- {
					col := c[1] + (c[0] - row)
					grid[row][col]++
				}
			}
		}
	}

	// count up intersections
	var ans int
	for _, rows := range grid {
		for _, v := range rows {
			if v >= 2 {
				ans++
			}
		}
	}

	return ans
}
