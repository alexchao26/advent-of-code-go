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
	parsed := parseInputs(input)

	seen := map[int]bool{}

	for _, n := range parsed {
		if seen[n] {
			return n * (2020 - n)
		}
		seen[2020-n] = true
	}

	return -1 // should not be hit
}

func part2(input string) int {
	parsed := parseInputs(input)

	// O(n^3) is fast enough
	for i := 0; i < len(parsed); i++ {
		for j := i + 1; j < len(parsed); j++ {
			for k := j + 1; k < len(parsed); k++ {
				if parsed[i]+parsed[j]+parsed[k] == 2020 {
					return parsed[i] * parsed[j] * parsed[k]
				}
			}
		}
	}

	return -1 // should not be hit
}

func parseInputs(input string) []int {
	split := strings.Split(input, "\n")

	nums := []int{}
	for _, n := range split {
		nums = append(nums, mathutil.StrToInt(n))
	}

	return nums
}
