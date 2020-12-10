package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algo"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	nums := parseInput(input)
	nums = append(nums, mathutil.MaxInt(nums...)+3)
	sort.Ints(nums)

	var oneDiff, threeDiff int
	var currentJoltage int
	for _, v := range nums {
		switch v - currentJoltage {
		case 1: // check for 1 diff first, so no adapters are skipped
			oneDiff++
		case 3:
			threeDiff++
		default:
			panic("adpaters not connected by 3 or 1")
		}
		currentJoltage = v
	}

	return oneDiff * threeDiff
}

func part2(input string) int {
	nums := parseInput(input)
	nums = append(nums, mathutil.MaxInt(nums...)+3)
	sort.Ints(nums)

	// return dynamicProgramming(input)
	return memoCountPossibilities(nums, 0)
}

func parseInput(input string) []int {
	var ans []int

	lines := strings.Split(input, "\n")
	for _, l := range lines {
		ans = append(ans, mathutil.StrToInt(l))
	}

	return ans
}

// storing memo in global state isn't ideal... but it's fastser to code
var memo = map[string]int{}

func memoCountPossibilities(nums []int, lastJolt int) int {
	// if in memo, return that value
	str := makeMemoKey(nums, lastJolt)
	if v, ok := memo[str]; ok {
		return v
	}

	// if all adapters used up, return 1
	if len(nums) == 0 {
		return 1
	}

	// create a recursive call for each adapter within 3 of the lastJoltage
	var count int
	for i, v := range nums {
		if v-lastJolt <= 3 {
			count += memoCountPossibilities(nums[i+1:], v)
		} else { // stop counting if the joltage diff is too larger (>3)
			break
		}
	}

	// update memo
	memo[str] = count

	return count
}
func makeMemoKey(nums []int, lastJolt int) string {
	ans := algo.IntToStr(lastJolt) + "x"
	for _, v := range nums {
		ans += algo.IntToStr(v)
	}
	return ans
}

func dynamicProgramming(input string) int {
	nums := parseInput(input)
	nums = append(nums, mathutil.MaxInt(nums...)+3, 0)
	sort.Ints(nums)

	// initialize table with "1 way" to get to zero jolts
	table := make([]int, len(nums))
	table[0] = 1

	for i := 1; i < len(nums); i++ {
		currentJolts := nums[i]
		for j := i - 1; j >= 0; j-- {
			// add the ways to get to currentJolts that are within 3 jolts
			if currentJolts-nums[j] <= 3 {
				table[i] += table[j]
			} else {
				break
			}
		}
	}

	return table[len(table)-1]
}
