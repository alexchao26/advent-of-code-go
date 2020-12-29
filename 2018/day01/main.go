package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := part1(util.ReadFile("./input.txt"))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sum += cast.ToInt(line)

	}

	return sum
}
func part2(input string) int {
	var sum int
	seen := map[int]bool{}
	for {
		for _, line := range strings.Split(input, "\n") {
			sum += cast.ToInt(line)

			if seen[sum] {
				return sum
			}
			seen[sum] = true
		}
	}

	panic("expect return from loop")
}
