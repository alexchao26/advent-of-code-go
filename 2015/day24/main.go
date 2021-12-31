package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	// first solution to use all the packages?!
	"github.com/alexchao26/advent-of-code-go/algos"
	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
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

	// make the gross assumption that if a group is found, and adds up to the
	// target, the remaining elements will be able to be split into two equal
	// groups. This is not always true, but the inputs are nicely generated
	var individualGroups [][]int
	for groupLen := 2; len(individualGroups) == 0; groupLen++ {
		individualGroups = algos.CombinationsInts(nums, groupLen)

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

func sortGroups(groups [][]int) [][]int {
	clone := append([][]int{}, groups...)
	sort.Slice(clone, func(i, j int) bool {
		return quantumEntanglement(clone[i]) < quantumEntanglement(clone[j])
	})
	return clone
}

var quantumEntanglement = mathy.MultiplyIntSlice
