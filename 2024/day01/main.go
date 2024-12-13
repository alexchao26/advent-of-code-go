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
	list1, list2 := parseInput(input)
	sort.Ints(list1)
	sort.Ints(list2)

	ans := 0
	for i := range len(list1) {
		ans += mathy.AbsInt(list2[i] - list1[i])
	}

	return ans
}

func part2(input string) int {
	list1, list2 := parseInput(input)

	countsList2 := map[int]int{}
	for _, v := range list2 {
		countsList2[v]++
	}

	ans := 0
	for _, v := range list1 {
		ans += v * countsList2[v]
	}
	return ans
}

func parseInput(input string) (list1, list2 []int) {
	for _, line := range strings.Split(input, "\n") {
		nums := strings.Split(line, "   ")
		list1 = append(list1, cast.ToInt(nums[0]))
		list2 = append(list2, cast.ToInt(nums[1]))
	}
	return list1, list2
}
