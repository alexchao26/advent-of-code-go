// TODO: pull combinations, multiplying a slice of ints, into algos and mathy packages?
package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := balancingPackages(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func balancingPackages(input string, part int) int {
	var nums []int
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, cast.ToInt(line))
	}

	sum := mathy.SumIntSlice(nums)

	target := sum / 3
	if part == 2 {
		target = sum / 4
	}

	var individualGroups [][]int
	for groupLen := 2; len(individualGroups) == 0; groupLen++ {
		for i := 0; i < len(nums); i++ {
			individualGroups = append(individualGroups, combinations(nums, groupLen, []int{}, i)...)
		}

		// validate that a combination adds up to the target sum
		var validGroups [][]int
		for _, gr := range individualGroups {
			if mathy.SumIntSlice(gr) == target {
				validGroups = append(validGroups, gr)
			}
		}
		// reassign individual groups, if len(validGroups) == 0; we need to rerun
		// with a larger groupLen
		individualGroups = validGroups
	}

	individualGroups = sortGroups(individualGroups)

	return quantumEntanglement(individualGroups[0])
}

func combinations(nums []int, length int, combo []int, index int) [][]int {
	if len(combo) == length {
		return [][]int{append([]int{}, combo...)}
	}
	var combos [][]int
	for i := index; i < len(nums); i++ {
		combo = append(combo, nums[i])
		with := combinations(nums, length, combo, i+1)
		combos = append(combos, with...)
		// backtrack
		combo = combo[:len(combo)-1]
	}

	return combos
}

func sortGroups(groups [][]int) [][]int {
	sort.Slice(groups, func(i, j int) bool {
		return quantumEntanglement(groups[i]) < quantumEntanglement(groups[j])
	})
	return groups
}

func quantumEntanglement(nums []int) int {
	prod := 1
	for _, n := range nums {
		prod *= n
	}
	return prod
}
