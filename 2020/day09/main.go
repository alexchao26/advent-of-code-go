package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"), 25)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"), 25)
		fmt.Println("Output:", ans)
	}
}

func part1(input string, preambleLength int) int {
	nums := parseInput(input)

	for i := preambleLength; i < len(nums); i++ {
		current := nums[i]

		var good bool
		seen := map[int]bool{}
		for back := i - preambleLength; back < i; back++ {
			wanted := current - nums[back]
			if seen[wanted] {
				good = true
				break
			}
			seen[nums[back]] = true
		}
		if !good {
			return nums[i]
		}
	}

	fmt.Println("ERROR: No invalid number found for part 1")
	return -1
}

// better time-complexity solution using a sliding window
func part2(input string, preambleLength int) int {
	numToFind := part1(input, preambleLength)
	nums := parseInput(input)

	var left, right, sum int
	for right < len(nums) {
		switch {
		case left == right:
			sum += nums[right]
			right++
		case sum > numToFind:
			sum -= nums[left]
			left++
		case sum < numToFind:
			sum += nums[right]
			right++
		}
		// if sum found, and more than 1 num is in it, break
		if sum == numToFind && left+1 != right {
			break
		}
	}

	// find smallest and largest values, wasted 20 minutes here when the smallest
	// number was greater than MaxInt16...
	smallest, largest := math.MaxInt32, -math.MaxInt32
	for _, v := range nums[left:right] {
		if smallest > v {
			smallest = v
		}
		if largest < v {
			largest = v
		}
	}

	return smallest + largest
}

func part2BruteForce(input string, preambleLength int) int {
	numToFind := part1(input, preambleLength)
	fmt.Println(numToFind)
	nums := parseInput(input)

	var left, right int
	for i, firstVal := range nums {
		sum := firstVal
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == numToFind {
				left = i
				right = j
				break
			}
		}
	}

	if left == 0 && right == 0 {
		fmt.Println("ERROR: No valid answer found for part 2")
		return -1
	}

	smallest, largest := nums[left], nums[left]
	for _, v := range nums[left:right] {
		if smallest > v {
			smallest = v
		}
		if largest < v {
			largest = v
		}
	}

	return smallest + largest
}

func parseInput(input string) []int {
	var ans []int

	lines := strings.Split(input, "\n")
	for _, l := range lines {
		num, err := strconv.Atoi(l)
		if err != nil {
			panic("parsing numbers in input " + err.Error())
		}
		ans = append(ans, num)
	}

	return ans
}
