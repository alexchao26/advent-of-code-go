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
	rows := parseInput(input)
	var checksum int
	for _, r := range rows {
		checksum += mathutil.MaxInt(r...) - mathutil.MinInt(r...)
	}
	return checksum
}

func part2(input string) int {
	rows := parseInput(input)
	var sumDivisible int
	for _, r := range rows {
		for i, val1 := range r {
			for _, val2 := range r[i+1:] {
				if val1%val2 == 0 {
					sumDivisible += val1 / val2
					break
				} else if val2%val1 == 0 {
					sumDivisible += val2 / val1
					break
				}
			}
		}
	}
	return sumDivisible
}

func parseInput(input string) (ans [][]int) {
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		ans = append(ans, []int{})
		// split by tabs
		for _, num := range strings.Split(l, "\t") {
			ans[i] = append(ans[i], mathutil.StrToInt(num))
		}
	}
	return ans
}
