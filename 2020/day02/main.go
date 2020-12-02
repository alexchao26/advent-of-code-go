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
	lines := strings.Split(input, "\n")

	var ans int

	for _, l := range lines {
		var lower, upper int
		var char, pw string

		_, err := fmt.Sscanf(l, "%d-%d %1s: %s", &lower, &upper, &char, &pw)
		if err != nil {
			panic("scanning line " + err.Error())
		}

		count := 0
		for _, v := range pw {
			if string(v) == char {
				count++
			}
		}
		if count <= upper && count >= lower {
			ans++
		}
	}

	return ans
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	var ans int
	for _, l := range lines {
		var lower, upper int
		var char, pw string

		_, err := fmt.Sscanf(l, "%d-%d %1s: %s", &lower, &upper, &char, &pw)
		if err != nil {
			panic("scanning line " + err.Error())
		}

		count := 0
		if string(pw[lower-1]) == char {
			count++
		}
		if string(pw[upper-1]) == char {
			count++
		}

		if count == 1 {
			ans++
		}
	}

	return ans
}
