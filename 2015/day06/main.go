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

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	// 1000x1000 grid
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "toggle"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] = !grid[i][j]
				}
			}
		case strings.HasPrefix(line, "turn on"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] = true
				}
			}
		case strings.HasPrefix(line, "turn off"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] = false
				}
			}
		default:
			panic("unhandled instruction")
		}
	}
	var count int
	for _, row := range grid {
		for _, b := range row {
			if b {
				count++
			}
		}
	}
	return count
}

func part2(input string) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "toggle"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] += 2
				}
			}
		case strings.HasPrefix(line, "turn on"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j]++
				}
			}
		case strings.HasPrefix(line, "turn off"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					if grid[i][j] > 0 {
						grid[i][j]--
					}
				}
			}
		default:
			panic("unhandled instruction")
		}
	}
	var brightness int
	for _, row := range grid {
		for _, v := range row {
			brightness += v
		}
	}
	return brightness
}
