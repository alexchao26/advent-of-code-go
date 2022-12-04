package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	lines := parseInput(input)
	ans := 0

	for _, l := range lines {
		if doesPair2ContainPair1(l[:2], l[2:]) || doesPair2ContainPair1(l[2:], l[:2]) {
			ans++
		}
	}

	return ans
}

func doesPair2ContainPair1(pair1, pair2 []int) bool {
	return pair1[0] >= pair2[0] && pair1[1] <= pair2[1]
}

func part2(input string) int {
	lines := parseInput(input)
	ans := 0

	for _, l := range lines {
		if doesOverlap(l[:2], l[2:]) {
			ans++
		}
	}

	return ans
}

func doesOverlap(pair1, pair2 []int) bool {
	// sort
	if pair1[0] > pair2[0] {
		pair1, pair2 = pair2, pair1
	}
	return pair1[1] >= pair2[0]
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		leftParts := strings.Split(parts[0], "-")
		rightParts := strings.Split(parts[1], "-")
		ans = append(ans, []int{
			cast.ToInt(leftParts[0]), cast.ToInt(leftParts[1]),
			cast.ToInt(rightParts[0]), cast.ToInt(rightParts[1]),
		})
	}
	return ans
}
