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
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)

	var increased int
	for i, num := range parsed {
		if i > 0 {
			if num > parsed[i-1] {
				increased++
			}
		}
	}

	return increased
}

func part2(input string) int {
	var increased int
	parsed := parseInput(input)

	for i := 0; i < len(parsed)-3; i++ {
		left := parsed[i] + parsed[i+1] + parsed[i+2]
		right := parsed[i+1] + parsed[i+2] + parsed[i+3]
		if right > left {
			increased++
		}
	}

	return increased
}

func parseInput(input string) (ans []int) {
	for _, l := range strings.Split(input, "\n") {
		ans = append(ans, cast.ToInt(l))
	}
	return ans
}
