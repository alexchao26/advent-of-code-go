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

	// may be visible from multiple angles
	visibleCoords := map[[2]int]string{}
	for r := 1; r < len(grid)-1; r++ {
		// from left
		highestFromLeft := -1
		for c := 0; c < len(grid[0])-1; c++ {
			height := grid[r][c]
			if height > highestFromLeft {
				visibleCoords[[2]int{r, c}] = "L"
				highestFromLeft = height
			}
		}
		// from right
		highestFromRight := -1
		for c := len(grid[0]) - 1; c > 0; c-- {
			height := grid[r][c]
			if height > highestFromRight {
				visibleCoords[[2]int{r, c}] = "R"
				highestFromRight = height
			}
		}
	}

	for c := 1; c < len(grid[0])-1; c++ {
		// from top
		highestFromTop := -1
		for r := 0; r < len(grid)-1; r++ {
			height := grid[r][c]
			if height > highestFromTop {
				visibleCoords[[2]int{r, c}] = "T"
				highestFromTop = height
			}
		}
		// from bottom
		highestFromBottom := -1
		for r := len(grid) - 1; r > 0; r-- {
			height := grid[r][c]
			if height > highestFromBottom {
				visibleCoords[[2]int{r, c}] = "B"
				highestFromBottom = height
			}
		}
	}

	return len(visibleCoords) + 4 // plus 4 for corners
}

func part2(input string) int {
	// multiply the four scores together... score = how many trees any tree can see
	// because trees on the edge will have a zero, just ignore them
	grid := parseInput(input)

	bestScore := 0
	// iterate through every eligible tree
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			score := visible(grid, r, c, -1, 0)
			score *= visible(grid, r, c, 1, 0)
			score *= visible(grid, r, c, 0, -1)
			score *= visible(grid, r, c, 0, 1)

			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func visible(grid [][]int, r, c, dr, dc int) int {
	count := 0
	startingHeight := grid[r][c]
	r += dr
	c += dc
	for r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
		height := grid[r][c]
		if height < startingHeight {
			count++
		} else {
			count++
			break
		}

		r += dr
		c += dc
	}

	return count
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, n := range strings.Split(line, "") {
			row = append(row, cast.ToInt(n))
		}
		ans = append(ans, row)
	}
	return ans
}
