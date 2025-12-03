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
	pairs := parseInput(input)

	total := 0
	for _, p := range pairs {
		for i := p[0]; i <= p[1]; i++ {
			str := cast.ToString(i)
			// check str is even length (probably not necessary...)
			if len(str)%2 != 0 {
				continue
			}
			// directly check first half to second half of str
			if str[:len(str)/2] == str[len(str)/2:] {
				total += i
			}
		}
	}

	return total
}

func part2(input string) int {
	pairs := parseInput(input)

	total := 0
	for _, p := range pairs {
		for i := p[0]; i <= p[1]; i++ {
			str := cast.ToString(i)

			for l := 1; l <= len(str)/2; l++ {
				// skip if this chunk size does not divide evenly into the entire str
				if len(str)%l != 0 {
					continue
				}

				// compare chunk repeated the correct number of times against the entire str
				chunk := str[:l]
				if str == strings.Repeat(chunk, len(str)/l) {
					total += i
					break
				}
			}
		}
	}

	return total
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, ",") {
		nums := strings.Split(line, "-")
		ans = append(ans, []int{
			cast.ToInt(nums[0]),
			cast.ToInt(nums[1]),
		})
	}
	return ans
}
