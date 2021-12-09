package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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
	grid := parseInput(input)

	diffs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var totalRisk int
	for r, rows := range grid {
		for c, v := range rows {
			lowerThanAll := true
			for _, d := range diffs {
				dr, dc := r+d[0], c+d[1]
				// in bounds
				if dr >= 0 && dr < len(grid) && dc >= 0 && dc < len(grid[0]) {
					// neighbor is higher or even, r,c is not a low point
					if grid[dr][dc] <= v {
						lowerThanAll = false
						break
					}
				}
			}

			if lowerThanAll {
				totalRisk += 1 + v
			}
		}
	}

	return totalRisk
}

func part2(input string) int {
	grid := parseInput(input)

	diffs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	// copy pasta but collecting the coordinates to check
	var lowPoints [][2]int
	for r, rows := range grid {
		for c, v := range rows {
			lowerThanAll := true
			for _, d := range diffs {
				dr, dc := r+d[0], c+d[1]
				if dr >= 0 && dr < len(grid) && dc >= 0 && dc < len(grid[0]) {
					if grid[dr][dc] <= v {
						lowerThanAll = false
						break
					}
				}
			}

			if lowerThanAll {
				lowPoints = append(lowPoints, [2]int{r, c})
			}
		}
	}

	// go through all lowpoints and get basin sizes via helper func
	var basins []int
	for _, lp := range lowPoints {
		basins = append(basins, getBasinSize(grid, lp[0], lp[1], map[[2]int]bool{}))
	}

	// return 3 largest basins multiplied together
	ans := 1
	sort.Ints(basins)
	for i := 0; i < 3; i++ {
		ans *= basins[len(basins)-1-i]
	}

	return ans
}

var diffs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func getBasinSize(grid [][]int, r, c int, basinCoords map[[2]int]bool) int {
	// assume that every cell will be involved in one basin, just have to stop at nines
	if grid[r][c] == 9 {
		return 0
	}

	coord := [2]int{r, c}
	// stop if already seen
	if basinCoords[coord] {
		return 0
	}
	// mark as seen
	basinCoords[coord] = true

	for _, d := range diffs {
		dr, dc := r+d[0], c+d[1]
		if dr >= 0 && dr < len(grid) && dc >= 0 && dc < len(grid[0]) {
			// recurse
			getBasinSize(grid, dr, dc, basinCoords)
		}
	}

	// final size of coords map is the basin size
	return len(basinCoords)
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var nums []int
		for _, v := range strings.Split(line, "") {
			nums = append(nums, cast.ToInt(v))
		}
		ans = append(ans, nums)
	}
	return ans
}
