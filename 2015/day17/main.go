package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := eggnogCombinations(util.ReadFile("./input.txt"), 150, part)
	fmt.Println("Output:", ans)
}

func eggnogCombinations(input string, target int, part int) int {
	var nums []int
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, cast.ToInt(line))
	}

	allIndexCombinations := backtrack(nums, 0, target, []int{})
	// part 1, just return len
	if part == 1 {
		return len(allIndexCombinations)
	}

	// part 2, get the number of combinations w/ the lowest length
	minLen := math.MaxInt32
	for _, comb := range allIndexCombinations {
		minLen = mathy.MinInt(minLen, len(comb))
	}

	var count int
	for _, comb := range allIndexCombinations {
		if len(comb) == minLen {
			count++
		}
	}

	return count
}

func backtrack(nums []int, startingIndex, remaining int, usedIndices []int) [][]int {
	if remaining == 0 {
		return [][]int{append([]int{}, usedIndices...)}
	}
	if remaining < 0 {
		return nil
	}

	var validReturns [][]int
	for i := startingIndex; i < len(nums); i++ {
		usedIndices = append(usedIndices, i)
		validReturns = append(validReturns, backtrack(nums, i+1, remaining-nums[i], usedIndices)...)
		usedIndices = usedIndices[:len(usedIndices)-1]
	}
	return validReturns
}
