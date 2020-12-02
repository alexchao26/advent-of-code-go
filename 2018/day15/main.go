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
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	strGrid := parseInputs(input)
	_ = strGrid
	return 0
}

func part2(input string) int {
	return 0
}

func parseInputs(input string) [][]string {
	lines := strings.Split(input, "\n")

	grid := [][]string{}
	for r, row := range lines {
		grid = append(grid, []string{})
		for _, v := range row {
			grid[r] = append(grid[r], string(v))
		}
	}

	return grid
}

type Board struct {
	grid    [][]string
	goblins []*Character
	elves   []*Character
}

type Character struct {
	HP   int
	Type string
}
