package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
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
	maze := parseInput(input)

	var index, steps int
	for index >= 0 && index < len(maze) {
		maze[index]++
		index += maze[index] - 1
		steps++
	}

	return steps
}

func part2(input string) int {
	maze := parseInput(input)

	var index, steps int
	for index >= 0 && index < len(maze) {
		nextIndex := maze[index]
		if maze[index] >= 3 {
			maze[index]--
		} else {
			maze[index]++
		}
		index += nextIndex
		steps++
	}

	return steps
}

func parseInput(input string) (ans []int) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		ans = append(ans, mathutil.StrToInt(l))
	}
	return ans
}
