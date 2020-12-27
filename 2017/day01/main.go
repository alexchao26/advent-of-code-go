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

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	digits := parseInput(input)
	var sum int
	for i := 0; i < len(digits); i++ {
		if digits[i] == digits[(i+1)%len(digits)] {
			sum += digits[i]
		}
	}
	return sum
}

func part2(input string) int {
	digits := parseInput(input)
	var sum int
	offset := len(digits) / 2
	for i := 0; i < len(digits); i++ {
		if digits[i] == digits[(i+offset)%len(digits)] {
			sum += digits[i]
		}
	}
	return sum
}

func parseInput(input string) (ans []int) {
	for _, num := range strings.Split(input, "") {
		ans = append(ans, cast.ToInt(num))
	}
	return ans
}
