package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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
	ranges, ingredients := parseInput(input)

	ans := 0
	for _, ing := range ingredients {
		for _, rng := range ranges {
			if ing >= rng[0] && ing <= rng[1] {
				ans++
				break
			}
		}
	}

	return ans
}

func part2(input string) int {
	ans := 0

	ranges, _ := parseInput(input)

	// have to merge ranges first... there's probably a better way to do this but this has always
	// been my first instinct...
	sort.Slice(ranges, func(i, j int) bool {
		r1, r2 := ranges[i], ranges[j]
		if r1[0] == r2[0] {
			return r1[1] < r2[1]
		}
		return r1[0] < r2[0]
	})

	merged := [][]int{}
	for i := 0; i < len(ranges); i++ {
		if i == 0 {
			merged = append(merged, ranges[i])
		} else {
			// if overlap, combine...
			if merged[len(merged)-1][1] >= ranges[i][0] {
				merged[len(merged)-1][1] = max(merged[len(merged)-1][1], ranges[i][1])
			} else {
				// otherwise start new range
				merged = append(merged, ranges[i])
			}
		}
	}

	for _, rng := range merged {
		// eg 5, 6, 7, 8 is 4 numbers 4 = 8-5+1
		ans += rng[1] - rng[0] + 1
	}

	return ans
}

func parseInput(input string) (ranges [][]int, ings []int) {
	parts := strings.Split(input, "\n\n")
	for _, rng := range strings.Split(parts[0], "\n") {
		nums := strings.Split(rng, "-")
		ranges = append(ranges, []int{
			cast.ToInt(nums[0]), cast.ToInt(nums[1]),
		})
	}
	for _, ing := range strings.Split(parts[1], "\n") {
		ings = append(ings, cast.ToInt(ing))
	}
	return ranges, ings
}
