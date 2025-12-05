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
	// fewer than four rolls of paper in the eight adjacent positions
	grid := [][]string{}
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}

	ans := 0
	for r := range grid {
		for c, v := range grid[r] {
			if v != "@" {
				continue
			}
			neighbors := 0
			for _, diff := range [][2]int{
				{-1, -1},
				{-1, 0},
				{-1, 1},
				{0, -1},
				{0, 1},
				{1, -1},
				{1, 0},
				{1, 1},
			} {
				nr := r + diff[0]
				nc := c + diff[1]
				if nr < 0 || nr == len(grid) || nc < 0 || nc == len(grid[0]) {
					continue
				}
				if grid[nr][nc] == "@" {
					neighbors++
				}
			}

			if neighbors < 4 {
				ans++
			}
		}
	}

	return ans
}

// modifies underlying grid
func removeRolls(grid [][]string) int {
	ans := 0
	for r := range grid {
		for c, v := range grid[r] {
			if v != "@" {
				continue
			}
			neighbors := 0
			for _, diff := range [][2]int{
				{-1, -1},
				{-1, 0},
				{-1, 1},
				{0, -1},
				{0, 1},
				{1, -1},
				{1, 0},
				{1, 1},
			} {
				nr := r + diff[0]
				nc := c + diff[1]
				if nr < 0 || nr == len(grid) || nc < 0 || nc == len(grid[0]) {
					continue
				}
				if grid[nr][nc] == "@" {
					neighbors++
				}
			}

			if neighbors < 4 {
				ans++
				grid[r][c] = "."
			}
		}
	}

	return ans
}

func part2(input string) int {
	grid := [][]string{}
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}

	ans := removeRolls(grid)
	for {
		removed := removeRolls(grid)
		if removed == 0 {
			break
		}
		ans += removed
	}
	return ans
}
