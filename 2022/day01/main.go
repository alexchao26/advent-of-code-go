package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
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
	elves := parseInput(input)

	totals := []int{}
	for _, items := range elves {
		totals = append(totals, mathy.SumIntSlice(items))
	}

	return mathy.MaxInt(totals...)
}

func part2(input string) int {
	elves := parseInput(input)

	totals := []int{}
	for _, items := range elves {
		totals = append(totals, mathy.SumIntSlice(items))
	}
	sort.Ints(totals)

	topThree := 0
	for i := 0; i < 3; i++ {
		topThree += totals[len(totals)-1-i]
	}
	return topThree
}

func parseInput(input string) (ans [][]int) {
	for _, group := range strings.Split(input, "\n\n") {
		row := []int{}
		for _, line := range strings.Split(group, "\n") {
			row = append(row, cast.ToInt(line))
		}
		ans = append(ans, row)
	}
	return ans
}
