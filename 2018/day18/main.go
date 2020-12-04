package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	grid := parseInput(input)

	for i := 0; i < 10; i++ {
		grid = step(grid)

		// fmt.Println(i)
		// for _, v := range grid {
		// 	fmt.Println(v)
		// }
		// fmt.Println("")
	}

	finalTrees, finalLumbers := countTreesAndLumber(grid)
	return finalTrees * finalLumbers
}

func part2(input string) int {
	grid := parseInput(input)
	stepsWanted := 1000000000
	cacheStringifiedGridToIndex := make(map[string]int)

	for i := 0; i < stepsWanted; i++ {
		grid = step(grid)

		stringified := stringify(grid)
		if lastIndex, ok := cacheStringifiedGridToIndex[stringified]; ok {
			freq := i - lastIndex
			// skip steps
			for i+freq < stepsWanted {
				i += freq
			}
		}

		cacheStringifiedGridToIndex[stringified] = i

	}
	// between 197616 and 209420
	finalTrees, finalLumbers := countTreesAndLumber(grid)
	return finalTrees * finalLumbers
}

func parseInput(input string) [][]string {
	var ans [][]string

	lines := strings.Split(input, "\n")

	// to pad either edge
	columnsPlusTwo := len(lines[0]) + 2

	ans = append(ans, make([]string, columnsPlusTwo))

	for _, l := range lines {
		rowSli := append([]string{""}, strings.Split(l, "")...)
		rowSli = append(rowSli, "")
		ans = append(ans, rowSli)
	}

	ans = append(ans, make([]string, columnsPlusTwo))

	return ans
}

func step(grid [][]string) [][]string {
	nextGrid := make([][]string, len(grid))
	for i := range grid {
		nextGrid[i] = make([]string, len(grid[0]))
	}

	dir := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[0])-1; col++ {
			var adjTrees, adjOpen, adjLumber int
			for _, d := range dir {
				neighborRow, neighborCol := row+d[0], col+d[1]
				switch grid[neighborRow][neighborCol] {
				case "#":
					adjLumber++
				case "|":
					adjTrees++
				case ".":
					adjOpen++
				}
			}

			switch grid[row][col] {
			case ".":
				if adjTrees >= 3 {
					nextGrid[row][col] = "|"
				} else {
					nextGrid[row][col] = "."
				}
			case "|":
				if adjLumber >= 3 {
					nextGrid[row][col] = "#"
				} else {
					nextGrid[row][col] = "|"
				}
			case "#":
				if adjLumber >= 1 && adjTrees >= 1 {
					nextGrid[row][col] = "#"
				} else {
					nextGrid[row][col] = "."
				}
			}
		}
	}

	return nextGrid
}

func countTreesAndLumber(grid [][]string) (int, int) {
	var trees, lumbers int
	for _, row := range grid {
		for _, v := range row {
			switch v {
			case "#":
				lumbers++
			case "|":
				trees++
			}
		}
	}
	return trees, lumbers
}

func stringify(grid [][]string) string {
	var ans string
	for _, row := range grid {
		for _, v := range row {
			ans += v
		}
	}
	return ans
}
