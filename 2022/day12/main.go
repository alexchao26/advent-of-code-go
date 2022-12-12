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

var diffs = [4][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func part1(input string) int {
	grid := parseInput(input)

	queue := [][3]int{}
label:
	for r, rows := range grid {
		for c, cell := range rows {
			if cell == "S" {
				queue = append(queue, [3]int{r, c, 0})
				break label
			}
		}
	}
	seen := map[[2]int]bool{}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if seen[[2]int{front[0], front[1]}] {
			continue
		}
		seen[[2]int{front[0], front[1]}] = true

		if grid[front[0]][front[1]] == "E" {
			return front[2]
		}
		for _, d := range diffs {
			nextR, nextC := front[0]+d[0], front[1]+d[1]
			if nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0]) {
				letterDiff := distanceBetweenLetters(grid[front[0]][front[1]], grid[nextR][nextC])

				if letterDiff <= 1 {
					queue = append(queue, [3]int{nextR, nextC, front[2] + 1})
				}
			}
		}
	}

	return -1
}

func part2(input string) int {
	grid := parseInput(input)

	queue := [][3]int{}
label:
	for r, rows := range grid {
		for c, cell := range rows {
			if cell == "E" {
				queue = append(queue, [3]int{r, c, 0})
				break label
			}
		}
	}
	seen := map[[2]int]bool{}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if seen[[2]int{front[0], front[1]}] {
			continue
		}
		seen[[2]int{front[0], front[1]}] = true

		if grid[front[0]][front[1]] == "a" {
			return front[2]
		}
		for _, d := range diffs {
			nextR, nextC := front[0]+d[0], front[1]+d[1]
			if nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0]) {
				letterDiff := distanceBetweenLetters(grid[front[0]][front[1]], grid[nextR][nextC])

				if letterDiff >= -1 {
					queue = append(queue, [3]int{nextR, nextC, front[2] + 1})
				}
			}
		}
	}

	return -1
}

func distanceBetweenLetters(x, y string) int {
	if x == "S" {
		x = "a"
	}
	if y == "S" {
		y = "a"
	}
	if y == "E" {
		y = "z"
	}
	if x == "E" {
		x = "z"
	}

	return cast.ToASCIICode(y) - cast.ToASCIICode(x)
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
